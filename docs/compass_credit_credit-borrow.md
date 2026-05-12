## compass credit credit-borrow

Borrow against collateral

### Synopsis

Borrow an asset from Aave using a Credit Account.

Bundles an optional swap, collateral supply, and borrow into a single atomic Safe transaction.

- If `token_in` equals `collateral_token`, the tokens are supplied directly as collateral.
- If `token_in` differs from `collateral_token`, a swap is performed first via 1inch.

The Credit Account must already be created via `/v2/credit/create_account` and funded with `token_in`.

```
compass credit credit-borrow [flags]
```

### Examples

```
  compass credit credit-borrow --owner 0xfeEb60eBf707DAA8248277058eb8D28EaeCF5425 --chain base --borrow-token WETH --borrow-amount 0.01
```

### Options

```
  -a, --amount-in string            JSON value (one of: number | string)
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --borrow-amount string        JSON value (one of: number | string)
      --borrow-token string         Asset to borrow from Aave. [required]
      --chain string                The chain to use. (options: base, ethereum, arbitrum, hyperevm) [required]
      --collateral-token string     Aave reserve token to supply as collateral. Omit together with token_in and amount_in for borrow-only mode.
  -f, --fee string                  Optional fee configuration. If provided, a fee will be deducted from the borrowed amount and sent to the specified recipient address.
  -g, --gas-sponsorship             If true, returns EIP-712 signature data instead of an unsigned transaction.
  -h, --help                        help for credit-borrow
  -i, --interest-rate-mode string   On AAVE there are 2 different interest modes.
                                    
                                    A stable (but typically higher rate), or a variable rate. (options: stable, variable)
      --owner string                The address that owns the Credit Account. [required]
      --permit2-deadline string     The deadline timestamp used in the Permit2 signature (from the signed typed data).
      --permit2-nonce string        The nonce used in the Permit2 signature (from the signed typed data).
      --permit2-signature string    The EOA owner's signature of the Permit2 PermitTransferFrom typed data. When provided, the borrow bundle will first pull token_in from the owner's EOA into the Credit Account via Permit2. Obtain by calling /v2/credit/transfer and signing the returned EIP-712 data.
  -s, --slippage string             JSON value (one of: number | string)
  -t, --token-in string             Token currently held in the Credit Account to use as input. If the same as collateral_token, no swap is performed. Omit together with amount_in and collateral_token to borrow against existing collateral.
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
