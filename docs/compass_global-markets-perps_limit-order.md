## compass global-markets-perps limit-order

Place limit order

### Synopsis

Prepare a limit order on a global markets perps market.

Returns EIP-712 typed data for the user to sign. After signing, submit
the signature via the /execute endpoint. Supports GTC and ALO time-in-force.

```
compass global-markets-perps limit-order [flags]
```

### Examples

```
  compass global-markets-perps limit-order --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --asset AAPL --side buy --size 10.0 --price 190.00
```

### Options

```
  -a, --asset string             Asset ticker symbol (e.g. AAPL, GOLD, EUR) [required]
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --builder string           Optional builder fee. When provided, the order action carries the builder address and fee — the end-user must have already approved this builder via /approve_builder_fee up to at least this rate. Omit (or pass null) to place the order with no builder fee.
  -c, --client-order-id string   Optional client order ID for idempotency (uint128)
  -h, --help                     help for limit-order
      --owner string             Owner of the global markets perps product account [required]
  -p, --price string             Limit price (human-readable) [required]
  -r, --reduce-only              If true, order can only reduce an existing position
      --side string              Order side: 'buy' or 'sell' [required]
      --size string              Number of contracts (human-readable) [required]
  -t, --time-in-force string     Time-in-force: 'gtc' (good til cancel) or 'alo' (add liquidity only)
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

* [compass global-markets-perps](compass_global-markets-perps.md)	 - Operations for global-markets-perps
