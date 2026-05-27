## compass tokenized-equities quote

Preview a buy/sell quote

### Synopsis

Preview the input/output amounts, fees, and slippage tolerance for an order.

Read-only relative to Fusion: hits ``/quote/receive`` only and does not
consume a ``quote_id`` or commit an order. Pair with `POST /order`:
surface this preview to the user, and on confirm pass the body plus
``recommended_slippage_bps`` to `/order`.

The response carries:

- **`quote`** — input/output token amounts, fees, and an
  ``est_fill_seconds`` upper bound.
- **`recommended_slippage_bps`** — system-derived slippage tolerance
  that clears Fusion's current auction floor; pass back as
  ``slippage_bps`` on `/order` so the build call validates against the
  same floor the user was shown.
- **`auction_range_bps`** — worst-case bps gap between the auction
  end amount and the reference quote amount. Use to surface a
  thin-liquidity warning to the user.

```
compass tokenized-equities quote [flags]
```

### Examples

```
  compass tokenized-equities quote --from-token <value> --to-token <value> --amount 936.80 --owner <value>
```

### Options

```
  -a, --amount from_token     Human-readable amount of from_token to swap (decimal string). Decimals are applied server-side. [required]
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -f, --from-token TSLAon     Token the sender is paying. Either an on-chain symbol (e.g. TSLAon), the literal `USDC`, or a 0x-prefixed token address. [required]
  -h, --help                  help for quote
      --owner string          Wallet that owns the Tokenized Equities Account. Used to verify the account is deployed before quoting; the account address is derived deterministically from this owner. [required]
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

* [compass tokenized-equities](compass_tokenized-equities.md)	 - Operations for tokenized-equities
