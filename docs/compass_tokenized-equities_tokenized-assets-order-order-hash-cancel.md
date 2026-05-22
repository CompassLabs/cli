## compass tokenized-equities tokenized-assets-order-order-hash-cancel

Cancel an unfilled order

### Synopsis

Build the EIP-712 payload to cancel an unfilled order on-chain.

Returns ``cancel_safe_tx_eip712``, an EIP-712 payload that authorizes
the on-chain cancellation. Sign with the Tokenized Equities Account's
owner via ``wallet.signTypedData(...)`` and relay via
``POST /v2/gas_sponsorship/prepare`` so the sponsor broadcasts the
cancellation on the product account. The owner can also broadcast
the resulting transaction directly without using gas sponsorship.

Cancellation works on `pending` and `expired` orders only. Only the
Tokenized Equities Account that placed the order can cancel it.

```
compass tokenized-equities tokenized-assets-order-order-hash-cancel [flags]
```

### Examples

```
  compass tokenized-equities tokenized-assets-order-order-hash-cancel --order-hash <value> --owner <value>
```

### Options

```
      --body string         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                help for tokenized-assets-order-order-hash-cancel
      --order-hash string   [required]
      --owner string        Wallet that owns the Tokenized Equities Account. The account address derived from this owner must match the order's on-chain maker; the API rejects otherwise (only the order's maker can cancel it). [required]
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

* [compass tokenized-equities](compass_tokenized-equities.md)	 - Operations for tokenized-equities
