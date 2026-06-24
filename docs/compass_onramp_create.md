## compass onramp create

Create a fiat on-ramp payment + get the checkout link

### Synopsis

Create a fiat on-ramp payment and return the browser-checkout handoff.

A fiat card + identity-verification flow can only render in a browser, so
this endpoint returns a **`checkout_url`** (a Compass-hosted page) together
with `payment_id`, the one-time-wallet `deposit_address`, an initial
`status`, and a `user_instructions` string.

**How to drive this (CLI / MCP agent / UI):**

1. **Present `checkout_url` to the user and open it in a browser.** It
   renders the card payment + identity verification and delivers USDC on
   Ethereum to the destination wallet. (An MCP agent cannot open a browser
   itself — show the link and ask the user to open it.)
2. **Poll `GET /v2/onramp/status?payment_id=...`** until `status` becomes
   `delivered` (or `failed`). `pending` = awaiting payment, `processing` =
   payment received and bridging/swapping to USDC.
3. Once `delivered`, the USDC is in the user's own wallet — continue with
   the tokenized buy.

Funds are never held by Compass: the `deposit_address` is a one-time wallet
owned solely by the user's wallet.

```
compass onramp create [flags]
```

### Examples

```
  compass onramp create --fiat-amount <value> --destination-address <value>
```

### Options

```
      --body string                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --destination-address string   The wallet that will receive the USDC on Ethereum. Funds are delivered here directly — Compass never holds them. [required]
      --fiat-amount 100              Amount of fiat to spend, as a decimal string (e.g. 100). [required]
      --fiat-currency USD            ISO-4217 fiat currency code. Defaults to USD.
  -h, --help                         help for create
      --quote-id string              Optional Halliday quote handle to confirm against. Omit to confirm against a fresh quote derived from the amount and currency.
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
