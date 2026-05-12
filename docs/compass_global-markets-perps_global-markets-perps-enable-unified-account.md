## compass global-markets-perps global-markets-perps-enable-unified-account

Enable unified account mode

### Synopsis

Check account mode and prepare the enable-unified-account action if needed.

If the account is already in unified mode (or portfolio margin), returns
the current mode with null typed_data — no signing needed. Otherwise,
returns EIP-712 typed data for the user to sign. After signing, submit
the signature via the /execute endpoint.

```
compass global-markets-perps global-markets-perps-enable-unified-account [flags]
```

### Examples

```
  compass global-markets-perps global-markets-perps-enable-unified-account --owner <value>
```

### Options

```
      --body string    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help           help for global-markets-perps-enable-unified-account
      --owner string   User's EOA address [required]
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
