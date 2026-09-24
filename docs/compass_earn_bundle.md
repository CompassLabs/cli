## compass earn bundle

Execute multiple earn actions

### Synopsis

Combine multiple actions into a single atomic transaction.

Bundle swaps and venue deposits/withdrawals into one transaction executed through the Earn Account. This saves gas compared to executing actions separately and ensures all actions succeed or fail together.

**Example:** Swap AUSD to USDC, then deposit USDC into a vault - all in one transaction.

**Fees:** Manage actions (deposits/withdrawals) support optional fee configuration, same as the standalone manage endpoint.

**Gas sponsorship:** Set `gas_sponsorship=true` to receive EIP-712 typed data. Owner signs the typed data, then submit to [/gas_sponsorship/prepare](https://docs.compasslabs.ai/v2/api-reference/gas-sponsorship/prepare-gas-sponsored-transaction).

```
compass earn bundle [flags]
```

### Examples

```
  compass earn bundle --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --chain ethereum --actions '[{"body":{"action_type":"V2_TRANSFER_FROM_EOA","token":"USDC","amount":"100","permit2_signature":"0x...","permit2_nonce":1706000000,"permit2_deadline":1706001800}},{"body":{"action_type":"V2_SWAP","token_in":"USDC","token_out":"AUSD","amount_in":"100","slippage":"0.5"}},{"body":{"action_type":"V2_MANAGE","venue":{"type":"VAULT","vault_address":"0x1B4cd53a1A8e5F50aB6320EF34E5fB4D3df7B6f6"},"action":"DEPOSIT","amount":"100"}}]'
```

### Options

```
  -a, --actions string    List of actions to bundle. Actions are executed in order. [required]
      --body string       Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -c, --chain string      Target blockchain network where the bundled actions will execute. (options: arbitrum, base, bsc, ethereum, hyperevm, tempo) [required]
  -g, --gas-sponsorship   If true, returns EIP-712 typed data for gas sponsorship. The owner must sign this data and submit to /gas_sponsorship/prepare.
  -h, --help              help for bundle
      --owner string      The owner's wallet address that controls the Earn Account. [required]
      --schema            Print the exact JSON Schema of the request body and exit
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

* `compass earn bundle --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass earn bundle --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `compass earn bundle --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
