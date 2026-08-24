## compass credit rebalance

Rebalance the leveraged credit book

### Synopsis

Move one or more leveraged positions to the state you want, in ONE atomic
transaction.

List the positions to change and the end state you want for each. The API
works out whether that means opening, growing, shrinking, delevering or
closing, and does them all together.

Money freed by shrinking or closing one position pays for growing or opening
another, so shifting funds between positions needs no new deposit. Any
shortfall is taken from the Credit Account's idle balance, and the call
returns 422 if that does not cover it either.

See the [Leveraged Looping guide](https://docs.compasslabs.ai/v2/Products/Looping)
for protocol and chain coverage, capital routing and the full error list.

```
compass credit rebalance [flags]
```

### Examples

```
  compass credit rebalance --owner 0x4D3c07d1db7E4A9E44285fae9810d6549655bc74 --chain ethereum --targets '[{"protocol":"AAVE","collateral_token":"WETH","borrow_token":"USDC","target_multiplier":2}]'
```

### Options

```
      --body string                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string                  Blockchain network. Rebalance is not available on HyperEVM yet; use loop/unloop there. (options: arbitrum, base, bsc, ethereum, tempo) [required]
  -g, --gas-sponsorship               If true, returns EIP-712 typed data for gas-sponsored execution instead of an unsigned transaction.
  -h, --help                          help for rebalance
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
