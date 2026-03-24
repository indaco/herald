// Using herald with chroma for syntax-highlighted code blocks.
//
// This example is a separate Go module with its own go.mod to keep
// github.com/alecthomas/chroma out of herald's core dependencies.
//
// Run from the repository root:
//
//	cd examples/200_chroma-syntax-highlighting && go run .
package main

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/alecthomas/chroma/v2/quick"
	"github.com/indaco/herald"
)

// chromaFormatter returns a CodeFormatter that uses chroma for highlighting.
func chromaFormatter(style string) func(code, language string) string {
	return func(code, language string) string {
		var buf strings.Builder
		err := quick.Highlight(&buf, code, language, "terminal256", style)
		if err != nil {
			return code
		}
		return strings.TrimRight(buf.String(), "\n")
	}
}

func main() {
	ty := herald.New(
		herald.WithCodeFormatter(chromaFormatter("catppuccin-mocha")),
		herald.WithCodeBlockStyle(lipgloss.NewStyle().
			Padding(1, 2).
			MarginBottom(1)),
	)

	fmt.Println(ty.H1("Syntax Highlighting Demo"))

	fmt.Println(ty.P("Inline code with highlighting: " + ty.Code("os.Exit(1)", "go")))
	fmt.Println()

	fmt.Println(ty.H3("Go"))
	fmt.Println(ty.CodeBlock(`package main

import "fmt"

func main() {
	items := []string{"hello", "world"}
	for i, item := range items {
		fmt.Printf("%d: %s\n", i, item)
	}
}`, "go"))

	fmt.Println(ty.H3("Python"))
	fmt.Println(ty.CodeBlock(`def fibonacci(n: int) -> list[int]:
    a, b = 0, 1
    result = []
    for _ in range(n):
        result.append(a)
        a, b = b, a + b
    return result

print(fibonacci(10))`, "python"))

	fmt.Println(ty.H3("JSON"))
	fmt.Println(ty.CodeBlock(`{
  "name": "herald",
  "description": "HTML-inspired typography for terminal UIs in Go",
  "version": "0.1.0"
}`, "json"))

	fmt.Println(ty.H3("Without language (no highlighting)"))
	fmt.Println(ty.CodeBlock("Just plain text in a code block, no formatter applied."))
}
