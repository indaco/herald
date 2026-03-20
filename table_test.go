package herald

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestTable(t *testing.T) {
	tests := []struct {
		name      string
		rows      [][]string
		wantEmpty bool
		contains  []string
	}{
		{
			name:      "nil input returns empty",
			rows:      nil,
			wantEmpty: true,
		},
		{
			name:      "empty slice returns empty",
			rows:      [][]string{},
			wantEmpty: true,
		},
		{
			name:      "single row with empty cells returns empty",
			rows:      [][]string{{}},
			wantEmpty: true,
		},
		{
			name: "single column table",
			rows: [][]string{
				{"Name"},
				{"Alice"},
				{"Bob"},
			},
			contains: []string{"Name", "Alice", "Bob"},
		},
		{
			name: "multi-column with header and body",
			rows: [][]string{
				{"Name", "Role", "Status"},
				{"Alice", "Admin", "Active"},
				{"Bob", "User", "Idle"},
			},
			contains: []string{
				"Name", "Role", "Status",
				"Alice", "Admin", "Active",
				"Bob", "User", "Idle",
				"┌", "┐", "└", "┘", "│", "─",
				"├", "┤", "┬", "┴", "┼",
			},
		},
		{
			name: "ragged rows padded with empty cells",
			rows: [][]string{
				{"A", "B", "C"},
				{"1"},
				{"x", "y"},
			},
			contains: []string{"A", "B", "C", "1", "x", "y"},
		},
		{
			name: "cells of varying widths",
			rows: [][]string{
				{"Short", "A much longer header"},
				{"X", "Y"},
			},
			contains: []string{"Short", "A much longer header", "X", "Y"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ty := New()
			result := ty.Table(tc.rows)

			if tc.wantEmpty {
				if result != "" {
					t.Errorf("expected empty string, got %q", result)
				}
				return
			}

			plain := stripANSI(result)
			for _, s := range tc.contains {
				if !strings.Contains(plain, s) {
					t.Errorf("expected output to contain %q, got:\n%s", s, plain)
				}
			}
		})
	}
}

func TestTableMinimalBorder(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()))
	result := stripANSI(ty.Table([][]string{
		{"Name", "Role"},
		{"Alice", "Admin"},
	}))

	// Minimal should NOT have outer corners.
	for _, ch := range []string{"┌", "┐", "└", "┘"} {
		if strings.Contains(result, ch) {
			t.Errorf("minimal border should not contain %q, got:\n%s", ch, result)
		}
	}

	// Should still have header separator.
	if !strings.Contains(result, "─") {
		t.Errorf("minimal border should have header separator ─, got:\n%s", result)
	}

	// Should contain cell content.
	if !strings.Contains(result, "Name") || !strings.Contains(result, "Alice") {
		t.Errorf("minimal border missing cell content, got:\n%s", result)
	}
}

func TestTableCustomCellPad(t *testing.T) {
	tests := []struct {
		name string
		pad  int
	}{
		{"zero padding", 0},
		{"padding of 2", 2},
		{"padding of 3", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ty := New(WithTableCellPad(tc.pad))
			result := stripANSI(ty.Table([][]string{
				{"A"},
				{"B"},
			}))

			if result == "" {
				t.Fatal("expected non-empty table output")
			}

			// Verify padding is reflected: the cell content should be surrounded
			// by the expected number of spaces.
			padStr := strings.Repeat(" ", tc.pad)
			if !strings.Contains(result, padStr+"A"+padStr) {
				t.Errorf("expected cell padding of %d spaces around 'A', got:\n%s", tc.pad, result)
			}
		})
	}
}

func TestTableColumnAlignment(t *testing.T) {
	ty := New()
	rows := [][]string{
		{"A", "Long Header"},
		{"Short", "B"},
	}
	result := stripANSI(ty.Table(rows))
	lines := strings.Split(result, "\n")

	// All lines in a bordered table should have the same visible width.
	if len(lines) < 2 {
		t.Fatalf("expected multiple lines, got %d", len(lines))
	}

	width := len([]rune(lines[0]))
	for i, line := range lines {
		if len([]rune(line)) != width {
			t.Errorf("line %d width %d != expected %d:\n%s", i, len([]rune(line)), width, result)
		}
	}
}

func TestTableOptions(t *testing.T) {
	t.Run("WithTableHeaderStyle", func(t *testing.T) {
		ty := New(WithTableHeaderStyle(lipgloss.NewStyle().Bold(true)))
		result := ty.Table([][]string{{"H"}, {"D"}})
		if result == "" {
			t.Error("expected non-empty output")
		}
	})

	t.Run("WithTableCellStyle", func(t *testing.T) {
		ty := New(WithTableCellStyle(lipgloss.NewStyle().Italic(true)))
		result := ty.Table([][]string{{"H"}, {"D"}})
		if result == "" {
			t.Error("expected non-empty output")
		}
	})

	t.Run("WithTableBorderStyle", func(t *testing.T) {
		ty := New(WithTableBorderStyle(lipgloss.NewStyle().Faint(true)))
		result := ty.Table([][]string{{"H"}, {"D"}})
		if result == "" {
			t.Error("expected non-empty output")
		}
	})

	t.Run("WithTableBorderSet", func(t *testing.T) {
		ty := New(WithTableBorderSet(MinimalBorderSet()))
		result := stripANSI(ty.Table([][]string{{"H"}, {"D"}}))
		if strings.Contains(result, "┌") {
			t.Error("expected minimal border without corners")
		}
	})

	t.Run("WithTableCellPad negative ignored", func(t *testing.T) {
		ty := New(WithTableCellPad(-1))
		if ty.theme.TableCellPad != DefaultTableCellPad {
			t.Errorf("negative pad should be ignored, got %d", ty.theme.TableCellPad)
		}
	})
}

func TestTableHeaderOnly(t *testing.T) {
	ty := New()
	result := stripANSI(ty.Table([][]string{
		{"Name", "Role"},
	}))

	// Should still render the header with borders.
	if !strings.Contains(result, "Name") {
		t.Errorf("header-only table should contain header text, got:\n%s", result)
	}
	if !strings.Contains(result, "┌") {
		t.Errorf("header-only table should have top border, got:\n%s", result)
	}
}

func TestTableSpecialCharacters(t *testing.T) {
	ty := New()
	// Should not panic on special characters.
	result := ty.Table([][]string{
		{"<script>", "a & b"},
		{"\t\n", ""},
	})
	if result == "" {
		t.Error("expected non-empty output for special characters")
	}
}

func TestTableConcurrency(t *testing.T) {
	ty := New()
	done := make(chan struct{})

	for range 10 {
		go func() {
			defer func() { done <- struct{}{} }()
			ty.Table([][]string{
				{"Name", "Role"},
				{"Alice", "Admin"},
			})
		}()
	}

	for range 10 {
		<-done
	}
}

func TestTableWithOptsEmpty(t *testing.T) {
	ty := New()
	if got := ty.TableWithOpts(nil); got != "" {
		t.Errorf("expected empty, got %q", got)
	}
	if got := ty.TableWithOpts([][]string{}); got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}

func TestTableWithOptsMatchesTable(t *testing.T) {
	ty := New()
	rows := [][]string{
		{"Name", "Role"},
		{"Alice", "Admin"},
	}
	want := ty.Table(rows)
	got := ty.TableWithOpts(rows)
	if got != want {
		t.Errorf("TableWithOpts with no opts should match Table\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestTableWithOptsAlignRight(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"X", "Y"},
		{"AB", "C"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithColumnAlign(0, AlignRight)))
	lines := strings.Split(result, "\n")
	if !strings.Contains(lines[0], " X") {
		t.Errorf("expected right-aligned 'X' with leading space, got line: %q", lines[0])
	}
}

func TestTableWithOptsAlignCenter(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"ABCDE"},
		{"X"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithColumnAlign(0, AlignCenter)))
	lines := strings.Split(result, "\n")
	bodyLine := lines[2]
	if !strings.Contains(bodyLine, "  X  ") {
		t.Errorf("expected centered 'X' as '  X  ', got line: %q", bodyLine)
	}
}

func TestTableWithOptsAlignLeftDefault(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"ABC"},
		{"X"},
	}
	withOpt := stripANSI(ty.TableWithOpts(rows, WithColumnAlign(0, AlignLeft)))
	without := stripANSI(ty.Table(rows))
	if withOpt != without {
		t.Errorf("explicit AlignLeft should match default\ngot:\n%s\nwant:\n%s", withOpt, without)
	}
}

func TestTableWithOptsMixedAlignments(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"Left", "Center", "Right"},
		{"A", "B", "C"},
	}
	result := stripANSI(ty.TableWithOpts(rows,
		WithColumnAlign(0, AlignLeft),
		WithColumnAlign(1, AlignCenter),
		WithColumnAlign(2, AlignRight),
	))
	lines := strings.Split(result, "\n")
	bodyLine := lines[2]

	if !strings.Contains(bodyLine, "A   ") {
		t.Errorf("expected left-aligned 'A' with trailing spaces, got: %q", bodyLine)
	}
	if !strings.Contains(bodyLine, "  B   ") {
		t.Errorf("expected centered 'B', got: %q", bodyLine)
	}
	if !strings.Contains(bodyLine, "    C") {
		t.Errorf("expected right-aligned 'C' with leading spaces, got: %q", bodyLine)
	}
}

func TestTableWithOptsColumnAligns(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"AAA", "BBB"},
		{"X", "Y"},
	}
	resultA := stripANSI(ty.TableWithOpts(rows,
		WithColumnAlign(0, AlignRight),
		WithColumnAlign(1, AlignCenter),
	))
	resultB := stripANSI(ty.TableWithOpts(rows,
		WithColumnAligns(AlignRight, AlignCenter),
	))
	if resultA != resultB {
		t.Errorf("WithColumnAligns should match individual WithColumnAlign calls\ngot:\n%s\nwant:\n%s", resultB, resultA)
	}
}

func TestTableWithOptsOutOfRange(t *testing.T) {
	ty := New()
	rows := [][]string{
		{"A", "B"},
		{"C", "D"},
	}
	result := ty.TableWithOpts(rows, WithColumnAlign(99, AlignRight))
	if result == "" {
		t.Error("expected non-empty output")
	}
}

func TestTableWithOptsBorderedAlignment(t *testing.T) {
	ty := New(WithTableCellPad(0))
	rows := [][]string{
		{"Name", "Score"},
		{"A", "100"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithColumnAlign(1, AlignRight)))
	if !strings.Contains(result, "  100") {
		t.Errorf("expected right-aligned '100' with leading spaces, got:\n%s", result)
	}
}

func TestTableWithOptsUniformWidth(t *testing.T) {
	ty := New()
	rows := [][]string{
		{"A", "Long Header"},
		{"Short", "B"},
	}
	result := stripANSI(ty.TableWithOpts(rows,
		WithColumnAlign(0, AlignRight),
		WithColumnAlign(1, AlignCenter),
	))
	lines := strings.Split(result, "\n")
	if len(lines) < 2 {
		t.Fatalf("expected multiple lines, got %d", len(lines))
	}
	width := len([]rune(lines[0]))
	for i, line := range lines {
		if len([]rune(line)) != width {
			t.Errorf("line %d width %d != expected %d:\n%s", i, len([]rune(line)), width, result)
		}
	}
}

func TestTableWithOptsRowSeparatorsBordered(t *testing.T) {
	ty := New(WithTableCellPad(0))
	rows := [][]string{
		{"A", "B"},
		{"C", "D"},
		{"E", "F"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithRowSeparators(true)))
	lines := strings.Split(result, "\n")
	// Bordered: top, header, header-sep, row1, row-sep, row2, bottom = 7 lines.
	if len(lines) != 7 {
		t.Fatalf("expected 7 lines, got %d:\n%s", len(lines), result)
	}
	// Row separator (line index 4) should contain junction characters.
	if !strings.Contains(lines[4], "─") {
		t.Errorf("expected row separator with ─, got: %q", lines[4])
	}
}

func TestTableWithOptsRowSeparatorsMinimal(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"A", "B"},
		{"C", "D"},
		{"E", "F"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithRowSeparators(true)))
	lines := strings.Split(result, "\n")
	// Minimal: header, header-sep, row1, row-sep, row2 = 5 lines.
	if len(lines) != 5 {
		t.Fatalf("expected 5 lines, got %d:\n%s", len(lines), result)
	}
}

func TestTableWithOptsRowSeparatorsDisabled(t *testing.T) {
	ty := New(WithTableCellPad(0))
	rows := [][]string{
		{"A", "B"},
		{"C", "D"},
		{"E", "F"},
	}
	withSep := ty.TableWithOpts(rows, WithRowSeparators(false))
	without := ty.Table(rows)
	if withSep != without {
		t.Errorf("WithRowSeparators(false) should match Table output")
	}
}

// ---------------------------------------------------------------------------
// Striped rows
// ---------------------------------------------------------------------------

func TestTableStripedRows(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"H"},
		{"A"},
		{"B"},
		{"C"},
		{"D"},
	}
	result := ty.TableWithOpts(rows, WithStripedRows(true))
	// Should produce output without panic.
	if result == "" {
		t.Error("expected non-empty output")
	}
	plain := ty.Table(rows)
	if result == plain {
		t.Error("striped output should differ from plain output")
	}
}

func TestTableStripedRowsDisabled(t *testing.T) {
	ty := New(WithTableCellPad(0))
	rows := [][]string{
		{"H"},
		{"A"},
		{"B"},
	}
	withOpt := ty.TableWithOpts(rows, WithStripedRows(false))
	without := ty.Table(rows)
	if withOpt != without {
		t.Error("WithStripedRows(false) should match Table output")
	}
}

// ---------------------------------------------------------------------------
// Caption
// ---------------------------------------------------------------------------

func TestTableCaptionTop(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"A"},
		{"B"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithCaption("My Table")))
	lines := strings.Split(result, "\n")
	if lines[0] != "My Table" {
		t.Errorf("expected first line to be caption, got: %q", lines[0])
	}
}

func TestTableCaptionBottom(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"A"},
		{"B"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithCaptionBottom("Source: test")))
	lines := strings.Split(result, "\n")
	last := lines[len(lines)-1]
	if last != "Source: test" {
		t.Errorf("expected last line to be caption, got: %q", last)
	}
}

func TestTableCaptionEmpty(t *testing.T) {
	ty := New(WithTableCellPad(0))
	rows := [][]string{
		{"A"},
		{"B"},
	}
	withCaption := ty.TableWithOpts(rows, WithCaption(""))
	without := ty.Table(rows)
	if withCaption != without {
		t.Error("empty caption should not change output")
	}
}

// ---------------------------------------------------------------------------
// Footer row
// ---------------------------------------------------------------------------

func TestTableFooterRow(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"Item", "Price"},
		{"Widget", "$10"},
		{"Gadget", "$20"},
		{"Total", "$30"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithFooterRow(true)))
	// Should contain a footer separator and the footer row.
	if !strings.Contains(result, "Total") {
		t.Error("expected footer row with 'Total'")
	}
	// Footer separator should appear (a second ─ line after body).
	lines := strings.Split(result, "\n")
	sepCount := 0
	for _, line := range lines {
		if strings.Contains(line, "─") {
			sepCount++
		}
	}
	// Header sep + footer sep = 2.
	if sepCount != 2 {
		t.Errorf("expected 2 separator lines, got %d", sepCount)
	}
}

func TestTableFooterRowBordered(t *testing.T) {
	ty := New(WithTableCellPad(0))
	rows := [][]string{
		{"H1", "H2"},
		{"A", "B"},
		{"Footer1", "Footer2"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithFooterRow(true)))
	if !strings.Contains(result, "Footer1") {
		t.Error("expected footer row content")
	}
	// All lines should be same width.
	lines := strings.Split(result, "\n")
	width := len([]rune(lines[0]))
	for i, line := range lines {
		if len([]rune(line)) != width {
			t.Errorf("line %d width %d != %d:\n%s", i, len([]rune(line)), width, result)
		}
	}
}

func TestTableFooterRowTooFewRows(t *testing.T) {
	ty := New(WithTableCellPad(0))
	// Only header + 1 body row: footer should not activate.
	rows := [][]string{
		{"H"},
		{"A"},
	}
	withFooter := ty.TableWithOpts(rows, WithFooterRow(true))
	without := ty.Table(rows)
	if withFooter != without {
		t.Error("footer should not activate with <= 2 rows")
	}
}

func TestTableFooterRowDisabled(t *testing.T) {
	ty := New(WithTableCellPad(0))
	rows := [][]string{
		{"H"},
		{"A"},
		{"B"},
	}
	withFooter := ty.TableWithOpts(rows, WithFooterRow(false))
	without := ty.Table(rows)
	if withFooter != without {
		t.Error("WithFooterRow(false) should match Table output")
	}
}

// ---------------------------------------------------------------------------
// Theme-level options for new styles
// ---------------------------------------------------------------------------

func TestTableStyleOptions(t *testing.T) {
	t.Run("WithTableStripedCellStyle", func(t *testing.T) {
		ty := New(WithTableStripedCellStyle(lipgloss.NewStyle().Faint(true)))
		result := ty.TableWithOpts([][]string{{"H"}, {"A"}, {"B"}}, WithStripedRows(true))
		if result == "" {
			t.Error("expected non-empty output")
		}
	})

	t.Run("WithTableFooterStyle", func(t *testing.T) {
		ty := New(WithTableFooterStyle(lipgloss.NewStyle().Bold(true)))
		result := ty.TableWithOpts([][]string{{"H"}, {"A"}, {"F"}}, WithFooterRow(true))
		if result == "" {
			t.Error("expected non-empty output")
		}
	})

	t.Run("WithTableCaptionStyle", func(t *testing.T) {
		ty := New(WithTableCaptionStyle(lipgloss.NewStyle().Italic(true)))
		result := ty.TableWithOpts([][]string{{"H"}, {"A"}}, WithCaption("cap"))
		if result == "" {
			t.Error("expected non-empty output")
		}
	})
}

// ---------------------------------------------------------------------------
// Auto-truncation
// ---------------------------------------------------------------------------

func TestTruncateCell(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxWidth int
		want     string
	}{
		{"no truncation needed", "Hello", 10, "Hello"},
		{"exact fit", "Hello", 5, "Hello"},
		{"truncate with ellipsis", "Hello World", 8, "Hello W…"},
		{"truncate to 1", "Hello", 1, "…"},
		{"zero max disables", "Hello", 0, "Hello"},
		{"negative max disables", "Hello", -1, "Hello"},
		{"empty string", "", 5, ""},
		{"single char truncate", "AB", 1, "…"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := truncateCell(tc.input, tc.maxWidth)
			if got != tc.want {
				t.Errorf("truncateCell(%q, %d) = %q, want %q", tc.input, tc.maxWidth, got, tc.want)
			}
			if tc.maxWidth > 0 && lipgloss.Width(got) > tc.maxWidth {
				t.Errorf("result width %d exceeds max %d", lipgloss.Width(got), tc.maxWidth)
			}
		})
	}
}

func TestTableWithOptsMaxColumnWidth(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"Name", "Description"},
		{"Go", "A statically typed compiled language"},
	}
	result := stripANSI(ty.TableWithOpts(rows, WithMaxColumnWidth(10)))
	// "A statically typed compiled language" should be truncated.
	if strings.Contains(result, "compiled language") {
		t.Error("expected long cell to be truncated")
	}
	if !strings.Contains(result, "…") {
		t.Error("expected ellipsis in truncated cell")
	}
}

func TestTableWithOptsColumnMaxWidth(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"Name", "Description"},
		{"Go", "A compiled language"},
	}
	// Only truncate column 1.
	result := stripANSI(ty.TableWithOpts(rows, WithColumnMaxWidth(1, 8)))
	lines := strings.Split(result, "\n")
	// Header "Description" should be truncated.
	if !strings.Contains(lines[0], "…") {
		t.Errorf("expected header cell truncated, got: %q", lines[0])
	}
	// Column 0 ("Name", "Go") should not be truncated.
	if strings.Contains(lines[0], "Na…") {
		t.Error("column 0 should not be truncated")
	}
}

func TestTableWithOptsColumnMaxWidthOverridesGlobal(t *testing.T) {
	ty := New(WithTableBorderSet(MinimalBorderSet()), WithTableCellPad(0))
	rows := [][]string{
		{"LongName", "VeryLongContent"},
	}
	// Global max=5, but column 1 has max=15 (no truncation for col 1).
	result := stripANSI(ty.TableWithOpts(rows,
		WithMaxColumnWidth(5),
		WithColumnMaxWidth(1, 15),
	))
	if !strings.Contains(result, "VeryLongContent") {
		t.Error("column 1 should not be truncated (per-column override)")
	}
	if !strings.Contains(result, "Long…") {
		t.Errorf("column 0 should be truncated by global max, got: %s", result)
	}
}

func TestTableWithOptsNoTruncation(t *testing.T) {
	ty := New(WithTableCellPad(0))
	rows := [][]string{
		{"A"},
		{"B"},
	}
	// MaxColumnWidth(0) should not change output.
	withOpt := ty.TableWithOpts(rows, WithMaxColumnWidth(0))
	without := ty.Table(rows)
	if withOpt != without {
		t.Error("WithMaxColumnWidth(0) should not change output")
	}
}

func TestAlignCell(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		cellWidth int
		total     int
		align     Alignment
		want      string
	}{
		{"left pad", "X", 1, 5, AlignLeft, "X    "},
		{"right pad", "X", 1, 5, AlignRight, "    X"},
		{"center even", "XX", 2, 6, AlignCenter, "  XX  "},
		{"center odd", "X", 1, 6, AlignCenter, "  X   "},
		{"no gap", "ABCDE", 5, 5, AlignLeft, "ABCDE"},
		{"no gap right", "ABCDE", 5, 5, AlignRight, "ABCDE"},
		{"no gap center", "ABCDE", 5, 5, AlignCenter, "ABCDE"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := alignCell(tc.text, tc.cellWidth, tc.total, tc.align)
			if got != tc.want {
				t.Errorf("alignCell(%q, %d, %d, %d) = %q, want %q",
					tc.text, tc.cellWidth, tc.total, tc.align, got, tc.want)
			}
		})
	}
}
