## compass risk-yield risk-yield-tokens

List tokens

### Synopsis

The assets that trade on Robinhood Chain.

Tokenized stocks carry their issuer metadata, including the corporate-action
multiplier — balances do not rebase on a split, so a quote for the
underlying only becomes a token price once that is applied.

Memecoins carry the launch they came from. Every price carries the source it
was resolved from and a confidence: below 0.6 it came from a pool shallow
enough that one trade moves it, and every dollar figure derived from it is
indicative.

```
compass risk-yield risk-yield-tokens [flags]
```

### Examples

```
  compass risk-yield risk-yield-tokens --chain robinhood
```

### Options

```
  -c, --chain string               options: robinhood [required]
      --direction string           options: asc, desc
      --has-price string           Only tokens the indexer can currently price.
  -h, --help                       help for risk-yield-tokens
  -k, --kind string                Filter to one kind of asset. (options: STOCK, STABLE, NATIVE, WRAPPED_NATIVE, MAJOR, LAUNCHPAD_MEME, OTHER)
      --launchpad string           options: PONS_V1, PONS_V2, POOLS_TRADE, BAGS
      --limit int                  The number of items to return.
  -m, --min-liquidity-usd string   JSON value (one of: number | string)
      --offset int                 The offset of the first item to return.
      --order-by string            options: liquidity_usd, volume_24h_usd, market_cap_usd, price_change_24h_pct, created_at
      --search string              Matches symbol, name or stock ticker.
      --stock-ticker string        string value
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
