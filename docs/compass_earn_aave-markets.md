## compass earn aave-markets

List aave markets

### Synopsis

List Aave lending markets with supply and borrow rates.

Returns rates organized by token symbol, with chain-specific data for each token. Each token includes rates for all chains where it's available, plus information about which chain offers the highest supply APY.

APY values are returned in percentage format (e.g., 4.5 means 4.5%). Tokens with zero APY on both supply and borrow are excluded.

To deposit into an Aave market, use the [manage endpoint](https://docs.compasslabs.ai/v2/api-reference/earn/manage-earn-position) with `venue_type=AAVE`.

```
compass earn aave-markets [flags]
```

### Examples

```
  compass earn aave-markets
```

### Options

```
  -c, --chain string          Optional chain filter. If not provided, returns rates for all chains. (options: base, ethereum, arbitrum, hyperevm)
      --days supply_apy_avg   Window in days used to compute supply_apy_avg / `borrow_apy_avg`. Mirrors the `days` parameter of the v1 `/v1/aave/avg_rate` endpoint.
  -h, --help                  help for aave-markets
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
