## compass credit morpho-markets

List curated Morpho markets

### Synopsis

List curated Morpho Blue lending markets for a chain.

Morpho Blue is permissionless, so credit actions identify a market by its
bytes32 market id. This returns the curated market set with live LLTV,
supply/borrow APY, utilization, and available liquidity -- read on-chain per
request -- so callers know which market_id to use and what it currently costs.

```
compass credit morpho-markets [flags]
```

### Examples

```
  compass credit morpho-markets --chain ethereum
```

### Options

```
  -c, --chain string       options: arbitrum, base, bsc, ethereum, tempo [required]
      --direction string   Order direction (asc/desc). (options: asc, desc)
  -h, --help               help for morpho-markets
  -l, --limit int          The number of items to return.
      --offset int         The offset of the first item to return.
      --order-by string    Field to order the markets by before paginating. (options: tvl_usd, liquidity_usd, lltv)
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
