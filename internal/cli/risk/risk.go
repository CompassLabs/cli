// Package risk holds hand-written CLI commands that don't come from
// the OpenAPI spec. Currently just `risk-recipes`, which prints the
// embedded RISK_RECIPES.md so agents using the installed binary can
// read it without a clone of the repo or web access.
//
// Speakeasy will not touch this package on regeneration (it only
// rewrites files it generated). The one fragile point is the
// registration call in internal/cli/root.go — that file IS regenerated
// and the call has to be re-injected. The
// .github/workflows/generate_combined_spec_and_sdks.yaml workflow
// patches it back in after `speakeasy run`.
package risk

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed RISK_RECIPES.md
var recipesMarkdown string

// InitRiskRoot registers the hand-written risk commands on the given
// parent (typically the cobra root command).
func InitRiskRoot(parent *cobra.Command) error {
	cmd := &cobra.Command{
		Use:   "risk-recipes",
		Short: "Print risk-analysis recipes (LLTV cascade, JTD, vault correlation)",
		Long: `Print formulas and worked recipes for computing the same risk metrics shown on compasslabs.ai/risk — borrower exposure (JTD), LLTV cascade stress test, vault correlation matrix, yield decomposition + withdrawable — using only the existing earn-vaults / earn-aave-markets / earn-pendle-markets endpoints.

The doc is embedded in the binary at compile time, so this command works offline. Source: cli-sdk/internal/cli/risk/RISK_RECIPES.md.

Typical usage:
  compass risk-recipes              # print to stdout
  compass risk-recipes | less       # paged
  compass risk-recipes > recipes.md # save to file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprint(cmd.OutOrStdout(), recipesMarkdown)
			return err
		},
	}
	parent.AddCommand(cmd)
	return nil
}
