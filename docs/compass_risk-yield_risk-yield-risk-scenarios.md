## compass risk-yield risk-yield-risk-scenarios

Model impermanent loss

### Synopsis

What this position would do if the price moved — and how likely that is.

A fee APR on a memecoin pair means nothing without the loss it is paid
against, which is what this measures.

Each scenario takes **two** independent moves, and keeping them apart is the
point. `ratio_move_pct` is how far the pair's ratio travels, and that is
what causes impermanent loss. `quote_usd_move_pct` is how far the stock
moves in dollars, which scales what the position is worth without touching
the ratio at all — a stock-only move produces an impermanent loss of exactly
zero. "NVDA drops 10%" and "the memecoin drops 10% against NVDA" are
different questions, and one number cannot ask both.

The two most useful columns are `breakeven_fee_apr_pct` — what this position
would have to earn for the loss to be worth taking, which you can compare
against the pool's actual fee APR — and `days_of_fees_to_recover_il`, which
is null when the pool pays liquidity providers nothing, because no amount of
time fixes that.

`probabilistic` runs the pair's own measured volatility through a driftless
simulation. Driftless deliberately: nobody knows which way a memecoin goes,
and a model that assumed one would be predicting returns rather than
measuring risk. It is absent when the pair has no measured volatility —
which a pool with a few hours of history does not — because a distribution
built on zero would report that nothing can happen.

```
compass risk-yield risk-yield-risk-scenarios [flags]
```

### Examples

```
  compass risk-yield risk-yield-risk-scenarios --pool '{"pool_id":1}' --range '{"type":"symmetric_pct","width_pct":"30"}' --deposit '{"type":"usd","value_usd":"1000"}'
```

### Options

```
      --body string                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string                  options: robinhood
      --deposit string                JSON value (variants: both: { amount0: value, amount1: value }, single: { token: string, amount: value }, usd: { value_usd: value })
      --deposit.both string           BothSidesDeposit variant as JSON
      --deposit.single string         SingleSidedDeposit variant as JSON
      --deposit.usd string            UsdDeposit variant as JSON
  -f, --fee-apr-pct-override string   JSON value (one of: number | string)
  -h, --help                          help for risk-yield-risk-scenarios
      --horizon-days int              integer value
  -m, --monte-carlo                   Also estimate how likely each outcome is, not just its size.
  -n, --n-paths int                   integer value
  -p, --pool string                   Which pool, by id or by its key.
                                      
                                      A pool id is stable and is what the list endpoints return. The long form
                                      exists so a caller who knows the pair can address a pool without a lookup —
                                      on a chain minting twenty thousand pools a day, a client should not have to
                                      search for one it just created. [required]
  -r, --range string                  JSON value (variants: full: object, prices: { price_lower: value, price_upper: value }, symmetric_pct: { width_pct: value }, ticks: { tick_lower: integer, tick_upper: integer })
      --range.full string             FullRange variant as JSON
      --range.prices string           PriceRange variant as JSON
      --range.symmetric-pct string    SymmetricRange variant as JSON
      --range.ticks string            TickRange variant as JSON
      --range.ticks.tick-lower int    [required]
      --range.ticks.tick-upper int    [required]
      --scenarios string              Leave empty for a default spread of moves.
      --seed int                      Fixed so the same request gives the same answer.
  -v, --vol-override-pct string       JSON value (one of: number | string)
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
