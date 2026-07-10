## compass tokenized-assets order

Build an order

### Synopsis

Build a tokenized-**equity** (Ondo) buy/sell order; the product account is the
maker.

Returns everything needed to place the order in one round-trip: a price quote,
a one-time token approval to sign (only until the settlement contract is
approved), and the order payload the owner signs off-chain. Only the approval
ever touches the chain — the signed order is relayed to market makers, never
broadcast. Equities only; RWA yield tokens use `/transact/buy` and
`/transact/sell`.

```
compass tokenized-assets order [flags]
```

### Examples

```
  compass tokenized-assets order --from-token <value> --to-token <value> --amount 853.30 --owner <value>
```

### Options

```
  -a, --amount from_token     Human-readable amount of from_token to swap (decimal string). Decimals are applied server-side. [required]
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -f, --from-token TSLAon     Token the sender is paying. Either an on-chain symbol (e.g. TSLAon), the literal `USDC`, or a 0x-prefixed token address. [required]
  -h, --help                  help for order
      --owner string          Wallet that owns the Tokenized Assets Account. The product account address is derived deterministically from this owner. The owner signs the EIP-712 payloads returned by this endpoint (the optional approval and the order itself). [required]
  -s, --slippage-bps int      Max acceptable slippage in basis points (1 bp = 0.01%). Range 1-5000 (0.01%-50%); defaults to 50 (0.5%). The upper bound is intentionally wide so callers can clear the wide auction floors quoted for thinly-traded tokenized stocks.
  -t, --to-token from_token   Token the sender is receiving. Same accepted forms as from_token. [required]
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

* [compass tokenized-assets](compass_tokenized-assets.md)	 - Operations for tokenized-assets
