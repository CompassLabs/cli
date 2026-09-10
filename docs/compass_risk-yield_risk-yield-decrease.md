## compass risk-yield risk-yield-decrease

Withdraw from a position

### Synopsis

Take liquidity back out of a position.

The collect is bundled in by default and should stay that way:
`decreaseLiquidity` only *credits* the tokens to the position, it does not
move them. Without the collect the position looks emptied while the funds
sit in the position manager.

```
compass risk-yield risk-yield-decrease [flags]
```

### Examples

```
  compass risk-yield risk-yield-decrease --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --position-id 1076416
```

### Options

```
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --burn-if-empty               Destroy the NFT once it holds nothing. Frees a little gas.
      --chain string                Risk Yield is available on Robinhood Chain only. (options: robinhood)
      --collect decreaseLiquidity   Sweep the withdrawn tokens and any fees in the same transaction. Leave this on: decreaseLiquidity only credits tokens, it does not move them, so without a collect nothing reaches your account.
      --deadline-seconds int        How long the transaction stays valid once built.
      --dex-version string          Which Uniswap deployment a pool belongs to.
                                    
                                    v3 pools hold their own tokens and pay LPs the fee tier. v4 pools live in a
                                    shared PoolManager and may route the fee through a hook, which is where the
                                    launchpad pools' zero LP yield comes from. (options: V3, V4)
  -g, --gas-sponsorship             Return EIP-712 typed data for the owner to sign instead of a transaction, so a sponsor can broadcast it.
  -h, --help                        help for risk-yield-decrease
      --owner string                The wallet that owns the Risk Yield Account. [required]
      --percent string              JSON value (one of: number | string)
      --position-id int             [required]
      --preview                     Return the plan and its simulation without building a transaction. Nothing is signed and nothing can be broadcast.
  -s, --slippage-pct string         JSON value (one of: number | string)
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
