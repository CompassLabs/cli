## compass credit bundle

Execute multiple credit actions

### Synopsis

Compose individual credit operations (supply, withdraw, borrow, repay, swap, transfer)
into a single atomic Safe transaction.

The Credit Account must already be created via `/v2/credit/create_account`.

```
compass credit bundle [flags]
```

### Examples

```
  compass credit bundle --owner <value> --chain arbitrum --actions '[{"body":{"action_type":"V2_TRANSFER_FROM_EOA","token":"<value>","amount":"709.93","permit2_signature":"<value>","permit2_nonce":210517,"permit2_deadline":120606} }]'
```

### Options

```
  -a, --actions string    List of actions to bundle. Actions are executed in order. [required]
      --body string       Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string      The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo) [required]
  -g, --gas-sponsorship   If true, returns EIP-712 typed data for gas-sponsored execution.
  -h, --help              help for bundle
      --owner string      The owner's wallet address that controls the Credit Account. [required]
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
