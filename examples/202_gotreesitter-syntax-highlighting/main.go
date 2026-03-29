// Using herald with gotreesitter for syntax-highlighted code blocks.
//
// This example is a separate Go module with its own go.mod to keep
// github.com/odvcencio/gotreesitter out of herald's core dependencies.
//
// Run from the repository root:
//
//	cd examples/202_gotreesitter-syntax-highlighting && go run .
package main

import (
	"fmt"
	"sort"

	"charm.land/lipgloss/v2"
	"github.com/odvcencio/gotreesitter"
	"github.com/odvcencio/gotreesitter/grammars"

	"github.com/indaco/herald"
)

// captureStyle maps tree-sitter capture names to lipgloss styles (Rose Pine).
var captureStyle = map[string]lipgloss.Style{
	"keyword":            lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"function":           lipgloss.NewStyle().Foreground(lipgloss.Color("#EB6F92")),
	"function.builtin":   lipgloss.NewStyle().Foreground(lipgloss.Color("#EB6F92")),
	"string":             lipgloss.NewStyle().Foreground(lipgloss.Color("#9CCFD8")),
	"number":             lipgloss.NewStyle().Foreground(lipgloss.Color("#F6C177")),
	"constant.builtin":   lipgloss.NewStyle().Foreground(lipgloss.Color("#F6C177")),
	"comment":            lipgloss.NewStyle().Foreground(lipgloss.Color("#6E6A86")).Italic(true),
	"type":               lipgloss.NewStyle().Foreground(lipgloss.Color("#EA9A97")),
	"type.builtin":       lipgloss.NewStyle().Foreground(lipgloss.Color("#EA9A97")),
	"variable":           lipgloss.NewStyle().Foreground(lipgloss.Color("#E0DEF4")),
	"variable.parameter": lipgloss.NewStyle().Foreground(lipgloss.Color("#E0DEF4")),
	"operator":           lipgloss.NewStyle().Foreground(lipgloss.Color("#908CAA")),
	"punctuation":        lipgloss.NewStyle().Foreground(lipgloss.Color("#908CAA")),
}

// gotreesitterFormatter returns a CodeFormatter that uses gotreesitter for
// syntax highlighting. Currently supports Go; other languages fall back
// to returning the code unstyled.
func gotreesitterFormatter() func(code, language string) string {
	return func(code, language string) string {
		if language != "go" {
			return code
		}

		lang := grammars.GoLanguage()

		entry := grammars.DetectLanguageByName("go")
		if entry == nil || entry.HighlightQuery == "" {
			return code
		}

		hl, err := gotreesitter.NewHighlighter(lang, entry.HighlightQuery)
		if err != nil {
			return code
		}

		ranges := hl.Highlight([]byte(code))

		// Sort by start position so we can walk left to right.
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i].StartByte < ranges[j].StartByte
		})

		// Build the highlighted output by interleaving styled and
		// unstyled spans.
		src := []byte(code)
		var result []byte
		var pos uint32

		for _, r := range ranges {
			if r.StartByte < pos {
				continue // overlapping range, skip
			}

			style, ok := captureStyle[r.Capture]
			if !ok {
				// Try a parent capture name (e.g. "function" from "function.call").
				for base, s := range captureStyle {
					if len(r.Capture) > len(base) && r.Capture[:len(base)] == base {
						style = s
						ok = true
						break
					}
				}
			}

			if !ok {
				// No style for this capture, emit as plain text.
				continue
			}

			// Append any unstyled gap before this range.
			if r.StartByte > pos {
				result = append(result, src[pos:r.StartByte]...)
			}
			// Append the styled span.
			text := string(src[r.StartByte:r.EndByte])
			result = append(result, []byte(style.Render(text))...)
			pos = r.EndByte
		}

		// Append any remaining unstyled tail.
		if int(pos) < len(src) {
			result = append(result, src[pos:]...)
		}

		return string(result)
	}
}

func main() {
	ty := herald.New(
		herald.WithCodeFormatter(gotreesitterFormatter()),
		herald.WithCodeBlockStyle(lipgloss.NewStyle().
			Padding(1, 2).
			MarginBottom(1)),
	)

	fmt.Println(ty.H1("gotreesitter Syntax Highlighting Demo"))

	fmt.Println(ty.P("Using gotreesitter (pure Go, no CGo) for AST-based syntax highlighting."))

	fmt.Println(ty.H3("Go - with gotreesitter highlighting"))
	fmt.Println(ty.CodeBlock(`package main

import "fmt"

// fibonacci returns the first n Fibonacci numbers.
func fibonacci(n int) []int {
	result := make([]int, 0, n)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		result = append(result, a)
		a, b = b, a+b
	}
	return result
}

func main() {
	nums := fibonacci(10)
	for _, v := range nums {
		fmt.Println(v)
	}
}`, "go"))

	fmt.Println(ty.H3("Without language (no highlighting)"))
	fmt.Println(ty.CodeBlock("Just plain text in a code block, no formatter applied."))
}
