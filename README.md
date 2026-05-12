# cli

Command-line interface for the *Compass* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=github.com/CompassLabs/cli&utm_campaign=cli)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This CLI is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/compasslabs/api). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Compass API: Compass Labs DeFi API
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [cli](#cli)
  * [CLI Installation](#cli-installation)
  * [Shell Completion](#shell-completion)
  * [CLI Example Usage](#cli-example-usage)
  * [Authentication](#authentication)
  * [Available Commands](#available-commands)
  * [Request Body Input](#request-body-input)
  * [Server Selection](#server-selection)
  * [Output Formats](#output-formats)
  * [Error Handling](#error-handling)
  * [Diagnostics](#diagnostics)
  * [Wallet Signing — Local Daemon](#wallet-signing-local-daemon)
  * [Common Pitfalls](#common-pitfalls)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start CLI Installation [installation] -->
## CLI Installation

### Quick Install (Linux/macOS)

```bash
curl -fsSL https://raw.githubusercontent.com/CompassLabs/cli/main/scripts/install.sh | bash
```

### Quick Install (Windows PowerShell)

```powershell
iwr -useb https://raw.githubusercontent.com/CompassLabs/cli/main/scripts/install.ps1 | iex
```

### Go Install

Alternatively, install directly via Go:

```bash
go install github.com/CompassLabs/cli/cmd/compass@latest
```

### Manual Download

Download pre-built binaries for your platform from the [releases page](https://github.com/CompassLabs/cli/releases).
<!-- End CLI Installation [installation] -->

<!-- Start Shell Completion [completion] -->
## Shell Completion

Shell completions are available for Bash, Zsh, Fish, and PowerShell.

### Bash

```bash
# Add to ~/.bashrc:
source <(compass completion bash)

# Or install permanently:
compass completion bash > /etc/bash_completion.d/compass
```

### Zsh

```zsh
# Add to ~/.zshrc:
source <(compass completion zsh)

# Or install permanently:
compass completion zsh > "${fpath[1]}/_compass"
```

### Fish

```fish
compass completion fish | source

# Or install permanently:
compass completion fish > ~/.config/fish/completions/compass.fish
```

### PowerShell

```powershell
compass completion powershell | Out-String | Invoke-Expression
```
<!-- End Shell Completion [completion] -->

<!-- Start CLI Example Usage [usage] -->
## CLI Example Usage

### Example

```bash
compass gas-sponsorship gas-sponsorship-prepare --api-key-auth test_api_key --owner 0xCE1A77F0abff993d6d3D04d44b70831c6924fb40 --chain arbitrum --eip-712 '{"domain":{"name":"USD Coin","version":"2","chainId":42161,"verifyingContract":"0xaf88d065e77c8cC2239327C5EDb3A432268e5831"},"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Permit":[{"name":"owner","type":"address"},{"name":"spender","type":"address"},{"name":"value","type":"uint256"},{"name":"nonce","type":"uint256"},{"name":"deadline","type":"uint256"}]},"primaryType":"Permit","message":{"owner":"0xCE1A77F0abff993d6d3D04d44b70831c6924fb40","spender":"0x000000000022D473030F116dDEE9F6B43aC78BA3","value":"115792089237316195423570985008687907853269984665640564039457584007913129639935","nonce":"0","deadline":"1762269774"\x7d\x7d' --signature 0x160d2709ae195f591daa33ad6ab1fb18b8762a39d8c4466c4cbe95cf6881fc3d54d469710ef0e7fd64ecff47c1ba5741d7254903bfaebdacea5aa8289f81ba9a1c --sender 0x02122Ac49b0Be2e0eAD957F2D080805A0127Aa9d

```
<!-- End CLI Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

Authentication credentials can be configured in four ways (in order of priority):

### 1. Command-line flags

Pass credentials directly as flags to any command:

```bash
compass --api-key-auth <value> <command> [arguments]
```

### 2. Environment variables

Set credentials via environment variables:

| Variable | Description |
|----------|-------------|
| `COMPASS_API_KEY_AUTH` | Your Compass API Key. Get your key [here](https://www.compasslabs.ai/dashboard). |

### 3. OS Keychain (recommended for workstations)

Credentials are stored securely in your operating system's keychain when you run:

```bash
compass configure
```

Secret credentials (tokens, API keys, passwords) are automatically stored in:
- **macOS**: Keychain
- **Linux**: GNOME Keyring / KWallet (via D-Bus Secret Service)
- **Windows**: Windows Credential Locker

If no keychain is available (e.g., in CI environments), credentials fall back to the config file.

### 4. Configuration file

Run the interactive `configure` command to store non-secret settings:

```bash
compass configure
```

Configuration is stored in `~/.config/compass/config.yaml`.
<!-- End Authentication [security] -->

<!-- Start Available Commands [operations] -->
## Available Commands

<details open>
<summary>Available commands</summary>

### [gas-sponsorship](docs/compass_gas-sponsorship.md)

* [`gas-sponsorship-prepare`](docs/compass_gas-sponsorship_gas-sponsorship-prepare.md) - Prepare gas-sponsored transaction
* [`gas-sponsorship-approve-transfer`](docs/compass_gas-sponsorship_gas-sponsorship-approve-transfer.md) - Approve token transfer

### [earn](docs/compass_earn.md)

* [`earn-positions`](docs/compass_earn_earn-positions.md) - List earn positions
* [`earn-positions-all`](docs/compass_earn_earn-positions-all.md) - List earn positions across all chains
* [`earn-vaults`](docs/compass_earn_earn-vaults.md) - List vaults
* [`earn-aave-markets`](docs/compass_earn_earn-aave-markets.md) - List aave markets
* [`earn-pendle-markets`](docs/compass_earn_earn-pendle-markets.md) - List pendle markets
* [`earn-balances`](docs/compass_earn_earn-balances.md) - Get token balances
* [`earn-create-account`](docs/compass_earn_earn-create-account.md) - Create earn account
* [`earn-transfer`](docs/compass_earn_earn-transfer.md) - Transfer tokens to/from account
* [`earn-manage`](docs/compass_earn_earn-manage.md) - Manage earn position
* [`earn-swap`](docs/compass_earn_earn-swap.md) - Swap tokens within Earn Account
* [`earn-bundle`](docs/compass_earn_earn-bundle.md) - Execute multiple earn actions

### [credit](docs/compass_credit.md)

* [`credit-positions`](docs/compass_credit_credit-positions.md) - List credit positions
* [`credit-balances`](docs/compass_credit_credit-balances.md) - Get credit account token balances
* [`credit-create-account`](docs/compass_credit_credit-create-account.md) - Create credit account
* [`credit-borrow`](docs/compass_credit_credit-borrow.md) - Borrow against collateral
* [`credit-transfer`](docs/compass_credit_credit-transfer.md) - Transfer tokens to/from Credit Account
* [`credit-repay`](docs/compass_credit_credit-repay.md) - Repay debt and withdraw collateral
* [`credit-bundle`](docs/compass_credit_credit-bundle.md) - Execute multiple credit actions

### [global-markets-perps](docs/compass_global-markets-perps.md)

* [`global-markets-perps-opportunities`](docs/compass_global-markets-perps_global-markets-perps-opportunities.md) - List global markets perps markets
* [`global-markets-perps-positions`](docs/compass_global-markets-perps_global-markets-perps-positions.md) - List global markets perps positions
* [`global-markets-perps-deposit`](docs/compass_global-markets-perps_global-markets-perps-deposit.md) - Deposit USDC to global markets perps account
* [`global-markets-perps-withdraw`](docs/compass_global-markets-perps_global-markets-perps-withdraw.md) - Withdraw USDC from global markets perps account
* [`global-markets-perps-market-order`](docs/compass_global-markets-perps_global-markets-perps-market-order.md) - Place market order
* [`global-markets-perps-limit-order`](docs/compass_global-markets-perps_global-markets-perps-limit-order.md) - Place limit order
* [`global-markets-perps-cancel-order`](docs/compass_global-markets-perps_global-markets-perps-cancel-order.md) - Cancel order
* [`global-markets-perps-execute`](docs/compass_global-markets-perps_global-markets-perps-execute.md) - Execute signed action
* [`global-markets-perps-approve-builder-fee`](docs/compass_global-markets-perps_global-markets-perps-approve-builder-fee.md) - Approve builder fee
* [`global-markets-perps-enable-unified-account`](docs/compass_global-markets-perps_global-markets-perps-enable-unified-account.md) - Enable unified account mode
* [`global-markets-perps-ensure-leverage`](docs/compass_global-markets-perps_global-markets-perps-ensure-leverage.md) - Ensure 1x cross leverage

### [tokenized-assets](docs/compass_tokenized-assets.md)

* [`tokenized-assets-markets`](docs/compass_tokenized-assets_tokenized-assets-markets.md) - List tokenized equity markets
* [`tokenized-assets-markets-symbol`](docs/compass_tokenized-assets_tokenized-assets-markets-symbol.md) - Get a single market
* [`tokenized-assets-positions`](docs/compass_tokenized-assets_tokenized-assets-positions.md) - Get tokenized-asset positions for a wallet
* [`tokenized-assets-order-order-hash`](docs/compass_tokenized-assets_tokenized-assets-order-order-hash.md) - Get order status
* [`tokenized-assets-create-account`](docs/compass_tokenized-assets_tokenized-assets-create-account.md) - Create a Tokenized Assets Account
* [`tokenized-assets-quote`](docs/compass_tokenized-assets_tokenized-assets-quote.md) - Preview a buy/sell quote
* [`tokenized-assets-order`](docs/compass_tokenized-assets_tokenized-assets-order.md) - Build a buy/sell order
* [`tokenized-assets-order-submit`](docs/compass_tokenized-assets_tokenized-assets-order-submit.md) - Submit a signed order
* [`tokenized-assets-order-order-hash-cancel`](docs/compass_tokenized-assets_tokenized-assets-order-order-hash-cancel.md) - Cancel an unfilled order

</details>
<!-- End Available Commands [operations] -->

<!-- Start Request Body Input [stdinpiping] -->
## Request Body Input

Operations that accept a request body support three input methods, with a clear priority chain:

### Individual flags (highest priority)

```bash
compass <command> --name "Jane" --age 30
```

### `--body` flag

Provide the entire request body as a JSON string:

```bash
compass <command> --body '{"name": "John", "age": 30}'
```

Individual flags override `--body` values:

```bash
# Result: {name: "Jane", age: 30}
compass <command> --body '{"name": "John", "age": 30}' --name "Jane"
```

### Stdin piping (lowest priority)

Pipe JSON into any command that accepts a request body:

```bash
echo '{"name": "John", "age": 30}' | compass <command>
```

Individual flags override stdin values:

```bash
# Result: {name: "Jane", age: 30}
echo '{"name": "John", "age": 30}' | compass <command> --name "Jane"
```

This is useful for chaining commands, reading from files, or scripting:

```bash
# Read body from a file
compass <command> < request.json

# Pipe from another command
curl -s https://example.com/data.json | compass <command>
```

### Priority

When multiple input methods are used, the priority is:

| Priority | Source | Description |
|----------|--------|-------------|
| 1 (highest) | Individual flags | `--name "Jane"` always wins |
| 2 | `--body` flag | Whole-body JSON via flag |
| 3 (lowest) | Stdin | Piped JSON input |
<!-- End Request Body Input [stdinpiping] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL

Use `--server-url` to override the server URL entirely, bypassing any named or indexed server selection:

```bash
compass --server-url https://custom-api.example.com <command> [arguments]
```

**Precedence**: `--server-url` > `--server` > default
<!-- End Server Selection [server] -->

<!-- Start Output Formats [output-formats] -->
## Output Formats

Every command supports a `--output-format` flag that controls how the response is rendered to stdout.

### Available formats

| Format | Flag | Description |
|--------|------|-------------|
| Pretty | `--output-format pretty` (default) | Aligned key-value pairs with color, nested indentation. Human-readable at a glance. |
| JSON | `--output-format json` | JSON output. Passthrough when the response is already JSON (preserves original field order and numeric precision). Falls back to typed marshaling otherwise. |
| YAML | `--output-format yaml` | YAML output via standard marshaling. |
| Table | `--output-format table` | Tabular output for array responses. |
| TOON | `--output-format toon` | [Token-Oriented Object Notation](https://github.com/toon-format/spec) — a compact, line-oriented format that typically uses 30–60% fewer tokens than JSON. Well-suited for piping responses into LLM prompts. |

```bash
# Default pretty output
compass <command>

# Machine-readable JSON
compass <command> --output-format json

# TOON for LLM-friendly compact output
compass <command> --output-format toon

# Pipe JSON to jq without using --output-format
compass <command> --output-format json | jq '.fieldName'
```

### jq filtering

Use `--jq` to filter or transform the response inline using a [jq](https://jqlang.org) expression. This always outputs JSON and overrides `--output-format`:

```bash
# Extract a single field
compass <command> --jq '.name'

# Filter an array
compass <command> --jq '.items[] | select(.active == true)'
```

### Color control

Use `--color` to control terminal colors:

| Value | Behavior |
|-------|----------|
| `auto` (default) | Color when stdout is a TTY, plain text otherwise |
| `always` | Always colorize |
| `never` | Never colorize |

The `NO_COLOR` and `FORCE_COLOR` environment variables are also respected.

### Streaming and pagination

When using `--all` (pagination) or streaming operations, output is written incrementally as items arrive:

| Format | Streaming behavior |
|--------|-------------------|
| `json` | One compact JSON object per line ([NDJSON](https://github.com/ndjson/ndjson-spec)) |
| `yaml` | YAML documents separated by `---` |
| `toon` | One TOON-encoded object per block, separated by blank lines |
| `pretty` (default) | Pretty-printed items separated by blank lines |
<!-- End Output Formats [output-formats] -->

<!-- Start Error Handling [errors] -->
## Error Handling

The CLI uses standard exit codes to indicate success or failure:

| Exit Code | Meaning |
|-----------|---------|
| `0` | Success |
| `1` | Error (API error, invalid input, etc.) |

On success, the response data is printed to **stdout** as JSON. On failure, error details are printed to **stderr**.

```bash
# Capture output and handle errors
compass ... > output.json 2> error.log
if [ $? -ne 0 ]; then
  echo "Error occurred, see error.log"
fi
```
<!-- End Error Handling [errors] -->

<!-- Start Diagnostics [diagnostics] -->
## Diagnostics

The CLI includes two diagnostic flags available on all commands:

### Dry Run

Preview what would be sent without making any network calls:

```bash
compass <command> --dry-run
```

Output goes to stderr and includes:
- HTTP method and URL
- Request headers (sensitive values redacted)
- Request body preview (sensitive fields redacted)

The command exits successfully without contacting the API. This is useful for verifying request construction before executing.

### Debug

Log request and response diagnostics while running normally:

```bash
compass <command> --debug
```

Debug output goes to stderr and includes:
- Request method, URL, headers, and body preview
- Response status, headers, and body preview
- Transport errors (if any)

The command still executes normally and produces its regular output on stdout.

### Flag Precedence

If both `--dry-run` and `--debug` are set, `--dry-run` takes precedence and no network calls are made.

### Security

Sensitive information is automatically redacted in diagnostic output:
- **Headers**: `Authorization`, `Cookie`, `Set-Cookie`, `X-API-Key`, and other security headers show `[REDACTED]`
- **Body**: JSON fields named `password`, `secret`, `token`, `api_key`, `client_secret`, etc. show `[REDACTED]`

Diagnostic output should still be treated as potentially sensitive operational data.
<!-- End Diagnostics [diagnostics] -->

## Wallet Signing — Local Daemon

The CLI is non-custodial: action endpoints return either an unsigned transaction or EIP-712 typed data, and the signing happens in the user's wallet. The CLI ships with a small **embedded signer daemon** that pairs with the `compass sign` subcommand. The daemon binary, page, and JS are all baked into `compass` itself — no Node, no separate install.

```bash
# 1. Start the daemon (foreground; opens http://127.0.0.1:3030 in your browser)
compass daemon start

# 2. In the browser: click "Connect wallet" — uses your installed wallet
#    (MetaMask, Rabby, Coinbase Wallet, Frame, etc. via window.ethereum).

# 3. In another terminal, pipe any signing-shaped output into `compass sign`
compass earn earn-manage --venue.vault.vault-address 0x... \
  --action DEPOSIT --amount 100 --owner 0x... --chain base \
  --output-format json --jq '.unsigned_tx' \
  | compass sign
```

`compass sign` reads JSON from stdin (or `--input <file>`), POSTs it to the daemon, and prints the signature or transaction hash. The browser tab shows an approval prompt for every incoming request.

| Setting | Default | Override |
|---------|---------|----------|
| Daemon port | `3030` | `compass daemon start --port <n>` |
| Signer URL (for `compass sign`) | `http://127.0.0.1:3030/api/sign` | `$COMPASS_SIGNER_URL` or `--signer-url` |
| Auto-open browser | enabled | `compass daemon start --no-open` |
| Approval timeout | 5 minutes | n/a |
| Display metadata | none | `--meta key=value` (repeatable, shown in the approval prompt) |

The HTTP server binds to `127.0.0.1` and rejects non-loopback connections — the daemon cannot be reached over the network.

## Common Pitfalls

These are real footguns surfaced during prod testing. AI coding agents should read this section before writing CLI invocations — most of these errors are not obvious from the per-command help text.

### Some optional flags require JSON-quoted values

Flags whose help text reads as a regular string but that are implemented as JSON-encoded query parameters require **JSON-quoted** values. Affected today (commands × flags):

| Command | Flag | Wrong | Right |
|---------|------|-------|-------|
| `earn earn-vaults` | `--chain`, `--asset-symbol` | `--chain base` | `--chain '"base"'` |
| `earn earn-aave-markets` | `--chain` | `--chain base` | `--chain '"base"'` |
| `earn earn-pendle-markets` | `--chain` | `--chain base` | `--chain '"base"'` |
| `tokenized-assets tokenized-assets-markets` | `--search` | `--search USDC` | `--search '"USDC"'` |

**Symptom:** an error like `invalid value for --chain: error unmarshalling json response body: invalid character 'b' looking for beginning of value`. The phrase "response body" is misleading — this is flag parsing, not an HTTP error.

**Workaround:** wrap the value in JSON-quoted form (`'"value"'` in zsh/bash). Or, if the flag is optional, omit it entirely.

**Why:** Speakeasy generates these as `FlagKindJSON` because they're optional query parameters; the generator does `json.Unmarshal` on the raw input. Required enum/string flags use a different code path and accept plain values normally. Tracked for upstream/overlay fix.

### `-o table` does not unwrap response envelopes

Endpoints that return paginated lists wrap the array in an envelope, e.g. `earn-vaults` returns `{ total, offset, limit, vaults: [...] }`. `--output-format table` renders the envelope, not the array.

```bash
# Tiny scalar table — not what you want
compass earn earn-vaults --order-by tvl_usd -o table

# Drill into the array (output is JSON; --jq forces JSON regardless of -o)
compass earn earn-vaults --order-by tvl_usd --jq '.vaults'
```

The same pattern applies to other list endpoints (`earn-aave-markets`, `tokenized-assets-markets`, etc.). If you want a table of the inner array specifically, that's not supported in a single command today — pipe the JSON result through another tool, or just consume the JSON.

### Flag metavars can be misleading

Some flags display unusual metavars in `--help` listings — these are placeholders inferred from the OpenAPI example, not part of the syntax:

```
--action venue              # "venue" is a metavar; pass DEPOSIT or WITHDRAW
--gas-sponsorship true      # "true" is a metavar; the flag is a bool
--amount from_token         # "from_token" is a metavar; pass a number string like "100"
--from-token TSLAon         # "TSLAon" is an example value; pass any token symbol or address
--to-token from_token       # "from_token" is a metavar; pass any token symbol or address
```

Trust the **Description** column over the metavar.

### Command names contain "stutter"

Several commands repeat their group name: `compass earn earn-vaults`, `compass credit credit-transfer`, etc. This is by current design (the OpenAPI tag and operationId both include the domain word). When piping `--help` output or relying on tab completion, expect the duplicated prefix. The full command tree is in `docs/compass.md`.

### Other quirks

- **Line continuations**: `compass <cmd> \` followed by a trailing space (especially after pasting) breaks zsh — strip trailing whitespace or use a one-liner.
- **No `--api-key` / `--api-key-auth-key` flag**: the auth flag is `--api-key-auth` (single token).
- **`COMPASS_API_KEY` doesn't work**; the env var is `COMPASS_API_KEY_AUTH`.

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This CLI is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this CLI, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### CLI Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github.com/CompassLabs/cli&utm_campaign=cli)
