## compass risk-yield risk-yield-positions

List LP positions

### Synopsis

Your liquidity positions: what they hold, what they cost, what they paid.

Only the first of those is on chain. A position reports its liquidity and
its range; it does not report the amounts that opened it or what has already
been collected. Those come from recorded events, and where there are none
`cost_basis.source` is `unavailable` and the PnL fields are null rather than
guessed — an invented cost basis would make every number derived from it
wrong in the same direction.

`hold_value_usd` is what the deposited amounts would be worth if they had
simply been held. The gap between that and `value_usd` is impermanent loss,
and it is the number that decides whether the fees were worth it.

`needs_rebalance` and `rebalance_reasons` say what the API would look at, not
what it will do: rebalancing realizes the loss, pays gas twice and a swap,
and puts the position back at risk from a new price.

```
compass risk-yield risk-yield-positions [flags]
```

### Examples

```
  compass risk-yield risk-yield-positions --chain robinhood --owner 0x06A9aF046187895AcFc7258450B15397CAc67400
```

### Options

```
  -c, --chain string         options: robinhood [required]
      --dex-version string   Which Uniswap deployment a pool belongs to.
                             
                             v3 pools hold their own tokens and pay LPs the fee tier. v4 pools live in a
                             shared PoolManager and may route the fee through a hook, which is where the
                             launchpad pools' zero LP yield comes from. (options: V3, V4)
  -h, --help                 help for risk-yield-positions
  -i, --include-closed       Include positions with no liquidity and nothing owed. Off by default: a closed position is history, not a holding.
      --owner string         The wallet that owns the Risk Yield Account. [required]
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

* [compass risk-yield](compass_risk-yield.md)	 - Operations for risk-yield
