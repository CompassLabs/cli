## compass tokenized-assets market

Get a market

### Synopsis

Get extended detail for a single market — an Ondo **equity** (e.g. `TSLAon`), a
Midas **RWA yield token** (e.g. `mTBILL`), or an IXS **managed vault** (e.g.
`ixv1`).

Adds richer market data on top of the `/markets` listing, plus an optional
OHLC candle series: pass matching `interval` and `range` query params to
include `candles` (available for equities, Midas tokens except `mBTC`, and IXS
vaults; omit both for detail without candles).

```
compass tokenized-assets market [symbol] [flags]
```

### Examples

```
  compass tokenized-assets market --symbol <value>
```

### Options

```
  -c, --chain string     Network the market is deployed on (defaults to Ethereum). A token deployed on multiple chains (e.g. Midas RWA on Ethereum and Base) is resolved per chain; 404 if the symbol isn't deployed there. Ondo equities are Ethereum-only. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc, robinhood, ethereum_sepolia)
  -h, --help             help for market
  -i, --interval range   Optional candle interval. Must be paired with range and form a valid `(interval, range)` pair to include OHLC candles in the response. (options: 1min, 5min, 15min, 1hour, 4hour, 12hour, 1day)
  -r, --range interval   Optional lookback window. Must be paired with interval and form a valid `(interval, range)` pair to include OHLC candles in the response. (options: 1day, 1month, 3month, 6month, 1year, all)
  -s, --symbol string    string value (or pass it as the [symbol] argument)
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key-auth string    Your Compass API Key. Get your key [here](https://www.compasslabs.ai/dashboard).
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
      --interactive            Prompt for missing inputs and open guided configure/auth forms (forms fall back to line prompts on stdin off-TTY) (default true)
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [compass tokenized-assets](compass_tokenized-assets.md)	 - Operations for tokenized-assets

### Machine interface

* `compass tokenized-assets market --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass tokenized-assets market --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
