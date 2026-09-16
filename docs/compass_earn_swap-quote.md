## compass earn swap-quote

Quote a swap

### Synopsis

Estimate the output of a swap without building a transaction.

Returns the expected amount of `token_out` received for selling `amount_in`
of `token_in`, routed via 1inch. This is read-only: it does not build a
transaction, require an account, or check balances.

Use it to gauge exit liquidity and price impact for a token before entering
a position — for example, to warn when a market's underlying asset cannot be
swapped back to a stablecoin without large slippage.

```
compass earn swap-quote [flags]
```

### Examples

```
  compass earn swap-quote --chain base --token-in WETH --amount-in 1
```

### Options

```
  -a, --amount-in string      JSON value (one of: number | string)
  -c, --chain string          Target blockchain network. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc, robinhood, ethereum_sepolia) [required]
  -h, --help                  help for swap-quote
      --slippage string       JSON value (one of: number | string)
      --sy-address token_in   Optional Pendle SY (Standardized Yield) address. When provided, token_in is overridden with the token the PT actually redeems into on withdrawal (the SY asset if it is a valid token-out, else the SY yield token) — use this to gauge a Pendle position's real exit liquidity rather than the reported underlying.
      --token-in string       Token to sell (input). A token symbol (e.g. 'WETH') or any token address. [required]
      --token-out string      Token to buy (output). A token symbol (e.g. 'USDC') or any token address.
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

* `compass earn swap-quote --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass earn swap-quote --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
