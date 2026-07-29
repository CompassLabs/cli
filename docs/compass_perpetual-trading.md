## compass perpetual-trading

Operations for perpetual-trading

### Synopsis

Operations for perpetual-trading

```
compass perpetual-trading [flags]
```

### Options

```
  -h, --help   help for perpetual-trading
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
* [compass perpetual-trading activity](compass_perpetual-trading_activity.md)	 - Aggregated Hyperliquid activity for a user
* [compass perpetual-trading approve-builder-fee](compass_perpetual-trading_approve-builder-fee.md)	 - Approve builder fee
* [compass perpetual-trading cancel-order](compass_perpetual-trading_cancel-order.md)	 - Cancel order
* [compass perpetual-trading candles](compass_perpetual-trading_candles.md)	 - Get OHLCV candles
* [compass perpetual-trading deposit](compass_perpetual-trading_deposit.md)	 - Deposit USDC to perpetual trading account
* [compass perpetual-trading deposit-sponsor-prepare](compass_perpetual-trading_deposit-sponsor-prepare.md)	 - Build the Bridge2 deposit tx from a signed permit
* [compass perpetual-trading enable-unified-account](compass_perpetual-trading_enable-unified-account.md)	 - Enable unified account mode
* [compass perpetual-trading ensure-leverage](compass_perpetual-trading_ensure-leverage.md)	 - Ensure 1x cross leverage
* [compass perpetual-trading execute](compass_perpetual-trading_execute.md)	 - Execute signed action
* [compass perpetual-trading limit-order](compass_perpetual-trading_limit-order.md)	 - Place limit order
* [compass perpetual-trading market-order](compass_perpetual-trading_market-order.md)	 - Place market order
* [compass perpetual-trading opportunities](compass_perpetual-trading_opportunities.md)	 - List perpetual trading markets
* [compass perpetual-trading positions](compass_perpetual-trading_positions.md)	 - List perpetual trading positions
* [compass perpetual-trading set-leverage](compass_perpetual-trading_set-leverage.md)	 - Set leverage (defaults to market maximum)
* [compass perpetual-trading withdraw](compass_perpetual-trading_withdraw.md)	 - Withdraw USDC from perpetual trading account
