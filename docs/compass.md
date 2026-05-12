## compass

Compass API: Compass Labs DeFi API

### Synopsis

Compass API: Compass Labs DeFi API

```
compass [flags]
```

### Options

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key-auth string    Your Compass API Key. Get your key [here](https://www.compasslabs.ai/dashboard).
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
  -h, --help                   help for compass
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

* [compass auth](compass_auth.md)	 - Manage authentication credentials
* [compass configure](compass_configure.md)	 - Configure authentication credentials and preferences
* [compass credit](compass_credit.md)	 - Operations for credit
* [compass earn](compass_earn.md)	 - Operations for earn
* [compass explore](compass_explore.md)	 - Interactively browse and run commands
* [compass gas-sponsorship](compass_gas-sponsorship.md)	 - Operations for gas-sponsorship
* [compass tokenized-assets](compass_tokenized-assets.md)	 - Operations for tokenized-assets
* [compass traditional-investing](compass_traditional-investing.md)	 - Operations for traditional-investing
* [compass version](compass_version.md)	 - Print the CLI version
* [compass whoami](compass_whoami.md)	 - Display current authentication configuration
