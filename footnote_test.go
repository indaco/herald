package herald

import (
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// FootnoteRef
// ---------------------------------------------------------------------------

func TestFootnoteRef(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		n        int
		contains string
	}{
		{"ref 1", 1, "[1]"},
		{"ref 99", 99, "[99]"},
		{"ref 0", 0, "[0]"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := stripANSI(ty.FootnoteRef(tc.n))
			if !strings.Contains(result, tc.contains) {
				t.Errorf("FootnoteRef(%d): expected %q in %q", tc.n, tc.contains, result)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// FootnoteSection
// ---------------------------------------------------------------------------

func TestFootnoteSection(t *testing.T) {
	ty := newTestTypography()

	t.Run("empty list returns empty string", func(t *testing.T) {
		result := ty.FootnoteSection(nil)
		if result != "" {
			t.Errorf("FootnoteSection(nil) should be empty, got %q", result)
		}
	})

	t.Run("empty slice returns empty string", func(t *testing.T) {
		result := ty.FootnoteSection([]string{})
		if result != "" {
			t.Errorf("FootnoteSection([]) should be empty, got %q", result)
		}
	})

	t.Run("single note", func(t *testing.T) {
		result := stripANSI(ty.FootnoteSection([]string{"First note"}))
		if !strings.Contains(result, "[1] First note") {
			t.Errorf("expected '[1] First note' in %q", result)
		}
		// Verify divider is present
		if !strings.Contains(result, strings.Repeat(DefaultFootnoteDividerChar, DefaultFootnoteDividerWidth)) {
			t.Errorf("expected divider in %q", result)
		}
	})

	t.Run("multiple notes", func(t *testing.T) {
		notes := []string{"Built on lipgloss v2", "Headings, lists, alerts, and more"}
		result := stripANSI(ty.FootnoteSection(notes))

		if !strings.Contains(result, "[1] Built on lipgloss v2") {
			t.Errorf("expected '[1] Built on lipgloss v2' in %q", result)
		}
		if !strings.Contains(result, "[2] Headings, lists, alerts, and more") {
			t.Errorf("expected '[2] Headings, lists, alerts, and more' in %q", result)
		}
		// Verify divider is present
		if !strings.Contains(result, strings.Repeat(DefaultFootnoteDividerChar, DefaultFootnoteDividerWidth)) {
			t.Errorf("expected divider in %q", result)
		}
	})

	t.Run("divider appears before notes", func(t *testing.T) {
		result := stripANSI(ty.FootnoteSection([]string{"A note"}))
		lines := strings.Split(result, "\n")
		if len(lines) < 2 {
			t.Fatalf("expected at least 2 lines, got %d", len(lines))
		}
		// First line should be the divider
		if !strings.Contains(lines[0], strings.Repeat(DefaultFootnoteDividerChar, DefaultFootnoteDividerWidth)) {
			t.Errorf("first line should be the divider, got %q", lines[0])
		}
		// Second line should be the note
		if !strings.Contains(lines[1], "[1] A note") {
			t.Errorf("second line should be the note, got %q", lines[1])
		}
	})
}

// ---------------------------------------------------------------------------
// FootnoteCustomDivider
// ---------------------------------------------------------------------------

func TestFootnoteCustomDivider(t *testing.T) {
	t.Run("custom char", func(t *testing.T) {
		ty := New(WithFootnoteDividerChar("="))
		result := stripANSI(ty.FootnoteSection([]string{"A note"}))
		if !strings.Contains(result, strings.Repeat("=", DefaultFootnoteDividerWidth)) {
			t.Errorf("expected custom divider char '=' in %q", result)
		}
	})

	t.Run("custom width", func(t *testing.T) {
		ty := New(WithFootnoteDividerWidth(10))
		result := stripANSI(ty.FootnoteSection([]string{"A note"}))
		if !strings.Contains(result, strings.Repeat(DefaultFootnoteDividerChar, 10)) {
			t.Errorf("expected divider width 10 in %q", result)
		}
	})

	t.Run("custom char and width", func(t *testing.T) {
		ty := New(WithFootnoteDividerChar("*"), WithFootnoteDividerWidth(5))
		result := stripANSI(ty.FootnoteSection([]string{"A note"}))
		if !strings.Contains(result, "*****") {
			t.Errorf("expected '*****' in %q", result)
		}
	})
}
