## compass tokenized-assets tokenized-assets-transfer

Deposit to / withdraw from a Tokenized Assets Account

### Synopsis

Move tokens between the owner's wallet and their Tokenized Assets Account.

Use `DEPOSIT` to fund the account from the owner's wallet, or `WITHDRAW` to
send tokens from the account back to the owner. Equity orders settle in
USDC; RWA yield assets trade against USDC on Ethereum and Base.

With `gas_sponsorship=true` the response is EIP-712 typed data the owner
signs off-chain, then submits to `POST /v2/gas_sponsorship/prepare` so the
sponsor broadcasts and pays the gas:

- **DEPOSIT** returns a Permit2 `PermitTransferFrom`. The owner must first
  grant a one-time token->Permit2 allowance (gaslessly via
  `POST /v2/gas_sponsorship/approve_transfer`).
- **WITHDRAW** returns a Safe transaction the account executes.

With `gas_sponsorship=false` a DEPOSIT returns an unsigned ERC-20 transfer
the owner broadcasts directly, and a WITHDRAW returns an unsigned Safe
`execTransaction` the owner signs and broadcasts.

```
compass tokenized-assets tokenized-assets-transfer [flags]
```

### Examples

```
  compass tokenized-assets tokenized-assets-transfer --owner 0x9bDC45AA15FdFFc52E103EA05c260c494A5638f7 --token USDC --amount 100 --action DEPOSIT
```

### Options

```
      --action string          Whether you are depositing to or withdrawing from your Tokenized Assets Account. (options: DEPOSIT, WITHDRAW) [required]
      --amount string          JSON value (one of: number | string)
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string           The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo)
  -g, --gas-sponsorship true   Optionally request gas sponsorship. If set to true, EIP-712 signature data will be returned that must be signed by the `owner` and submitted to the `/gas_sponsorship/prepare` endpoint.
  -h, --help                   help for tokenized-assets-transfer
      --owner string           The owner's wallet address (EOA). [required]
  -s, --spender action         The address that will call Permit2's permitTransferFrom to execute the transfer. When action is 'DEPOSIT' and `gas_sponsorship` is `true`: - If provided, the signature will authorize this address (typically a gas sponsor) to pull tokens. - If not provided, defaults to the Tokenized Assets Account (Safe) address, allowing the transfer to be included in a bundle transaction where the Safe pulls the tokens itself.
  -t, --token string           The token to transfer. Tokenized-asset orders settle in USDC. [required]
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
