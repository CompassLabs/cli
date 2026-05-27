## compass global-markets-perps candles

Get OHLCV candles

### Synopsis

Return OHLCV candles for a single asset and interval.

Wraps Hyperliquid's ``candleSnapshot`` info type and adds the xyz: HIP-3 DEX
prefix server-side so the SDK only needs the bare ticker (e.g. ``AAPL``).
Candle ``time`` is returned in unix seconds, ready for TradingView
lightweight-charts. Cached for 15 seconds per (symbol, interval, window).

```
compass global-markets-perps candles [flags]
```

### Examples

```
  compass global-markets-perps candles --symbol AAPL --interval 1h
```

### Options

```
  -e, --end-time string     Optional end of the candle window in unix milliseconds. Defaults to now.
  -h, --help                help for candles
  -i, --interval string     Candle interval: 1m, 5m, 15m, 1h, 4h, 1d, 1w (options: 1m, 5m, 15m, 1h, 4h, 1d, 1w) [required]
  -l, --limit int           Number of candles to return (max 5000, capped by Hyperliquid).
      --start-time string   Optional start of the candle window in unix milliseconds. If omitted, computed as end_time - limit * interval.
      --symbol string       Asset ticker (e.g. AAPL, GOLD, EUR). The xyz: HIP-3 DEX prefix is added server-side if not already present. [required]
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

* [compass global-markets-perps](compass_global-markets-perps.md)	 - Operations for global-markets-perps
