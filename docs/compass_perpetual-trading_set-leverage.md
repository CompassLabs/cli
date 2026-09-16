## compass perpetual-trading set-leverage

Set leverage (defaults to market maximum)

### Synopsis

Check leverage and prepare an updateLeverage action to the requested value.

If `leverage` is omitted, targets the asset's maximum leverage for that
market. If the asset is already at the requested leverage, returns
leverage_ok=true with null typed_data — no signing needed. Otherwise,
returns EIP-712 typed data for the user to sign. After signing, submit the
signature via the /execute endpoint.

```
compass perpetual-trading set-leverage [flags]
```

### Examples

```
  compass perpetual-trading set-leverage --owner <value> --asset <value>
```

### Options

```
  -a, --asset string      Asset ticker (e.g. 'AAPL', 'CL') [required]
      --body string       Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -h, --help              help for set-leverage
  -l, --leverage string   Target leverage as a whole-number multiplier. Omit (or pass null) to use the asset's maximum leverage. Must be between 1 and the asset's max leverage, which is available from the /opportunities endpoint.
      --owner string      User's EOA address [required]
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

* [compass perpetual-trading](compass_perpetual-trading.md)	 - Operations for perpetual-trading

### Machine interface

* `compass perpetual-trading set-leverage --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass perpetual-trading set-leverage --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `compass perpetual-trading set-leverage --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
