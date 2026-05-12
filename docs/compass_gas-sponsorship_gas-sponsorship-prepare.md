## compass gas-sponsorship gas-sponsorship-prepare

Prepare gas-sponsored transaction

### Synopsis

Prepare a gas-sponsored transaction from signed EIP-712 typed data.

Submit the `owner`'s off-chain signature along with the EIP-712 typed data that was signed. Returns an unsigned transaction for the `sender` to sign and broadcast.

**How gas sponsorship works:**
1. Call an endpoint with `gas_sponsorship=true` (e.g., [/earn/transfer](https://docs.compasslabs.ai/v2/api-reference/earn/transfer-tokens-tofrom-account), [/earn/manage](https://docs.compasslabs.ai/v2/api-reference/earn/manage-earn-position), [/credit/transfer](https://docs.compasslabs.ai/v2/api-reference/credit/transfer-tokens-tofrom-account)) to get EIP-712 typed data
2. Owner signs the typed data off-chain
3. Submit signature + typed data to this endpoint
4. Sender signs and broadcasts the returned transaction, paying gas on behalf of the owner

**Note:** For gas-sponsored deposits via [/earn/transfer](https://docs.compasslabs.ai/v2/api-reference/earn/transfer-tokens-tofrom-account) or [/credit/transfer](https://docs.compasslabs.ai/v2/api-reference/credit/transfer-tokens-tofrom-account), the owner must first set up a Permit2 allowance using [/approve_transfer](https://docs.compasslabs.ai/v2/api-reference/gas-sponsorship/approve-token-transfer) (once per token).

```
compass gas-sponsorship gas-sponsorship-prepare [flags]
```

### Examples

```
  compass gas-sponsorship gas-sponsorship-prepare --owner 0xCE1A77F0abff993d6d3D04d44b70831c6924fb40 --chain arbitrum --eip-712 '{"domain":{"name":"USD Coin","version":"2","chainId":42161,"verifyingContract":"0xaf88d065e77c8cC2239327C5EDb3A432268e5831"},"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Permit":[{"name":"owner","type":"address"},{"name":"spender","type":"address"},{"name":"value","type":"uint256"},{"name":"nonce","type":"uint256"},{"name":"deadline","type":"uint256"}]},"primaryType":"Permit","message":{"owner":"0xCE1A77F0abff993d6d3D04d44b70831c6924fb40","spender":"0x000000000022D473030F116dDEE9F6B43aC78BA3","value":"115792089237316195423570985008687907853269984665640564039457584007913129639935","nonce":"0","deadline":"1762269774"} }' --signature 0x160d2709ae195f591daa33ad6ab1fb18b8762a39d8c4466c4cbe95cf6881fc3d54d469710ef0e7fd64ecff47c1ba5741d7254903bfaebdacea5aa8289f81ba9a1c --sender 0x02122Ac49b0Be2e0eAD957F2D080805A0127Aa9d
```

### Options

```
      --body string        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string       The chain to use. (options: base, ethereum, arbitrum, hyperevm) [required]
  -e, --eip-712 string     JSON value (one of: { domain: object, types: object, message: object } | { types: object, domain: object, message: object })
  -h, --help               help for gas-sponsorship-prepare
      --owner string       The wallet address that owns the Product Account. [required]
  -p, --product string     Which product the gas sponsorship is for. Determines which Product Account (Safe) address to use. (options: earn, credit, tokenized_assets)
      --sender string      The address of the wallet which will send the transaction. [required]
      --signature string   The EIP-712 signed typed data signature. [required]
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
