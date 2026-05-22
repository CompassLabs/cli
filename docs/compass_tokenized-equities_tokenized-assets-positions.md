## compass tokenized-equities tokenized-assets-positions

Get tokenized-asset positions for an owner

### Synopsis

Get the tokenized-asset holdings for an owner.

The owner's Tokenized Equities Account address is derived deterministically
from the `owner` query param; balances are read from that account (proceeds
from filled orders settle there). The response returns the balance of every
listed tokenized equity, plus the latest USD price and a USD-valued balance
when pricing is available. Zero balances are omitted, and a `total_usd`
aggregate is returned across all priced positions.

Returns 400 `ACCOUNT_NOT_DEPLOYED` if the owner has no Tokenized Equities
Account deployed yet — create one via `/create_account` first.

```
compass tokenized-equities tokenized-assets-positions [flags]
```

### Examples

```
  compass tokenized-equities tokenized-assets-positions --owner 0x29F20a192328eF1aD35e1564aBFf4Be9C5ce5f7B
```

### Options

```
  -h, --help           help for tokenized-assets-positions
      --owner string   The address of the owner of the Tokenized Equities Account to get positions for. The account address is derived deterministically from this owner; balances are read from the derived account. [required]
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

* [compass tokenized-equities](compass_tokenized-equities.md)	 - Operations for tokenized-equities
