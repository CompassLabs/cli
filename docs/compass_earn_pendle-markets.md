## compass earn pendle-markets

List pendle markets

### Synopsis

List Pendle yield trading markets with TVL and implied APY.

Returns Pendle market data including Principal Token (PT), Standardized Yield (SY),
and Yield Token (YT) addresses, along with market expiry, TVL, and implied APY.

Use this endpoint to discover yield trading opportunities, compare rates across
Pendle markets, or build market selection interfaces.

APY values are returned in percentage format (e.g., 5.25 means 5.25%). Markets
without complete metadata or statistics are excluded.

```
compass earn pendle-markets [flags]
```

### Examples

```
  compass earn pendle-markets --order-by tvl_usd
```

### Options

```
  -c, --chain string               Optional chain filter. If not provided, returns markets for all chains. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc)
      --direction string           The direction to order the results by. (options: asc, desc)
  -h, --help                       help for pendle-markets
  -l, --limit int                  The number of items to return.
  -m, --min-tvl-usd string         JSON value (one of: number | string)
      --offset int                 The offset of the first item to return.
      --order-by string            The field to order the results by. [required]
  -u, --underlying-symbol string   Filter markets by underlying asset symbol (e.g., 'USDC', 'WETH').
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

* [compass earn](compass_earn.md)	 - Operations for earn
