## compass credit swap

Swap tokens

### Synopsis

Swap one token for another inside a Credit Account.

Exchanges tokens the Credit Account already holds in a single atomic
transaction, so idle or borrowed balances can be converted without moving
funds out first. Returns an unsigned transaction to sign, the expected
output amount, and which route priced the swap.

```
compass credit swap [flags]
```

### Examples

```
  compass credit swap --token-in USDC --token-out USDT --amount-in 0.01 --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --chain base
```

### Options

```
  -a, --amount-in string                  JSON value (one of: number | string)
      --body string                       Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -c, --chain string                      Target blockchain network where the swap will execute. (options: arbitrum, base, bsc, ethereum, hyperevm, tempo) [required]
  -g, --gas-sponsorship true              Optionally request gas sponsorship. If true, EIP-712 typed data will be returned that must be signed by the `owner` and submitted to the 'Prepare gas-sponsored transaction' endpoint (`/gas_sponsorship/prepare`). Gas-sponsored builds always execute at market rate.
  -h, --help                              help for swap
      --owner /v2/credit/create_account   The owner's wallet address. Their Credit Account must already exist (create it with /v2/credit/create_account) and hold `token_in` (deposit with `/v2/credit/transfer`). [required]
      --preview pricing                   If true, build a display ESTIMATE: no firm RFQ quote is ever requested (quote_expires_at stays null). NOTE that this guarantees only that no firm quote was spent — it does not guarantee an absent payload: under 'auto' when the firm provider does not cover or cannot currently price the pair, and always under pricing=market, the call falls through to the aggregator and returns a signable market build (an unsigned transaction, or EIP-712 typed data when gas_sponsorship=true), without requiring the account to hold token_in yet. How the estimate is priced follows pricing: on a firm-covered pair whose size the firm provider's live price levels can serve, 'auto' and 'firm' price it from those levels (indicative, transaction stays null — re-call with preview=false for the signable build); otherwise it comes from the market build above (pricing='firm' instead refuses with a typed error). Set it on every call made while a user is exploring parameters, and leave it false only for the build they actually intend to sign — firm quotes are single-use maker commitments, and requesting them for displays that are never executed degrades the pricing this API is offered.
      --pricing slippage                  Swap routing policy. 'auto': a firm zero-slippage quote when a firm venue covers the pair, transparent fallback to the market aggregator otherwise. 'firm': never price on the market route; an uncovered pair fails with a typed error instead of silently substituting market pricing — previews included: a preview the firm provider's live price levels cannot price returns the same typed error rather than market numbers. 'market': never route through the firm venue; the swap is priced by the aggregator and bounded by slippage (which firm fills ignore). 'firm' is incompatible with gas_sponsorship (sponsored swaps force market routing). (options: auto, firm, market)
      --schema                            Print the exact JSON Schema of the request body and exit
  -s, --slippage string                   JSON value (one of: number | string)
      --token-in string                   Token to sell (input). Provide a token symbol from a limited set (e.g., 'USDC') or any token address. [required]
      --token-out string                  Token to buy (output). Provide a token symbol from a limited set (e.g., 'USDT') or any token address. [required]
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key-auth string    Your Compass API Key. Get your key [here](https://www.compasslabs.ai/dashboard).
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
      --interactive            Prompt for missing inputs and open guided configure/auth forms (forms fall back to line prompts on stdin off-TTY) (default true)
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [compass credit](compass_credit.md)	 - Operations for credit

### Machine interface

* `compass credit swap --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass credit swap --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `compass credit swap --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
