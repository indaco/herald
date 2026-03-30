package herald

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestDraculaTheme(t *testing.T) {
	assertThemeValid(t, DraculaTheme())
}

func TestCatppuccinTheme(t *testing.T) {
	assertThemeValid(t, CatppuccinTheme())
}

func TestBase16Theme(t *testing.T) {
	assertThemeValid(t, Base16Theme())
}

func TestCharmTheme(t *testing.T) {
	assertThemeValid(t, CharmTheme())
}

// assertThemeValid verifies that a Theme has all style fields populated,
// configurable tokens matching defaults, and a nil CodeFormatter.
func assertThemeValid(t *testing.T, theme Theme) {
	t.Helper()

	t.Run("all style fields are populated", func(t *testing.T) {
		styles := []struct {
			name  string
			style lipgloss.Style
		}{
			{"H1", theme.H1},
			{"H2", theme.H2},
			{"H3", theme.H3},
			{"H4", theme.H4},
			{"H5", theme.H5},
			{"H6", theme.H6},
			{"Paragraph", theme.Paragraph},
			{"Blockquote", theme.Blockquote},
			{"CodeInline", theme.CodeInline},
			{"CodeBlock", theme.CodeBlock},
			{"HR", theme.HR},
			{"ListBullet", theme.ListBullet},
			{"ListItem", theme.ListItem},
			{"Bold", theme.Bold},
			{"Italic", theme.Italic},
			{"Underline", theme.Underline},
			{"Strikethrough", theme.Strikethrough},
			{"Small", theme.Small},
			{"Mark", theme.Mark},
			{"Link", theme.Link},
			{"Kbd", theme.Kbd},
			{"Abbr", theme.Abbr},
			{"Sub", theme.Sub},
			{"Sup", theme.Sup},
			{"DT", theme.DT},
			{"DD", theme.DD},
			{"CodeLineNumber", theme.CodeLineNumber},
		}

		for i := range styles {
			tc := &styles[i]
			t.Run(tc.name, func(t *testing.T) {
				result := tc.style.Render("test")
				if result == "" {
					t.Errorf("%s style rendered empty string", tc.name)
				}
			})
		}
	})

	t.Run("configurable tokens match defaults", func(t *testing.T) {
		if theme.H1UnderlineChar != DefaultH1UnderlineChar {
			t.Errorf("H1UnderlineChar: expected %q, got %q", DefaultH1UnderlineChar, theme.H1UnderlineChar)
		}
		if theme.H2UnderlineChar != DefaultH2UnderlineChar {
			t.Errorf("H2UnderlineChar: expected %q, got %q", DefaultH2UnderlineChar, theme.H2UnderlineChar)
		}
		if theme.H3UnderlineChar != DefaultH3UnderlineChar {
			t.Errorf("H3UnderlineChar: expected %q, got %q", DefaultH3UnderlineChar, theme.H3UnderlineChar)
		}
		if theme.HeadingBarChar != DefaultHeadingBarChar {
			t.Errorf("HeadingBarChar: expected %q, got %q", DefaultHeadingBarChar, theme.HeadingBarChar)
		}
		if theme.BulletChar != DefaultBulletChar {
			t.Errorf("BulletChar: expected %q, got %q", DefaultBulletChar, theme.BulletChar)
		}
		if theme.HRChar != DefaultHRChar {
			t.Errorf("HRChar: expected %q, got %q", DefaultHRChar, theme.HRChar)
		}
		if theme.HRWidth != DefaultHRWidth {
			t.Errorf("HRWidth: expected %d, got %d", DefaultHRWidth, theme.HRWidth)
		}
		if theme.BlockquoteBar != DefaultBlockquoteBar {
			t.Errorf("BlockquoteBar: expected %q, got %q", DefaultBlockquoteBar, theme.BlockquoteBar)
		}
		if theme.CodeLineNumberSep != DefaultCodeLineNumberSep {
			t.Errorf("CodeLineNumberSep: expected %q, got %q", DefaultCodeLineNumberSep, theme.CodeLineNumberSep)
		}
		if theme.ShowLineNumbers {
			t.Error("ShowLineNumbers should be false by default")
		}
		if theme.CodeLineNumberOffset != DefaultCodeLineNumberOffset {
			t.Errorf("CodeLineNumberOffset: expected %d, got %d", DefaultCodeLineNumberOffset, theme.CodeLineNumberOffset)
		}
	})

	t.Run("CodeFormatter is nil", func(t *testing.T) {
		if theme.CodeFormatter != nil {
			t.Error("expected CodeFormatter to be nil")
		}
	})
}

func TestDraculaThemeRenderAllElements(t *testing.T) {
	ty := New(WithTheme(DraculaTheme()))
	assertRenderAllElements(t, ty)
}

func TestCatppuccinThemeRenderAllElements(t *testing.T) {
	ty := New(WithTheme(CatppuccinTheme()))
	assertRenderAllElements(t, ty)
}

func TestBase16ThemeRenderAllElements(t *testing.T) {
	ty := New(WithTheme(Base16Theme()))
	assertRenderAllElements(t, ty)
}

func TestCharmThemeRenderAllElements(t *testing.T) {
	ty := New(WithTheme(CharmTheme()))
	assertRenderAllElements(t, ty)
}

// assertRenderAllElements verifies that rendering every element does not panic
// and produces non-empty output.
func assertRenderAllElements(t *testing.T, ty *Typography) {
	t.Helper()

	tests := []struct {
		name string
		fn   func() string
	}{
		{"H1", func() string { return ty.H1("Heading 1") }},
		{"H2", func() string { return ty.H2("Heading 2") }},
		{"H3", func() string { return ty.H3("Heading 3") }},
		{"H4", func() string { return ty.H4("Heading 4") }},
		{"H5", func() string { return ty.H5("Heading 5") }},
		{"H6", func() string { return ty.H6("Heading 6") }},
		{"P", func() string { return ty.P("A paragraph.") }},
		{"Blockquote", func() string { return ty.Blockquote("A wise quote.") }},
		{"UL", func() string { return ty.UL("Apples", "Bananas") }},
		{"OL", func() string { return ty.OL("First", "Second") }},
		{"Code", func() string { return ty.Code("x := 1") }},
		{"CodeBlock", func() string { return ty.CodeBlock("func main() {}") }},
		{"HR", func() string { return ty.HR() }},
		{"Bold", func() string { return ty.Bold("bold") }},
		{"Italic", func() string { return ty.Italic("italic") }},
		{"Underline", func() string { return ty.Underline("underline") }},
		{"Strikethrough", func() string { return ty.Strikethrough("strike") }},
		{"Small", func() string { return ty.Small("small") }},
		{"Mark", func() string { return ty.Mark("highlight") }},
		{"Link", func() string { return ty.Link("Go", "https://go.dev") }},
		{"Kbd", func() string { return ty.Kbd("Ctrl+C") }},
		{"Abbr", func() string { return ty.Abbr("HTML", "HyperText Markup Language") }},
		{"Sub", func() string { return ty.Sub("2") }},
		{"Sup", func() string { return ty.Sup("2") }},
		{"DT", func() string { return ty.DT("Term") }},
		{"DD", func() string { return ty.DD("Description") }},
		{"DL", func() string { return ty.DL([][2]string{{"Go", "A language"}}) }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn()
			plain := stripANSI(result)
			if strings.TrimSpace(plain) == "" {
				t.Errorf("%s rendered empty", tc.name)
			}
		})
	}
}
