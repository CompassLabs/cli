## compass tokenized-assets buy

Buy a swap-traded tokenized asset

### Synopsis

Buy an RWA yield asset (e.g. `mTBILL`) inside the product account.

Swaps `token_in` (already held by the Tokenized Assets Account — fund it
with a plain transfer first) into `token_out` via the 1inch Aggregation
Router, executed by the account. Returns an unsigned transaction for the
owner to sign, or an EIP-712 payload when `gas_sponsorship` is true.

`token_out` must be a swap-traded tokenized asset; equities trade via the
order endpoints (`/quote`, `/order`, `/order/submit`).

```
compass tokenized-assets buy [flags]
```

### Examples

```
  compass tokenized-assets buy --token-in <value> --token-out <value> --amount-in 4533.23 --owner <value> --chain ethereum
```

### Options

```
  -a, --amount-in string   JSON value (one of: number | string)
      --body string        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string       The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo) [required]
  -g, --gas-sponsorship    When true, returns an EIP-712 payload for gas-sponsored execution instead of an unsigned transaction.
  -h, --help               help for buy
      --owner string       The owner's wallet address. [required]
  -s, --slippage string    JSON value (one of: number | string)
      --token-in string    Token to spend. For a buy this is any supported token (e.g. 'USDC'); for a sell it must be a swap-traded tokenized asset (e.g. 'mTBILL'). [required]
      --token-out string   Token to receive. For a buy this must be a swap-traded tokenized asset (e.g. 'mTBILL'); for a sell it is any supported token. [required]
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

* [compass tokenized-assets](compass_tokenized-assets.md)	 - Operations for tokenized-assets
