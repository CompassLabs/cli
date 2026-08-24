## compass credit unloop

Unwind a leveraged loop

### Synopsis

Unwind a leveraged position in ONE atomic transaction.

Withdraw collateral, swap it back to the borrow token, repay, repeat — until
the position reaches target_multiplier, or until the debt is cleared exactly
if you omit it. If any step fails the whole transaction reverts.

See the [Leveraged Looping guide](https://docs.compasslabs.ai/v2/Products/Looping)
for chain coverage, vault-share collateral, shared Aave reserves and the full
error list.

```
compass credit unloop [flags]
```

### Examples

```
  compass credit unloop --owner 0x0E407CdeBD8e078E6966ef6740540d25F5897082 --chain ethereum --collateral-token WETH --borrow-token USDC
```

### Options

```
  -a, --allow-partial                 If the target cannot be reached in one transaction, return the maximum-progress plan (preview.fully_unwound=false) instead of a 400. This can happen because every withdrawal is sized to keep the health factor at or above 1.02 at that step, so a position opened very close to its liquidation threshold runs out of headroom before the target is met. A second unloop call, now from a much lower leverage, finishes the job.
      --body string                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --borrow-token string           Token borrowed in the loop; withdrawn collateral is swapped back to it and used to repay. For MORPHO it must be the market's loan token. [required]
      --borrow-vault string           Euler only: the EVK vault the loop borrowed from (the sub-account's controller). Required when protocol=EULER.
      --chain string                  Blockchain network. Not every protocol is deployed on every chain — see the protocol field — and a chain with no credit venue at all returns a 422 naming the chains that do. (options: arbitrum, base, bsc, ethereum, hyperevm, tempo) [required]
      --collateral-token string       Token supplied as collateral in the loop being unwound. For MORPHO it must be the market's collateral token. [required]
      --collateral-vault string       Euler only: the EVK vault the loop's collateral is in. Required when protocol=EULER.
  -g, --gas-sponsorship               If true, returns EIP-712 typed data for gas-sponsored execution instead of an unsigned transaction.
  -h, --help                          help for unloop
      --market-id string              Morpho only: the bytes32 market id (from /v2/credit/morpho_markets). Required when protocol=MORPHO.
      --max-slippage-percent string   JSON value (one of: number | string)
      --owner string                  The address that owns the Credit Account. [required]
      --preview                       If true, build a display estimate only — no single-use firm quote is ever spent. NOTE that this guarantees only that no firm quote was spent — it does not guarantee an absent transaction: on an unwind no firm provider covers, and under pricing=market, the call still falls through to the aggregator and returns a signable transaction. On a firm-covered unwind (pricing 'auto' or 'firm') the numbers are firm-INDICATIVE, priced off the firm provider's live maker levels; otherwise they are market-priced with slippage-bounded floors. Under pricing='firm' a preview no firm provider can serve returns the advisory alone (preview=null + firm_available). Set it on every parameter-exploration call and omit it only on the build the user is about to sign.
      --pricing string                Swap-leg routing policy. 'auto': firm quotes where a firm venue covers the pair, transparent fallback to the market aggregator otherwise. 'firm': never price on the market route — previews the firm venue cannot serve return the firm_available advisory alone (preview=null, zero aggregator calls), and executions fail with a typed error instead of silently substituting market pricing. 'market': never route through the firm venue; every leg is priced by the aggregator. max_slippage_percent applies on EVERY policy — unlike the loop, firm unwinds consume it to size the guaranteed withdraw/repay floors (the fills themselves are exact). 'firm' is incompatible with gas_sponsorship (sponsored unwinds force market routing). (options: auto, firm, market)
      --protocol                      Which lending protocol a credit action targets.
                                      
                                      AAVE`` is the default so existing callers (which never send a ``protocol``
                                      field) keep hitting the unchanged Aave code path; markets are named by token
                                      symbol. ``MORPHO`` identifies Morpho Blue lending markets by their bytes32
                                      ``market_id``. ``EULER`` identifies Euler V2 markets by their EVK
                                      ``collateral_vault`` + ``borrow_vault`` addresses and supports isolated
                                      per-sub-account positions (``sub_account_id``).
                                      
                                      Deployment is per chain, so a valid protocol can still 422 on a given chain:
                                      AAVE on Ethereum, Base, Arbitrum, BSC and HyperEVM (where it is Hyperlend, the
                                      chain's Aave V3 deployment); MORPHO on Ethereum, Base, Arbitrum and HyperEVM
                                      (where it is Felix); EULER on Ethereum, Base, Arbitrum and BSC.
                                      
                                      All three support ``/v2/credit/loop`` and ``/v2/credit/unloop``. EULER does
                                      NOT: ``/v2/credit/rebalance`` rejects it with a 422, and
                                      ``/v2/credit/looped_positions`` covers only AAVE and MORPHO — an Euler loop is
                                      silently absent there rather than an error, so read it from
                                      ``/v2/credit/positions`` instead. (EULER still appears in the
                                      ``looped_positions`` response enum because this enum is shared; it is never
                                      emitted.) (options: AAVE, EULER, MORPHO)
  -s, --sub-account-id int            Euler only: the EVC sub-account (0-255) holding the looped position to unwind. 0 is the Credit Account itself.
  -t, --target-multiplier string      JSON value (one of: number | string)
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key-auth string    Your Compass API Key. Get your key [here](https://www.compasslabs.ai/dashboard).
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [compass credit](compass_credit.md)	 - Operations for credit
