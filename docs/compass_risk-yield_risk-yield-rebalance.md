## compass risk-yield risk-yield-rebalance

Rebalance a position

### Synopsis

Move a position to a new range, in one transaction.

**`NOOP` is a real answer and often the right one.** Rebalancing realizes
the impermanent loss the position has taken so far, pays gas twice and a
swap, and puts the position back at risk from a new price. If the position
is already centred at the requested width, that is what you will get.

When there is something to do, the transaction unwinds the old position,
sweeps it together with the fees, opens the new range and burns the empty
NFT. Every leg is sized from the **floor** the previous one guarantees
rather than from its quote — they execute in the same transaction, and a
mint sized from a quote reverts when the decrease comes back a wei short,
leaving the owner having paid gas to dismantle a position. `estimated_dust`
is what that safety costs.

The default `vol_scaled` strategy sizes the new range from the pair's own
measured volatility: two standard deviations over the horizon. A range
narrower than the asset's ordinary movement spends most of its life out of
range earning nothing, which is the most common way to lose money doing
this on purpose.

```
compass risk-yield risk-yield-rebalance [flags]
```

### Examples

```
  compass risk-yield risk-yield-rebalance --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --position-id 1076416
```

### Options

```
      --body string               Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --burn-old                  Destroy the emptied NFT once it is unwound.
  -c, --chain string              Risk Yield is available on Robinhood Chain only. (options: robinhood)
      --deadline-seconds int      How long the transaction stays valid once built.
      --dex-version string        Which Uniswap deployment a pool belongs to.
                                  
                                  v3 pools hold their own tokens and pay LPs the fee tier. v4 pools live in a
                                  shared PoolManager and may route the fee through a hook, which is where the
                                  launchpad pools' zero LP yield comes from. (options: V3, V4)
  -g, --gas-sponsorship           Return EIP-712 typed data for the owner to sign instead of a transaction, so a sponsor can broadcast it.
  -h, --help                      help for risk-yield-rebalance
      --horizon-days vol_scaled   For vol_scaled: how long the range should survive an ordinary move. A range narrower than the asset's usual movement spends most of its life out of range earning nothing.
      --owner string              The wallet that owns the Risk Yield Account. [required]
      --position-id int           [required]
      --preview                   Return the plan and its simulation without building a transaction. Nothing is signed and nothing can be broadcast.
      --slippage-pct string       JSON value (one of: number | string)
      --strategy string           How to choose the new range. (options: recenter_symmetric_pct, vol_scaled, manual)
      --tick-lower manual         For manual.
      --tick-upper manual         For manual.
  -w, --width-pct string          JSON value (one of: number | string)
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
