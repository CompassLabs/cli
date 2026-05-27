## compass global-markets-perps

Operations for global-markets-perps

### Synopsis

Operations for global-markets-perps

```
compass global-markets-perps [flags]
```

### Options

```
  -h, --help   help for global-markets-perps
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
* [compass global-markets-perps activity](compass_global-markets-perps_activity.md)	 - Aggregated Hyperliquid activity for a user
* [compass global-markets-perps approve-builder-fee](compass_global-markets-perps_approve-builder-fee.md)	 - Approve builder fee
* [compass global-markets-perps cancel-order](compass_global-markets-perps_cancel-order.md)	 - Cancel order
* [compass global-markets-perps candles](compass_global-markets-perps_candles.md)	 - Get OHLCV candles
* [compass global-markets-perps deposit](compass_global-markets-perps_deposit.md)	 - Deposit USDC to global markets perps account
* [compass global-markets-perps deposit-sponsor-prepare](compass_global-markets-perps_deposit-sponsor-prepare.md)	 - Build the Bridge2 deposit tx from a signed permit
* [compass global-markets-perps enable-unified-account](compass_global-markets-perps_enable-unified-account.md)	 - Enable unified account mode
* [compass global-markets-perps ensure-leverage](compass_global-markets-perps_ensure-leverage.md)	 - Ensure 1x cross leverage
* [compass global-markets-perps execute](compass_global-markets-perps_execute.md)	 - Execute signed action
* [compass global-markets-perps limit-order](compass_global-markets-perps_limit-order.md)	 - Place limit order
* [compass global-markets-perps market-order](compass_global-markets-perps_market-order.md)	 - Place market order
* [compass global-markets-perps opportunities](compass_global-markets-perps_opportunities.md)	 - List global markets perps markets
* [compass global-markets-perps positions](compass_global-markets-perps_positions.md)	 - List global markets perps positions
* [compass global-markets-perps withdraw](compass_global-markets-perps_withdraw.md)	 - Withdraw USDC from global markets perps account
