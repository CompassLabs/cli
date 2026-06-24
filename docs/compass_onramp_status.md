## compass onramp status

Get the status of a fiat on-ramp payment

### Synopsis

Get the current status of a fiat on-ramp payment.

Poll this after `create` until `status` is `delivered`. The `status` field
is one of:

- `pending` — awaiting the card payment in the hosted checkout.
- `processing` — payment received; bridging/swapping the proceeds to USDC.
- `delivered` — USDC has been delivered to the destination wallet
  (`delivery_tx_hash` and `output_amount_delivered` are populated).
- `failed` — the payment failed, expired, or was refunded.

State is re-read live from the on-ramp provider on every call — Compass
stores nothing of its own.

```
compass onramp status [flags]
```

### Examples

```
  compass onramp status --payment-id <id>
```

### Options

```
  -h, --help                    help for status
  -p, --payment-id payment_id   The payment_id returned by `POST /v2/onramp/create`. [required]
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

* [compass onramp](compass_onramp.md)	 - Operations for onramp
