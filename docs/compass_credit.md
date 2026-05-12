## compass credit

Operations for credit

### Synopsis

Operations for credit

```
compass credit [flags]
```

### Options

```
  -h, --help   help for credit
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
* [compass credit credit-balances](compass_credit_credit-balances.md)	 - Get credit account token balances
* [compass credit credit-borrow](compass_credit_credit-borrow.md)	 - Borrow against collateral
* [compass credit credit-bundle](compass_credit_credit-bundle.md)	 - Execute multiple credit actions
* [compass credit credit-create-account](compass_credit_credit-create-account.md)	 - Create credit account
* [compass credit credit-positions](compass_credit_credit-positions.md)	 - List credit positions
* [compass credit credit-repay](compass_credit_credit-repay.md)	 - Repay debt and withdraw collateral
* [compass credit credit-transfer](compass_credit_credit-transfer.md)	 - Transfer tokens to/from Credit Account
