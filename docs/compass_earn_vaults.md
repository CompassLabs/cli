## compass earn vaults

List vaults

### Synopsis

List ERC-4626 yield vaults across DeFi venues.

Returns vault data including APY, TVL, and underlying asset information. Use this endpoint to discover yield opportunities, compare rates across venues, or build vault selection interfaces.

Supports dozens of vaults and markets like Morpho and other ERC-4626 compatible yield venues.

To deposit into a vault, use the [manage endpoint](https://docs.compasslabs.ai/v2/api-reference/earn/manage-earn-position) with `venue_type=VAULTS`.

```
compass earn vaults [flags]
```

### Examples

```
  compass earn vaults --order-by tvl_usd
```

### Options

```
  -a, --asset-symbol string          Filter vaults by underlying asset symbol (e.g., 'USDC', 'WETH').
  -c, --chain string                 Optional chain filter. If not provided, returns vaults for all chains. (options: base, ethereum, arbitrum, hyperevm, tempo, bsc, robinhood, ethereum_sepolia)
      --direction string             The direction to order the results by. (options: asc, desc)
  -h, --help                         help for vaults
  -l, --limit int                    The number of items to return.
      --min-deposit-cap-usd string   JSON value (one of: number | string)
      --min-liquidity-usd string     JSON value (one of: number | string)
      --min-tvl-usd string           JSON value (one of: number | string)
      --offset int                   The offset of the first item to return.
      --order-by string              The field to order the results by. [required]
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --api-key-auth string    Your Compass API Key. Get your key [here](https://www.compasslabs.ai/dashboard).
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
      --interactive            Prompt for missing inputs and open guided configure/auth forms (forms fall back to line prompts on stdin off-TTY) (default true)
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [compass earn](compass_earn.md)	 - Operations for earn

### Machine interface

* `compass earn vaults --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `compass earn vaults --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
