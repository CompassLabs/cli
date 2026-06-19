## compass tokenized-assets

Operations for tokenized-assets

### Synopsis

Operations for tokenized-assets

```
compass tokenized-assets [flags]
```

### Options

```
  -h, --help   help for tokenized-assets
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
* [compass tokenized-assets buy](compass_tokenized-assets_buy.md)	 - Buy an RWA yield token (Midas: mTBILL/mBASIS/mBTC)
* [compass tokenized-assets create-account](compass_tokenized-assets_create-account.md)	 - Create a Tokenized Assets Account
* [compass tokenized-assets market](compass_tokenized-assets_market.md)	 - Get a single market
* [compass tokenized-assets markets](compass_tokenized-assets_markets.md)	 - List tokenized asset markets
* [compass tokenized-assets order](compass_tokenized-assets_order.md)	 - Build a tokenized-equity buy/sell order (Ondo)
* [compass tokenized-assets order-cancel](compass_tokenized-assets_order-cancel.md)	 - Cancel an unfilled tokenized-equity order (Ondo)
* [compass tokenized-assets order-status](compass_tokenized-assets_order-status.md)	 - Get tokenized-equity order status (Ondo)
* [compass tokenized-assets order-submit](compass_tokenized-assets_order-submit.md)	 - Submit a signed tokenized-equity order (Ondo)
* [compass tokenized-assets positions](compass_tokenized-assets_positions.md)	 - Get tokenized-asset positions for an owner
* [compass tokenized-assets quote](compass_tokenized-assets_quote.md)	 - Preview a tokenized-equity buy/sell quote (Ondo)
* [compass tokenized-assets sell](compass_tokenized-assets_sell.md)	 - Sell an RWA yield token (Midas: mTBILL/mBASIS/mBTC)
* [compass tokenized-assets tokenized-assets-balances](compass_tokenized-assets_tokenized-assets-balances.md)	 - Get account balances
* [compass tokenized-assets tokenized-assets-order-order-hash-charge-fee](compass_tokenized-assets_tokenized-assets-order-order-hash-charge-fee.md)	 - Charge a partner fee on a filled sell order's USDC proceeds
* [compass tokenized-assets transfer](compass_tokenized-assets_transfer.md)	 - Deposit to / withdraw from a Tokenized Assets Account
