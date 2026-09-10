## compass risk-yield risk-yield-mint

Open a position

### Synopsis

Provide liquidity to a pool over a price range.

The range you ask for is not quite the range you get: ticks are discrete, so
a 30% band becomes the usable ticks that bracket it — snapped outward, never
inward, because a position that stops earning sooner than you expected is
the worse surprise. The `range` in the response is what will actually be
opened.

A range also fixes the proportion of the two assets, so a deposit rarely
uses all of both. What does not fit is reported as `leftover` and stays in
your account rather than being quietly swept.

Set `preview` to see the plan and its simulation without building anything.

```
compass risk-yield risk-yield-mint [flags]
```

### Examples

```
  compass risk-yield risk-yield-mint --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --pool '{"pool_id":1}' --range '{"type":"symmetric_pct","width_pct":"30"}' --deposit '{"type":"both","amount0":"100","amount1":"0.04"}'
```

### Options

```
      --body string                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string                 Risk Yield is available on Robinhood Chain only. (options: robinhood)
      --deadline-seconds int         How long the transaction stays valid once built.
      --deposit string               JSON value (variants: both: { amount0: value, amount1: value }, single: { token: string, amount: value }, usd: { value_usd: value })
      --deposit.both string          BothSidesDeposit variant as JSON
      --deposit.single string        SingleSidedDeposit variant as JSON
      --deposit.usd string           UsdDeposit variant as JSON
  -g, --gas-sponsorship              Return EIP-712 typed data for the owner to sign instead of a transaction, so a sponsor can broadcast it.
  -h, --help                         help for risk-yield-mint
      --owner string                 The wallet that owns the Risk Yield Account. [required]
      --pool string                  Which pool, by id or by its key.
                                     
                                     A pool id is stable and is what the list endpoints return. The long form
                                     exists so a caller who knows the pair can address a pool without a lookup —
                                     on a chain minting twenty thousand pools a day, a client should not have to
                                     search for one it just created. [required]
      --preview                      Return the plan and its simulation without building a transaction. Nothing is signed and nothing can be broadcast.
      --range string                 JSON value (variants: full: object, prices: { price_lower: value, price_upper: value }, symmetric_pct: { width_pct: value }, ticks: { tick_lower: integer, tick_upper: integer })
      --range.full string            FullRange variant as JSON
      --range.prices string          PriceRange variant as JSON
      --range.symmetric-pct string   SymmetricRange variant as JSON
      --range.ticks string           TickRange variant as JSON
      --range.ticks.tick-lower int   [required]
      --range.ticks.tick-upper int   [required]
      --recipient string             Who receives the position NFT. Defaults to the Risk Yield Account, which is what makes it manageable through this API.
  -s, --slippage-pct string          JSON value (one of: number | string)
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
