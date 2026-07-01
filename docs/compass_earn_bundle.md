## compass earn bundle

Execute multiple earn actions

### Synopsis

Combine multiple actions into a single atomic transaction.

Bundle swaps and venue deposits/withdrawals into one transaction executed through the Earn Account. This saves gas compared to executing actions separately and ensures all actions succeed or fail together.

**Example:** Swap AUSD to USDC, then deposit USDC into a vault - all in one transaction.

**Fees:** Manage actions (deposits/withdrawals) support optional fee configuration, same as the standalone manage endpoint.

**Gas sponsorship:** Set `gas_sponsorship=true` to receive EIP-712 typed data. Owner signs the typed data, then submit to [/gas_sponsorship/prepare](https://docs.compasslabs.ai/v2/api-reference/gas-sponsorship/prepare-gas-sponsored-transaction).

```
compass earn bundle [flags]
```

### Examples

```
  compass earn bundle --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --chain ethereum --actions '[{"body":{"action_type":"V2_TRANSFER_FROM_EOA","token":"USDC","amount":"100","permit2_signature":"0x...","permit2_nonce":1706000000,"permit2_deadline":1706001800} },{"body":{"action_type":"V2_SWAP","token_in":"USDC","token_out":"AUSD","amount_in":"100","slippage":"0.5"} },{"body":{"action_type":"V2_MANAGE","venue":{"type":"VAULT","vault_address":"0x1B4cd53a1A8e5F50aB6320EF34E5fB4D3df7B6f6"},"action":"DEPOSIT","amount":"100"} }]'
```

### Options

```
  -a, --actions string    List of actions to bundle. Actions are executed in order. [required]
      --body string       Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string      The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc) [required]
  -g, --gas-sponsorship   If true, returns EIP-712 typed data for gas sponsorship. The owner must sign this data and submit to /gas_sponsorship/prepare.
  -h, --help              help for bundle
      --owner string      The owner's wallet address that controls the Earn Account. [required]
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
