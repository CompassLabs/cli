## compass tokenized-assets tokenized-assets-order

Build a buy/sell order

### Synopsis

Build a buy or sell order whose maker is the Tokenized Assets Account.

Returns up to three pieces in a single round-trip:

- **`quote`** — preview of the input/output amounts and fees.
- **`approval_safe_tx_eip712`** — only present when the account's
  allowance to the settlement contract is below `amount`. The owner
  signs this EIP-712 payload, then it is broadcast via
  `POST /v2/gas_sponsorship/prepare` (or the owner can broadcast
  directly) to set the on-chain allowance. Wait for that transaction
  to confirm before signing the order.
- **`order`** — the order metadata (`order_hash`, `extension`,
  `quote_id`, `order_message`) plus `safe_message_eip712`, an EIP-712
  payload the owner signs off-chain to authorize the order. The
  signature is submitted to `/order/submit` and is **never** broadcast
  on-chain.

The owner never broadcasts the order itself — only the (one-time)
approval transaction ever hits the chain.

```
compass tokenized-assets tokenized-assets-order [flags]
```

### Examples

```
  compass tokenized-assets tokenized-assets-order --from-token <value> --to-token <value> --amount 853.30 --owner <value>
```

### Options

```
  -a, --amount from_token     Human-readable amount of from_token to swap (decimal string). Decimals are applied server-side. [required]
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -f, --from-token TSLAon     Token the sender is paying. Either an on-chain symbol (e.g. TSLAon), the literal `USDC`, or a 0x-prefixed token address. [required]
  -h, --help                  help for tokenized-assets-order
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
