## compass tokenized-assets tokenized-assets-positions

Get tokenized-asset positions for a wallet

### Synopsis

Get the tokenized-asset holdings for a wallet.

Returns the balance of every listed tokenized equity at the queried
address, plus the latest USD price and a USD-valued balance when
pricing is available. Zero balances are omitted, and a `total_usd`
aggregate is returned across all priced positions.

Pass the **Tokenized Assets Account address** (returned by
`/create_account`), not the owner's wallet — proceeds from filled
orders settle into the Tokenized Assets Account.

```
compass tokenized-assets tokenized-assets-positions [flags]
```

### Examples

```
  compass tokenized-assets tokenized-assets-positions --wallet 0x29F20a192328eF1aD35e1564aBFf4Be9C5ce5f7B
```

### Options

```
  -h, --help            help for tokenized-assets-positions
  -w, --wallet string   Wallet address to read on-chain ERC-20 balances for. [required]
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
