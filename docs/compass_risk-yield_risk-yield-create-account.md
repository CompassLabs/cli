## compass risk-yield risk-yield-create-account

Create account

### Synopsis

Create the Risk Yield Account that will hold your LP positions.

Liquidity is provided from an account, not from your wallet, so this comes
first. The address is deterministic from the `owner`, and it is the **same
account as your Earn Account** on this chain — if you already have one,
this returns it with `transaction: null` and nothing needs to be signed.

**If the owner pays gas:** set `sender` to the owner's address.

**If someone else pays:** set `sender` to whoever will broadcast.

```
compass risk-yield risk-yield-create-account [flags]
```

### Examples

```
  compass risk-yield risk-yield-create-account --sender 0xdB035cb494EEE30996fbDD25a7b1Fa38D795bdC6 --owner 0xdB035cb494EEE30996fbDD25a7b1Fa38D795bdC6
```

### Options

```
      --body string     Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --chain string    Risk Yield is available on Robinhood Chain only. (options: robinhood)
  -e, --estimate-gas    Estimate gas, which also proves the transaction can execute.
  -h, --help            help for risk-yield-create-account
      --owner string    The address that will own and control the Risk Yield Account. [required]
  -s, --sender string   The address that signs and broadcasts this transaction. [required]
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

* [compass risk-yield](compass_risk-yield.md)	 - Operations for risk-yield
