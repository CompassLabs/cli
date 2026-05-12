## compass tokenized-assets tokenized-assets-order-submit

Submit a signed order

### Synopsis

Submit a signed order to the resolver network.

The body echoes the `order` fields from `/order` (`signed_order`,
`extension`, `quote_id`, optionally `order_hash`) plus the owner's
signature over `order.safe_message_eip712`. The maker on the order
struct is the Tokenized Assets Account, not the owner's wallet —
pass `signed_order` back unchanged.

Returns the order hash and a server-side ISO 8601 timestamp.
Subsequent calls to `GET /order/{order_hash}` track the lifecycle
(`pending` → `filled` / `expired` / `cancelled`).

```
compass tokenized-assets tokenized-assets-order-submit [flags]
```

### Examples

```
  compass tokenized-assets tokenized-assets-order-submit --signed-order '{"key":"<value>"}' --signature <value> --extension m2a --quote-id <id>
```

### Options

```
      --body string                           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --extension /order                      Opaque hex blob from the /order response — pass back unchanged. [required]
  -h, --help                                  help for tokenized-assets-order-submit
      --order-hash order.order_hash           order.order_hash from the `/order` response. Optional but recommended: the upstream relayer occasionally returns a 2xx with an empty body, and supplying the hash lets the API still return a usable handle for status and cancel lookups instead of failing.
      --quote-id order.quote_id               order.quote_id from the `/order` response — pass back unchanged. [required]
      --signature order.safe_message_eip712   Owner's EIP-712 signature over order.safe_message_eip712 from the `/order` response. The signature is validated against the Tokenized Assets Account at fill time, so it must be a signature over the typed-data hash, not the raw order hash. [required]
      --signed-order /order                   The order struct returned by /order (`order.order_message`). `maker` is the Tokenized Assets Account, not the owner's wallet — pass this dict back to the API verbatim. [required]
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
