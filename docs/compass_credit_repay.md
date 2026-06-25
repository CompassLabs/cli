## compass credit repay

Repay debt and withdraw collateral

### Synopsis

Repay an Aave debt and withdraw collateral from a Credit Account.

Bundles repayment, collateral withdrawal, and an optional swap into a single atomic Safe transaction.

- If `token_out` is None or equals `withdraw_token`, the withdrawn collateral is kept as-is.
- If `token_out` differs from `withdraw_token`, a swap is performed after withdrawal via 1inch.

The Credit Account must already have a borrow position created via `/v2/credit/borrow`.
The repay_token must be available in the Credit Account (or pulled from EOA via Permit2).

```
compass credit repay [flags]
```

### Examples

```
  compass credit repay --owner 0x831Ad0C52C77708DA5D49dc8278C966dfdD4ddA1 --chain base --repay-token WETH --repay-amount 0.01
```

### Options

```
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --borrow-vault string         Euler only: the EVK vault the debt is owed to (repay target). Required when protocol=EULER.
      --chain string                The chain to use. (options: base, ethereum, arbitrum, hyperevm, tempo) [required]
      --collateral-vault string     Euler only: the EVK collateral vault to withdraw from. Required when protocol=EULER and withdrawing collateral.
  -g, --gas-sponsorship             If true, returns EIP-712 signature data instead of an unsigned transaction.
  -h, --help                        help for repay
  -i, --interest-rate-mode string   On AAVE there are 2 different interest modes.
                                    
                                    A stable (but typically higher rate), or a variable rate. (options: stable, variable)
      --owner string                The address that owns the Credit Account. [required]
      --permit2-deadline string     The deadline timestamp used in the Permit2 signature (from the signed typed data).
      --permit2-nonce string        The nonce used in the Permit2 signature (from the signed typed data).
      --permit2-signature string    The EOA owner's signature of the Permit2 PermitTransferFrom typed data. When provided, the repay bundle will first pull repay_token from the owner's EOA into the Credit Account via Permit2. Obtain by calling /v2/credit/transfer and signing the returned EIP-712 data.
      --protocol                    Which lending protocol a credit action targets.
                                    
                                    AAVE`` is the default so existing callers (which never send a ``protocol``
                                    field) keep hitting the unchanged Aave code path. ``EULER`` opts in to the
                                    Euler V2 path, where the market is identified by EVK vault address(es). (options: AAVE, EULER)
      --repay-amount string         JSON value (one of: number | string)
      --repay-token string          The borrowed asset to repay (e.g. WETH). Must match the debt position's token. [required]
  -s, --slippage string             JSON value (one of: number | string)
  -t, --token-out string            Desired output token. If different from withdraw_token, a swap is performed after withdrawal. If None, the withdrawn collateral is kept as-is.
      --withdraw-amount string      JSON value (one of: number | string)
      --withdraw-token string       Collateral token to withdraw from Aave after repaying debt. Omit together with withdraw_amount for repay-only mode.
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
