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
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string           The chain to use. (options: base, ethereum, arbitrum, hyperevm) [required]
  -g, --gas-sponsorship true   Optionally request gas sponsorship. If set to true, EIP-712 signature data will be returned that must be signed by the `owner` and submitted to the `/gas_sponsorship/prepare` endpoint.
  -h, --help                   help for transfer
      --owner string           The owner's wallet address (EOA). [required]
  -s, --spender action         The address that will call Permit2's permitTransferFrom to execute the transfer. When action is 'DEPOSIT' and `gas_sponsorship` is `true`: - If provided, the signature will authorize this address (typically a gas sponsor) to pull tokens. - If not provided, defaults to the Credit Account (Safe) address, allowing the transfer to be included in a bundle transaction where the Safe pulls the tokens itself.
  -t, --token string           The token to transfer. [required]
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

* [compass credit](compass_credit.md)	 - Operations for credit
