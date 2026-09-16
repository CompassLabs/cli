## compass tokenized-assets buy

Buy an RWA yield token

### Synopsis

Buy an RWA yield token, deposit into an IXS managed vault, or buy a Centrifuge
deRWA token, with a stablecoin in one transaction.

Set `token_out` to a Midas symbol (`mTBILL`, `mBASIS`, `mBTC`), a Centrifuge
deRWA symbol (e.g. `deSPXA`), or an IXS **vault address** (its shares aren't a
registered symbol). The account spends a stablecoin it already holds (fund it
with a plain transfer first) and settles inside the product account — an
unsigned transaction the owner signs, or EIP-712 with `gas_sponsorship`. All
three settle instantly. Equities use the order flow (`/quote`, `/order`).

```
compass tokenized-assets buy [flags]
```

### Examples

```
  compass tokenized-assets buy --token-in <value> --token-out <value> --amount-in 4533.23 --owner <value> --chain ethereum
```

### Options

```
  -a, --amount-in string    JSON value (one of: number | string)
      --body string         Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -c, --chain string        The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc, robinhood, ethereum_sepolia) [required]
  -f, --fee string          Optional partner fee charged when selling (exiting). It is taken from the payout-token (USDC) proceeds and sent to your fee recipient inside the same execution.
  -g, --gas-sponsorship     When true, returns an EIP-712 payload for gas-sponsored execution instead of an unsigned transaction.
  -h, --help                help for buy
      --owner string        The owner's wallet address. [required]
      --schema              Print the exact JSON Schema of the request body and exit
  -s, --slippage string     JSON value (one of: number | string)
      --token-in string     Token to spend. For a buy this must be a stablecoin the Midas issuance vault accepts (USDC on every supported network; mBASIS also accepts USDT/DAI on Ethereum). For a sell it is the Midas RWA asset to redeem (e.g. 'mTBILL'). [required]
      --token-out string    Token to receive. For a buy this is the Midas RWA asset to mint (e.g. 'mTBILL'); for a sell it is the payout stablecoin (USDC). [required]
  -w, --wisdomtree string   WisdomTree Connect API credentials, required only when trading a WisdomTree money-market fund. Each organization authenticates with its own credentials, which are exchanged for a short-lived token to look up the settlement wallet for this trade. They are never stored and are masked in logs.
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

* `compass tokenized-assets buy --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass tokenized-assets buy --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `compass tokenized-assets buy --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
