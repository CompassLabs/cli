## compass tokenized-assets tokenized-assets-create-account

Create a Tokenized Assets Account

### Synopsis

Create a Tokenized Assets Account for a wallet address.

Before placing orders, the owner must create a Tokenized Assets Account.
Each wallet address has one Tokenized Assets Account, isolated from the
owner's Earn, Credit, and other product accounts.

The account address is deterministic. If it already exists, the
response returns `transaction: null` and you can skip straight to
building orders.

Returns an unsigned transaction to create the account. The `sender`
signs and broadcasts this transaction.

**If owner pays gas:** Set `sender` to the owner's address.

**If someone else pays gas:** Set `sender` to the wallet that will
sign and broadcast the transaction on behalf of the owner.

```
compass tokenized-assets tokenized-assets-create-account [flags]
```

### Examples

```
  compass tokenized-assets tokenized-assets-create-account --sender 0x18b42407AbC163f595410Ffe773BB98Db40B48F7 --owner 0x18b42407AbC163f595410Ffe773BB98Db40B48F7
```

### Options

```
      --body string     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string    The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo)
  -e, --estimate-gas    Determines whether to estimate gas costs for transactions, also verifying that the transaction can be successfully executed.
  -h, --help            help for tokenized-assets-create-account
      --owner string    The address that will own and control the compass Tokenized Assets Account [required]
  -s, --sender string   The address of the transaction sender. [required]
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
