## compass gas-sponsorship approve-transfer

Approve token transfer

### Synopsis

Set up a one-time Permit2 allowance for gas-sponsored token transfers.

Required when using [/earn/transfer](https://docs.compasslabs.ai/v2/api-reference/earn/transfer-tokens-tofrom-account) or [/credit/transfer](https://docs.compasslabs.ai/v2/api-reference/credit/transfer-tokens-tofrom-account) with `gas_sponsorship=true`. This allowance only needs to be set up once per token.

**With gas sponsorship (`gas_sponsorship=true`):**
- Returns EIP-712 typed data for the owner to sign off-chain
- Submit signature + typed data to [/prepare](https://docs.compasslabs.ai/v2/api-reference/gas-sponsorship/prepare-gas-sponsored-transaction)
- Only works for tokens that support EIP-2612 permit (e.g., USDC)

Some tokens, like USDT and WETH, do not support EIP-2612 permit. For these tokens, set `gas_sponsorship=false` to receive an unsigned transaction that the owner must sign, submit, and pay gas for.

```
compass gas-sponsorship approve-transfer [flags]
```

### Examples

```
  compass gas-sponsorship approve-transfer --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --chain base --token USDT
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -c, --chain string           The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc, robinhood, ethereum_sepolia) [required]
  -g, --gas-sponsorship true   Optionally request gas sponsorship. If set to true, EIP-712 signature data will be returned that must be signed by the `owner` and submitted to the `/gas_sponsorship/prepare` endpoint.
  -h, --help                   help for approve-transfer
      --owner string           The wallet address that owns the Earn Account. [required]
      --schema                 Print the exact JSON Schema of the request body and exit
  -t, --token string           The token you would like to transfer with gas sponsorship. [required]
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

* [compass gas-sponsorship](compass_gas-sponsorship.md)	 - Operations for gas-sponsorship

### Machine interface

* `compass gas-sponsorship approve-transfer --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass gas-sponsorship approve-transfer --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `compass gas-sponsorship approve-transfer --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
