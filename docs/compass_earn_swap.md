## compass earn swap

Swap tokens

### Synopsis

Swap one token for another inside an Earn Account.

Exchanges tokens the Earn Account already holds in a single atomic
transaction, so funds never leave it, and can be chained with other actions
through the [bundle endpoint](https://docs.compasslabs.ai/v2/api-reference/earn/execute-multiple-earn-actions)
(for example, swap ETH to USDC and deposit the USDC into a vault at once).
Returns an unsigned transaction to sign and the expected output amount.

```
compass earn swap [flags]
```

### Examples

```
  compass earn swap --token-in USDC --token-out USDT --amount-in 0.01 --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --chain base
```

### Options

```
  -a, --amount-in string       JSON value (one of: number | string)
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -c, --chain string           Target blockchain network where the swap will execute. (options: arbitrum, base, bsc, ethereum, hyperevm, tempo) [required]
  -g, --gas-sponsorship true   Optionally request gas sponsorship. If true, EIP-712 typed data will be returned that must be signed by the `owner` and submitted to the 'Prepare gas-sponsored transaction' endpoint (`/gas_sponsorship/prepare`). Firm-priced builds may be sponsored: the sponsor must broadcast before `quote_expires_at`.
  -h, --help                   help for swap
      --owner string           The owner's wallet address. [required]
  -p, --pricing slippage       Swap routing policy. 'auto': on chains with a market route the swap is priced by the market aggregator and bounded by slippage (the default behaviour); on HyperEVM, where only firm pricing is available, it is a firm zero-slippage quote that fills exactly or reverts. 'firm': always a firm quote; refused with a typed 409 when no firm quote covers the pair right now, never silently priced at market. 'market': always the market route; refused with 422 on HyperEVM. Read `swap_provider` on the response for the route that actually priced the build, and `quote_expires_at` for a firm quote's deadline. (options: auto, firm, market)
      --schema                 Print the exact JSON Schema of the request body and exit
  -s, --slippage string        JSON value (one of: number | string)
      --token-in string        Token to sell (input). Provide a token symbol from a limited set (e.g., 'USDC') or any token address. [required]
      --token-out string       Token to buy (output). Provide a token symbol from a limited set (e.g., 'USDT') or any token address. [required]
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

* [compass earn](compass_earn.md)	 - Operations for earn

### Machine interface

* `compass earn swap --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass earn swap --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `compass earn swap --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
