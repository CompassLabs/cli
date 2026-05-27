# compass CLI — Guide for AI Agents

This file is for AI coding agents (Claude Code, Cursor, Aider, OpenAI Codex, etc.) that interact with the Compass DeFi API through the `compass` CLI in a shell.

If that's you: read this top-to-bottom once before issuing commands. Most of the errors that bite first-time runs are documented here.

---

## What `compass` is

`compass` is a thin command-line wrapper over the Compass Labs DeFi API. The API is **non-custodial** — it returns unsigned transactions or EIP-712 typed data. Your wallet (or the user's) signs and broadcasts. The CLI never touches private keys.

Two execution patterns:

| Pattern | What `compass` returns | Who broadcasts |
|---------|------------------------|----------------|
| **Standalone** | Unsigned transaction (`{to, data, value, chainId}`) | The owner's wallet |
| **Gas-sponsored** | EIP-712 typed data | A sponsor wallet, after the owner signs the typed data and the signature is submitted to `gas-sponsorship-prepare` |

You'll usually pick standalone unless the user explicitly wants gas sponsorship.

---

## Setup (do this once)

```bash
# 1. Set the API key (env var is the most reliable for agents)
export COMPASS_API_KEY_AUTH=ck_...   # the env var name has _AUTH suffix

# 2. Verify
compass whoami
```

**Notes for agents:**
- The env var is `COMPASS_API_KEY_AUTH`, not `COMPASS_API_KEY`. Mixing these up is a common error.
- `compass configure` opens an interactive TUI — don't call it from non-interactive contexts. Set the env var instead.
- The CLI auto-detects agent environments (`CLAUDE_CODE`, `CURSOR_AGENT`) and switches to a structured-error / TOON-output mode. You normally don't need `--agent-mode` explicitly.

---

## Read the per-command doc before composing a command

There are 30+ commands and many use **nested flags** (`--venue.vault.vault-address`) that aren't obvious from endpoint names. Always read the relevant doc file first:

```bash
cat cli-sdk/docs/compass_<group>_<command>.md
# e.g.
cat cli-sdk/docs/compass_earn_earn-manage.md
```

Or run `compass <group> <command> --help` if the binary is installed.

**Do not infer flag names from the API URL or endpoint name.** Speakeasy's CLI generator produces nested flags from request body schemas (e.g., `--venue.vault.vault-address`, not `--vault-address`).

---

## Critical rules — internalize these before issuing commands

### Rule 1 — Use `--dry-run` for safe exploration

`--dry-run` prints the request that would be sent (URL, headers, body) without contacting the API. Use this when you're constructing a command from user instructions and want to verify shape before spending an API call:

```bash
compass earn manage \
  --venue.vault.vault-address 0x7BfA7C4f149E7415b73bdeDfe609237e29CBF34A \
  --action DEPOSIT --amount 0.01 --owner 0x... --chain base \
  --dry-run
```

### Rule 2 — Prefer TOON output to save tokens

If the user is going to read the result, `--output-format pretty` is fine. If the result will go back into your context for another step, prefer `--output-format toon` (30–60% fewer tokens than JSON) or `--jq '.field'` to extract the precise field you need.

In auto-detected agent mode, TOON is already the default — no flag needed.

### Rule 3 — Trust the Description, ignore the metavar

Some flag listings show metavars like `--action venue`, `--gas-sponsorship true`, `--amount from_token`, `--from-token TSLAon`. These are placeholders the generator inferred from OpenAPI examples; they are **not** part of the syntax. Always read the **Description** column for the actual semantics.

---

## Recipes

### Recipe 1 — Discover the highest-yielding USDC vault on Base, then deposit

```bash
# Step 1: find candidates (note the JSON-quoted optional --chain)
compass earn vaults \
  --chain base \
  --asset-symbol USDC \
  --order-by tvl_usd \
  --direction desc \
  --limit 5 \
  --jq '.vaults[] | {name, vault_address, apy, tvl_usd}'

# Step 2: pick a vault address from the output, then prepare the deposit
compass earn manage \
  --venue.vault.vault-address 0x<chosen> \
  --action DEPOSIT \
  --amount 100 \
  --owner 0x<user-address> \
  --chain base
# This returns an unsigned transaction. Hand it to the user's wallet to sign.
```

### Recipe 2 — List a user's earn positions across all chains

```bash
compass earn positions-all --owner 0x<user-address>
```

(`earn positions` requires `--chain`; `earn-positions-all` doesn't.)

### Recipe 3 — Gas-sponsored deposit (advanced)

```bash
# Step 1: prepare with gas_sponsorship=true → returns EIP-712 typed data
compass earn manage \
  --venue.vault.vault-address 0x<vault> \
  --action DEPOSIT --amount 100 --owner 0x<owner> --chain base \
  --gas-sponsorship true \
  --output-format json \
  --jq '.eip712'

# Step 2: owner signs the typed data off-chain (in their wallet)

# Step 3: submit signature + typed data to gas-sponsorship-prepare
compass gas-sponsorship prepare \
  --owner 0x<owner> --chain base --product earn \
  --eip-712 '<the typed data from step 1>' \
  --signature 0x<signature from step 2> \
  --sender 0x<sponsor>
# Returns an unsigned tx. Sponsor signs and broadcasts; sponsor pays gas.
```

### Recipe 4 — Quote and place a tokenized-equity order

```bash
# Step 1: preview
compass tokenized-equities quote \
  --chain base --owner 0x<owner> \
  --from-token USDC --to-token TSLAon --amount 100

# Step 2: build the order (carry recommended_slippage_bps from quote)
compass tokenized-equities order \
  --chain base --owner 0x<owner> \
  --from-token USDC --to-token TSLAon --amount 100 \
  --slippage-bps <value-from-step-1>

# Step 3: owner signs the returned order_message off-chain

# Step 4: submit signed order
compass tokenized-equities order-submit \
  --signed-order '<order_message from step 2>' \
  --order-hash '<order_hash from step 2>'
```

### Recipe 5 — Bundle multiple actions atomically

```bash
# Bundle a swap + deposit in one transaction. The --chain here is REQUIRED
# (FlagKindEnum), so it accepts plain values without JSON quotes.
compass earn bundle \
  --owner 0x<owner> --chain ethereum \
  --actions '[
    {"body":{"action_type":"V2_SWAP","token_in":"USDC","token_out":"AUSD","amount_in":"100","slippage":"0.5"}},
    {"body":{"action_type":"V2_MANAGE","venue":{"type":"VAULT","vault_address":"0x..."},"action":"DEPOSIT","amount":"100"}}
  ]'
```

### Recipe 6 — Compute risk metrics from existing data

Agents asked about portfolio risk, vault correlation, cross-protocol exposure, or "what breaks at -X%" can answer those questions from the existing `earn-vaults` / `earn-aave-markets` / `earn-pendle-markets` endpoints — no specialized risk command needed. The math mirrors the dashboard at <https://compasslabs.ai/risk>: borrower exposure (JTD), LLTV cascade stress test, vault correlation matrix, yield decomposition, withdrawable share of TVL.

See [`internal/cli/risk/RISK_RECIPES.md`](internal/cli/risk/RISK_RECIPES.md) for the formulas, jq snippets, and worked example recipes. Or — if you only have the installed binary and no repo/web access — run `compass risk-recipes` to print the same doc to stdout (it's embedded at compile time). The one piece worth reading carefully is the **LLTV cascade** — Morpho ships LLTV as a uint256 scaled by 1e18, and the bad-debt formula is continuous (`1 − 1/newLtv`), not a binary trip-flag. Getting either of those wrong gives plausible-looking numbers that are off by a lot.

---

## Error-recovery cheat sheet

| Symptom | Diagnosis | Fix |
|---------|-----------|-----|
| `unknown flag: --foo` | Inferred flag name doesn't exist | Re-read `cli-sdk/docs/compass_<command>.md`, look for nested form (`--venue.vault.foo`) |
| `missing required flag: --foo` | Required flag missing | Add it; check the doc for nested required flags |
| `HTTP 401 ... API key missing or invalid` | Auth | `export COMPASS_API_KEY_AUTH=...` then re-run |
| `HTTP 422` | Request body validation | Read the response body for which field failed; nested objects often need explicit type discriminator (e.g., `{"type":"VAULT", ...}`) |
| `HTTP 4xx` other than above | Domain logic (insufficient balance, invalid address, etc.) | Surface to user verbatim — usually actionable |
| `HTTP 5xx` | Backend issue | Retry once with `--debug`; if persistent, surface to user |
| `command not found: --foo` (zsh) | Trailing whitespace after `\` line continuation broke parsing | Re-run as a single line, or strip trailing spaces |

---

## Output format selection cheat sheet

| Goal | Flag combination |
|------|------------------|
| Show user a result they'll read | default (`pretty`) — already enabled |
| Get one field or array for the next step | `--jq '.path'` (always emits JSON; ignores `-o`) |
| Render an envelope as a table | `--output-format table` (accepts the envelope; nested arrays not supported in one shot) |
| Compact result for your own context | `--output-format toon` (or rely on agent-mode default) |
| Stream a paginated list | `--all` (with `--output-format json` for NDJSON or `toon` for blocks) |

---

## What `compass` does NOT do

- **Sign or broadcast transactions.** That's the wallet's job (or `gas-sponsorship-prepare` flow).
- **Track positions over time.** It's a stateless API client. For positions, query `earn positions`/`credit positions`/`tokenized-equities positions` each time.
- **Manage on-chain approvals.** Allowance setup is the wallet's responsibility (or use Permit2 via gas-sponsorship `approve-transfer`).

---

## Reference links

- Full command list & global flag reference: [`docs/compass.md`](docs/compass.md)
- Per-command docs: [`docs/`](docs/)
- API docs: <https://docs.compasslabs.ai>
- README (human-oriented overview): [`README.md`](README.md)
