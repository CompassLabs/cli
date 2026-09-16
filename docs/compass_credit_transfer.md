## compass credit transfer

Transfer tokens to/from Credit Account

### Synopsis

Transfer tokens between the owner's EOA and their Credit Account.

**DEPOSIT** (EOA → Credit Account):
- With `gas_sponsorship=true`: returns Permit2 EIP-712 typed data to sign. The gas sponsor
  calls `permitTransferFrom` to pull tokens (1 signature).
- With `gas_sponsorship=false`: returns an unsigned ERC-20 transfer transaction.

**WITHDRAW** (Credit Account → EOA):
- With `gas_sponsorship=true`: returns SafeTx EIP-712 typed data to sign. The gas sponsor
  broadcasts the `execTransaction` (1 signature).
- With `gas_sponsorship=false`: returns an unsigned `execTransaction`.

```
compass credit transfer [flags]
```

### Examples

```
  compass credit transfer --owner 0x4A83fec8c6A9A25Be28f3242a16dBaD0ab00f3a6 --chain base --token USDC --amount 100 --action DEPOSIT
```

### Options

```
      --action string          Whether you are depositing to or withdrawing from your credit account. (options: DEPOSIT, WITHDRAW) [required]
      --amount string          JSON value (one of: number | string)
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -c, --chain string           Blockchain network. (options: arbitrum, base, bsc, ethereum, hyperevm, tempo) [required]
  -g, --gas-sponsorship true   Optionally request gas sponsorship. If set to true, EIP-712 signature data will be returned that must be signed by the `owner` and submitted to the `/gas_sponsorship/prepare` endpoint.
  -h, --help                   help for transfer
      --owner string           The owner's wallet address (EOA). [required]
      --schema                 Print the exact JSON Schema of the request body and exit
  -s, --spender action         The address that will call Permit2's permitTransferFrom to execute the transfer. When action is 'DEPOSIT' and `gas_sponsorship` is `true`: - If provided, the signature will authorize this address (typically a gas sponsor) to pull tokens. - If not provided, defaults to the Credit Account (Safe) address, allowing the transfer to be included in a bundle transaction where the Safe pulls the tokens itself.
  -t, --token string           The token to transfer. [required]
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

* `compass credit transfer --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass credit transfer --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `compass credit transfer --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
