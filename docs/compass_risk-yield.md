## compass risk-yield

Operations for risk-yield

### Synopsis

Operations for risk-yield

```
compass risk-yield [flags]
```

### Options

```
  -h, --help   help for risk-yield
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
* [compass risk-yield risk-yield-collect](compass_risk-yield_risk-yield-collect.md)	 - Collect fees
* [compass risk-yield risk-yield-create-account](compass_risk-yield_risk-yield-create-account.md)	 - Create account
* [compass risk-yield risk-yield-decrease](compass_risk-yield_risk-yield-decrease.md)	 - Withdraw from a position
* [compass risk-yield risk-yield-increase](compass_risk-yield_risk-yield-increase.md)	 - Add to a position
* [compass risk-yield risk-yield-mint](compass_risk-yield_risk-yield-mint.md)	 - Open a position
* [compass risk-yield risk-yield-pools](compass_risk-yield_risk-yield-pools.md)	 - List pools
* [compass risk-yield risk-yield-pools-pool-id](compass_risk-yield_risk-yield-pools-pool-id.md)	 - Get pool detail
* [compass risk-yield risk-yield-positions](compass_risk-yield_risk-yield-positions.md)	 - List LP positions
* [compass risk-yield risk-yield-rebalance](compass_risk-yield_risk-yield-rebalance.md)	 - Rebalance a position
* [compass risk-yield risk-yield-risk-scenarios](compass_risk-yield_risk-yield-risk-scenarios.md)	 - Model impermanent loss
* [compass risk-yield risk-yield-swap](compass_risk-yield_risk-yield-swap.md)	 - Swap tokens
* [compass risk-yield risk-yield-tokens](compass_risk-yield_risk-yield-tokens.md)	 - List tokens
* [compass risk-yield risk-yield-transfer](compass_risk-yield_risk-yield-transfer.md)	 - Transfer tokens to/from account
