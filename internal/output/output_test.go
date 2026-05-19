package output

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/CompassLabs/cli/internal/sdk/models/components"
	"github.com/CompassLabs/cli/internal/sdk/optionalnullable"
)

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	cmd.PersistentFlags().StringP("output-format", "o", "pretty", "")
	cmd.PersistentFlags().String("color", "never", "")
	cmd.PersistentFlags().StringP("jq", "q", "", "")
	cmd.PersistentFlags().Bool("include-headers", false, "")
	cmd.PersistentFlags().String("server", "", "")
	return cmd
}

// testEnvelope mirrors the shape Speakeasy generates for SDK responses
// so output.Result's reflection-based tryReadRawBody / extractResultContent
// find the same fields they would in production.
type testEnvelope struct {
	HTTPMeta components.HTTPMetadata
	Data     any
}

func envelopeWithRawBody(body string) *testEnvelope {
	return &testEnvelope{
		HTTPMeta: components.HTTPMetadata{
			Response: &http.Response{
				Body: io.NopCloser(strings.NewReader(body)),
			},
		},
	}
}

// TestResult_JSONOutputIsIndented covers Fix #3:
// When -o json receives minified JSON from the API, the CLI should pretty-print
// the body so a TTY user can read it.
func TestResult_JSONOutputIsIndented(t *testing.T) {
	cmd := newTestCmd()
	require.NoError(t, cmd.PersistentFlags().Set("output-format", "json"))

	minified := `{"total":2,"vaults":[{"name":"A"},{"name":"B"}]}`
	res := envelopeWithRawBody(minified)

	var buf bytes.Buffer
	cmd.SetOut(&buf)

	require.NoError(t, Result(cmd, res))

	got := buf.String()
	assert.Contains(t, got, "\n", "JSON output should be multi-line, got: %q", got)
	assert.Contains(t, got, "  ", "JSON output should be indented with two spaces")

	var parsed map[string]any
	require.NoError(t, json.Unmarshal([]byte(got), &parsed), "indented output must still parse as JSON")
	assert.EqualValues(t, 2, parsed["total"])
}

// TestResult_JSONPassesThroughNonJSONBody covers the existing fallback:
// if the raw body isn't valid JSON, the CLI should still print it (don't break).
func TestResult_JSONPassesThroughNonJSONBody(t *testing.T) {
	cmd := newTestCmd()
	require.NoError(t, cmd.PersistentFlags().Set("output-format", "json"))

	res := envelopeWithRawBody("not json at all")

	var buf bytes.Buffer
	cmd.SetOut(&buf)

	require.NoError(t, Result(cmd, res))
	assert.Contains(t, buf.String(), "not json at all")
}

// --- Fix #1: table envelope unwrap ------------------------------------------

type fakeVault struct {
	Name    string `json:"name"`
	Chain   string `json:"chain"`
	TvlUsd  string `json:"tvl_usd"`
}

type fakeVaultsEnvelope struct {
	Total  int64       `json:"total"`
	Offset int64       `json:"offset"`
	Limit  int64       `json:"limit"`
	Vaults []fakeVault `json:"vaults"`
}

type fakeMultiList struct {
	Aave    []fakeVault `json:"aave"`
	Pendle  []fakeVault `json:"pendle"`
}

// TestPrintTable_UnwrapsPaginatedEnvelope covers Fix #1:
// `-o table` on a {total, offset, limit, items[]} envelope must render the
// items as rows, not the envelope metadata.
func TestPrintTable_UnwrapsPaginatedEnvelope(t *testing.T) {
	content := fakeVaultsEnvelope{
		Total: 1701, Offset: 0, Limit: 2,
		Vaults: []fakeVault{
			{Name: "Steakhouse", Chain: "base", TvlUsd: "468M"},
			{Name: "Gauntlet", Chain: "ethereum", TvlUsd: "120M"},
		},
	}

	var buf bytes.Buffer
	require.NoError(t, printTable(&buf, content))
	got := buf.String()

	// Items appear as rows.
	assert.Contains(t, got, "Steakhouse")
	assert.Contains(t, got, "Gauntlet")
	assert.Contains(t, got, "ethereum")

	// Per-item columns (uppercased headers).
	assert.Contains(t, got, "NAME")
	assert.Contains(t, got, "CHAIN")
	assert.Contains(t, got, "TVL_USD")

	// Pagination metadata is surfaced too, but not as the only output.
	assert.Contains(t, got, "1701")
}

// TestPrintTable_EmptyListSaysSo covers the edge case where the slice is
// empty — we should clearly say so instead of silently printing only metadata.
func TestPrintTable_EmptyListSaysSo(t *testing.T) {
	content := fakeVaultsEnvelope{Total: 0, Vaults: []fakeVault{}}

	var buf bytes.Buffer
	require.NoError(t, printTable(&buf, content))
	got := buf.String()

	assert.Contains(t, strings.ToLower(got), "empty",
		"empty list should be explicit; got: %q", got)
}

// TestPrintTable_MultiListRendersEachSection covers responses that contain
// multiple slice-of-struct fields (e.g. earn-positions returns aave + vaults + pendle_pt).
func TestPrintTable_MultiListRendersEachSection(t *testing.T) {
	content := fakeMultiList{
		Aave:   []fakeVault{{Name: "aave-1", Chain: "base"}},
		Pendle: []fakeVault{{Name: "pendle-1", Chain: "ethereum"}},
	}

	var buf bytes.Buffer
	require.NoError(t, printTable(&buf, content))
	got := buf.String()

	assert.Contains(t, got, "aave-1")
	assert.Contains(t, got, "pendle-1")
}

// TestPrintTable_SingleStructPreserved guards the existing behavior for
// non-list responses: a plain struct still renders as a vertical key/value table.
func TestPrintTable_SingleStructPreserved(t *testing.T) {
	single := fakeVault{Name: "Steakhouse", Chain: "base", TvlUsd: "468M"}

	var buf bytes.Buffer
	require.NoError(t, printTable(&buf, single))
	got := buf.String()

	assert.Contains(t, got, "NAME")
	assert.Contains(t, got, "Steakhouse")
	assert.Contains(t, got, "CHAIN")
	assert.Contains(t, got, "base")
}

// Vault-style struct using OptionalNullable[string] for fields the SDK
// models as nullable (apy, tvl, names, ...). These are the columns most
// users actually want to see, so the table must surface them.
type fakeVaultWithNullable struct {
	Chain  string                                    `json:"chain"`
	Name   optionalnullable.OptionalNullable[string] `json:"name"`
	Apy30d optionalnullable.OptionalNullable[string] `json:"apy_30d"`
}

// TestPrintTable_UnwrapsOptionalNullableColumns ensures the most useful
// columns of SDK responses (which model nullable fields as
// optionalnullable.OptionalNullable[T]) actually appear in the table.
func TestPrintTable_UnwrapsOptionalNullableColumns(t *testing.T) {
	name := "Steakhouse"
	apy := "4.29"
	rows := []fakeVaultWithNullable{
		{
			Chain:  "base",
			Name:   optionalnullable.From(&name),
			Apy30d: optionalnullable.From(&apy),
		},
	}

	var buf bytes.Buffer
	require.NoError(t, printTable(&buf, rows))
	got := buf.String()

	assert.Contains(t, got, "NAME", "name column should appear in header")
	assert.Contains(t, got, "APY_30D", "apy_30d column should appear in header")
	assert.Contains(t, got, "Steakhouse", "name value should appear in row")
	assert.Contains(t, got, "4.29", "apy value should appear in row")
}

// TestPrintTable_OptionalNullableNullPrintsBlank confirms both representations
// of "no value" render as blank cells, not the underlying map.
//
// OptionalNullable[T] = map[bool]*T has three states:
//  1. UNSET           — zero value (nil map)                  | `{}`
//  2. EXPLICIT NULL   — set, value pointer is nil             | `From[string](nil)`
//  3. SET TO VALUE    — set, value pointer non-nil            | `From(&s)`
//
// unwrapOptionalNullable takes different reflect branches for (1) vs (2)
// (`!inner.IsValid()` vs `inner.IsNil()`), so both must be tested or a
// regression in one path could slip through.
func TestPrintTable_OptionalNullableNullPrintsBlank(t *testing.T) {
	rows := []fakeVaultWithNullable{
		{Chain: "row-unset"}, // Name + Apy30d unset (zero-value map)
		{
			Chain:  "row-explicit",
			Name:   optionalnullable.From[string](nil), // explicit null
			Apy30d: optionalnullable.From[string](nil),
		},
	}

	var buf bytes.Buffer
	require.NoError(t, printTable(&buf, rows))
	got := buf.String()

	assert.Contains(t, got, "row-unset")
	assert.Contains(t, got, "row-explicit")
	assert.NotContains(t, got, "map[", "OptionalNullable should never surface as map[...]")
	assert.NotContains(t, got, "0xc0", "OptionalNullable should never surface as a pointer address")
	assert.NotContains(t, got, "null", "explicit-null cells should render as blank, not the literal 'null'")
}

// --- Fix #2: pretty default for list responses ------------------------------

// SDK-shaped response wrappers: HTTPMeta + one payload field. Result() uses
// extractResultContent to peel the envelope before printing, so tests must
// mirror that shape.
type fakeVaultsResp struct {
	HTTPMeta       components.HTTPMetadata
	VaultsResponse *fakeVaultsEnvelope
}

type fakeVaultResp struct {
	HTTPMeta components.HTTPMetadata
	Vault    *fakeVault
}

// TestResult_PrettyDefaultsToTableForLists covers Fix #2:
// The default human-readable output for a list/envelope response should be the
// compact table, not the full nested key:value dump.
func TestResult_PrettyDefaultsToTableForLists(t *testing.T) {
	cmd := newTestCmd()
	// Default format is "pretty".
	res := &fakeVaultsResp{
		VaultsResponse: &fakeVaultsEnvelope{
			Total: 1701, Offset: 0, Limit: 2,
			Vaults: []fakeVault{
				{Name: "Steakhouse", Chain: "base", TvlUsd: "468M"},
				{Name: "Gauntlet", Chain: "ethereum", TvlUsd: "120M"},
			},
		},
	}

	var buf bytes.Buffer
	cmd.SetOut(&buf)

	require.NoError(t, Result(cmd, res))
	got := buf.String()

	// Compact, table-style: uppercased column headers, one line per row.
	assert.Contains(t, got, "NAME")
	assert.Contains(t, got, "CHAIN")
	assert.Contains(t, got, "Steakhouse")
	assert.Contains(t, got, "Gauntlet")

	// Crucially: NOT the verbose nested key:value form (would render
	// "- name:" / "name: " on its own line per item).
	assert.NotContains(t, got, "- name:")
}

// TestPrettyPrint_EmptyTopLevelSliceSaysEmpty covers a consistency gap:
// when the response is a top-level empty slice-of-struct, pretty mode should
// render "(empty)" — same as the envelope path — instead of falling through
// to the JSON walk and printing "[]".
func TestPrettyPrint_EmptyTopLevelSliceSaysEmpty(t *testing.T) {
	empty := []fakeVault{}

	var buf bytes.Buffer
	require.NoError(t, prettyPrint(&buf, empty, false))
	got := buf.String()

	assert.Contains(t, strings.ToLower(got), "empty",
		"top-level empty slice should render as '(empty)', got: %q", got)
	assert.NotContains(t, got, "[]",
		"should not fall back to the JSON-walk '[]' for empty slices")
}

// TestResult_PrettySingleItemStillVerbose makes sure non-list responses
// keep the existing key:value pretty rendering.
func TestResult_PrettySingleItemStillVerbose(t *testing.T) {
	cmd := newTestCmd()
	res := &fakeVaultResp{
		Vault: &fakeVault{Name: "Steakhouse", Chain: "base", TvlUsd: "468M"},
	}

	var buf bytes.Buffer
	cmd.SetOut(&buf)

	require.NoError(t, Result(cmd, res))
	got := buf.String()

	// Verbose key:value, lower-case keys, one field per line.
	assert.Contains(t, got, "name:")
	assert.Contains(t, got, "chain:")
	assert.Contains(t, got, "Steakhouse")
}
