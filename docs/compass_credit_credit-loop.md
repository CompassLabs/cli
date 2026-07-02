## compass credit credit-loop

Open a leveraged loop

### Synopsis

Open a leveraged loop into an Aave or Morpho market in ONE atomic transaction.

Repeatedly supplies collateral, borrows at the requested loan-to-value, and
swaps the borrow back to collateral. Each iteration's supply uses the swap's
GUARANTEED minimum output (enforced on-chain), so a fill anywhere within the
slippage tolerance can never break a later step; any positive surplus stays
in the Credit Account (the preview reports the bound as estimated_max_dust).
Very large loops relative to pool depth can still exceed the slippage
tolerance through their own cumulative price impact — size accordingly or
raise max_slippage_percent.

The Credit Account must already hold initial_collateral_amount of
collateral_token. For protocol=MORPHO pass a market_id from
/v2/credit/morpho_markets.

```
compass credit credit-loop [flags]
```

### Examples

```
  compass credit credit-loop --owner 0x5e5b00ed886A6879C2B934612D2312975427fcAf --chain ethereum --collateral-token WETH --borrow-token USDC --initial-collateral-amount 1
```

### Options

```
      --body string                        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --borrow-token string                Token borrowed each iteration and swapped back to the collateral token. For MORPHO it must be the market's loan token. [required]
      --chain string                       The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc) [required]
      --collateral-token string            Token supplied as collateral each iteration. Must already be in the Credit Account for the initial amount. For MORPHO it must be the market's collateral token. [required]
  -e, --emode-category string              Aave only: e-mode category to enable before looping (higher LTV for correlated pairs, e.g. ETH-correlated).
  -g, --gas-sponsorship                    If true, returns EIP-712 typed data for gas-sponsored execution instead of an unsigned transaction.
  -h, --help                               help for credit-loop
  -i, --initial-collateral-amount string   JSON value (one of: number | string)
  -l, --loan-to-value string               JSON value (one of: number | string)
      --market-id string                   Morpho only: the bytes32 market id (from /v2/credit/morpho_markets). Required when protocol=MORPHO.
      --max-slippage-percent string        JSON value (one of: number | string)
      --multiplier string                  JSON value (one of: number | string)
      --owner string                       The address that owns the Credit Account. [required]
  -p, --protocol                           Which lending protocol a credit action targets.
                                           
                                           AAVE`` is the default so existing callers (which never send a ``protocol``
                                           field) keep hitting the unchanged Aave code path. ``EULER`` opts in to the
                                           Euler V2 path, where the market is identified by EVK vault address(es).
                                           ``MORPHO`` identifies Morpho Blue lending markets (bytes32 market id) and is
                                           currently read-only: positions and market discovery only — transaction
                                           builders land with the looping work (COM-7106/7107/7108), so transact
                                           endpoints reject it with a 422. (options: AAVE, EULER, MORPHO)
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
