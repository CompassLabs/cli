## compass earn transfer

Transfer tokens to/from account

### Synopsis

Transfer tokens between an owner's wallet and their Earn Account.

Use `DEPOSIT` to move tokens from the owner's wallet into their Earn Account. Use `WITHDRAW` to move tokens from the Earn Account back to the owner's wallet.

**Gas-sponsored deposits** require two steps:

**Step 1 (once per token):** Set up Permit2 allowance
1. Call [/gas_sponsorship/approve_transfer](https://docs.compasslabs.ai/v2/api-reference/gas-sponsorship/approve-token-transfer) to get EIP-712 typed data
2. Owner signs the typed data
3. Submit signature + typed data to [/gas_sponsorship/prepare](https://docs.compasslabs.ai/v2/api-reference/gas-sponsorship/prepare-gas-sponsored-transaction)

**Step 2 (each transfer):** Execute the gas-sponsored transfer
1. Call this endpoint with `gas_sponsorship=true` to get EIP-712 typed data
2. Owner signs the typed data
3. Submit signature + typed data to [/gas_sponsorship/prepare](https://docs.compasslabs.ai/v2/api-reference/gas-sponsorship/prepare-gas-sponsored-transaction)

**Note:** This endpoint moves tokens to/from the Earn Account itself—not into yield venues. To deposit into a vault or Aave market, use the [manage endpoint](https://docs.compasslabs.ai/v2/api-reference/earn/manage-earn-position).

```
compass earn transfer [flags]
```

### Examples

```
  compass earn transfer --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --chain base --token USDC --amount 0.01 --action DEPOSIT
```

### Options

```
      --action string          Whether you are depositing to or withdrawing from your earn account. (options: DEPOSIT, WITHDRAW) [required]
      --amount string          JSON value (one of: number | string)
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string           The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo) [required]
  -g, --gas-sponsorship true   Optionally request gas sponsorship. If set to true, EIP-712 signature data will be returned that must be signed by the `owner` and submitted to the `/gas_sponsorship/prepare` endpoint.
  -h, --help                   help for transfer
      --owner string           The owner's wallet address. [required]
  -r, --recipient action       Optional recipient address for withdrawals. When action is 'WITHDRAW': - If provided, tokens will be sent to this address instead of the owner. - If not provided, defaults to the owner's address. 
  -s, --spender action         The address that will call Permit2's permitTransferFrom to execute the transfer. When action is 'DEPOSIT' and `gas_sponsorship` is `true`: - If provided, the signature will authorize this address (typically a gas sponsor) to pull tokens. - If not provided, defaults to the Earn Account (Safe) address, allowing the transfer to be included in a bundle transaction where the Safe pulls the tokens itself.
  -t, --token string           The token you would like to transfer. [required]
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
