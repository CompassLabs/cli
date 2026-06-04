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
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string           The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo) [required]
  -g, --gas-sponsorship true   Optionally request gas sponsorship. If set to true, EIP-712 signature data will be returned that must be signed by the `owner` and submitted to the `/gas_sponsorship/prepare` endpoint.
  -h, --help                   help for approve-transfer
      --owner string           The wallet address that owns the Earn Account. [required]
  -t, --token string           The token you would like to transfer with gas sponsorship. [required]
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

* [compass gas-sponsorship](compass_gas-sponsorship.md)	 - Operations for gas-sponsorship
