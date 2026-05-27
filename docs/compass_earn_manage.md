## compass earn manage

Manage earn position

### Synopsis

Deposit into or withdraw from a yield venue.

Use `DEPOSIT` to move tokens from the Earn Account into a vault, Aave market, or Pendle PT position. Use `WITHDRAW` to move tokens back from a venue into the Earn Account.

**Venue types:**
- `VAULT`: ERC-4626 vaults (see [/vaults](https://docs.compasslabs.ai/v2/api-reference/earn/list-vaults) for available options)
- `AAVE`: Aave lending pools (see [/aave_markets](https://docs.compasslabs.ai/v2/api-reference/earn/list-aave-markets) for available options)
- `PENDLE_PT`: Pendle Principal Tokens (see [/pendle_markets](https://docs.compasslabs.ai/v2/api-reference/earn/list-pendle-markets) for available options)

**Fees:** A fee can be configured and deducted from the transaction amount (not supported for PENDLE_PT):
- `FIXED`: A fixed token amount (e.g., 0.1 means 0.1 tokens)
- `PERCENTAGE`: A percentage of the amount (e.g., 1.5 means 1.5%)
- `PERFORMANCE`: A percentage of realized profit (withdrawals only)

**Gas sponsorship:** Set `gas_sponsorship=true` to receive EIP-712 typed data. Owner signs the typed data, then submit to [/gas_sponsorship/prepare](https://docs.compasslabs.ai/v2/api-reference/gas-sponsorship/prepare-gas-sponsored-transaction).

```
compass earn manage [flags]
```

### Examples

```
  compass earn manage --venue '{"type":"VAULT","vault_address":"0x7BfA7C4f149E7415b73bdeDfe609237e29CBF34A"}' --action DEPOSIT --amount 0.01 --owner 0x01E62835dd7F52173546A325294762143eE4a882 --chain base
```

### Options

```
      --action venue                                 Whether you are depositing into or withdrawing from the given Earn venue. (options: DEPOSIT, WITHDRAW) [required]
      --amount string                                JSON value (one of: number | string)
      --body string                                  Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string                                 The chain to use. (options: base, ethereum, arbitrum, hyperevm) [required]
  -f, --fee string                                   Optional fee configuration. If provided, a fee will be applied to the transaction amount and sent to the specified recipient address. The fee can be specified as a percentage of the transaction amount, as a fixed token amount, or as a percentage of realized profit (PERFORMANCE).
  -g, --gas-sponsorship true                         Optionally request gas sponsorship. If set to true, EIP-712 typed data will be returned that must be signed by the `owner` and submitted to the 'Prepare gas-sponsored transaction' endpoint (`/gas_sponsorship/prepare`).
  -h, --help                                         help for manage
      --owner string                                 The primary wallet address that owns and controls the Earn Account. [required]
  -v, --venue string                                 JSON value (variants: AAVE: { token: string }, PENDLE_PT: { market_address: string, token: string, max_slippage_percent: number }, VAULT: { vault_address: string })
      --venue.aave string                            AaveVenue variant as JSON
      --venue.aave.token string                      The reserve token for the Aave pool you are interacting with. [required]
      --venue.pendle-pt string                       PendlePTVenue variant as JSON
      --venue.pendle-pt.market-address string        The Pendle market address identifying which PT to trade. [required]
      --venue.pendle-pt.max-slippage-percent float   Maximum slippage allowed in percent (e.g., 1 means 1% slippage).
      --venue.pendle-pt.token string                 The token to exchange for PT (only used for DEPOSIT/buy). For WITHDRAW/sell, this field must not be provided - withdrawals always return the underlying asset to ensure accurate yield fee calculation.
      --venue.vault string                           VaultVenue variant as JSON
      --venue.vault.vault-address string             The vault address you are interacting with for this action. [required]
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
