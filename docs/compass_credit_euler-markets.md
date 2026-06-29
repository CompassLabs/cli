## compass credit euler-markets

List curated Euler markets

### Synopsis

List curated Euler V2 credit markets for a chain.

Euler is permissionless, so credit borrow/repay identify a market by EVK vault
address. This returns the borrow markets from Euler's Governed Perspective (the
DAO-vetted vault set) -- each with the collateral vaults it accepts and their
LTVs -- so callers know which borrow_vault / collateral_vault to use.

```
compass credit euler-markets [flags]
```

### Examples

```
  compass credit euler-markets --chain ethereum
```

### Options

```
  -c, --chain string   options: arbitrum, base, ethereum, tempo [required]
  -h, --help           help for euler-markets
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
