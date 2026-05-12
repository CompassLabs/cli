## compass traditional-investing traditional-investing-execute

Execute signed action

### Synopsis

Submit a signed Hyperliquid action for execution.

Accepts the signature from any prepare endpoint (market_order, limit_order,
cancel_order, withdraw, approve_builder_fee) and POSTs it to the Hyperliquid
exchange API.

```
compass traditional-investing traditional-investing-execute [flags]
```

### Examples

```
  compass traditional-investing traditional-investing-execute --action '{"key":"<value>","key1":"<value>","key2":"<value>"}' --nonce 31516 --signature <value>
```

### Options

```
  -a, --action string          Raw Hyperliquid action from the prepare step [required]
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                   help for traditional-investing-execute
  -n, --nonce int              Nonce from the prepare step [required]
  -s, --signature string       User's EIP-712 signature (hex, 65 bytes) [required]
  -v, --vault-address string   Optional vault address
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
