# Risk analysis recipes

Reproduce the metrics from compasslabs.ai/risk using existing CLI commands. No new endpoints, no specialized tooling — the dashboard is a thin client over three public endpoints, and the math is reproducible from the same data this CLI already exposes.

Source of truth for the math: `landing_page/lib/risk/calculations.ts` and `landing_page/lib/risk/cascade.ts`. If a formula here ever drifts from those files, **those files win**.

---

## The three data sources

All risk metrics derive from these three endpoints. The CLI flags below match what the landing-page dashboard uses by default (chain=Base, no min-TVL filter; tighten with `--min-tvl-usd` for faster turns on big datasets).

| Data | Command | Returns |
|------|---------|---------|
| Morpho vaults — TVL, APY, per-market allocations | `compass earn earn-vaults --order-by tvl_usd --chain '"base"' --jq '.vaults'` | Array of vaults. Each has `tvl_usd`, `apy_7d`, `apy_30d`, `apy_90d`, `liquidity_usd`, `factory_version`, and `markets[]` — per-market allocations with `collateral_token_symbol`, `loan_token_symbol`, `lltv` (wei-scaled 1e18), `allocation_pct`, `market_id`. |
| Aave markets — for cross-protocol overlap + supply APY | `compass earn earn-aave-markets --chain '"base"' --jq '.markets'` | Object keyed by token symbol. Each entry has `chains.<chain>.supply_apy`, `chains.<chain>.borrow_apy` (already in percent-scale, e.g. `"3.4"` = 3.4%). |
| Pendle markets — for yield decomposition | `compass earn earn-pendle-markets --chain '"base"' --jq '.markets'` | Array with `pt_symbol`, `implied_apy`, `tvl_usd`, `expiry`. |

Reminder (Rule 1 in AGENTS.md): `--chain` on these three commands is JSON-quoted (`'"base"'`), not plain. Forgetting the quotes gives `invalid character 'b' looking for beginning of value`.

---

## Metric 1 — Borrower exposure (JTD)

**Question agents are asked:** "What is this USDC actually lent against?" / "Which collaterals dominate Morpho borrower exposure on Base?"

**Formula** (from `computeJtd` in `calculations.ts`):

For each vault, for each entry in `vault.markets[]`:
- Skip if `allocation_pct <= 0`.
- Skip if `collateral_token_symbol == "0x0000000000000000000000000000000000000000"` (zero address = idle/cash, not real collateral).
- `exposure_usd = vault.tvl_usd × (allocation_pct / 100)`

Group by `collateral_token_symbol`. Per group: sum `exposure_usd`, count distinct vaults, record the top vault by exposure. Sort by total exposure descending.

**One-shot jq aggregate:**

```bash
compass earn earn-vaults --order-by tvl_usd --chain '"base"' \
  --jq '
    .vaults
    | [.[]
       | (.tvl_usd | tonumber) as $tvl
       | (.name // .vault_address) as $vault
       | .markets[]
       | select((.allocation_pct // 0) > 0
                and .collateral_token_symbol != "0x0000000000000000000000000000000000000000")
       | {sym: .collateral_token_symbol,
          vault: $vault,
          exposure: ($tvl * (.allocation_pct / 100))}]
    | group_by(.sym)
    | [.[] | {collateral: .[0].sym,
              exposed_tvl: ([.[].exposure] | add),
              vault_count: ([.[].vault] | unique | length),
              top_vault: (max_by(.exposure) | .vault)}]
    | sort_by(-.exposed_tvl)
  '
```

---

## Metric 2 — LLTV cascade (price-shock stress test)

**Question:** "What's the bad-debt exposure if cbBTC drops 30%?" / "Which markets cross their liquidation threshold at a -X% shock?"

**This is the only metric where math correctness materially matters.** Agents commonly mis-handle the wei-scaling or fall back to a binary LLTV trip-flag — both produce wrong answers. Use the formula below exactly.

**Constants:**
- `LLTV_DIVISOR = 1e18` — Morpho ships LLTV as a uint256 string scaled by 1e18.
- `CURRENT_LTV_FACTOR = 0.85` — assumed average borrower utilization. The API does not yet expose `average_borrow_ltv` per market, so we approximate every borrower as sitting at 85% of LLTV. **This inflates the impact at high shock sizes — surface the caveat in your answer.**
- `MIN_HIT_USD = 1` — suppress sub-dollar hits as rounding noise.

**Formula** (from `simulateShock` in `cascade.ts`):

For each vault, for each `vault.markets[]` entry where `collateral_token_symbol == <target>`:
1. `lltv_ratio = parseFloat(market.lltv) / 1e18` — yields a 0..1 ratio.
2. `current_ltv = lltv_ratio × 0.85`
3. `new_ltv = current_ltv / (1 − shock_pct)`, where `shock_pct ∈ [0, 0.99]`. A `-30%` price drop is `shock_pct = 0.30`.
4. `bad_debt_fraction = new_ltv > 1 ? (1 − 1 / new_ltv) : 0`
   - **Common bug to avoid:** do NOT use a binary `new_ltv ≥ lltv_ratio` trip-flag and report full exposure when tripped. That over-estimates and makes every result identical past the trip point. Use the continuous fraction above.
5. `exposure_usd = vault.tvl_usd × (allocation_pct / 100)`
6. `hit_usd = exposure_usd × bad_debt_fraction`. Skip rows where `hit_usd < 1`.

Aggregate `hit_usd` per vault (sum) and per `market_id` (sum). The total across all vaults is "vault capital under-collateralised at this shock".

**Always disclose the methodology when you return cascade numbers:**
> Back-of-envelope estimate. Assumes all loans sit at 85% of LLTV across the board (`average_borrow_ltv` per market not yet exposed in the API). No TWAP smoothing, oracle deviation, partial liquidations, or liquidation incentives modelled. Read as worst-case bad debt, not "first liquidation triggered".

---

## Metric 3 — Vault correlation matrix

**Question:** "Do vaults A and B actually diversify, or do they share the same underlying markets?"

**Formula** (from `computeCorrelationMatrix` in `calculations.ts`). The landing page restricts this to the top 10 active vaults (those with `allocation_pct > 0` on at least one market) — do the same for a comparable result, or extend to all vaults if asked.

For each vault `v`, build a collateral-weight map:
- For each `v.markets[m]` where `allocation_pct > 0`:
  - `w_v[m.collateral_token_symbol] += v.tvl_usd × (m.allocation_pct / 100)`

For each pair `(i, j)`:
1. `total_i = sum(w_i.values())`, `total_j = sum(w_j.values())`.
2. For each `sym` in `w_i ∩ w_j` (collaterals both vaults hold), if `total_i > 0` and `total_j > 0`:
   - `shared += min(w_i[sym] / total_i, w_j[sym] / total_j)`
3. `overlap_pct = round(shared × 100)`

Interpretation buckets used by the heatmap (from `landing_page/lib/risk/format.ts` `heatBucket6`):

| overlap_pct | Bucket |
|------|--------|
| ≥ 90 | Critical — "different vault" is essentially rebranding |
| 70–89 | High |
| 50–69 | Medium-high |
| 30–49 | Medium |
| 10–29 | Low |
| < 10 | Minimal — meaningful diversification |

The 4-label legend on `/risk` (Low / Medium / High / Critical) is a visual simplification; the actual heatmap uses the 6 buckets above.

---

## Metric 4 — Vault economics (yield decomposition + withdrawable)

**Question:** "Where does this vault's APY come from?" / "How much of my deposit can I pull right now?"

### Yield decomposition (from `computeVaultEconomics` in `calculations.ts`)

Pendle PT collateral is the only mechanically-distinct yield source the API exposes today. Everything else is bucketed as "borrower demand" — including idle Aave reserves (an approximation; the API doesn't yet tag idle reserves separately).

For each vault `v`:
1. Build a lookup `pt_apy_by_symbol`: from `earn-pendle-markets` response, map each `pt_symbol` → `implied_apy` (normalize: if raw value < 1, multiply by 100; Pendle ships percent-scale like `"12.5"`, but guard against decimal-scale).
2. For each `v.markets[m]` where `collateral_token_symbol` starts with `PT-` and `allocation_pct > 0`:
   - `capital_m = v.tvl_usd × (m.allocation_pct / 100)`
   - If `pt_apy_by_symbol[m.collateral_token_symbol]` exists:
     - `pt_apy_contribution_usd += capital_m × pt_apy_by_symbol[m.collateral_token_symbol]`
3. `pendle_apy_contribution = pt_apy_contribution_usd / v.tvl_usd`
4. `headline_apy = v.apy_7d ?? v.apy_30d ?? 0`
5. If `headline_apy > 0 && pendle_apy_contribution > 0`:
   - `pendle_share_pct = min(100, round(pendle_apy_contribution / headline_apy × 100))`
6. Else `pendle_share_pct = 0`
7. `borrower_share_pct = 100 − pendle_share_pct`

**Disclose with the numbers:** vaults parking idle capital in Aave look like "borrower" in this split because the API doesn't tag idle reserves. Treat the split as directional, not exact.

### Withdrawable right now

- `withdrawable_pct_of_tvl = (v.liquidity_usd / v.tvl_usd) × 100`

A value of 100% means you can pull your entire deposit immediately; lower values mean some capital is currently lent and you'd need to wait for repayments or accept a partial withdrawal.

---

## Worked recipes

### Recipe A — "What breaks if cbBTC drops 30%?"

```bash
# 1. Pull vault data once, cache it
compass earn earn-vaults --order-by tvl_usd --chain '"base"' -o json > /tmp/vaults.json

# 2. Extract the cbBTC markets across all vaults with the fields the cascade needs
jq '
  .vaults[]
  | (.tvl_usd | tonumber) as $tvl
  | (.name // .vault_address) as $vault
  | .markets[]
  | select(.collateral_token_symbol == "cbBTC" and (.allocation_pct // 0) > 0)
  | {vault: $vault,
     market_id,
     pair: (.collateral_token_symbol + "/" + .loan_token_symbol),
     lltv_ratio: ((.lltv | tonumber) / 1e18),
     exposure_usd: ($tvl * (.allocation_pct / 100))}
' /tmp/vaults.json
```

Then apply Metric 2's cascade formula with `shock_pct = 0.30` and sum `hit_usd` per vault and per market. Report the total bad-debt USD + the top affected vaults + the markets that crossed LLTV (`new_ltv > lltv_ratio` means "at the trip"; `new_ltv > 1` means "underwater"). Always include the 85%-LLTV caveat.

### Recipe B — "Are Steakhouse Prime USDC and Gauntlet USDC Prime actually diversified?"

```bash
# 1. Pull both vaults' market allocations
compass earn earn-vaults --order-by tvl_usd --chain '"base"' \
  --jq '.vaults[] | select(.name == "Steakhouse Prime USDC" or .name == "Gauntlet USDC Prime")
        | {name, tvl_usd, markets: [.markets[] | select((.allocation_pct // 0) > 0)
                                  | {sym: .collateral_token_symbol, alloc: .allocation_pct}]}'
```

Compute Metric 3's overlap between the two collateral-weight maps. Return the single overlap %, the bucket (Low/Medium/High/Critical), and the top 3 collaterals driving the shared exposure.

### Recipe C — "How concentrated is USDC lending on Base against a single collateral?"

```bash
# 1. JTD aggregation, filtered to USDC vaults
compass earn earn-vaults --order-by tvl_usd --chain '"base"' --asset-symbol '"USDC"' \
  --jq '
    .vaults
    | (map(.tvl_usd | tonumber) | add) as $total
    | [.[] | (.tvl_usd | tonumber) as $tvl
            | .markets[]
            | select((.allocation_pct // 0) > 0
                     and .collateral_token_symbol != "0x0000000000000000000000000000000000000000")
            | {sym: .collateral_token_symbol, exposure: ($tvl * (.allocation_pct / 100))}]
    | group_by(.sym)
    | map({collateral: .[0].sym, exposed_tvl: ([.[].exposure] | add)})
    | sort_by(-.exposed_tvl)
    | {total_usdc_tvl: $total,
       top_collateral: .[0].collateral,
       top_collateral_share_pct: (.[0].exposed_tvl / $total * 100),
       breakdown: .}
  '
```

A `top_collateral_share_pct` over 40% means most USDC borrower demand on Base flows through a single asset — flag it as concentration risk.

---

## When to recommend the dashboard instead

If the user wants a visual representation (Sankey flow, heatmap, slider, drag-to-explore), point them at <https://compasslabs.ai/risk> — the CLI gives identical raw data but no visualization. Use the CLI when they want answers in chat, scripts, or agentic workflows; use the dashboard for visual exploration.

---

## Source of truth

If anything in this file looks stale or contradicts the dashboard, check:
- `landing_page/lib/risk/calculations.ts` — JTD, correlation, yield decomposition, rate sensitivity.
- `landing_page/lib/risk/cascade.ts` — LLTV cascade simulator.
- `landing_page/lib/risk/types.ts` — exact field shapes for vault / market / pendle responses.

Those TypeScript files are the canonical implementation. This doc rephrases them in formula form so an agent can compute without reading the TS source. When the landing page evolves its math (e.g., when `average_borrow_ltv` lands and the 85% assumption is dropped), update this file.
