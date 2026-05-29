## compass earn create-account

Create earn account

### Synopsis

Create an Earn Account for a wallet address.

Before depositing into venues or managing positions, the owner must create an Earn Account. Each wallet address has one Earn Account per chain.

Returns an unsigned transaction to create the account. The `sender` signs and broadcasts this transaction.

**If owner pays gas:** Set `sender` to the owner's address.

**If someone else pays gas:** Set `sender` to the wallet that will sign and broadcast the transaction on behalf of the owner.

```
compass earn create-account [flags]
```

### Examples

```
  compass earn create-account --chain base --sender 0x01E62835dd7F52173546A325294762143eE4a882 --owner 0x01E62835dd7F52173546A325294762143eE4a882
```

### Options

```
      --body string     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string    options: arbitrum, base, ethereum, tempo [required]
  -e, --estimate-gas    Determines whether to estimate gas costs for transactions, also verifying that the transaction can be successfully executed.
  -h, --help            help for create-account
      --owner string    The address that will own and control the compass account [required]
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

* [compass earn](compass_earn.md)	 - Operations for earn
