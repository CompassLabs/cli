## compass earn

Operations for earn

### Synopsis

Operations for earn

```
compass earn [flags]
```

### Options

```
  -h, --help   help for earn
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

* [compass](compass.md)	 - Compass API: Compass Labs DeFi API
* [compass earn aave-markets](compass_earn_aave-markets.md)	 - List aave markets
* [compass earn balances](compass_earn_balances.md)	 - Get token balances
* [compass earn bundle](compass_earn_bundle.md)	 - Execute multiple earn actions
* [compass earn create-account](compass_earn_create-account.md)	 - Create earn account
* [compass earn manage](compass_earn_manage.md)	 - Manage earn position
* [compass earn pendle-markets](compass_earn_pendle-markets.md)	 - List pendle markets
* [compass earn positions](compass_earn_positions.md)	 - List earn positions
* [compass earn positions-all](compass_earn_positions-all.md)	 - List earn positions across all chains
* [compass earn swap](compass_earn_swap.md)	 - Swap tokens within Earn Account
* [compass earn swap-quote](compass_earn_swap-quote.md)	 - Quote a swap (read-only)
* [compass earn transfer](compass_earn_transfer.md)	 - Transfer tokens to/from account
* [compass earn vaults](compass_earn_vaults.md)	 - List vaults
