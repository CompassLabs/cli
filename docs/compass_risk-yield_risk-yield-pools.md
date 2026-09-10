## compass risk-yield risk-yield-pools

List pools

### Synopsis

Pools you could provide liquidity to, ranked by what you would keep.

**Read `lp_earns_fees` first.** The launchpad that dominates this chain
graduates tokens into pools whose hook takes the swap fee, so they trade
enormously and pay an external liquidity provider nothing. Those pools are
listed — they carry the memecoin's price and volume — with a fee APR of
exactly zero and a note saying why. Pass `lp_earns_fees=true` to hide them.

The default ordering is `il_adjusted_apr_7d`: the fee APR less the expected
impermanent-loss drag from the pair's own volatility. A pool paying 200% on
an asset that moves 400% is not a better position than one paying 20% on an
asset that moves 30%, and sorting on the raw APR would say it was.

Pools too thin for a rate to mean anything — under 20 swaps a day, or under
$1,000 of liquidity — are served with `ranked: false` and always sort last.

```
compass risk-yield risk-yield-pools [flags]
```

### Examples

```
  compass risk-yield risk-yield-pools --chain robinhood
```

### Options

```
  -c, --chain string                options: robinhood [required]
      --dex-version string          Which Uniswap deployment a pool belongs to.
                                    
                                    v3 pools hold their own tokens and pay LPs the fee tier. v4 pools live in a
                                    shared PoolManager and may route the fee through a hook, which is where the
                                    launchpad pools' zero LP yield comes from. (options: V3, V4)
      --direction string            options: asc, desc
  -h, --help                        help for risk-yield-pools
      --launchpad string            options: PONS_V1, PONS_V2, POOLS_TRADE, BAGS
      --limit int                   The number of items to return.
      --lp-earns-fees string        Leave unset to see everything. Set true to hide the pools whose hook keeps the fee — most of this chain's volume, and none of its yield.
      --lp-open string              Whether liquidity may be added at all.
      --min-tvl-usd string          JSON value (one of: number | string)
      --min-volume-24h-usd string   JSON value (one of: number | string)
      --offset int                  The offset of the first item to return.
      --order-by string             options: il_adjusted_apr_7d, fee_apr_24h, fee_apr_7d, tvl_usd, volume_24h_usd, volume_7d_usd, realized_vol_24h, realized_vol_7d, created_at
  -p, --pair-class string           What kind of pair a pool is, which is what the product is sorted by. (options: MEME_STOCK, STOCK_STABLE, STOCK_ETH, STOCK_STOCK, MEME_ETH, MEME_STABLE, ETH_STABLE, OTHER)
  -s, --stock-ticker string         string value
  -t, --token string                Pools holding this token, on either side.
  -w, --watched string              Only pools the indexer follows closely, which are the only ones with trailing metrics.
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
