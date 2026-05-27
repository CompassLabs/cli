## compass earn swap

Swap tokens within Earn Account

### Synopsis

Swap tokens within an Earn Account.

Use this endpoint to exchange one token for another without transferring funds out of the Earn Account.

The swap executes atomically within the Earn Account and can be combined with other actions using the [bundle endpoint](https://docs.compasslabs.ai/v2/api-reference/earn/execute-multiple-earn-actions). For example, swap ETH to USDC, then deposit USDC into a vault—all in one transaction.

Returns either an unsigned transaction (when `gas_sponsorship=false`) or EIP-712 typed data for off-chain signing (when `gas_sponsorship=true`). For gas-sponsored swaps, submit the signed typed data to `/gas_sponsorship/prepare`.

```
compass earn swap [flags]
```

### Examples

```
  compass earn swap --token-in USDC --token-out USDT --amount-in 0.01 --owner 0x01E62835dd7F52173546A325294762143eE4a882 --chain base
```

### Options

```
  -a, --amount-in string       JSON value (one of: number | string)
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string           The chain to use. (options: base, ethereum, arbitrum, hyperevm) [required]
  -g, --gas-sponsorship true   Optionally request gas sponsorship. If true, EIP-712 typed data will be returned that must be signed by the `owner` and submitted to the 'Prepare gas-sponsored transaction' endpoint (`/gas_sponsorship/prepare`).
  -h, --help                   help for swap
      --owner string           The owner's wallet address. [required]
  -s, --slippage string        JSON value (one of: number | string)
      --token-in string        Token to sell (input). Provide a token symbol from a limited set (e.g., 'USDC') or any token address. [required]
      --token-out string       Token to buy (output). Provide a token symbol from a limited set (e.g., 'USDT') or any token address. [required]
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

* [compass earn](compass_earn.md)	 - Operations for earn
