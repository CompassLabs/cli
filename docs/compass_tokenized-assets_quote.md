## compass tokenized-assets quote

Quote an order

### Synopsis

Preview a buy/sell quote for a tokenized **equity** (Ondo, e.g. `TSLAon`).

Read-only price preview: returns the expected input/output amounts and a
system-recommended slippage tolerance to carry into `/order`. Equities only —
RWA yield tokens (Midas) have no quote step and are rejected with `422 Wrong
trade flow`; trade them via `/transact/buy` and `/transact/sell`.

```
compass tokenized-assets quote [flags]
```

### Examples

```
  compass tokenized-assets quote --from-token <value> --to-token <value> --amount 226.42 --owner <value>
```

### Options

```
  -a, --amount from_token     Human-readable amount of from_token to swap (decimal string). Decimals are applied server-side. [required]
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -f, --from-token TSLAon     Token the sender is paying. Either an on-chain symbol (e.g. TSLAon), the literal `USDC`, or a 0x-prefixed token address. [required]
  -h, --help                  help for quote
      --owner string          Wallet that owns the Tokenized Assets Account. Used to verify the account is deployed before quoting; the account address is derived deterministically from this owner. [required]
      --schema                Print the exact JSON Schema of the request body and exit
  -t, --to-token from_token   Token the sender is receiving. Same accepted forms as from_token. [required]
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

* [compass tokenized-assets](compass_tokenized-assets.md)	 - Operations for tokenized-assets

### Machine interface

* `compass tokenized-assets quote --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass tokenized-assets quote --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `compass tokenized-assets quote --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
