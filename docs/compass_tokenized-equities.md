## compass tokenized-equities

Operations for tokenized-equities

### Synopsis

Operations for tokenized-equities

```
compass tokenized-equities [flags]
```

### Options

```
  -h, --help   help for tokenized-equities
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

* [compass](compass.md)	 - Compass API: Compass Labs DeFi API
* [compass tokenized-equities create-account](compass_tokenized-equities_create-account.md)	 - Create a Tokenized Equities Account
* [compass tokenized-equities market](compass_tokenized-equities_market.md)	 - Get a single market
* [compass tokenized-equities markets](compass_tokenized-equities_markets.md)	 - List tokenized equity markets
* [compass tokenized-equities order](compass_tokenized-equities_order.md)	 - Build a buy/sell order
* [compass tokenized-equities order-cancel](compass_tokenized-equities_order-cancel.md)	 - Cancel an unfilled order
* [compass tokenized-equities order-status](compass_tokenized-equities_order-status.md)	 - Get order status
* [compass tokenized-equities order-submit](compass_tokenized-equities_order-submit.md)	 - Submit a signed order
* [compass tokenized-equities positions](compass_tokenized-equities_positions.md)	 - Get tokenized-equity positions for an owner
* [compass tokenized-equities quote](compass_tokenized-equities_quote.md)	 - Preview a buy/sell quote
