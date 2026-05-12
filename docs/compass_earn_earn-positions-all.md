## compass earn earn-positions-all

List earn positions across all chains

### Synopsis

List all Earn positions across all supported chains (Ethereum, Base, Arbitrum).

Returns positions grouped by chain, with per-chain and total USD values.
Each chain includes Aave, vault, and Pendle PT positions. Chains where
the user has no earn account return empty position lists.

Use this endpoint for a cross-chain portfolio overview instead of making
separate calls per chain to /positions.

```
compass earn earn-positions-all [flags]
```

### Examples

```
  compass earn earn-positions-all --owner 0x01E62835dd7F52173546A325294762143eE4a882
```

### Options

```
  -h, --help           help for earn-positions-all
      --owner string   The address of the owner of the earn account. [required]
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

* [compass earn](compass_earn.md)	 - Operations for earn
