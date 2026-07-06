## compass credit credit-unloop

Unwind a leveraged loop

### Synopsis

Unwind an Aave or Morpho loop in ONE atomic transaction.

Repeatedly withdraws collateral, swaps it to the borrow token at a GUARANTEED
minimum output (enforced on-chain), and repays. The floor discipline means a
swap filling anywhere within the slippage tolerance can never break a later
step; any positive surplus stays in the Credit Account as borrow-token dust
(the preview reports the bound as estimated_max_dust).

Omit target_multiplier for a full close: the debt is cleared exactly —
accrued interest included — and the pair collateral is returned to the Credit
Account. Pass 1 to clear the debt but keep the collateral supplied, or a value
above 1 to delever to that multiplier (it must be below the position's current
multiplier).

Each withdrawal is sized to keep the position's health factor ≥ 1.02 at that
step, so a position opened very close to the liquidation threshold may need
more than one transaction to fully close — set allow_partial=true to return
the maximum single-transaction progress (preview.fully_unwound=false), then
call unloop again to finish. Very large unwinds relative to pool depth can
still exceed the slippage tolerance through their own cumulative price impact.

For protocol=MORPHO pass a market_id from /v2/credit/morpho_markets; inspect
open loops via /v2/credit/looped_positions.

```
compass credit credit-unloop [flags]
```

### Examples

```
  compass credit credit-unloop --owner 0x0E407CdeBD8e078E6966ef6740540d25F5897082 --chain ethereum --collateral-token WETH --borrow-token USDC
```

### Options

```
  -a, --allow-partial                 If the target cannot be reached in one transaction (e.g. a position opened very close to the liquidation threshold), return the maximum-progress plan (preview.fully_unwound=false) instead of a 400. A second unloop call, now from a much lower leverage, finishes the job.
      --body string                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --borrow-token string           Token borrowed in the loop; withdrawn collateral is swapped back to it and used to repay. For MORPHO it must be the market's loan token. [required]
      --chain string                  The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc) [required]
      --collateral-token string       Token supplied as collateral in the loop being unwound. For MORPHO it must be the market's collateral token. [required]
  -g, --gas-sponsorship               If true, returns EIP-712 typed data for gas-sponsored execution instead of an unsigned transaction.
  -h, --help                          help for credit-unloop
      --market-id string              Morpho only: the bytes32 market id (from /v2/credit/morpho_markets). Required when protocol=MORPHO.
      --max-slippage-percent string   JSON value (one of: number | string)
      --owner string                  The address that owns the Credit Account. [required]
  -p, --protocol                      Which lending protocol a credit action targets.
                                      
                                      AAVE`` is the default so existing callers (which never send a ``protocol``
                                      field) keep hitting the unchanged Aave code path. ``EULER`` opts in to the
                                      Euler V2 path, where the market is identified by EVK vault address(es).
                                      ``MORPHO`` identifies Morpho Blue lending markets (bytes32 market id) and is
                                      currently read-only: positions and market discovery only — transaction
                                      builders land with the looping work (COM-7106/7107/7108), so transact
                                      endpoints reject it with a 422. (options: AAVE, EULER, MORPHO)
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
