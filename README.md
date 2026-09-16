# cli

Command-line interface for the *Compass* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=github.com/CompassLabs/cli&utm_campaign=cli)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)

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
  * [For AI agents](#for-ai-agents)
  * [Authentication](#authentication)
  * [Commands](#commands)
  * [Request Body Input](#request-body-input)
  * [Server Selection](#server-selection)
  * [Output Formats](#output-formats)
  * [Error Handling](#error-handling)
  * [Diagnostics](#diagnostics)
  * [Common Pitfalls](#common-pitfalls)
  * [Executing transactions (signing & broadcasting)](#executing-transactions-signing-broadcasting)
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
### Homebrew (macOS/Linux)

```bash
brew install CompassLabs/tap/compass
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
compass gas-sponsorship prepare --api-key-auth test_api_key --owner 0xCE1A77F0abff993d6d3D04d44b70831c6924fb40 --chain arbitrum --eip-712 '{"domain":{"name":"USD Coin","version":"2","chainId":42161,"verifyingContract":"0xaf88d065e77c8cC2239327C5EDb3A432268e5831"},"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Permit":[{"name":"owner","type":"address"},{"name":"spender","type":"address"},{"name":"value","type":"uint256"},{"name":"nonce","type":"uint256"},{"name":"deadline","type":"uint256"}]},"primaryType":"Permit","message":{"owner":"0xCE1A77F0abff993d6d3D04d44b70831c6924fb40","spender":"0x000000000022D473030F116dDEE9F6B43aC78BA3","value":"115792089237316195423570985008687907853269984665640564039457584007913129639935","nonce":"0","deadline":"1762269774"}}' --signature 0x160d2709ae195f591daa33ad6ab1fb18b8762a39d8c4466c4cbe95cf6881fc3d54d469710ef0e7fd64ecff47c1ba5741d7254903bfaebdacea5aa8289f81ba9a1c --sender 0x02122Ac49b0Be2e0eAD957F2D080805A0127Aa9d

```
<!-- End CLI Example Usage [usage] -->

<!-- Start For AI agents [agents] -->
## For AI agents

This CLI is built to be driven by AI coding agents as well as people: everything an agent needs is discoverable from the binary itself, and every command can be validated without credentials. Work down this ladder:

| Run | You get |
|-----|---------|
| `compass --help`, `compass earn aave-markets --help` | Commands by category, runnable examples, flags |
| `compass --usage`, `compass earn aave-markets --usage` | The command surface as machine-readable [KDL](https://kdl.dev): commands, aliases, flags, defaults, env vars, config keys |
| `compass credit create-account --schema` | The exact JSON Schema of the command's request body (all `$ref`s bundled) — build a valid `--body` from it |
| `compass earn aave-markets --dry-run` | The exact HTTP request (method, URL, headers, body), with no credentials or network call |
| `compass earn aave-markets --output-format json` (or `--jq`) | Machine-readable output |

### Discover the command surface

```bash
# Every command, flag, default, env var and config key, as KDL
compass --usage

# One command's subtree only
compass earn aave-markets --usage
```

### Read the exact request schema

`--schema` is available on every command that accepts a request body (`--body`, stdin, or a whole-body flag where the command has one), including intent commands. It prints the JSON Schema the request is validated against and exits without calling the API.

```bash
# JSON Schema (draft 2020-12) of the request body, with every $ref bundled under $defs
compass credit create-account --schema
```

### Probe before you spend

Start quota-spending commands with `--dry-run`. It validates inputs, resolves the request, redacts secrets and binary payloads, makes no network call, and exits 0. It never reads the OS keychain; credentials supplied by flag, environment, or config file are included only as `[REDACTED]`.

```bash
# Human preview: the [DRY-RUN] block is on stderr and stdout is empty
compass earn aave-markets --dry-run

# Machine preview: compact JSON on stdout and silent stderr
compass earn aave-markets --dry-run --output-format json
```

The machine form writes one object per would-be request, one per line (NDJSON for multi-request commands), with exactly this shape:

```json
{"dry_run":true,"request":{"method":"POST","url":"https://…","headers":{"Accept":["application/json"],…},"body":<JSON value | string | null>}}
```

`body` is a parsed JSON value when the body is JSON, a string for text, `"<bytes:N>"` for binary data, and `null` when absent. An explicit caller `--jq` also selects this JSON preview protocol, but the filter is not applied to preview objects. Command-declared jq presets do not select or filter the preview.

Local mutation commands make no request under `--dry-run`: instead of a preview they emit one `{"dry_run":true,"local":true,"command":"…","message":"…"}` object. `select(.request)` keeps only would-be requests; `select(.local)` keeps the local no-ops.

### Machine-readable output

```bash
# JSON on stdout
compass earn aave-markets --output-format json

# Filter or reshape with a jq expression (always emits JSON, overrides --output-format)
compass earn aave-markets --jq '.'

# Print jq string results as plain text instead of JSON strings (like jq -r)
compass earn aave-markets --jq '.' --raw-output
```

`--output-format toon` emits [TOON](https://github.com/toon-format/spec), a compact line-oriented format that uses fewer tokens than JSON; it is the default in agent mode.

### Interactive mode
Required-input prompts and guided `configure` / `auth login` forms are enabled by default. Required-input prompts require an interactive terminal; off-TTY forms read line input from stdin. Use `--no-interactive` to force flag-only execution.

```bash
# Prompt for missing command inputs
compass earn aave-markets --interactive

# Open the guided configuration form
compass configure --interactive

# Explicitly launch the terminal command explorer
compass explore
```

### Agent mode and structured errors
Agent mode turns on automatically when a known agent environment is detected (`CLAUDECODE`, `CURSOR_AGENT`, `CODEX`, `AIDER`, `CLINE`, `WINDSURF_AGENT`, `GITHUB_COPILOT`, `AMAZON_Q`, `GEMINI_CODE_ASSIST`, `SRC_CODY`) or with `--agent-mode` (`--agent-mode=false` disables detection).
In agent mode interactive prompts never launch, output defaults to TOON, and every failure — API errors and CLI usage errors alike — is one JSON envelope on stderr:
Outside agent mode, explicit JSON and `--jq` preserve the compatibility envelope without classification; enable agent mode to request the classified contract.

```json
{
  "error": "...",
  "error_type": "validation_error",
  "error_reason": "CLI_VALIDATION",
  "exit_code": 2,
  "message": "human-readable message",
  "hints": ["what to try next"]
}
```

`error_type` is one of `authentication_error`, `authorization_error`, `not_found`, `validation_error`, `rate_limit_error`, `server_error`, `api_error`, `connection_error`, `protocol_error`, `runtime_error`, `unsupported_error`, `async_failed`, `async_timeout`, `async_unknown_state`. Classification derives from the HTTP status and transport evidence; `error_reason` is absent for API errors. Status-less local failures may use `CLI_VALIDATION`, `CLI_CONNECTION`, `CLI_PROTOCOL`, `CLI_RUNTIME`, `CLI_UNAVAILABLE`, `CLI_AUTHENTICATION`, or the async polling reasons `CLI_ASYNC_FAILED`, `CLI_ASYNC_TIMEOUT`, and `CLI_ASYNC_UNKNOWN_STATE`. `hints` preserves server guidance first, adds the most specific local taxonomy guidance, then typed CLI and command-specific guidance, removing exact duplicates. `exit_code` is always the code for the final `error_type` shown in the envelope: 1 runtime, 2 usage, or 3 authentication/authorization.
<!-- End For AI agents [agents] -->

<!-- Start Authentication [security] -->
## Authentication

Authentication credentials can be configured in four ways (in order of priority):

### 1. Command-line flags

Pass credentials directly as flags to any command:

```bash
compass --api-key-auth "$COMPASS_API_KEY_AUTH" earn aave-markets
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

<!-- Start Commands [operations] -->
## Commands

<details open>
<summary>Available commands</summary>

* [`gas-sponsorship`](docs/compass_gas-sponsorship.md) - Operations for gas-sponsorship
  * [`prepare`](docs/compass_gas-sponsorship_prepare.md) - Prepare gas-sponsored transaction
  * [`approve-transfer`](docs/compass_gas-sponsorship_approve-transfer.md) - Approve token transfer
* [`earn`](docs/compass_earn.md) - Operations for earn
  * [`positions`](docs/compass_earn_positions.md) - List earn positions
  * [`positions-all`](docs/compass_earn_positions-all.md) - List earn positions across all chains
  * [`vaults`](docs/compass_earn_vaults.md) - List vaults
  * [`aave-markets`](docs/compass_earn_aave-markets.md) - List aave markets
  * [`pendle-markets`](docs/compass_earn_pendle-markets.md) - List pendle markets
  * [`swap-quote`](docs/compass_earn_swap-quote.md) - Quote a swap
  * [`balances`](docs/compass_earn_balances.md) - Get token balances
  * [`create-account`](docs/compass_earn_create-account.md) - Create earn account
  * [`transfer`](docs/compass_earn_transfer.md) - Transfer tokens to/from account
  * [`manage`](docs/compass_earn_manage.md) - Manage earn position
  * [`swap`](docs/compass_earn_swap.md) - Swap tokens
  * [`bundle`](docs/compass_earn_bundle.md) - Execute multiple earn actions
* [`credit`](docs/compass_credit.md) - Operations for credit
  * [`positions`](docs/compass_credit_positions.md) - List credit positions
  * [`balances`](docs/compass_credit_balances.md) - Get credit account token balances
  * [`looped-positions`](docs/compass_credit_looped-positions.md) - List looped (leveraged) credit positions
  * [`euler-markets`](docs/compass_credit_euler-markets.md) - List curated Euler markets
  * [`morpho-markets`](docs/compass_credit_morpho-markets.md) - List curated Morpho markets
  * [`create-account`](docs/compass_credit_create-account.md) - Create credit account
  * [`borrow`](docs/compass_credit_borrow.md) - Borrow against collateral
  * [`loop`](docs/compass_credit_loop.md) - Open a leveraged loop
  * [`unloop`](docs/compass_credit_unloop.md) - Unwind a leveraged loop
  * [`rebalance`](docs/compass_credit_rebalance.md) - Rebalance the leveraged credit book
  * [`transfer`](docs/compass_credit_transfer.md) - Transfer tokens to/from Credit Account
  * [`swap`](docs/compass_credit_swap.md) - Swap tokens
  * [`repay`](docs/compass_credit_repay.md) - Repay debt and withdraw collateral
  * [`bundle`](docs/compass_credit_bundle.md) - Execute multiple credit actions
* [`perpetual-trading`](docs/compass_perpetual-trading.md) - Operations for perpetual-trading
  * [`opportunities`](docs/compass_perpetual-trading_opportunities.md) - List perpetual trading markets
  * [`positions`](docs/compass_perpetual-trading_positions.md) - List perpetual trading positions
  * [`candles`](docs/compass_perpetual-trading_candles.md) - Get OHLCV candles
  * [`activity`](docs/compass_perpetual-trading_activity.md) - Aggregated Hyperliquid activity for a user
  * [`deposit`](docs/compass_perpetual-trading_deposit.md) - Deposit USDC to perpetual trading account
  * [`deposit-sponsor-prepare`](docs/compass_perpetual-trading_deposit-sponsor-prepare.md) - Build the Bridge2 deposit tx from a signed permit
  * [`withdraw`](docs/compass_perpetual-trading_withdraw.md) - Withdraw USDC from perpetual trading account
  * [`market-order`](docs/compass_perpetual-trading_market-order.md) - Place market order
  * [`limit-order`](docs/compass_perpetual-trading_limit-order.md) - Place limit order
  * [`cancel-order`](docs/compass_perpetual-trading_cancel-order.md) - Cancel order
  * [`execute`](docs/compass_perpetual-trading_execute.md) - Execute signed action
  * [`approve-builder-fee`](docs/compass_perpetual-trading_approve-builder-fee.md) - Approve builder fee
  * [`enable-unified-account`](docs/compass_perpetual-trading_enable-unified-account.md) - Enable unified account mode
  * [`ensure-leverage`](docs/compass_perpetual-trading_ensure-leverage.md) - Ensure 1x cross leverage
  * [`set-leverage`](docs/compass_perpetual-trading_set-leverage.md) - Set leverage (defaults to market maximum)
* [`tokenized-assets`](docs/compass_tokenized-assets.md) - Operations for tokenized-assets
  * [`markets`](docs/compass_tokenized-assets_markets.md) - List markets
  * [`market`](docs/compass_tokenized-assets_market.md) - Get a market
  * [`positions`](docs/compass_tokenized-assets_positions.md) - List positions
  * [`balances`](docs/compass_tokenized-assets_balances.md) - Get token balances
  * [`order-status`](docs/compass_tokenized-assets_order-status.md) - Get order status
  * [`redemptions`](docs/compass_tokenized-assets_redemptions.md) - List redemption requests
  * [`create-account`](docs/compass_tokenized-assets_create-account.md) - Create account
  * [`transfer`](docs/compass_tokenized-assets_transfer.md) - Transfer tokens to/from account
  * [`quote`](docs/compass_tokenized-assets_quote.md) - Quote an order
  * [`order`](docs/compass_tokenized-assets_order.md) - Build an order
  * [`order-submit`](docs/compass_tokenized-assets_order-submit.md) - Submit a signed order
  * [`order-cancel`](docs/compass_tokenized-assets_order-cancel.md) - Cancel an order
  * [`order-charge-fee`](docs/compass_tokenized-assets_order-charge-fee.md) - Charge a partner fee
  * [`buy`](docs/compass_tokenized-assets_buy.md) - Buy an RWA yield token
  * [`sell`](docs/compass_tokenized-assets_sell.md) - Sell an RWA yield token

</details>
<!-- End Commands [operations] -->

<!-- Start Request Body Input [stdinpiping] -->
## Request Body Input

Commands that accept a request body take it three ways, with a clear priority chain. The examples use `compass credit create-account`; every body-bearing command works the same way and prints its exact request schema with `--schema`.

### Individual flags (highest priority)

Each top-level body field is a flag:

```bash
compass credit create-account --chain 'base' --sender '0xc98949522db2eE403d6c75E91DDEe875a824bB10' --owner '0xc98949522db2eE403d6c75E91DDEe875a824bB10'
```

### `--body` flag

Provide the entire request body as a JSON string:

```bash
compass credit create-account --body '{"chain":"base","sender":"0xc98949522db2eE403d6c75E91DDEe875a824bB10","owner":"0xc98949522db2eE403d6c75E91DDEe875a824bB10"}'
```

Individual flags override `--body` values:

```bash
# Sends {"chain":"base (updated)","sender":"0xc98949522db2eE403d6c75E91DDEe875a824bB10","owner":"0xc98949522db2eE403d6c75E91DDEe875a824bB10"}
compass credit create-account --body '{"chain":"base","sender":"0xc98949522db2eE403d6c75E91DDEe875a824bB10","owner":"0xc98949522db2eE403d6c75E91DDEe875a824bB10"}' --chain 'base (updated)'
```

### Stdin piping (lowest priority)

Pipe JSON into any command that accepts a request body:

```bash
echo '{"chain":"base","sender":"0xc98949522db2eE403d6c75E91DDEe875a824bB10","owner":"0xc98949522db2eE403d6c75E91DDEe875a824bB10"}' | compass credit create-account
```

Individual flags override stdin values:

```bash
# Sends {"chain":"base (updated)","sender":"0xc98949522db2eE403d6c75E91DDEe875a824bB10","owner":"0xc98949522db2eE403d6c75E91DDEe875a824bB10"}
echo '{"chain":"base","sender":"0xc98949522db2eE403d6c75E91DDEe875a824bB10","owner":"0xc98949522db2eE403d6c75E91DDEe875a824bB10"}' | compass credit create-account --chain 'base (updated)'
```

This is useful for chaining commands, reading from files, or scripting:

```bash
# Read body from a file
compass credit create-account < request.json

# Pipe from another command
curl -s https://example.com/request.json | compass credit create-account
```

### Priority

When multiple input methods are used, the priority is:

| Priority | Source | Description |
|----------|--------|-------------|
| 1 (highest) | Individual flags | `--chain ...` always wins |
| 2 | `--body` flag | Whole-body JSON via flag |
| 3 (lowest) | Stdin | Piped JSON input |
<!-- End Request Body Input [stdinpiping] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL

Use `--server-url` to override the server URL entirely, bypassing any named or indexed server selection:

```bash
compass --server-url https://custom-api.example.com earn aave-markets
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
compass earn aave-markets

# Machine-readable JSON
compass earn aave-markets --output-format json

# TOON for LLM-friendly compact output
compass earn aave-markets --output-format toon

# Pipe JSON to jq without using --output-format
compass earn aave-markets --output-format json | jq '.'
```

### jq filtering

Use `--jq` to filter or transform the response inline using a [jq](https://jqlang.org) expression. This always outputs JSON and overrides `--output-format`:

```bash
# Extract a single field
compass earn aave-markets --jq '.'

# Reshape with any jq program; --raw-output prints string results as plain text (like jq -r)
compass earn aave-markets --jq '.' --raw-output
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
| `1` | Runtime/API failure |
| `2` | Usage or input failure |
| `3` | Authentication or authorization failure |

On success, the response data is printed to **stdout** as JSON. On failure, error details are printed to **stderr**.

```bash
# Capture output and handle errors
compass earn aave-markets --output-format json > output.json 2> error.log
if [ $? -ne 0 ]; then
  echo "Error occurred, see error.log"
fi
```
This CLI uses unclassified error rendering outside agent mode: pretty and TOON print the API error text as received, while `--output-format json` and `--jq` emit the unclassified envelope (including the configure `_hint` for HTTP 401/403) plus `exit_code`. Agent mode always emits the classified JSON envelope with `exit_code`, `error_type`, optional `error_reason`, `message`, `hints`, and optional `status_code` — see [For AI agents](#for-ai-agents).
<!-- End Error Handling [errors] -->

<!-- Start Diagnostics [diagnostics] -->
## Diagnostics

The CLI includes two diagnostic flags available on all commands:

### Dry Run

Preview what would be sent without making any network calls:

```bash
compass earn aave-markets --dry-run
```

In human output modes, stdout is empty and the `[DRY-RUN]` block goes to stderr. It includes:
- HTTP method and URL
- Request headers (sensitive values redacted)
- Request body preview (sensitive fields redacted)

With `--output-format json`, or with a caller-explicit `--jq`, stderr is silent and stdout is NDJSON: one compact preview object per would-be request. The jq filter is not applied, and command-declared jq presets do not select the JSON protocol.

```json
{"dry_run":true,"request":{"method":"POST","url":"https://…","headers":{"Accept":["application/json"],…},"body":<JSON value | string | null>}}
```

JSON bodies remain structured; text bodies are strings; binary bodies are `"<bytes:N>"`; absent bodies are `null`. Headers retain all values as arrays, with credentials replaced by `[REDACTED]`. Dry-run never reads the OS keychain, but credentials supplied by flag, environment, or config file still appear redacted. The command exits successfully without contacting the API.

Local mutation commands emit one `{"dry_run":true,"local":true,"command":"…","message":"…"}` object in place of a preview; filter with `select(.request)` or `select(.local)`.

### Debug

Log request and response diagnostics while running normally:

```bash
compass earn aave-markets --debug
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
- **Binary data**: binary media and canonical base64 strings are replaced with `<bytes:N>`
- **URL query**: credential-like query parameters are replaced with `[REDACTED]`

Diagnostic output should still be treated as potentially sensitive operational data.
<!-- End Diagnostics [diagnostics] -->

## Common Pitfalls

These are real footguns surfaced during prod testing. AI coding agents should read this section before writing CLI invocations — most of these errors are not obvious from the per-command help text.

### Optional string flags accept plain values

Historically, optional string-typed query-param flags (which Speakeasy generates as `FlagKindJSON`) needed a **JSON-quoted** value — `--chain '"base"'` — and a bare `--chain base` failed with a misleading `error unmarshalling json response body: invalid character 'b'`. Current versions auto-quote bare string values for these flags, so **plain values just work**:

```bash
compass earn vaults --order-by tvl_usd --chain base                     # ✓
compass earn aave-markets --chain base                                  # ✓
compass earn pendle-markets --order-by tvl_usd --underlying-symbol WETH # ✓
```

If you pin an **older** build and still hit the `unmarshalling` error, wrap the value in JSON quotes (`'"base"'`) or omit the optional flag. (That error says "response body" but it's flag parsing, not HTTP.)

### `-o table` unwraps flat list envelopes, not nested ones

List endpoints wrap their array in an envelope, e.g. `earn vaults` returns `{ total, offset, limit, vaults: [...] }`. `-o table` unwraps that automatically — it prints the envelope scalars as a header and renders the inner array as a table:

```bash
compass earn vaults --order-by tvl_usd -o table   # header + a real table
```

The table renderer only handles rows of flat scalars. On endpoints whose rows embed nested objects (e.g. `earn aave-markets`, whose rows carry reserve objects) it prints `no displayable fields found` — use `-o json`, `-o toon`, or `--jq` there instead.

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

### Other quirks

- **Line continuations**: `compass <cmd> \` followed by a trailing space (especially after pasting) breaks zsh — strip trailing whitespace or use a one-liner.
- **No `--api-key` / `--api-key-auth-key` flag**: the auth flag is `--api-key-auth` (single token).
- **`COMPASS_API_KEY` doesn't work**; the env var is `COMPASS_API_KEY_AUTH`.

## Executing transactions (signing & broadcasting)

The CLI is **non-custodial**: it never holds keys and never broadcasts. Read commands return data directly; **action** commands (`earn manage`, `credit borrow`, `credit loop`, `tokenized-assets order`, …) return an **unsigned transaction** — or EIP-712 typed data for gas-sponsored and order flows — for *you* to sign and submit. Every action is therefore a two-step "build → sign & send"; there is intentionally no `compass sign` / `compass send` (keeping signing in your own wallet is what makes the CLI non-custodial).

### Self-paid EOA transaction (build → `cast send`)

Action responses carry a `transaction` object with hex-encoded `to`, `data`, `value`, `gas`, and `chainId`. Sign and broadcast it with your own wallet — e.g. Foundry's `cast`:

```bash
# 1. Build the unsigned tx and keep the JSON (any action command works;
#    see `compass credit borrow --help` for the exact flags).
compass credit borrow --chain base --owner "$ADDR" \
  --borrow-token USDC --amount-in 100 -o json > tx.json

# 2. cast signs with your key and broadcasts.
cast send \
  "$(jq -r .transaction.to   tx.json)" \
  "$(jq -r .transaction.data tx.json)" \
  --value     "$(jq -r .transaction.value tx.json)" \
  --gas-limit "$(jq -r .transaction.gas   tx.json)" \
  --rpc-url "$RPC_URL" --private-key "$PRIVATE_KEY"
```

Pass the API's `gas` through as an explicit `--gas-limit` — `cast`'s own estimate can undershoot for Safe/bundled calls and revert with `GS013`.

### Gas-sponsored transaction (sign EIP-712 → a sponsor pays)

Add `--gas-sponsorship` to an action to get EIP-712 typed data instead of an unsigned tx. Sign it with your wallet, then submit the signature so a sponsor broadcasts and pays the gas:

```bash
compass gas-sponsorship prepare --owner "$ADDR" --chain arbitrum \
  --eip-712 '<typed-data-json>' --signature 0x<sig> --sender "$SPONSOR"
```

### Hyperliquid perps & tokenized-asset orders

These sign an exchange / EIP-712 payload rather than an EVM tx. Build with the command, sign the returned typed data, then submit the signature via `perpetual-trading execute` (perps) or `tokenized-assets order-submit` (equity orders) — no raw broadcast on your side.

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
