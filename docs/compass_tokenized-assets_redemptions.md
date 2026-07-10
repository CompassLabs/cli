## compass tokenized-assets redemptions

List redemption requests

### Synopsis

Get an owner's IXS managed-vault redemption requests and their status.

IXS vault sells are asynchronous — `/transact/sell` files a `requestRedeem`
that the vault operator settles off-chain later. This reconstructs the owner's
requests live from the vault (Compass stores no async state), reporting each as
`pending`, `finalized`, or `rejected`. IXS vaults live on BNB Smart Chain
(`chain=bsc`, vault `ixv1`).

```
compass tokenized-assets redemptions [flags]
```

### Examples

```
  compass tokenized-assets redemptions --owner 0x29F20a192328eF1aD35e1564aBFf4Be9C5ce5f7B
```

### Options

```
  -c, --chain string   Network the IXS vault lives on (IXS managed vaults are on BNB Smart Chain). (options: base, ethereum, arbitrum, hyperevm, tempo, bsc)
  -h, --help           help for redemptions
      --owner string   The owner of the Tokenized Assets Account whose IXS redemption requests to read. Requests are reconstructed on-chain from the derived product account (the request receiver). [required]
  -v, --vault string   The IXS vault handle to read redemptions for (e.g. 'ixv1').
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
