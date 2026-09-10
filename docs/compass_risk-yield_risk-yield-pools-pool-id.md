## compass risk-yield risk-yield-pools-pool-id

Get pool detail

### Synopsis

One pool, with its price history and the caveats that apply to it.

For a pool holding a tokenized stock, `stock_reference` compares the pool's
price against the stock's own. They drift: the pool trades around the clock
while the stock trades in a session, and that gap is what a liquidity
provider is quoting against overnight.

```
compass risk-yield risk-yield-pools-pool-id [flags]
```

### Examples

```
  compass risk-yield risk-yield-pools-pool-id --pool-id 174842 --chain robinhood
```

### Options

```
  -c, --chain string         options: robinhood [required]
      --depth-words int      How many 256-tick words either side of the current price to read for the liquidity distribution. More is a wider picture and more node calls.
  -h, --help                 help for risk-yield-pools-pool-id
      --history string       options: 24h, 7d, 30d
  -p, --pool-id GET /pools   The pool id, as returned by GET /pools. [required]
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
