## compass global-markets-perps activity

Aggregated Hyperliquid activity for a user

### Synopsis

Return positions, fills, open orders, and (optionally) builder approval
state for an end-user in one normalized payload.

Each section is fetched in parallel from the Hyperliquid `info` API.
If a single upstream call fails the corresponding section returns ``null``
and an entry is added to ``partial_errors``; if every section fails the
endpoint responds with 502.

Pass ``builder`` to additionally include the user's current approved max
fee rate for that builder (used by dashboards to decide whether to prompt
the user to sign an `approveBuilderFee` action).

```
compass global-markets-perps activity [flags]
```

### Examples

```
  compass global-markets-perps activity --owner 0x01E62835dd7F52173546A325294762143eE4a882
```

### Options

```
  -b, --builder string   Optional builder address. When provided, the response includes the current builder-fee approval state for this (owner, builder) pair.
  -h, --help             help for activity
      --owner string     End-user EOA whose activity should be fetched. [required]
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
