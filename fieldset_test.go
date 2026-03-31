package herald

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

// ---------------------------------------------------------------------------
// Fieldset
// ---------------------------------------------------------------------------

func TestFieldsetBasic(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Fieldset("Legend", "content"))

	if !strings.Contains(result, "Legend") {
		t.Errorf("expected legend in result, got %q", result)
	}
	if !strings.Contains(result, "content") {
		t.Errorf("expected content in result, got %q", result)
	}
	if !strings.Contains(result, "╭") {
		t.Errorf("expected top-left corner, got %q", result)
	}
	if !strings.Contains(result, "╯") {
		t.Errorf("expected bottom-right corner, got %q", result)
	}
}

func TestFieldsetEmptyLegend(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Fieldset("", "content"))

	if !strings.Contains(result, "╭─") {
		t.Errorf("expected plain top border, got %q", result)
	}
	if !strings.Contains(result, "content") {
		t.Errorf("expected content in result, got %q", result)
	}
}

func TestFieldsetEmptyContent(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Fieldset("Legend", ""))

	if !strings.Contains(result, "Legend") {
		t.Errorf("expected legend in result, got %q", result)
	}
	lines := strings.Split(result, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines (top, content, bottom), got %d", len(lines))
	}
}

func TestFieldsetMultiLine(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Fieldset("Title", "line one\nline two\nline three"))

	if !strings.Contains(result, "line one") {
		t.Errorf("expected first line, got %q", result)
	}
	if !strings.Contains(result, "line three") {
		t.Errorf("expected third line, got %q", result)
	}
	lines := strings.Split(result, "\n")
	// top + 3 content + bottom = 5 lines
	if len(lines) != 5 {
		t.Errorf("expected 5 lines, got %d", len(lines))
	}
}

func TestFieldsetEmptyLegendAndContent(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Fieldset("", ""))
	lines := strings.Split(result, "\n")

	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines, got %d", len(lines))
	}
	if !strings.Contains(lines[0], "╭") {
		t.Errorf("expected top border, got %q", lines[0])
	}
	if !strings.Contains(lines[len(lines)-1], "╰") {
		t.Errorf("expected bottom border, got %q", lines[len(lines)-1])
	}
}

func TestFieldsetUniformWidth(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name    string
		legend  string
		content string
		width   []int
	}{
		{"basic", "Title", "short\nthis is a longer line\nmed", nil},
		{"empty legend", "", "aaa\nbbbb", nil},
		{"explicit width", "X", "hello", []int{30}},
		{"empty content", "Legend", "", nil},
		{"unicode legend", "Bonjour", "content line", nil},
		{"wide content narrow legend", "X", "this is a fairly wide content line", nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.Fieldset(tc.legend, tc.content, tc.width...)
			lines := strings.Split(result, "\n")
			if len(lines) < 3 {
				t.Fatalf("expected at least 3 lines, got %d", len(lines))
			}
			firstWidth := lipgloss.Width(lines[0])
			for i, line := range lines {
				w := lipgloss.Width(line)
				if w != firstWidth {
					t.Errorf("line %d width %d != first line width %d\nline: %q\nfull:\n%s",
						i, w, firstWidth, stripANSI(line), stripANSI(result))
				}
			}
		})
	}
}

func TestFieldsetExplicitWidth(t *testing.T) {
	ty := newTestTypography()

	t.Run("explicit width overrides auto", func(t *testing.T) {
		result := stripANSI(ty.Fieldset("L", "hi", 40))
		lines := strings.Split(result, "\n")
		if lipgloss.Width(lines[0]) != 40 {
			t.Errorf("expected width 40, got %d", lipgloss.Width(lines[0]))
		}
	})

	t.Run("theme width used when no override", func(t *testing.T) {
		custom := DefaultTheme()
		custom.FieldsetWidth = 30
		tyCustom := New(WithTheme(custom))
		result := stripANSI(tyCustom.Fieldset("L", "hi"))
		lines := strings.Split(result, "\n")
		if lipgloss.Width(lines[0]) != 30 {
			t.Errorf("expected width 30, got %d", lipgloss.Width(lines[0]))
		}
	})

	t.Run("zero width means auto-fit", func(t *testing.T) {
		result := stripANSI(ty.Fieldset("L", "hi", 0))
		lines := strings.Split(result, "\n")
		if lipgloss.Width(lines[0]) < 6 {
			t.Errorf("auto-fit should produce reasonable width, got %d", lipgloss.Width(lines[0]))
		}
	})
}

func TestFieldsetVaryingContentWidths(t *testing.T) {
	ty := newTestTypography()
	content := "a\nabcdef\nab"
	result := stripANSI(ty.Fieldset("Box", content))
	lines := strings.Split(result, "\n")

	for i := 1; i < len(lines)-1; i++ {
		if !strings.HasPrefix(lines[i], "│") {
			t.Errorf("content line %d should start with border, got %q", i, lines[i])
		}
		if !strings.HasSuffix(lines[i], "│") {
			t.Errorf("content line %d should end with border, got %q", i, lines[i])
		}
	}
}

func TestFieldsetTopBorderWithLegend(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Fieldset("Test", "content"))
	lines := strings.Split(result, "\n")
	top := lines[0]

	if !strings.HasPrefix(top, "╭─ ") {
		t.Errorf("top border should start with '╭─ ', got %q", top)
	}
	if !strings.Contains(top, "Test") {
		t.Errorf("top border should contain legend, got %q", top)
	}
	if !strings.HasSuffix(top, "╮") {
		t.Errorf("top border should end with '╮', got %q", top)
	}
}

func TestFieldsetTopBorderNoLegend(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Fieldset("", "content"))
	lines := strings.Split(result, "\n")
	top := lines[0]

	if !strings.HasPrefix(top, "╭─") {
		t.Errorf("plain top border should start with '╭─', got %q", top)
	}
	// Should be only dashes between corners (no spaces)
	inner := top[len("╭") : len(top)-len("╮")]
	for _, r := range inner {
		if r != '─' {
			t.Errorf("plain top border should only contain dashes, found %q in %q", string(r), top)
		}
	}
}

func TestFieldsetBottomBorderStructure(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Fieldset("X", "content"))
	lines := strings.Split(result, "\n")
	bottom := lines[len(lines)-1]

	if !strings.HasPrefix(bottom, "╰") {
		t.Errorf("bottom border should start with '╰', got %q", bottom)
	}
	if !strings.HasSuffix(bottom, "╯") {
		t.Errorf("bottom border should end with '╯', got %q", bottom)
	}
}

func TestFieldsetLegendOverflow(t *testing.T) {
	ty := newTestTypography()

	// Legend is wider than the box: triggers rightDashes < 1 guard
	result := stripANSI(ty.Fieldset("VeryLongLegendText", "x", 10))
	if !strings.Contains(result, "VeryLongLegendText") {
		t.Errorf("expected legend in result, got %q", result)
	}
	if !strings.Contains(result, "╭") || !strings.Contains(result, "╯") {
		t.Errorf("expected box corners, got %q", result)
	}
}

func TestFieldsetContentOverflow(t *testing.T) {
	ty := newTestTypography()

	// Content wider than the explicit width: triggers pad < 0 guard
	result := stripANSI(ty.Fieldset("L", "this content is wider than the box", 10))
	if !strings.Contains(result, "this content is wider than the box") {
		t.Errorf("expected content in result, got %q", result)
	}
	if !strings.Contains(result, "│") {
		t.Errorf("expected border chars, got %q", result)
	}
}

func TestFieldsetWithCustomStyles(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	ty := New(
		WithFieldsetStyle(style),
		WithFieldsetBorderStyle(style),
		WithFieldsetLegendStyle(style),
	)

	result := ty.Fieldset("Legend", "content")
	plain := stripANSI(result)
	if !strings.Contains(plain, "Legend") {
		t.Errorf("expected legend in styled result, got %q", plain)
	}
	if !strings.Contains(plain, "content") {
		t.Errorf("expected content in styled result, got %q", plain)
	}
}
