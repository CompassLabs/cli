## compass credit loop

Open a leveraged loop

### Synopsis

Open a leveraged position in ONE atomic transaction.

Supply collateral, borrow against it, swap the borrowed token back to
collateral, repeat — until total collateral reaches multiplier x
initial_collateral_amount. If any step fails the whole transaction reverts,
so a position is never left half-built.

See the [Leveraged Looping guide](https://docs.compasslabs.ai/v2/Products/Looping)
for chain coverage, market discovery, vault-share collateral, HyperEVM
specifics and the full error list.

```
compass credit loop [flags]
```

### Examples

```
  compass credit loop --owner 0x5e5b00ed886A6879C2B934612D2312975427fcAf --chain ethereum --collateral-token WETH --borrow-token USDC --initial-collateral-amount 1
```

### Options

```
      --body string                        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --borrow-token string                Token borrowed each iteration and swapped back to the collateral token. For MORPHO it must be the market's loan token. [required]
      --borrow-vault string                Euler only: the EVK vault address borrowed from (the sub-account's controller). Required when protocol=EULER.
      --chain string                       Blockchain network. Not every protocol is deployed on every chain — see the protocol field — and a chain with no credit venue at all returns a 422 naming the chains that do. (options: arbitrum, base, bsc, ethereum, hyperevm, tempo) [required]
      --collateral-token string            Token supplied as collateral each iteration. Must already be in the Credit Account for the initial amount. For MORPHO it must be the market's collateral token. [required]
      --collateral-vault string            Euler only: the EVK vault address collateral is supplied to (from /v2/credit/euler_markets). Required when protocol=EULER.
  -e, --emode-category string              Aave only: e-mode category to enable before looping (higher LTV for correlated pairs, e.g. ETH-correlated). Both tokens must belong to the category or the build returns a 422. On Hyperlend (chain=hyperevm) category 1 is the HYPE-correlated set covering wstHYPE and WHYPE.
  -g, --gas-sponsorship                    If true, returns EIP-712 typed data for gas-sponsored execution instead of an unsigned transaction.
  -h, --help                               help for loop
  -i, --initial-collateral-amount string   JSON value (one of: number | string)
  -l, --loan-to-value string               JSON value (one of: number | string)
      --market-id string                   Morpho only: the bytes32 market id (from /v2/credit/morpho_markets). Required when protocol=MORPHO.
      --max-slippage-percent string        JSON value (one of: number | string)
      --multiplier string                  JSON value (one of: number | string)
      --owner string                       The address that owns the Credit Account. [required]
      --preview pricing                    If true, build a display ESTIMATE: no firm RFQ quote is ever requested (quote_expires_at stays null). NOTE that this guarantees only that no firm quote was spent — it does not guarantee an absent transaction: on a pair no firm provider covers, and under pricing=market, the call still falls through to the aggregator and returns a signable transaction. How the estimate is priced follows pricing: on a firm-covered pair under 'auto' or 'firm' it is computed from the firm provider's live price levels (indicative); otherwise swap legs are priced by the default aggregator. Set it on every call made while a user is exploring parameters, and leave it false only for the build they actually intend to sign — firm quotes are single-use maker commitments, and requesting them for displays that are never executed degrades the pricing this API is offered.
      --pricing string                     Swap-leg routing policy. 'auto': firm quotes where a firm venue covers the pair, transparent fallback to the market aggregator otherwise. 'firm': never price on the market route — previews whose target the firm venue cannot serve return the coverage advisory alone (preview=null, zero aggregator calls), and executions fail with a typed error instead of silently substituting market pricing. 'market': never route through the firm venue; every leg is priced by the aggregator and bounded by max_slippage_percent (which firm legs ignore). 'firm' is incompatible with gas_sponsorship (sponsored loops force market routing). (options: auto, firm, market)
      --protocol                           Which lending protocol a credit action targets.
                                           
                                           AAVE`` is the default so existing callers (which never send a ``protocol``
                                           field) keep hitting the unchanged Aave code path; markets are named by token
                                           symbol. ``MORPHO`` identifies Morpho Blue lending markets by their bytes32
                                           ``market_id``. ``EULER`` identifies Euler V2 markets by their EVK
                                           ``collateral_vault`` + ``borrow_vault`` addresses and supports isolated
                                           per-sub-account positions (``sub_account_id``).
                                           
                                           Deployment is per chain, so a valid protocol can still 422 on a given chain:
                                           AAVE on Ethereum, Base, Arbitrum, BSC and HyperEVM (where it is Hyperlend, the
                                           chain's Aave V3 deployment); MORPHO on Ethereum, Base, Arbitrum and HyperEVM
                                           (where it is Felix); EULER on Ethereum, Base, Arbitrum and BSC.
                                           
                                           All three support ``/v2/credit/loop`` and ``/v2/credit/unloop``. EULER does
                                           NOT: ``/v2/credit/rebalance`` rejects it with a 422, and
                                           ``/v2/credit/looped_positions`` covers only AAVE and MORPHO — an Euler loop is
                                           silently absent there rather than an error, so read it from
                                           ``/v2/credit/positions`` instead. (EULER still appears in the
                                           ``looped_positions`` response enum because this enum is shared; it is never
                                           emitted.) (options: AAVE, EULER, MORPHO)
  -s, --sub-account-id int                 Euler only: the EVC sub-account (0-255) holding this isolated looped position. 0 is the Credit Account itself.
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
