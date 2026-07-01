## compass tokenized-assets market

Get a single market

### Synopsis

Get extended detail for a single tokenized market.

Works for both asset families: an Ondo **equity** (e.g. `TSLAon`) or a
Midas **RWA yield token** (e.g. `mTBILL`). Equities add 52-week range,
volume, market cap, holder count, and tradable sessions on top of the
`/markets` fields; RWA-yield entries instead carry `apy_7d`/`apy_30d` and
`tvl_usd`.

**OHLC candles are an equities-only feature** — opt in by passing both
`interval` and `range` query params to include a `candles` array. They
must be provided together and must form one of the supported pairs:

- `1min` / `5min` / `15min` with `range=1day`
- `1hour` / `4hour` with `range=1month`
- `12hour` with `range=3month`
- `1day` with `range=3month` / `6month` / `1year` / `all`

Omitting both returns the market detail without `candles`.

```
compass tokenized-assets market [flags]
```

### Examples

```
  compass tokenized-assets market --symbol <value>
```

### Options

```
  -c, --chain string     Network the market is deployed on (defaults to Ethereum). A token deployed on multiple chains (e.g. Midas RWA on Ethereum and Base) is resolved per chain; 404 if the symbol isn't deployed there. Ondo equities are Ethereum-only. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc)
  -h, --help             help for market
  -i, --interval range   Optional candle interval. Must be paired with range and form a valid `(interval, range)` pair to include OHLC candles in the response. (options: 1min, 5min, 15min, 1hour, 4hour, 12hour, 1day)
  -r, --range interval   Optional lookback window. Must be paired with interval and form a valid `(interval, range)` pair to include OHLC candles in the response. (options: 1day, 1month, 3month, 6month, 1year, all)
  -s, --symbol string    [required]
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

* [compass tokenized-assets](compass_tokenized-assets.md)	 - Operations for tokenized-assets
