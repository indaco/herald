package herald

import (
	"regexp"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

// stripANSI removes all ANSI escape sequences from a string so that
// plain-text assertions work regardless of styling.
var ansiRe = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRe.ReplaceAllString(s, "")
}

func newTestTypography() *Typography {
	return New()
}

func TestNew(t *testing.T) {
	ty := New()
	if ty == nil {
		t.Fatal("New() returned nil")
	}
}

func TestNewWithOptions(t *testing.T) {
	ty := New(WithHRWidth(60), WithBulletChar("-"))
	if ty.theme.HRWidth != 60 {
		t.Errorf("expected HRWidth 60, got %d", ty.theme.HRWidth)
	}
}

func TestThemeAccessor(t *testing.T) {
	ty := New(WithHRWidth(99))
	theme := ty.Theme()
	if theme.HRWidth != 99 {
		t.Errorf("Theme() HRWidth: expected 99, got %d", theme.HRWidth)
	}
}

// ---------------------------------------------------------------------------
// Headings
// ---------------------------------------------------------------------------

func TestHeadings(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name string
		fn   func(string) string
		text string
	}{
		{"H1", ty.H1, "Heading 1"},
		{"H2", ty.H2, "Heading 2"},
		{"H3", ty.H3, "Heading 3"},
		{"H4", ty.H4, "Heading 4"},
		{"H5", ty.H5, "Heading 5"},
		{"H6", ty.H6, "Heading 6"},
		{"H1 empty", ty.H1, ""},
		{"H1 special chars", ty.H1, "<script>alert('xss')</script>"},
		{"H1 unicode", ty.H1, "Bonjour le monde"},
		{"H1 long", ty.H1, strings.Repeat("A", 500)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn(tc.text)
			if tc.text != "" && !strings.Contains(stripANSI(result), tc.text) {
				t.Errorf("expected result to contain %q, got %q", tc.text, stripANSI(result))
			}
		})
	}
}

func TestHeadingUnderlines(t *testing.T) {
	ty := newTestTypography()

	t.Run("H1 has double-line underline", func(t *testing.T) {
		result := stripANSI(ty.H1("Hello"))
		if !strings.Contains(result, strings.Repeat("═", 5)) {
			t.Errorf("H1 should contain ═ underline, got %q", result)
		}
	})

	t.Run("H2 has single-line underline", func(t *testing.T) {
		result := stripANSI(ty.H2("Hello"))
		if !strings.Contains(result, strings.Repeat("─", 5)) {
			t.Errorf("H2 should contain ─ underline, got %q", result)
		}
	})

	t.Run("H3 has dotted underline", func(t *testing.T) {
		result := stripANSI(ty.H3("Hello"))
		if !strings.Contains(result, strings.Repeat("·", 5)) {
			t.Errorf("H3 should contain · underline, got %q", result)
		}
	})

	t.Run("H4 has bar prefix", func(t *testing.T) {
		result := stripANSI(ty.H4("Hello"))
		if !strings.Contains(result, "▎ Hello") {
			t.Errorf("H4 should have bar prefix, got %q", result)
		}
	})

	t.Run("H5 has bar prefix", func(t *testing.T) {
		result := stripANSI(ty.H5("Hello"))
		if !strings.Contains(result, "▎ Hello") {
			t.Errorf("H5 should have bar prefix, got %q", result)
		}
	})

	t.Run("H6 has bar prefix", func(t *testing.T) {
		result := stripANSI(ty.H6("Hello"))
		if !strings.Contains(result, "▎ Hello") {
			t.Errorf("H6 should have bar prefix, got %q", result)
		}
	})
}

func TestCustomHeadingDecorations(t *testing.T) {
	ty := New(
		WithH1UnderlineChar("="),
		WithH2UnderlineChar("-"),
		WithH3UnderlineChar("."),
		WithHeadingBarChar("|"),
	)

	t.Run("custom H1 underline", func(t *testing.T) {
		result := stripANSI(ty.H1("Hi"))
		if !strings.Contains(result, "==") {
			t.Errorf("expected custom H1 underline '==', got %q", result)
		}
	})

	t.Run("custom H2 underline", func(t *testing.T) {
		result := stripANSI(ty.H2("Hi"))
		if !strings.Contains(result, "--") {
			t.Errorf("expected custom H2 underline '--', got %q", result)
		}
	})

	t.Run("custom H3 underline", func(t *testing.T) {
		result := stripANSI(ty.H3("Hi"))
		if !strings.Contains(result, "..") {
			t.Errorf("expected custom H3 underline '..', got %q", result)
		}
	})

	t.Run("custom bar prefix", func(t *testing.T) {
		result := stripANSI(ty.H4("Hi"))
		if !strings.Contains(result, "| Hi") {
			t.Errorf("expected custom bar prefix '| Hi', got %q", result)
		}
	})
}

// ---------------------------------------------------------------------------
// Paragraph
// ---------------------------------------------------------------------------

func TestP(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name string
		text string
	}{
		{"normal", "Hello, world."},
		{"empty", ""},
		{"multiline", "Line one.\nLine two."},
		{"special chars", "a < b && c > d"},
		{"long", strings.Repeat("word ", 200)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.P(tc.text)
			if tc.text != "" && !strings.Contains(stripANSI(result), tc.text) {
				t.Errorf("P(%q) missing text in output", tc.text)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Blockquote
// ---------------------------------------------------------------------------

func TestBlockquote(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		text     string
		contains string
	}{
		{"simple", "A wise quote.", ty.theme.BlockquoteBar},
		{"multiline", "Line 1\nLine 2", ty.theme.BlockquoteBar},
		{"empty", "", ty.theme.BlockquoteBar},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.Blockquote(tc.text)
			if !strings.Contains(stripANSI(result), tc.contains) {
				t.Errorf("Blockquote should contain %q, got %q", tc.contains, stripANSI(result))
			}
		})
	}
}

func TestBlockquoteMultilineHasBars(t *testing.T) {
	ty := New(WithBlockquoteBar("|"))
	result := ty.Blockquote("Line 1\nLine 2\nLine 3")
	// Each line should have the bar
	if strings.Count(stripANSI(result), "|") < 3 {
		t.Errorf("expected at least 3 bars in multiline blockquote, got %q", stripANSI(result))
	}
}

// ---------------------------------------------------------------------------
// Lists
// ---------------------------------------------------------------------------

func TestUL(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name  string
		items []string
	}{
		{"three items", []string{"Apples", "Bananas", "Cherries"}},
		{"single item", []string{"Only one"}},
		{"empty list", nil},
		{"empty strings", []string{"", "", ""}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.UL(tc.items...)
			if len(tc.items) == 0 {
				if result != "" {
					t.Errorf("UL with no items should be empty, got %q", result)
				}
				return
			}
			for _, item := range tc.items {
				if item != "" && !strings.Contains(stripANSI(result), item) {
					t.Errorf("UL missing item %q in %q", item, stripANSI(result))
				}
			}
		})
	}
}

func TestULCustomBullet(t *testing.T) {
	ty := New(WithBulletChar("-"))
	result := ty.UL("Item")
	if !strings.Contains(stripANSI(result), "-") {
		t.Errorf("expected custom bullet '-' in %q", stripANSI(result))
	}
}

func TestOL(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name  string
		items []string
	}{
		{"three items", []string{"First", "Second", "Third"}},
		{"single item", []string{"Only"}},
		{"empty list", nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.OL(tc.items...)
			if len(tc.items) == 0 {
				if result != "" {
					t.Errorf("OL with no items should be empty, got %q", result)
				}
				return
			}
			// Check numbering
			if !strings.Contains(stripANSI(result), "1.") {
				t.Errorf("OL should contain '1.' in %q", stripANSI(result))
			}
			for _, item := range tc.items {
				if item != "" && !strings.Contains(stripANSI(result), item) {
					t.Errorf("OL missing item %q in %q", item, stripANSI(result))
				}
			}
		})
	}
}

func TestOLNumbering(t *testing.T) {
	ty := newTestTypography()
	result := ty.OL("A", "B", "C")
	for _, n := range []string{"1.", "2.", "3."} {
		if !strings.Contains(stripANSI(result), n) {
			t.Errorf("OL missing number %q in %q", n, stripANSI(result))
		}
	}
}

// ---------------------------------------------------------------------------
// Code
// ---------------------------------------------------------------------------

func TestCode(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name string
		text string
	}{
		{"simple", "fmt.Println()"},
		{"empty", ""},
		{"special", "x := map[string]int{}"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.Code(tc.text)
			if tc.text != "" && !strings.Contains(stripANSI(result), tc.text) {
				t.Errorf("Code(%q) missing text in %q", tc.text, stripANSI(result))
			}
		})
	}
}

func TestCodeWithLanguage(t *testing.T) {
	formatter := func(code, language string) string {
		return "«" + language + ":" + code + "»"
	}

	tests := []struct {
		name      string
		formatter func(code, language string) string
		lang      []string
		text      string
		contains  string
		excludes  string
	}{
		{
			name:      "no formatter no lang",
			formatter: nil,
			lang:      nil,
			text:      "x := 1",
			contains:  "x := 1",
		},
		{
			name:      "formatter but no lang",
			formatter: formatter,
			lang:      nil,
			text:      "x := 1",
			contains:  "x := 1",
			excludes:  "«",
		},
		{
			name:      "formatter with empty lang",
			formatter: formatter,
			lang:      []string{""},
			text:      "x := 1",
			contains:  "x := 1",
			excludes:  "«",
		},
		{
			name:      "formatter with lang",
			formatter: formatter,
			lang:      []string{"go"},
			text:      "x := 1",
			contains:  "«go:x := 1»",
		},
		{
			name:      "nil formatter with lang",
			formatter: nil,
			lang:      []string{"go"},
			text:      "x := 1",
			contains:  "x := 1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var opts []Option
			if tc.formatter != nil {
				opts = append(opts, WithCodeFormatter(tc.formatter))
			}
			ty := New(opts...)
			result := stripANSI(ty.Code(tc.text, tc.lang...))
			if !strings.Contains(result, tc.contains) {
				t.Errorf("Code: expected %q in %q", tc.contains, result)
			}
			if tc.excludes != "" && strings.Contains(result, tc.excludes) {
				t.Errorf("Code: did not expect %q in %q", tc.excludes, result)
			}
		})
	}
}

func TestCodeBlock(t *testing.T) {
	ty := newTestTypography()

	code := "func main() {\n\tfmt.Println(\"hello\")\n}"
	result := ty.CodeBlock(code)
	if !strings.Contains(stripANSI(result), "func main()") {
		t.Errorf("CodeBlock should contain source code, got %q", stripANSI(result))
	}
}

func TestCodeBlockWithLanguage(t *testing.T) {
	formatter := func(code, language string) string {
		return "«" + language + ":" + code + "»"
	}

	tests := []struct {
		name      string
		formatter func(code, language string) string
		lang      []string
		text      string
		contains  string
		excludes  string
	}{
		{
			name:      "no formatter no lang",
			formatter: nil,
			lang:      nil,
			text:      "fmt.Println()",
			contains:  "fmt.Println()",
		},
		{
			name:      "formatter but no lang",
			formatter: formatter,
			lang:      nil,
			text:      "fmt.Println()",
			contains:  "fmt.Println()",
			excludes:  "«",
		},
		{
			name:      "formatter with lang",
			formatter: formatter,
			lang:      []string{"go"},
			text:      "fmt.Println()",
			contains:  "«go:fmt.Println()»",
		},
		{
			name:      "nil formatter with lang",
			formatter: nil,
			lang:      []string{"go"},
			text:      "fmt.Println()",
			contains:  "fmt.Println()",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var opts []Option
			if tc.formatter != nil {
				opts = append(opts, WithCodeFormatter(tc.formatter))
			}
			ty := New(opts...)
			result := stripANSI(ty.CodeBlock(tc.text, tc.lang...))
			if !strings.Contains(result, tc.contains) {
				t.Errorf("CodeBlock: expected %q in %q", tc.contains, result)
			}
			if tc.excludes != "" && strings.Contains(result, tc.excludes) {
				t.Errorf("CodeBlock: did not expect %q in %q", tc.excludes, result)
			}
		})
	}
}

func TestCodeBlockLineNumbers(t *testing.T) {
	t.Run("disabled by default", func(t *testing.T) {
		ty := New()
		result := stripANSI(ty.CodeBlock("line one\nline two"))
		if strings.Contains(result, "1│") || strings.Contains(result, "1"+DefaultCodeLineNumberSep) {
			t.Errorf("line numbers should be off by default, got %q", result)
		}
	})

	t.Run("enabled shows line numbers", func(t *testing.T) {
		ty := New(WithCodeLineNumbers(true))
		result := stripANSI(ty.CodeBlock("aaa\nbbb\nccc"))
		if !strings.Contains(result, "1"+DefaultCodeLineNumberSep+" aaa") {
			t.Errorf("expected line 1 with separator, got %q", result)
		}
		if !strings.Contains(result, "3"+DefaultCodeLineNumberSep+" ccc") {
			t.Errorf("expected line 3 with separator, got %q", result)
		}
	})

	t.Run("multi-digit line numbers are right-aligned", func(t *testing.T) {
		lines := make([]string, 12)
		for i := range lines {
			lines[i] = "x"
		}
		ty := New(WithCodeLineNumbers(true))
		result := stripANSI(ty.CodeBlock(strings.Join(lines, "\n")))
		// Single-digit lines should be padded: " 1│"
		if !strings.Contains(result, " 1"+DefaultCodeLineNumberSep) {
			t.Errorf("expected padded single-digit line number, got %q", result)
		}
		if !strings.Contains(result, "12"+DefaultCodeLineNumberSep) {
			t.Errorf("expected line 12, got %q", result)
		}
	})

	t.Run("custom separator", func(t *testing.T) {
		ty := New(WithCodeLineNumbers(true), WithCodeLineNumberSep(":"))
		result := stripANSI(ty.CodeBlock("hello"))
		if !strings.Contains(result, "1: hello") {
			t.Errorf("expected custom separator ':', got %q", result)
		}
	})

	t.Run("with formatter", func(t *testing.T) {
		formatter := func(code, lang string) string {
			return "[" + code + "]"
		}
		ty := New(WithCodeLineNumbers(true), WithCodeFormatter(formatter))
		result := stripANSI(ty.CodeBlock("x := 1", "go"))
		if !strings.Contains(result, "1"+DefaultCodeLineNumberSep+" [x := 1]") {
			t.Errorf("line numbers should wrap formatted content, got %q", result)
		}
	})

	t.Run("single line", func(t *testing.T) {
		ty := New(WithCodeLineNumbers(true))
		result := stripANSI(ty.CodeBlock("only one line"))
		if !strings.Contains(result, "1"+DefaultCodeLineNumberSep+" only one line") {
			t.Errorf("expected single line number, got %q", result)
		}
	})

}

func TestCodeBlockLineNumberOffset(t *testing.T) {
	t.Run("custom offset", func(t *testing.T) {
		ty := New(WithCodeLineNumbers(true), WithCodeLineNumberOffset(42))
		result := stripANSI(ty.CodeBlock("aaa\nbbb\nccc"))
		if !strings.Contains(result, "42"+DefaultCodeLineNumberSep+" aaa") {
			t.Errorf("expected line 42, got %q", result)
		}
		if !strings.Contains(result, "44"+DefaultCodeLineNumberSep+" ccc") {
			t.Errorf("expected line 44, got %q", result)
		}
	})

	t.Run("offset widens gutter", func(t *testing.T) {
		ty := New(WithCodeLineNumbers(true), WithCodeLineNumberOffset(99))
		result := stripANSI(ty.CodeBlock("x\ny\nz"))
		// Lines 99, 100, 101 -> width 3 -> "99" should be padded to " 99"
		if !strings.Contains(result, " 99"+DefaultCodeLineNumberSep) {
			t.Errorf("expected padded line 99, got %q", result)
		}
		if !strings.Contains(result, "101"+DefaultCodeLineNumberSep) {
			t.Errorf("expected line 101, got %q", result)
		}
	})

	t.Run("offset ignored when line numbers disabled", func(t *testing.T) {
		ty := New(WithCodeLineNumberOffset(10))
		result := stripANSI(ty.CodeBlock("hello"))
		if strings.Contains(result, "10") {
			t.Errorf("offset should have no effect when line numbers disabled, got %q", result)
		}
	})
}

// ---------------------------------------------------------------------------
// HR
// ---------------------------------------------------------------------------

func TestHR(t *testing.T) {
	ty := newTestTypography()
	result := ty.HR()
	if result == "" {
		t.Error("HR should not be empty")
	}
}

func TestHRCustomWidth(t *testing.T) {
	ty := New(WithHRWidth(10), WithHRChar("-"))
	result := ty.HR()
	if !strings.Contains(stripANSI(result), "----------") {
		t.Errorf("HR with width 10 and char '-' should contain 10 dashes, got %q", stripANSI(result))
	}
}

func TestHRWithLabel(t *testing.T) {
	tests := []struct {
		name     string
		label    string
		hrWidth  int
		hrChar   string
		wantHR   bool // true if we expect HR chars in output
		wantText string
	}{
		{
			name:     "basic label",
			label:    "Section",
			hrWidth:  40,
			hrChar:   "-",
			wantHR:   true,
			wantText: "Section",
		},
		{
			name:     "empty label falls back to HR",
			label:    "",
			hrWidth:  40,
			hrChar:   "-",
			wantHR:   true,
			wantText: "",
		},
		{
			name:     "label longer than width",
			label:    "This is a very long label that exceeds the HR width",
			hrWidth:  10,
			hrChar:   "-",
			wantHR:   false,
			wantText: "This is a very long label that exceeds the HR width",
		},
		{
			name:     "HR chars on both sides",
			label:    "Mid",
			hrWidth:  20,
			hrChar:   "=",
			wantHR:   true,
			wantText: "Mid",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ty := New(WithHRWidth(tc.hrWidth), WithHRChar(tc.hrChar))
			result := stripANSI(ty.HRWithLabel(tc.label))

			if tc.wantText != "" && !strings.Contains(result, tc.wantText) {
				t.Errorf("expected result to contain %q, got %q", tc.wantText, result)
			}

			if tc.wantHR && !strings.Contains(result, tc.hrChar) {
				t.Errorf("expected HR char %q in result, got %q", tc.hrChar, result)
			}
		})
	}

	t.Run("empty label matches HR output", func(t *testing.T) {
		ty := New(WithHRWidth(20), WithHRChar("-"))
		hrResult := stripANSI(ty.HR())
		labelResult := stripANSI(ty.HRWithLabel(""))
		if hrResult != labelResult {
			t.Errorf("empty label should match HR(), got HR=%q, HRWithLabel=%q", hrResult, labelResult)
		}
	})

	t.Run("HR chars appear on both sides of label", func(t *testing.T) {
		ty := New(WithHRWidth(20), WithHRChar("-"))
		result := stripANSI(ty.HRWithLabel("X"))
		// Split around the label
		parts := strings.SplitN(result, "X", 2)
		if len(parts) != 2 {
			t.Fatalf("expected label 'X' in result, got %q", result)
		}
		if !strings.Contains(parts[0], "-") {
			t.Errorf("expected HR chars before label, got %q", parts[0])
		}
		if !strings.Contains(parts[1], "-") {
			t.Errorf("expected HR chars after label, got %q", parts[1])
		}
	})
}

// ---------------------------------------------------------------------------
// Inline styles
// ---------------------------------------------------------------------------

func TestInlineStyles(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name string
		fn   func(string) string
		text string
	}{
		{"Bold", ty.Bold, "bold text"},
		{"Bold empty", ty.Bold, ""},
		{"Italic", ty.Italic, "italic text"},
		{"Underline", ty.Underline, "underlined"},
		{"Strikethrough", ty.Strikethrough, "removed"},
		{"Small", ty.Small, "fine print"},
		{"Mark", ty.Mark, "highlighted"},
		{"Kbd", ty.Kbd, "Ctrl+C"},
		{"Kbd empty", ty.Kbd, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn(tc.text)
			if tc.text != "" && !strings.Contains(stripANSI(result), tc.text) {
				t.Errorf("%s(%q) missing text in output %q", tc.name, tc.text, stripANSI(result))
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Link
// ---------------------------------------------------------------------------

func TestLink(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		label    string
		url      []string
		contains string
	}{
		{"label only", "Click here", nil, "Click here"},
		{"label and url", "Go", []string{"https://go.dev"}, "Go"},
		{"same label and url", "https://go.dev", []string{"https://go.dev"}, "https://go.dev"},
		{"empty url", "Link", []string{""}, "Link"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.Link(tc.label, tc.url...)
			if !strings.Contains(stripANSI(result), tc.contains) {
				t.Errorf("Link: expected %q in %q", tc.contains, stripANSI(result))
			}
		})
	}
}

func TestLinkWithURL(t *testing.T) {
	ty := newTestTypography()
	result := ty.Link("Go website", "https://go.dev")
	if !strings.Contains(stripANSI(result), "https://go.dev") {
		t.Errorf("Link with separate URL should contain the URL, got %q", stripANSI(result))
	}
}

// ---------------------------------------------------------------------------
// Abbr
// ---------------------------------------------------------------------------

func TestAbbr(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		abbr     string
		desc     []string
		contains string
	}{
		{"no desc", "HTML", nil, "HTML"},
		{"with desc", "CSS", []string{"Cascading Style Sheets"}, "Cascading Style Sheets"},
		{"empty desc", "JS", []string{""}, "JS"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.Abbr(tc.abbr, tc.desc...)
			if !strings.Contains(stripANSI(result), tc.contains) {
				t.Errorf("Abbr: expected %q in %q", tc.contains, stripANSI(result))
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Sub / Sup
// ---------------------------------------------------------------------------

func TestSubSup(t *testing.T) {
	ty := newTestTypography()

	t.Run("Sub", func(t *testing.T) {
		result := ty.Sub("2")
		if !strings.Contains(stripANSI(result), "_2") {
			t.Errorf("Sub should contain '_2', got %q", result)
		}
	})

	t.Run("Sup", func(t *testing.T) {
		result := ty.Sup("2")
		if !strings.Contains(stripANSI(result), "^2") {
			t.Errorf("Sup should contain '^2', got %q", result)
		}
	})

	t.Run("Sub empty", func(t *testing.T) {
		result := ty.Sub("")
		if !strings.Contains(stripANSI(result), "_") {
			t.Errorf("Sub('') should contain '_', got %q", result)
		}
	})

	t.Run("Sup empty", func(t *testing.T) {
		result := ty.Sup("")
		if !strings.Contains(stripANSI(result), "^") {
			t.Errorf("Sup('') should contain '^', got %q", result)
		}
	})
}

// ---------------------------------------------------------------------------
// Ins / Del
// ---------------------------------------------------------------------------

func TestInsDel(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		fn       func(string) string
		text     string
		contains string
	}{
		{"Ins basic", ty.Ins, "added line", "+added line"},
		{"Del basic", ty.Del, "removed line", "-removed line"},
		{"Ins empty", ty.Ins, "", "+"},
		{"Del empty", ty.Del, "", "-"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := stripANSI(tc.fn(tc.text))
			if !strings.Contains(result, tc.contains) {
				t.Errorf("expected %q in %q", tc.contains, result)
			}
		})
	}
}

func TestInsDelCustomPrefix(t *testing.T) {
	ty := New(WithInsPrefix("++ "), WithDelPrefix("-- "))

	t.Run("custom ins prefix", func(t *testing.T) {
		result := stripANSI(ty.Ins("new"))
		if !strings.Contains(result, "++ new") {
			t.Errorf("expected '++ new' in %q", result)
		}
	})

	t.Run("custom del prefix", func(t *testing.T) {
		result := stripANSI(ty.Del("old"))
		if !strings.Contains(result, "-- old") {
			t.Errorf("expected '-- old' in %q", result)
		}
	})
}

// ---------------------------------------------------------------------------
// Definition List
// ---------------------------------------------------------------------------

func TestDL(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name  string
		pairs [][2]string
		empty bool
	}{
		{
			"two pairs",
			[][2]string{{"Go", "A programming language"}, {"Rust", "Another language"}},
			false,
		},
		{
			"single pair",
			[][2]string{{"Term", "Definition"}},
			false,
		},
		{
			"empty",
			nil,
			true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ty.DL(tc.pairs)
			if tc.empty {
				if result != "" {
					t.Errorf("DL with no pairs should be empty, got %q", result)
				}
				return
			}
			for _, pair := range tc.pairs {
				if !strings.Contains(stripANSI(result), pair[0]) {
					t.Errorf("DL missing term %q in %q", pair[0], result)
				}
				if !strings.Contains(stripANSI(result), pair[1]) {
					t.Errorf("DL missing description %q in %q", pair[1], result)
				}
			}
		})
	}
}

func TestDTDD(t *testing.T) {
	ty := newTestTypography()

	t.Run("DT", func(t *testing.T) {
		result := ty.DT("Term")
		if !strings.Contains(stripANSI(result), "Term") {
			t.Errorf("DT should contain 'Term', got %q", result)
		}
	})

	t.Run("DD", func(t *testing.T) {
		result := ty.DD("Description")
		if !strings.Contains(stripANSI(result), "Description") {
			t.Errorf("DD should contain 'Description', got %q", result)
		}
	})
}

// ---------------------------------------------------------------------------
// Address
// ---------------------------------------------------------------------------

func TestAddress(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		text     string
		contains string
	}{
		{"basic", "Jane Doe", "Jane Doe"},
		{"multi-line", "Jane Doe\njane@example.com\nSan Francisco, CA", "jane@example.com"},
		{"empty", "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := stripANSI(ty.Address(tc.text))
			if tc.contains != "" && !strings.Contains(result, tc.contains) {
				t.Errorf("Address: expected %q in %q", tc.contains, result)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// AddressCard
// ---------------------------------------------------------------------------

func TestAddressCard(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		text     string
		contains string
	}{
		{"basic", "Jane Doe", "Jane Doe"},
		{"multi-line", "Jane Doe\njane@example.com\nSan Francisco, CA", "jane@example.com"},
		{"empty", "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := stripANSI(ty.AddressCard(tc.text))
			if tc.contains != "" && !strings.Contains(result, tc.contains) {
				t.Errorf("AddressCard: expected %q in %q", tc.contains, result)
			}
		})
	}

	t.Run("has rounded border characters", func(t *testing.T) {
		result := stripANSI(ty.AddressCard("Test"))
		if !strings.Contains(result, "\u256d") && !strings.Contains(result, "\u2570") {
			t.Errorf("AddressCard should contain rounded border chars, got %q", result)
		}
	})
}

// ---------------------------------------------------------------------------
// Badge
// ---------------------------------------------------------------------------

func TestBadge(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		text     string
		contains string
	}{
		{"basic", "SUCCESS", "SUCCESS"},
		{"another label", "BETA", "BETA"},
		{"empty", "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := stripANSI(ty.Badge(tc.text))
			if tc.contains != "" && !strings.Contains(result, tc.contains) {
				t.Errorf("Badge: expected %q in %q", tc.contains, result)
			}
		})
	}
}

func TestBadgeWithStyle(t *testing.T) {
	ty := newTestTypography()
	custom := lipgloss.NewStyle().Bold(true)

	t.Run("one-off style override", func(t *testing.T) {
		result := stripANSI(ty.BadgeWithStyle("OK", custom))
		if !strings.Contains(result, "OK") {
			t.Errorf("BadgeWithStyle: expected %q in %q", "OK", result)
		}
	})

	t.Run("empty text", func(t *testing.T) {
		result := ty.BadgeWithStyle("", custom)
		// Should not panic; result may be empty or contain only ANSI codes.
		_ = result
	})
}

// ---------------------------------------------------------------------------
// Tag
// ---------------------------------------------------------------------------

func TestTag(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name     string
		text     string
		contains string
	}{
		{"basic", "v2.0", "v2.0"},
		{"another label", "go", "go"},
		{"empty", "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := stripANSI(ty.Tag(tc.text))
			if tc.contains != "" && !strings.Contains(result, tc.contains) {
				t.Errorf("Tag: expected %q in %q", tc.contains, result)
			}
		})
	}
}

func TestTagWithStyle(t *testing.T) {
	ty := newTestTypography()
	custom := lipgloss.NewStyle().Italic(true)

	t.Run("one-off style override", func(t *testing.T) {
		result := stripANSI(ty.TagWithStyle("go", custom))
		if !strings.Contains(result, "go") {
			t.Errorf("TagWithStyle: expected %q in %q", "go", result)
		}
	})

	t.Run("empty text", func(t *testing.T) {
		result := ty.TagWithStyle("", custom)
		// Should not panic; result may be empty or contain only ANSI codes.
		_ = result
	})
}

// ---------------------------------------------------------------------------
// Semantic Badges
// ---------------------------------------------------------------------------

func TestSemanticBadges(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name string
		fn   func(string) string
		text string
	}{
		{"SuccessBadge", ty.SuccessBadge, "running"},
		{"WarningBadge", ty.WarningBadge, "expiring"},
		{"ErrorBadge", ty.ErrorBadge, "failed"},
		{"InfoBadge", ty.InfoBadge, "pending"},
		{"SuccessBadge empty", ty.SuccessBadge, ""},
		{"ErrorBadge special", ty.ErrorBadge, "<error>"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn(tc.text)
			if tc.text != "" && !strings.Contains(stripANSI(result), tc.text) {
				t.Errorf("%s(%q) missing text in output %q", tc.name, tc.text, stripANSI(result))
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Semantic Tags
// ---------------------------------------------------------------------------

func TestSemanticTags(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name string
		fn   func(string) string
		text string
	}{
		{"SuccessTag", ty.SuccessTag, "healthy"},
		{"WarningTag", ty.WarningTag, "degraded"},
		{"ErrorTag", ty.ErrorTag, "critical"},
		{"InfoTag", ty.InfoTag, "maintenance"},
		{"SuccessTag empty", ty.SuccessTag, ""},
		{"InfoTag unicode", ty.InfoTag, "Bonjour"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn(tc.text)
			if tc.text != "" && !strings.Contains(stripANSI(result), tc.text) {
				t.Errorf("%s(%q) missing text in output %q", tc.name, tc.text, stripANSI(result))
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Edge cases
// ---------------------------------------------------------------------------

func TestSpecialCharacters(t *testing.T) {
	ty := newTestTypography()
	special := "\t\n\r\x00 \u00e9\u00e8\u00ea"

	// Should not panic on any method
	ty.H1(special)
	ty.P(special)
	ty.Bold(special)
	ty.Code(special)
	ty.CodeBlock(special)
	ty.Blockquote(special)
}

func TestVeryLongText(t *testing.T) {
	ty := newTestTypography()
	long := strings.Repeat("abcdefghij ", 10000)

	result := ty.P(long)
	if !strings.Contains(stripANSI(result), "abcdefghij") {
		t.Error("very long text should still be present")
	}
}

// ---------------------------------------------------------------------------
// Key-value pairs
// ---------------------------------------------------------------------------

func TestKV(t *testing.T) {
	ty := newTestTypography()

	t.Run("basic pair", func(t *testing.T) {
		got := stripANSI(ty.KV("Name", "Alice"))
		if got != "Name: Alice" {
			t.Errorf("KV basic pair = %q, want %q", got, "Name: Alice")
		}
	})

	t.Run("empty key", func(t *testing.T) {
		got := stripANSI(ty.KV("", "value"))
		if got != ": value" {
			t.Errorf("KV empty key = %q, want %q", got, ": value")
		}
	})

	t.Run("empty value", func(t *testing.T) {
		got := stripANSI(ty.KV("Key", ""))
		if got != "Key: " {
			t.Errorf("KV empty value = %q, want %q", got, "Key: ")
		}
	})

	t.Run("custom separator", func(t *testing.T) {
		ty2 := New(WithKVSeparator(" ->"))
		got := stripANSI(ty2.KV("Name", "Alice"))
		if got != "Name -> Alice" {
			t.Errorf("KV custom sep = %q, want %q", got, "Name -> Alice")
		}
	})

	t.Run("custom styles", func(t *testing.T) {
		ty2 := New(
			WithKVKeyStyle(lipgloss.NewStyle().Bold(true)),
			WithKVValueStyle(lipgloss.NewStyle().Italic(true)),
		)
		got := ty2.KV("Name", "Alice")
		if got == "" {
			t.Error("KV custom styles returned empty string")
		}
	})
}

func TestKVGroup(t *testing.T) {
	ty := newTestTypography()

	t.Run("empty input", func(t *testing.T) {
		got := ty.KVGroup(nil)
		if got != "" {
			t.Errorf("KVGroup empty = %q, want empty", got)
		}
	})

	t.Run("single pair", func(t *testing.T) {
		got := stripANSI(ty.KVGroup([][2]string{{"Name", "Alice"}}))
		if got != "Name: Alice" {
			t.Errorf("KVGroup single = %q, want %q", got, "Name: Alice")
		}
	})

	t.Run("multiple pairs aligned", func(t *testing.T) {
		pairs := [][2]string{
			{"Name", "Alice"},
			{"Age", "30"},
			{"Location", "Wonderland"},
		}
		got := stripANSI(ty.KVGroup(pairs))
		lines := strings.Split(got, "\n")
		if len(lines) != 3 {
			t.Fatalf("KVGroup lines = %d, want 3", len(lines))
		}
		// All colons should be at the same position (after "Location" length = 8).
		for i, line := range lines {
			idx := strings.Index(line, ":")
			if idx != 8 {
				t.Errorf("line %d colon at %d, want 8: %q", i, idx, line)
			}
		}
	})

	t.Run("empty key in group", func(t *testing.T) {
		pairs := [][2]string{
			{"Name", "Alice"},
			{"", "orphan"},
		}
		got := stripANSI(ty.KVGroup(pairs))
		if !strings.Contains(got, "orphan") {
			t.Errorf("KVGroup should contain orphan value, got %q", got)
		}
	})

	t.Run("custom separator in group", func(t *testing.T) {
		ty2 := New(WithKVSeparator(" ="))
		pairs := [][2]string{
			{"host", "localhost"},
			{"port", "8080"},
		}
		got := stripANSI(ty2.KVGroup(pairs))
		if !strings.Contains(got, "=") {
			t.Errorf("KVGroup custom sep should contain '=', got %q", got)
		}
	})
}

// ---------------------------------------------------------------------------
// Compose
// ---------------------------------------------------------------------------

func TestCompose(t *testing.T) {
	ty := newTestTypography()

	t.Run("joins blocks with double newline", func(t *testing.T) {
		result := ty.Compose("aaa", "bbb", "ccc")
		if result != "aaa\n\nbbb\n\nccc" {
			t.Errorf("Compose = %q, want %q", result, "aaa\n\nbbb\n\nccc")
		}
	})

	t.Run("skips empty blocks", func(t *testing.T) {
		result := ty.Compose("aaa", "", "bbb", "", "")
		if result != "aaa\n\nbbb" {
			t.Errorf("Compose with empties = %q, want %q", result, "aaa\n\nbbb")
		}
	})

	t.Run("single block", func(t *testing.T) {
		result := ty.Compose("only")
		if result != "only" {
			t.Errorf("Compose single = %q, want %q", result, "only")
		}
	})

	t.Run("no blocks", func(t *testing.T) {
		result := ty.Compose()
		if result != "" {
			t.Errorf("Compose empty = %q, want empty", result)
		}
	})

	t.Run("all empty blocks", func(t *testing.T) {
		result := ty.Compose("", "", "")
		if result != "" {
			t.Errorf("Compose all empty = %q, want empty", result)
		}
	})

	t.Run("strips trailing newlines", func(t *testing.T) {
		result := ty.Compose("aaa\n\n", "bbb\n")
		if result != "aaa\n\nbbb" {
			t.Errorf("Compose trailing newlines = %q, want %q", result, "aaa\n\nbbb")
		}
	})

	t.Run("whitespace-only blocks skipped", func(t *testing.T) {
		result := ty.Compose("aaa", "\n\n", "bbb")
		if result != "aaa\n\nbbb" {
			t.Errorf("Compose whitespace-only = %q, want %q", result, "aaa\n\nbbb")
		}
	})

	t.Run("with rendered elements", func(t *testing.T) {
		result := ty.Compose(
			ty.H1("Title"),
			ty.P("Body text"),
			ty.UL("one", "two"),
		)
		plain := stripANSI(result)
		if !strings.Contains(plain, "Title") {
			t.Error("Compose: missing heading")
		}
		if !strings.Contains(plain, "Body text") {
			t.Error("Compose: missing paragraph")
		}
		if !strings.Contains(plain, "one") {
			t.Error("Compose: missing list item")
		}
		if strings.Count(result, "\n\n") < 2 {
			t.Error("Compose: expected at least 2 double-newline separators")
		}
	})
}

// ---------------------------------------------------------------------------
// Concurrency safety (for -race detector)
// ---------------------------------------------------------------------------

func TestConcurrentAccess(t *testing.T) {
	ty := newTestTypography()
	done := make(chan struct{})

	for range 10 {
		go func() {
			defer func() { done <- struct{}{} }()
			ty.H1("concurrent heading")
			ty.P("concurrent paragraph")
			ty.UL("a", "b", "c")
			ty.OL("x", "y", "z")
			ty.Bold("bold")
			ty.HR()
			ty.HRWithLabel("Section")
			ty.Blockquote("quote")
			ty.Code("code")
			ty.CodeBlock("block")
			ty.Link("label", "url")
			ty.Abbr("abbr", "description")
			ty.Sub("sub")
			ty.Sup("sup")
			ty.Ins("added")
			ty.Del("removed")
			ty.DL([][2]string{{"term", "def"}})
			ty.Address("Jane Doe\njane@example.com")
			ty.AddressCard("Jane Doe\njane@example.com")
			ty.Badge("SUCCESS")
			ty.BadgeWithStyle("BETA", lipgloss.NewStyle().Bold(true))
			ty.Tag("v2.0")
			ty.TagWithStyle("go", lipgloss.NewStyle().Italic(true))
			ty.FootnoteRef(1)
			ty.FootnoteSection([]string{"note one", "note two"})
			ty.KV("key", "value")
			ty.KVGroup([][2]string{{"a", "1"}, {"bb", "2"}})
			ty.SuccessBadge("ok")
			ty.WarningBadge("warn")
			ty.ErrorBadge("err")
			ty.InfoBadge("info")
			ty.SuccessTag("ok")
			ty.WarningTag("warn")
			ty.ErrorTag("err")
			ty.InfoTag("info")
			ty.Compose(
				ty.H1("heading"),
				ty.P("paragraph"),
				ty.UL("a", "b"),
			)
		}()
	}

	for range 10 {
		<-done
	}
}
