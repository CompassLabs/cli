## compass credit credit-looped-positions

List looped (leveraged) credit positions

### Synopsis

List looped (leveraged) positions for a credit account owner.

Detects loops from the account's on-chain history: a transaction containing
lending + borrowing + swap legs is a loop transaction. Returns one position
per Morpho market / Aave collateral+debt reserve pair, each with its complete
per-transaction history, lifetime totals, and live on-chain state (health
factor, USD values, leverage) for open positions. Covers Aave V3 and Morpho
Blue.

```
compass credit credit-looped-positions [flags]
```

### Examples

```
  compass credit credit-looped-positions --chain base --owner 0x06A9aF046187895AcFc7258450B15397CAc67400
```

### Options

```
  -c, --chain string   options: arbitrum, base, bsc, ethereum, tempo [required]
  -h, --help           help for credit-looped-positions
      --owner string   The address of the owner of the credit account to get looped positions for. [required]
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
