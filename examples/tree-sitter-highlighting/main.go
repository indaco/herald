// Using herald with tree-sitter for syntax-highlighted code blocks.
//
// This example is a separate Go module with its own go.mod to keep
// github.com/smacker/go-tree-sitter out of herald's core dependencies.
//
// Run from the repository root:
//
//	cd examples/tree-sitter-highlighting && go run .
package main

import (
	"context"
	"fmt"
	"sort"

	"charm.land/lipgloss/v2"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"

	"github.com/indaco/herald"
)

// nodeStyle maps tree-sitter node types to lipgloss styles for Go syntax.
var nodeStyle = map[string]lipgloss.Style{
	// Keywords
	"func":        lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"package":     lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"import":      lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"return":      lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"if":          lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"else":        lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"for":         lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"range":       lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"var":         lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"const":       lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"type":        lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"struct":      lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"interface":   lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"defer":       lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"go":          lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"select":      lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"case":        lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"switch":      lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"default":     lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"map":         lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"chan":        lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"break":       lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"continue":    lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"fallthrough": lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),
	"goto":        lipgloss.NewStyle().Foreground(lipgloss.Color("#C4A7E7")),

	// String literals
	"interpreted_string_literal": lipgloss.NewStyle().Foreground(lipgloss.Color("#9CCFD8")),
	"raw_string_literal":         lipgloss.NewStyle().Foreground(lipgloss.Color("#9CCFD8")),
	"rune_literal":               lipgloss.NewStyle().Foreground(lipgloss.Color("#9CCFD8")),

	// Number literals
	"int_literal":   lipgloss.NewStyle().Foreground(lipgloss.Color("#F6C177")),
	"float_literal": lipgloss.NewStyle().Foreground(lipgloss.Color("#F6C177")),

	// Boolean and nil
	"true":  lipgloss.NewStyle().Foreground(lipgloss.Color("#F6C177")),
	"false": lipgloss.NewStyle().Foreground(lipgloss.Color("#F6C177")),
	"nil":   lipgloss.NewStyle().Foreground(lipgloss.Color("#F6C177")),

	// Comments
	"comment": lipgloss.NewStyle().Foreground(lipgloss.Color("#6E6A86")).Italic(true),

	// Types
	"type_identifier": lipgloss.NewStyle().Foreground(lipgloss.Color("#EA9A97")),
}

// coloredRange tracks a styled span within the source code.
type coloredRange struct {
	start uint32
	end   uint32
	style lipgloss.Style
}

// treeSitterFormatter returns a CodeFormatter that uses tree-sitter for
// syntax highlighting. Currently supports Go; other languages fall back
// to returning the code unstyled.
func treeSitterFormatter() func(code, language string) string {
	return func(code, language string) string {
		if language != "go" {
			return code // unsupported language, return as-is
		}

		parser := sitter.NewParser()
		parser.SetLanguage(golang.GetLanguage())

		tree, err := parser.ParseCtx(context.Background(), nil, []byte(code))
		if err != nil {
			return code
		}
		defer tree.Close()

		// Collect styled ranges from the AST.
		var ranges []coloredRange
		collectRanges(tree.RootNode(), []byte(code), &ranges)

		// Sort by start position so we can walk left to right.
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i].start < ranges[j].start
		})

		// Build the highlighted output by interleaving styled and
		// unstyled spans.
		src := []byte(code)
		var result []byte
		var pos uint32

		for _, r := range ranges {
			if r.start < pos {
				continue // overlapping range, skip
			}
			// Append any unstyled gap before this range.
			if r.start > pos {
				result = append(result, src[pos:r.start]...)
			}
			// Append the styled span.
			text := string(src[r.start:r.end])
			result = append(result, []byte(r.style.Render(text))...)
			pos = r.end
		}

		// Append any remaining unstyled tail.
		if int(pos) < len(src) {
			result = append(result, src[pos:]...)
		}

		return string(result)
	}
}

// collectRanges walks the tree-sitter AST and records styled ranges for
// leaf nodes whose type matches a known style.
func collectRanges(node *sitter.Node, src []byte, ranges *[]coloredRange) {
	nodeType := node.Type()

	// Check if this node type has a style. For leaf nodes (no children)
	// we record the range directly. For keyword tokens inside larger
	// nodes we also record them.
	if style, ok := nodeStyle[nodeType]; ok && node.ChildCount() == 0 {
		*ranges = append(*ranges, coloredRange{
			start: node.StartByte(),
			end:   node.EndByte(),
			style: style,
		})
		return
	}

	// Recurse into children.
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		collectRanges(child, src, ranges)
	}
}

func main() {
	ty := herald.New(
		herald.WithCodeFormatter(treeSitterFormatter()),
		herald.WithCodeBlockStyle(lipgloss.NewStyle().
			Padding(1, 2).
			MarginBottom(1)),
	)

	fmt.Println(ty.H1("Tree-sitter Syntax Highlighting Demo"))

	fmt.Println(ty.P("Using go-tree-sitter to parse Go code and apply Rose Pine colors."))

	fmt.Println(ty.H3("Go — with tree-sitter highlighting"))
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
