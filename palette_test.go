package herald

import (
	"image/color"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

// testPalette returns a ColorPalette with distinct colors for testing.
func testPalette() ColorPalette {
	return ColorPalette{
		Primary:   lipgloss.Color("#FF0000"),
		Secondary: lipgloss.Color("#00FF00"),
		Tertiary:  lipgloss.Color("#0000FF"),
		Accent:    lipgloss.Color("#FFFF00"),
		Highlight: lipgloss.Color("#FF00FF"),
		Muted:     lipgloss.Color("#888888"),
		Text:      lipgloss.Color("#FFFFFF"),
		Surface:   lipgloss.Color("#333333"),
		Base:      lipgloss.Color("#111111"),
	}
}

func TestThemeFromPalette(t *testing.T) {
	theme := ThemeFromPalette(testPalette())

	t.Run("all style fields are populated", func(t *testing.T) {
		// Every style field should produce a non-empty render for non-empty input.
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
		}

		for _, tc := range styles {
			t.Run(tc.name, func(t *testing.T) {
				result := tc.style.Render("test")
				if result == "" {
					t.Errorf("%s style rendered empty string", tc.name)
				}
			})
		}
	})

	t.Run("configurable tokens match defaults", func(t *testing.T) {
		def := DefaultTheme()

		if theme.H1UnderlineChar != def.H1UnderlineChar {
			t.Errorf("H1UnderlineChar: expected %q, got %q", def.H1UnderlineChar, theme.H1UnderlineChar)
		}
		if theme.H2UnderlineChar != def.H2UnderlineChar {
			t.Errorf("H2UnderlineChar: expected %q, got %q", def.H2UnderlineChar, theme.H2UnderlineChar)
		}
		if theme.H3UnderlineChar != def.H3UnderlineChar {
			t.Errorf("H3UnderlineChar: expected %q, got %q", def.H3UnderlineChar, theme.H3UnderlineChar)
		}
		if theme.HeadingBarChar != def.HeadingBarChar {
			t.Errorf("HeadingBarChar: expected %q, got %q", def.HeadingBarChar, theme.HeadingBarChar)
		}
		if theme.BulletChar != def.BulletChar {
			t.Errorf("BulletChar: expected %q, got %q", def.BulletChar, theme.BulletChar)
		}
		if theme.HRChar != def.HRChar {
			t.Errorf("HRChar: expected %q, got %q", def.HRChar, theme.HRChar)
		}
		if theme.HRWidth != def.HRWidth {
			t.Errorf("HRWidth: expected %d, got %d", def.HRWidth, theme.HRWidth)
		}
		if theme.BlockquoteBar != def.BlockquoteBar {
			t.Errorf("BlockquoteBar: expected %q, got %q", def.BlockquoteBar, theme.BlockquoteBar)
		}
	})

	t.Run("CodeFormatter is nil", func(t *testing.T) {
		if theme.CodeFormatter != nil {
			t.Error("expected CodeFormatter to be nil")
		}
	})
}

func TestThemeFromPaletteColorMapping(t *testing.T) {
	p := testPalette()
	theme := ThemeFromPalette(p)

	tests := []struct {
		name       string
		style      lipgloss.Style
		wantFg     color.Color
		wantBg     color.Color
		checkBg    bool
		wantBold   bool
		checkBold  bool
		wantUline  bool
		checkUline bool
	}{
		{
			name: "H1 uses Primary", style: theme.H1,
			wantFg: p.Primary, wantBold: true, checkBold: true,
		},
		{
			name: "H2 uses Secondary", style: theme.H2,
			wantFg: p.Secondary, wantBold: true, checkBold: true,
		},
		{
			name: "H3 uses Tertiary", style: theme.H3,
			wantFg: p.Tertiary, wantBold: true, checkBold: true,
		},
		{
			name: "H4 uses Accent", style: theme.H4,
			wantFg: p.Accent, wantBold: true, checkBold: true,
		},
		{
			name: "H5 uses Highlight", style: theme.H5,
			wantFg: p.Highlight, wantBold: true, checkBold: true,
		},
		{
			name: "H6 uses Muted", style: theme.H6,
			wantFg: p.Muted, wantBold: true, checkBold: true,
		},
		{
			name: "Blockquote uses Muted", style: theme.Blockquote,
			wantFg: p.Muted,
		},
		{
			name: "CodeInline uses Text on Base", style: theme.CodeInline,
			wantFg: p.Text, wantBg: p.Base, checkBg: true,
		},
		{
			name: "CodeBlock uses Text on Base", style: theme.CodeBlock,
			wantFg: p.Text, wantBg: p.Base, checkBg: true,
		},
		{
			name: "HR uses Muted", style: theme.HR,
			wantFg: p.Muted,
		},
		{
			name: "ListBullet uses Secondary", style: theme.ListBullet,
			wantFg: p.Secondary,
		},
		{
			name: "Mark uses Accent bg and Base fg", style: theme.Mark,
			wantFg: p.Base, wantBg: p.Accent, checkBg: true,
		},
		{
			name: "Link uses Tertiary with underline", style: theme.Link,
			wantFg: p.Tertiary, wantUline: true, checkUline: true,
		},
		{
			name: "Kbd uses Text on Surface", style: theme.Kbd,
			wantFg: p.Text, wantBg: p.Surface, checkBg: true,
			wantBold: true, checkBold: true,
		},
		{
			name: "Abbr uses Highlight with underline", style: theme.Abbr,
			wantFg: p.Highlight, wantUline: true, checkUline: true,
		},
		{
			name: "Sub uses Muted", style: theme.Sub,
			wantFg: p.Muted,
		},
		{
			name: "Sup uses Muted", style: theme.Sup,
			wantFg: p.Muted,
		},
		{
			name: "DT uses Text bold", style: theme.DT,
			wantFg: p.Text, wantBold: true, checkBold: true,
		},
		{
			name: "DD uses Muted", style: theme.DD,
			wantFg: p.Muted,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.style.GetForeground(); got != tc.wantFg {
				t.Errorf("foreground: expected %v, got %v", tc.wantFg, got)
			}
			if tc.checkBg {
				if got := tc.style.GetBackground(); got != tc.wantBg {
					t.Errorf("background: expected %v, got %v", tc.wantBg, got)
				}
			}
			if tc.checkBold {
				if got := tc.style.GetBold(); got != tc.wantBold {
					t.Errorf("bold: expected %v, got %v", tc.wantBold, got)
				}
			}
			if tc.checkUline {
				if got := tc.style.GetUnderline(); got != tc.wantUline {
					t.Errorf("underline: expected %v, got %v", tc.wantUline, got)
				}
			}
		})
	}
}

func TestWithPalette(t *testing.T) {
	p := testPalette()
	ty := New(WithPalette(p))

	t.Run("sets theme from palette", func(t *testing.T) {
		theme := ty.Theme()
		if theme.HRWidth != DefaultHRWidth {
			t.Errorf("expected default HRWidth %d, got %d", DefaultHRWidth, theme.HRWidth)
		}
		if got := theme.H1.GetForeground(); got != p.Primary {
			t.Errorf("H1 foreground: expected %v, got %v", p.Primary, got)
		}
	})

	t.Run("can be combined with other options", func(t *testing.T) {
		ty := New(WithPalette(p), WithHRWidth(80), WithBulletChar("-"))
		theme := ty.Theme()
		if theme.HRWidth != 80 {
			t.Errorf("expected HRWidth 80, got %d", theme.HRWidth)
		}
		if theme.BulletChar != "-" {
			t.Errorf("expected BulletChar %q, got %q", "-", theme.BulletChar)
		}
		// Palette colors should still be applied.
		if got := theme.H2.GetForeground(); got != p.Secondary {
			t.Errorf("H2 foreground after combined options: expected %v, got %v", p.Secondary, got)
		}
	})
}

func TestWithPaletteRenderAllElements(t *testing.T) {
	ty := New(WithPalette(testPalette()))

	// Rendering every element should not panic.
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
