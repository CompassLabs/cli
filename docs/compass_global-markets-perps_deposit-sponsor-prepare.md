## compass global-markets-perps deposit-sponsor-prepare

Build the Bridge2 deposit tx from a signed permit

### Synopsis

Build the Arbitrum tx that completes a USDC deposit to HL.

Takes the EIP-2612 permit signature returned by /deposit and returns
a fully-encoded `Bridge2.batchedDepositWithPermit` call. The integrator's
sponsor wallet (`sender`) broadcasts the returned tx — Compass does not
broadcast and does not hold gas keys.

```
compass global-markets-perps deposit-sponsor-prepare [flags]
```

### Examples

```
  compass global-markets-perps deposit-sponsor-prepare --owner 0x06A9aF046187895AcFc7258450B15397CAc67400 --amount-raw 100000000 --deadline 1747097383 --signature 0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 --sender 0x4A83b4413CF41C3244027e1590E35a0F48403F0c
```

### Options

```
  -a, --amount-raw int     USDC amount in raw 6-decimal units, as returned by /deposit. [required]
      --body string        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --deadline int       Permit deadline (unix seconds), as returned in the permit.message from /deposit. [required]
  -h, --help               help for deposit-sponsor-prepare
      --owner string       The user's EOA address on Arbitrum (the permit signer). [required]
      --sender from        Sponsor wallet address that will broadcast the returned tx (used for from and nonce). [required]
      --signature string   The user's EIP-2612 permit signature (65-byte hex, 0x-prefixed or raw). [required]
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

* [compass global-markets-perps](compass_global-markets-perps.md)	 - Operations for global-markets-perps
