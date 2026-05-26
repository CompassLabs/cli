## compass tokenized-equities tokenized-equities-markets

List tokenized equity markets

### Synopsis

List the tokenized US equities available on Ethereum.

Each entry includes the symbol, the underlying ticker, the on-chain
contract address, the latest USD price, and 24h price change. Filter
by `category` (sector tag) or `search` (substring match against symbol,
ticker, or name).

Only Ethereum-deployed tokens are returned; assets that exist only on
other chains are omitted.

```
compass tokenized-equities tokenized-equities-markets [flags]
```

### Examples

```
  compass tokenized-equities tokenized-equities-markets
```

### Options

```
  -c, --category string   Filter markets by category (e.g. 'tech', 'finance').
  -h, --help              help for tokenized-equities-markets
  -s, --search string     Case-insensitive substring match against the on-chain symbol, underlying ticker, and underlying name.
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

* [compass tokenized-equities](compass_tokenized-equities.md)	 - Operations for tokenized-equities
