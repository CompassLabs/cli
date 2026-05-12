## compass traditional-investing traditional-investing-opportunities

List traditional investing perp markets

### Synopsis

List available traditional investing perpetual futures markets with key metrics.

Returns perp markets (stocks, commodities, forex) with open interest,
24h volume, funding rate, and max leverage. Supports filtering by category,
minimum OI/volume, and sorting.

Only traditional investing assets are returned — crypto perps are excluded.

```
compass traditional-investing traditional-investing-opportunities [flags]
```

### Examples

```
  compass traditional-investing traditional-investing-opportunities
```

### Options

```
  -c, --category string            Filter by asset category: stock, commodity, forex (options: stock, commodity, forex)
  -h, --help                       help for traditional-investing-opportunities
      --min-open-interest string   Filter by minimum open interest in USD
      --min-volume-24h string      Filter by minimum 24h volume in USD
      --sort-by string             Sort results by this field (options: open_interest, volume_24h, funding_rate)
      --sort-order string          Sort direction (options: asc, desc)
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

* [compass traditional-investing](compass_traditional-investing.md)	 - Operations for traditional-investing
