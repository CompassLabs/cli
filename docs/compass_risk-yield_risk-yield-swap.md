## compass risk-yield risk-yield-swap

Swap tokens

### Synopsis

Trade one token for another from inside your Risk Yield Account.

Mostly used to reach the ratio a range needs before opening a position:
a range fixes the proportion of the two assets, and a caller holding one of
them has to swap to it.

There is no aggregator on this chain, so the route is a Uniswap v3 pool. The
API quotes every enabled fee tier and takes whichever returns the most — a
1% pool with depth beats a 0.05% pool with none, which here is the common
case rather than the corner one.

The transaction is simulated before it is returned. A failing simulation is
a 422 rather than something to sign: on this chain a token may tax
transfers, block an address or be paused, and none of that is visible until
the call is actually made.

```
compass risk-yield risk-yield-swap [flags]
```

### Examples

```
  compass risk-yield risk-yield-swap --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --token-in USDG --token-out WETH --amount-in 100
```

### Options

```
  -a, --amount-in string       JSON value (one of: number | string)
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string           Risk Yield is available on Robinhood Chain only. (options: robinhood)
      --deadline-seconds int   How long the transaction stays valid once built.
  -f, --fee-ppm string         Pin the route to one fee tier. Leave unset to let the API pick the tier that returns the most.
  -g, --gas-sponsorship        Return EIP-712 typed data for the owner to sign instead of a transaction, so a sponsor can broadcast it.
  -h, --help                   help for risk-yield-swap
      --owner string           The wallet that owns the Risk Yield Account. [required]
  -p, --preview                Return the plan and its simulation without building a transaction. Nothing is signed and nothing can be broadcast.
  -s, --slippage-pct string    JSON value (one of: number | string)
      --token-in string        What you are spending. [required]
      --token-out string       What you want. [required]
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

* [compass risk-yield](compass_risk-yield.md)	 - Operations for risk-yield
