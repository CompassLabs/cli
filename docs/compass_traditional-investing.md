## compass traditional-investing

Operations for traditional-investing

### Synopsis

Operations for traditional-investing

```
compass traditional-investing [flags]
```

### Options

```
  -h, --help   help for traditional-investing
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
* [compass traditional-investing traditional-investing-approve-builder-fee](compass_traditional-investing_traditional-investing-approve-builder-fee.md)	 - Approve builder fee
* [compass traditional-investing traditional-investing-cancel-order](compass_traditional-investing_traditional-investing-cancel-order.md)	 - Cancel order
* [compass traditional-investing traditional-investing-deposit](compass_traditional-investing_traditional-investing-deposit.md)	 - Deposit USDC to traditional investments account
* [compass traditional-investing traditional-investing-enable-unified-account](compass_traditional-investing_traditional-investing-enable-unified-account.md)	 - Enable unified account mode
* [compass traditional-investing traditional-investing-ensure-leverage](compass_traditional-investing_traditional-investing-ensure-leverage.md)	 - Ensure 1x cross leverage
* [compass traditional-investing traditional-investing-execute](compass_traditional-investing_traditional-investing-execute.md)	 - Execute signed action
* [compass traditional-investing traditional-investing-limit-order](compass_traditional-investing_traditional-investing-limit-order.md)	 - Place limit order
* [compass traditional-investing traditional-investing-market-order](compass_traditional-investing_traditional-investing-market-order.md)	 - Place market order
* [compass traditional-investing traditional-investing-opportunities](compass_traditional-investing_traditional-investing-opportunities.md)	 - List traditional investing perp markets
* [compass traditional-investing traditional-investing-positions](compass_traditional-investing_traditional-investing-positions.md)	 - List traditional investing positions
* [compass traditional-investing traditional-investing-withdraw](compass_traditional-investing_traditional-investing-withdraw.md)	 - Withdraw USDC from traditional investments account
