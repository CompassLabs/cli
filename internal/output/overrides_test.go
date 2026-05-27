package output

import (
	"errors"
	"strings"
	"testing"

	"github.com/CompassLabs/cli/internal/sdk/optionalnullable"
)

// vault is a deliberately-small mirror of the Speakeasy-generated struct
// shape (`,omitzero` tags + OptionalNullable[T]) so we can lock in the
// YAML/TOON encoder behavior without depending on the entire SDK model.
type vault struct {
	VaultAddress string                                    `json:"vault_address"`
	Name         optionalnullable.OptionalNullable[string] `json:"name,omitzero"`
	Apy30D       optionalnullable.OptionalNullable[string] `json:"apy_30d,omitzero"`
}

func TestNormalizeForEncoding_stripsOmitzeroAndUnwrapsOptionalNullable(t *testing.T) {
	t.Parallel()
	name := "Steakhouse Prime USDC"
	apy := "4.1917"
	in := vault{
		VaultAddress: "0xBEEF",
		Name:         optionalnullable.From(&name),
		Apy30D:       optionalnullable.From(&apy),
	}

	out, err := normalizeForEncoding(in)
	if err != nil {
		t.Fatalf("normalizeForEncoding returned err: %v", err)
	}
	m, ok := out.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", out)
	}
	// Key MUST be the bare json name, not "name,omitzero".
	if _, leaked := m["name,omitzero"]; leaked {
		t.Fatalf("comma-suffix tag leaked into output keys: %v", m)
	}
	if got, want := m["name"], "Steakhouse Prime USDC"; got != want {
		// Catches the YAML `{true: "..."}` regression — OptionalNullable
		// must be unwrapped to the underlying string value.
		t.Fatalf("name not unwrapped: got %#v want %q", got, want)
	}
	if got, want := m["apy_30d"], "4.1917"; got != want {
		t.Fatalf("apy_30d not unwrapped: got %#v want %q", got, want)
	}
	if got, want := m["vault_address"], "0xBEEF"; got != want {
		t.Fatalf("vault_address mismatch: got %#v want %q", got, want)
	}
}

func TestNormalizeForEncoding_preservesInt64Precision(t *testing.T) {
	t.Parallel()
	// 2^60 + 1 — well above float64's 53-bit integer precision, well below
	// int64 max. Without UseNumber() this round-trips as float64 and the
	// trailing digits are silently lost, which would mangle DeFi amounts in
	// YAML/TOON output.
	type row struct {
		Amount int64 `json:"amount"`
	}
	in := row{Amount: 1152921504606846977}
	out, err := normalizeForEncoding(in)
	if err != nil {
		t.Fatalf("normalizeForEncoding returned err: %v", err)
	}
	m, ok := out.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", out)
	}
	got, ok := m["amount"].(int64)
	if !ok {
		t.Fatalf("amount must round-trip as int64, got %T (value %v)", m["amount"], m["amount"])
	}
	if got != 1152921504606846977 {
		t.Fatalf("int64 precision lost: got %d want 1152921504606846977", got)
	}
}

func TestNormalizeForEncoding_preservesUint64Precision(t *testing.T) {
	t.Parallel()
	// 12.3 ETH expressed in wei = 1.23e19 — above int64 max (~9.22e18) but
	// well within uint64 max (~1.84e19). Without the uint64 step this falls
	// back to float64 and the trailing digits get rounded, garbling a typical
	// vault-TVL or large-balance value in YAML/TOON output.
	type row struct {
		WeiAmount uint64 `json:"wei_amount"`
	}
	in := row{WeiAmount: 12_300_000_000_000_000_007}
	out, err := normalizeForEncoding(in)
	if err != nil {
		t.Fatalf("normalizeForEncoding returned err: %v", err)
	}
	m, ok := out.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any, got %T", out)
	}
	got, ok := m["wei_amount"].(uint64)
	if !ok {
		t.Fatalf("wei_amount must round-trip as uint64, got %T (value %v)", m["wei_amount"], m["wei_amount"])
	}
	if got != 12_300_000_000_000_000_007 {
		t.Fatalf("uint64 precision lost: got %d want 12300000000000000007", got)
	}
}

func TestAlreadyPrinted_roundtripsAndDetects(t *testing.T) {
	t.Parallel()
	base := errors.New("api boom")
	wrapped := AlreadyPrinted(base)

	if wrapped == nil {
		t.Fatalf("AlreadyPrinted returned nil for non-nil err")
	}
	if !IsAlreadyPrinted(wrapped) {
		t.Fatalf("IsAlreadyPrinted should detect the wrapped marker")
	}
	if !errors.Is(wrapped, base) {
		t.Fatalf("errors.Is should still find the original err through Unwrap")
	}
	if got := wrapped.Error(); got != "api boom" {
		t.Fatalf("wrapped.Error() should pass through: got %q", got)
	}
	// Double-wrap must be a no-op so we don't bury the marker.
	if got := AlreadyPrinted(wrapped); got != wrapped {
		t.Fatalf("double-wrap should be identity")
	}
	if AlreadyPrinted(nil) != nil {
		t.Fatalf("AlreadyPrinted(nil) must stay nil so callers can pass it through")
	}
}

func TestFormatFastAPIDetail_collapsesAndDedupes(t *testing.T) {
	t.Parallel()
	// Mirrors the real API payload that produced "Error: …" lines twice
	// in v0.1.8: identical entries duplicated under .detail.
	payload := map[string]any{
		"detail": []any{
			map[string]any{
				"loc":  []any{"query", "owner"},
				"msg":  "Value error, Addresses must start with 0x",
				"type": "value_error",
			},
			map[string]any{
				"loc":  []any{"query", "owner"},
				"msg":  "Value error, Addresses must start with 0x",
				"type": "value_error",
			},
		},
	}
	got, ok := formatFastAPIDetail(payload)
	if !ok {
		t.Fatalf("expected FastAPI shape to be detected")
	}
	if strings.Count(got, "query.owner") != 1 {
		t.Fatalf("duplicate detail entries were not deduped:\n%s", got)
	}
	if !strings.Contains(got, "Validation error:") {
		t.Fatalf("singular header missing for single unique entry:\n%s", got)
	}
}

func TestFormatFastAPIDetail_ignoresNonFastAPIShape(t *testing.T) {
	t.Parallel()
	// A `{error, message}` body (the shape Compass uses for application
	// errors) must NOT match — otherwise we'd hide caller-facing fields.
	payload := map[string]any{
		"error":   "Insufficient balance",
		"message": "Insufficient USDC balance in account",
	}
	if _, ok := formatFastAPIDetail(payload); ok {
		t.Fatalf("non-FastAPI shape was incorrectly matched")
	}
}
