## compass tokenized-assets tokenized-assets-order-order-hash-charge-fee

Charge a partner fee on a filled sell order's USDC proceeds

### Synopsis

Build a USDC fee transfer on a filled equity sell order's proceeds.

Equity orders fill off-chain via a third-party resolver, so the fee can't be
bundled into the trade. Once the sell order has filled, call this with the
order hash and your `fee` (recipient + percentage/fixed); it reads the actual
filled USDC proceeds and returns a `transfer(recipient, fee)` executed by the
product account — an unsigned transaction the owner signs, or an EIP-712
payload when `gas_sponsorship` is true.

```
compass tokenized-assets tokenized-assets-order-order-hash-charge-fee [flags]
```

### Examples

```
  compass tokenized-assets tokenized-assets-order-order-hash-charge-fee --order-hash <value> --owner <value> --fee '{"recipient":"<value>","amount":8619.46,"denomination":"PERCENTAGE"}'
```

### Options

```
      --body string         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -f, --fee Fee             Fee configuration for tokenized-asset trades.
                            
                            Same shape as Fee, but narrows `denomination` to the values supported by
                            tokenized-asset trades. PERFORMANCE is not supported because realized-profit
                            fees are not computed for spot buy/sell trades. [required]
  -g, --gas-sponsorship     When true, returns an EIP-712 payload for gas-sponsored execution instead of an unsigned transaction.
  -h, --help                help for tokenized-assets-order-order-hash-charge-fee
      --order-hash string   [required]
      --owner string        The owner's wallet address; the product account is derived from it. [required]
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
