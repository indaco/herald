package herald

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestWithTheme(t *testing.T) {
	custom := DefaultTheme()
	custom.HRWidth = 80
	custom.BulletChar = "-"

	ty := New(WithTheme(custom))

	if ty.theme.HRWidth != 80 {
		t.Errorf("expected HRWidth 80, got %d", ty.theme.HRWidth)
	}
	if ty.theme.BulletChar != "-" {
		t.Errorf("expected BulletChar %q, got %q", "-", ty.theme.BulletChar)
	}
}

func TestWithHeadingStyles(t *testing.T) {
	style := lipgloss.NewStyle().Bold(true)

	tests := []struct {
		name string
		opt  Option
		get  func(*Typography) lipgloss.Style
	}{
		{"H1", WithH1Style(style), func(ty *Typography) lipgloss.Style { return ty.theme.H1 }},
		{"H2", WithH2Style(style), func(ty *Typography) lipgloss.Style { return ty.theme.H2 }},
		{"H3", WithH3Style(style), func(ty *Typography) lipgloss.Style { return ty.theme.H3 }},
		{"H4", WithH4Style(style), func(ty *Typography) lipgloss.Style { return ty.theme.H4 }},
		{"H5", WithH5Style(style), func(ty *Typography) lipgloss.Style { return ty.theme.H5 }},
		{"H6", WithH6Style(style), func(ty *Typography) lipgloss.Style { return ty.theme.H6 }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ty := New(tc.opt)
			got := tc.get(ty)
			// Verify the style was applied by rendering
			result := got.Render("test")
			if result == "" {
				t.Error("expected non-empty render")
			}
		})
	}
}

func TestWithBlockStyles(t *testing.T) {
	style := lipgloss.NewStyle().Italic(true)

	tests := []struct {
		name   string
		opt    Option
		render func(*Typography) string
	}{
		{"Paragraph", WithParagraphStyle(style), func(ty *Typography) string { return ty.P("test") }},
		{"Blockquote", WithBlockquoteStyle(style), func(ty *Typography) string { return ty.Blockquote("test") }},
		{"CodeInline", WithCodeInlineStyle(style), func(ty *Typography) string { return ty.Code("test") }},
		{"CodeBlock", WithCodeBlockStyle(style), func(ty *Typography) string { return ty.CodeBlock("test", "") }},
		{"HR", WithHRStyle(style), func(ty *Typography) string { return ty.HR() }},
		{"HRLabel", WithHRLabelStyle(style), func(ty *Typography) string { return ty.HRWithLabel("test") }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ty := New(tc.opt)
			result := stripANSI(tc.render(ty))
			if len(result) == 0 {
				t.Errorf("expected non-empty rendered output")
			}
		})
	}
}

func TestWithInlineStyles(t *testing.T) {
	style := lipgloss.NewStyle()

	tests := []struct {
		name   string
		opt    Option
		render func(*Typography) string
	}{
		{"Bold", WithBoldStyle(style), func(ty *Typography) string { return ty.Bold("test") }},
		{"Italic", WithItalicStyle(style), func(ty *Typography) string { return ty.Italic("test") }},
		{"Underline", WithUnderlineStyle(style), func(ty *Typography) string { return ty.Underline("test") }},
		{"Strikethrough", WithStrikethroughStyle(style), func(ty *Typography) string { return ty.Strikethrough("test") }},
		{"Small", WithSmallStyle(style), func(ty *Typography) string { return ty.Small("test") }},
		{"Mark", WithMarkStyle(style), func(ty *Typography) string { return ty.Mark("test") }},
		{"Link", WithLinkStyle(style), func(ty *Typography) string { return ty.Link("test", "url") }},
		{"Kbd", WithKbdStyle(style), func(ty *Typography) string { return ty.Kbd("test") }},
		{"Abbr", WithAbbrStyle(style), func(ty *Typography) string { return ty.Abbr("test", "title") }},
		{"Ins", WithInsStyle(style), func(ty *Typography) string { return ty.Ins("test") }},
		{"Del", WithDelStyle(style), func(ty *Typography) string { return ty.Del("test") }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ty := New(tc.opt)
			result := stripANSI(tc.render(ty))
			if !strings.Contains(result, "test") {
				t.Errorf("expected output to contain %q, got %q", "test", result)
			}
		})
	}
}

func TestWithListStyles(t *testing.T) {
	style := lipgloss.NewStyle()

	ty := New(
		WithListBulletStyle(style),
		WithListItemStyle(style),
	)
	result := stripANSI(ty.UL("item1", "item2"))
	if !strings.Contains(result, "item1") || !strings.Contains(result, "item2") {
		t.Errorf("expected list items in output, got %q", result)
	}
}

func TestWithDLStyles(t *testing.T) {
	style := lipgloss.NewStyle()

	ty := New(
		WithDTStyle(style),
		WithDDStyle(style),
	)
	result := stripANSI(ty.DL([][2]string{{"Term", "Definition"}}))
	if !strings.Contains(result, "Term") || !strings.Contains(result, "Definition") {
		t.Errorf("expected DL content in output, got %q", result)
	}
}

func TestWithAddressStyle(t *testing.T) {
	style := lipgloss.NewStyle().Italic(true)
	ty := New(WithAddressStyle(style))
	result := ty.theme.Address.Render("test")
	if result == "" {
		t.Error("expected non-empty render from Address style")
	}
}

func TestWithAddressCardStyle(t *testing.T) {
	style := lipgloss.NewStyle().Italic(true)
	ty := New(WithAddressCardStyle(style))
	result := ty.theme.AddressCard.Render("test")
	if result == "" {
		t.Error("expected non-empty render from AddressCard style")
	}
}

func TestWithAddressCardBorderStyle(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	ty := New(WithAddressCardBorderStyle(style))
	result := ty.theme.AddressCardBorder.Render("test")
	if result == "" {
		t.Error("expected non-empty render from AddressCardBorder style")
	}
}

func TestWithBadgeStyle(t *testing.T) {
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF00"))
	ty := New(WithBadgeStyle(style))
	result := ty.theme.Badge.Render("test")
	if result == "" {
		t.Error("expected non-empty render from Badge style")
	}
}

func TestWithTagStyle(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	ty := New(WithTagStyle(style))
	result := ty.theme.Tag.Render("test")
	if result == "" {
		t.Error("expected non-empty render from Tag style")
	}
}

func TestWithFootnoteRefStyle(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	ty := New(WithFootnoteRefStyle(style))
	result := ty.theme.FootnoteRef.Render("test")
	if result == "" {
		t.Error("expected non-empty render from FootnoteRef style")
	}
}

func TestWithFootnoteItemStyle(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	ty := New(WithFootnoteItemStyle(style))
	result := ty.theme.FootnoteItem.Render("test")
	if result == "" {
		t.Error("expected non-empty render from FootnoteItem style")
	}
}

func TestWithFootnoteDividerStyle(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF"))
	ty := New(WithFootnoteDividerStyle(style))
	result := ty.theme.FootnoteDivider.Render("test")
	if result == "" {
		t.Error("expected non-empty render from FootnoteDivider style")
	}
}

func TestWithBlockquoteBarStyle(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	ty := New(WithBlockquoteBarStyle(style))
	result := ty.theme.BlockquoteBarStyle.Render("│")
	if result == "" {
		t.Error("expected non-empty render from BlockquoteBarStyle")
	}
}

func TestWithTokenOptions(t *testing.T) {
	t.Run("BulletChar", func(t *testing.T) {
		ty := New(WithBulletChar("*"))
		if ty.theme.BulletChar != "*" {
			t.Errorf("expected %q, got %q", "*", ty.theme.BulletChar)
		}
	})

	t.Run("HRChar", func(t *testing.T) {
		ty := New(WithHRChar("="))
		if ty.theme.HRChar != "=" {
			t.Errorf("expected %q, got %q", "=", ty.theme.HRChar)
		}
	})

	t.Run("HRWidth positive", func(t *testing.T) {
		ty := New(WithHRWidth(60))
		if ty.theme.HRWidth != 60 {
			t.Errorf("expected 60, got %d", ty.theme.HRWidth)
		}
	})

	t.Run("HRWidth zero ignored", func(t *testing.T) {
		ty := New(WithHRWidth(0))
		if ty.theme.HRWidth != 40 {
			t.Errorf("expected default 40, got %d", ty.theme.HRWidth)
		}
	})

	t.Run("HRWidth negative ignored", func(t *testing.T) {
		ty := New(WithHRWidth(-5))
		if ty.theme.HRWidth != 40 {
			t.Errorf("expected default 40, got %d", ty.theme.HRWidth)
		}
	})

	t.Run("BlockquoteBar", func(t *testing.T) {
		ty := New(WithBlockquoteBar("|"))
		if ty.theme.BlockquoteBar != "|" {
			t.Errorf("expected %q, got %q", "|", ty.theme.BlockquoteBar)
		}
	})

	t.Run("InsPrefix", func(t *testing.T) {
		ty := New(WithInsPrefix("++ "))
		if ty.theme.InsPrefix != "++ " {
			t.Errorf("expected %q, got %q", "++ ", ty.theme.InsPrefix)
		}
	})

	t.Run("DelPrefix", func(t *testing.T) {
		ty := New(WithDelPrefix("-- "))
		if ty.theme.DelPrefix != "-- " {
			t.Errorf("expected %q, got %q", "-- ", ty.theme.DelPrefix)
		}
	})

	t.Run("FootnoteDividerChar", func(t *testing.T) {
		ty := New(WithFootnoteDividerChar("="))
		if ty.theme.FootnoteDividerChar != "=" {
			t.Errorf("expected %q, got %q", "=", ty.theme.FootnoteDividerChar)
		}
	})

	t.Run("FootnoteDividerWidth positive", func(t *testing.T) {
		ty := New(WithFootnoteDividerWidth(30))
		if ty.theme.FootnoteDividerWidth != 30 {
			t.Errorf("expected 30, got %d", ty.theme.FootnoteDividerWidth)
		}
	})

	t.Run("FootnoteDividerWidth zero ignored", func(t *testing.T) {
		ty := New(WithFootnoteDividerWidth(0))
		if ty.theme.FootnoteDividerWidth != DefaultFootnoteDividerWidth {
			t.Errorf("expected default %d, got %d", DefaultFootnoteDividerWidth, ty.theme.FootnoteDividerWidth)
		}
	})

	t.Run("FootnoteDividerWidth negative ignored", func(t *testing.T) {
		ty := New(WithFootnoteDividerWidth(-5))
		if ty.theme.FootnoteDividerWidth != DefaultFootnoteDividerWidth {
			t.Errorf("expected default %d, got %d", DefaultFootnoteDividerWidth, ty.theme.FootnoteDividerWidth)
		}
	})
}

func TestWithHeadingDecorationOptions(t *testing.T) {
	t.Run("H1UnderlineChar", func(t *testing.T) {
		ty := New(WithH1UnderlineChar("="))
		if ty.theme.H1UnderlineChar != "=" {
			t.Errorf("expected %q, got %q", "=", ty.theme.H1UnderlineChar)
		}
	})

	t.Run("H2UnderlineChar", func(t *testing.T) {
		ty := New(WithH2UnderlineChar("-"))
		if ty.theme.H2UnderlineChar != "-" {
			t.Errorf("expected %q, got %q", "-", ty.theme.H2UnderlineChar)
		}
	})

	t.Run("H3UnderlineChar", func(t *testing.T) {
		ty := New(WithH3UnderlineChar("."))
		if ty.theme.H3UnderlineChar != "." {
			t.Errorf("expected %q, got %q", ".", ty.theme.H3UnderlineChar)
		}
	})

	t.Run("HeadingBarChar", func(t *testing.T) {
		ty := New(WithHeadingBarChar("|"))
		if ty.theme.HeadingBarChar != "|" {
			t.Errorf("expected %q, got %q", "|", ty.theme.HeadingBarChar)
		}
	})
}

func TestWithListIndent(t *testing.T) {
	t.Run("positive value", func(t *testing.T) {
		ty := New(WithListIndent(4))
		if ty.theme.ListIndent != 4 {
			t.Errorf("expected 4, got %d", ty.theme.ListIndent)
		}
	})

	t.Run("zero ignored", func(t *testing.T) {
		ty := New(WithListIndent(0))
		if ty.theme.ListIndent != DefaultListIndent {
			t.Errorf("expected default %d, got %d", DefaultListIndent, ty.theme.ListIndent)
		}
	})

	t.Run("negative ignored", func(t *testing.T) {
		ty := New(WithListIndent(-1))
		if ty.theme.ListIndent != DefaultListIndent {
			t.Errorf("expected default %d, got %d", DefaultListIndent, ty.theme.ListIndent)
		}
	})
}

func TestWithNestedBulletChars(t *testing.T) {
	t.Run("custom chars", func(t *testing.T) {
		chars := []string{"*", "o", "-"}
		ty := New(WithNestedBulletChars(chars))
		if len(ty.theme.NestedBulletChars) != 3 {
			t.Fatalf("expected 3 chars, got %d", len(ty.theme.NestedBulletChars))
		}
		for i, c := range chars {
			if ty.theme.NestedBulletChars[i] != c {
				t.Errorf("char[%d]: expected %q, got %q", i, c, ty.theme.NestedBulletChars[i])
			}
		}
	})

	t.Run("empty slice ignored", func(t *testing.T) {
		ty := New(WithNestedBulletChars([]string{}))
		if len(ty.theme.NestedBulletChars) != len(DefaultNestedBulletChars()) {
			t.Errorf("expected default chars to be preserved")
		}
	})
}

func TestWithHierarchicalNumbers(t *testing.T) {
	t.Run("enabled", func(t *testing.T) {
		ty := New(WithHierarchicalNumbers(true))
		if !ty.theme.HierarchicalNumbers {
			t.Error("expected HierarchicalNumbers to be true")
		}
	})

	t.Run("disabled by default", func(t *testing.T) {
		ty := New()
		if ty.theme.HierarchicalNumbers {
			t.Error("expected HierarchicalNumbers to be false by default")
		}
	})
}

func TestWithCodeFormatter(t *testing.T) {
	formatter := func(code, language string) string {
		return "[" + language + "]" + code
	}

	t.Run("sets formatter", func(t *testing.T) {
		ty := New(WithCodeFormatter(formatter))
		if ty.theme.CodeFormatter == nil {
			t.Fatal("expected CodeFormatter to be set, got nil")
		}
		got := ty.theme.CodeFormatter("hello", "go")
		if got != "[go]hello" {
			t.Errorf("expected %q, got %q", "[go]hello", got)
		}
	})

	t.Run("nil by default", func(t *testing.T) {
		ty := New()
		if ty.theme.CodeFormatter != nil {
			t.Error("expected CodeFormatter to be nil by default")
		}
	})
}

func TestWithCodeLineNumbers(t *testing.T) {
	t.Run("enables line numbers", func(t *testing.T) {
		ty := New(WithCodeLineNumbers(true))
		if !ty.theme.ShowLineNumbers {
			t.Error("expected ShowLineNumbers to be true")
		}
	})

	t.Run("disabled by default", func(t *testing.T) {
		ty := New()
		if ty.theme.ShowLineNumbers {
			t.Error("expected ShowLineNumbers to be false by default")
		}
	})
}

func TestWithCodeLineNumberStyle(t *testing.T) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	ty := New(WithCodeLineNumberStyle(style))
	result := ty.theme.CodeLineNumber.Render("1")
	if result == "" {
		t.Error("expected non-empty render from CodeLineNumber style")
	}
}

func TestWithCodeLineNumberSep(t *testing.T) {
	ty := New(WithCodeLineNumberSep(":"))
	if ty.theme.CodeLineNumberSep != ":" {
		t.Errorf("expected separator %q, got %q", ":", ty.theme.CodeLineNumberSep)
	}
}

func TestMultipleOptions(t *testing.T) {
	ty := New(
		WithHRWidth(80),
		WithBulletChar("-"),
		WithHRChar("="),
		WithBlockquoteBar("|"),
	)

	if ty.theme.HRWidth != 80 {
		t.Errorf("HRWidth: expected 80, got %d", ty.theme.HRWidth)
	}
	if ty.theme.BulletChar != "-" {
		t.Errorf("BulletChar: expected %q, got %q", "-", ty.theme.BulletChar)
	}
	if ty.theme.HRChar != "=" {
		t.Errorf("HRChar: expected %q, got %q", "=", ty.theme.HRChar)
	}
	if ty.theme.BlockquoteBar != "|" {
		t.Errorf("BlockquoteBar: expected %q, got %q", "|", ty.theme.BlockquoteBar)
	}
}
