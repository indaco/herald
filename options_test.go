package herald

import (
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
		name string
		opt  Option
	}{
		{"Paragraph", WithParagraphStyle(style)},
		{"Blockquote", WithBlockquoteStyle(style)},
		{"CodeInline", WithCodeInlineStyle(style)},
		{"CodeBlock", WithCodeBlockStyle(style)},
		{"HR", WithHRStyle(style)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ty := New(tc.opt)
			_ = ty // option applied without panic
		})
	}
}

func TestWithInlineStyles(t *testing.T) {
	style := lipgloss.NewStyle()

	tests := []struct {
		name string
		opt  Option
	}{
		{"Bold", WithBoldStyle(style)},
		{"Italic", WithItalicStyle(style)},
		{"Underline", WithUnderlineStyle(style)},
		{"Strikethrough", WithStrikethroughStyle(style)},
		{"Small", WithSmallStyle(style)},
		{"Mark", WithMarkStyle(style)},
		{"Link", WithLinkStyle(style)},
		{"Kbd", WithKbdStyle(style)},
		{"Abbr", WithAbbrStyle(style)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ty := New(tc.opt)
			_ = ty
		})
	}
}

func TestWithListStyles(t *testing.T) {
	style := lipgloss.NewStyle()

	ty := New(
		WithListBulletStyle(style),
		WithListItemStyle(style),
	)
	_ = ty
}

func TestWithDLStyles(t *testing.T) {
	style := lipgloss.NewStyle()

	ty := New(
		WithDTStyle(style),
		WithDDStyle(style),
	)
	_ = ty
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
