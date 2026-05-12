## compass tokenized-assets tokenized-assets-order-order-hash

Get order status

### Synopsis

Get the lifecycle state of a submitted order.

The `status` field is one of `pending`, `filled`, `expired`, or
`cancelled`. Partial fills stay in `pending` while `filled_amount` is
populated as fills come in; once an order fully fills, `fill_tx_hash`
is also returned.

Upstream protocol states beyond these four (e.g. `partially-filled`,
`refunded`) are mapped onto this set.

```
compass tokenized-assets tokenized-assets-order-order-hash [flags]
```

### Examples

```
  compass tokenized-assets tokenized-assets-order-order-hash --order-hash <value>
```

### Options

```
  -h, --help                help for tokenized-assets-order-order-hash
      --order-hash string   [required]
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
