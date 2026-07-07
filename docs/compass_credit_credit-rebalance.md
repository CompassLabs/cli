## compass credit credit-rebalance

Rebalance the leveraged credit book

### Synopsis

Rebalance the leveraged credit book in ONE atomic transaction.

List only the positions to change — anything not named is left untouched;
remove a position with close=true. Each target states an end state:
target_equity_usd (net USD committed) × target_multiplier.

Releasing targets run first (each unwind/delever frees tokens into the Credit
Account), the freed tokens are then routed by swaps at a GUARANTEED minimum
output (enforced on-chain), and consuming targets run last — so moving a
levered position between markets, token pairs, or protocols (Aave ↔ Morpho) is
simply a close plus an open in the same transaction.

Net book growth is funded from the Credit Account's existing idle balance —
fund it first via /v2/credit/transfer; a net release stays in the Credit
Account as idle balance. Any swap surplus above the guaranteed floors also
stays in the Credit Account (preview.estimated_max_dust) — recoverable, never
lost.

A book already at its target returns transaction: null with the preview — the
call is idempotent and safe to drive from a converge-to-target loop.

A rebalance too large for one transaction is rejected with a 422 — split it
into two calls. Every deleveraging step keeps the health factor ≥ 1.02 and
every leveraging step respects the protocol's borrow limits; Aave targets share
one account-level health factor, which the preview reports.

For protocol=MORPHO pass a market_id from /v2/credit/morpho_markets; inspect the
current book via /v2/credit/looped_positions.

```
compass credit credit-rebalance [flags]
```

### Examples

```
  compass credit credit-rebalance --owner 0x4D3c07d1db7E4A9E44285fae9810d6549655bc74 --chain ethereum --targets '[{"protocol":"AAVE","collateral_token":"WETH","borrow_token":"USDC","target_multiplier":2}]'
```

### Options

```
      --body string                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string                  The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc) [required]
  -g, --gas-sponsorship               If true, returns EIP-712 typed data for gas-sponsored execution instead of an unsigned transaction.
  -h, --help                          help for credit-rebalance
  -m, --max-slippage-percent string   JSON value (one of: number | string)
      --owner string                  The address that owns the Credit Account. [required]
  -t, --targets string                The positions to change. Scoped: only positions named here are touched; any position not listed is left untouched. [required]
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

* [compass credit](compass_credit.md)	 - Operations for credit
