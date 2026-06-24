## compass onramp quote

Get a fiat on-ramp quote (card -> USDC on Ethereum)

### Synopsis

Get an indicative quote for buying USDC on Ethereum with a card.

Returns the estimated USDC the user will receive for a given fiat amount,
the exchange rate, fee breakdown (`business_fees` is `0`), and the provider
min/max limits. The output is always **USDC on Ethereum** delivered to the
`destination_address` — this is a non-custodial on-ramp, so Compass never
holds the funds.

This quote is indicative and may expire (`expires_at`). To actually start a
payment, call `POST /v2/onramp/create`.

```
compass onramp quote [flags]
```

### Examples

```
  compass onramp quote --fiat-amount <value> --destination-address <value>
```

### Options

```
      --body string                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --destination-address string   The wallet that will receive the USDC on Ethereum. Funds are delivered here directly — Compass never holds them. [required]
      --fiat-amount 100              Amount of fiat to spend, as a decimal string (e.g. 100). [required]
      --fiat-currency USD            ISO-4217 fiat currency code. Defaults to USD.
  -h, --help                         help for quote
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
