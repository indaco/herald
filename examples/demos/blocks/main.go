// Block-level elements: paragraph, blockquote, code block, HR, and definition list.
// Run: go run ./examples/demos/blocks/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	fmt.Println(ty.P("This is a paragraph. It wraps text with the paragraph style."))
	fmt.Println(ty.Blockquote("A wise person once said something profound.\nAnd then said more on a second line."))
	fmt.Println()

	fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"Hello, World!\")\n}"))

	// Code block with line numbers
	tyLN := herald.New(herald.WithCodeLineNumbers(true))
	fmt.Println(tyLN.CodeBlock("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"))

	// Code block with line number offset (e.g. snippet starting at line 42)
	tyOffset := herald.New(herald.WithCodeLineNumbers(true), herald.WithCodeLineNumberOffset(42))
	fmt.Println(tyOffset.CodeBlock("func greet(name string) string {\n\treturn \"Hello, \" + name\n}"))
	fmt.Println()

	fmt.Println(ty.HR())
	fmt.Println(ty.HRWithLabel("Section"))
	fmt.Println()

	fmt.Println(ty.DL([][2]string{
		{"Go", "A statically typed, compiled language"},
		{"Rust", "A systems programming language"},
		{"Python", "A dynamic, interpreted language"},
	}))
	fmt.Println()

	fmt.Println(ty.DT("Manual Term"))
	fmt.Println(ty.DD("Manual description using DT/DD directly"))
	fmt.Println()

	// Key-Value
	fmt.Println(ty.KV("Name", "Alice"))
	fmt.Println()
	fmt.Println(ty.KVGroup([][2]string{
		{"Name", "Alice"},
		{"Role", "Admin"},
		{"Status", "Active"},
	}))
	fmt.Println()

	// KVGroupWithOpts
	fmt.Println(ty.KVGroupWithOpts([][2]string{
		{ty.Var("--output") + " " + ty.Small("string"), "Output destination"},
		{ty.Var("--verbose"), "Enable verbose output"},
	}, herald.WithKVGroupSeparator(""), herald.WithKVRawKeys(true), herald.WithKVRawValues(true), herald.WithKVIndent(2)))
	fmt.Println()

	fmt.Println(ty.Address("Jane Doe\njane@example.com\nSan Francisco, CA"))
	fmt.Println()
	fmt.Println(ty.AddressCard("Jane Doe\njane@example.com\nSan Francisco, CA"))
	fmt.Println()

	// Fieldset
	fmt.Println(ty.Fieldset("Server Config", "Host:  localhost\nPort:  8080\nTLS:   enabled"))
	fmt.Println()

	// Fieldset without legend
	fmt.Println(ty.Fieldset("", "Plain bordered box\nwithout a legend"))
	fmt.Println()

	// Figure
	fmt.Println(ty.Figure(
		ty.CodeBlock("SELECT * FROM users WHERE active = true"),
		"Figure 1: Active users query",
	))
	fmt.Println()
	fmt.Println(ty.FigureTop(
		ty.Table([][]string{
			{"Name", "Role"},
			{"Alice", "Admin"},
			{"Bob", "Editor"},
		}),
		"Table 1: User roles",
	))
	fmt.Println()

	// Footnote section
	fmt.Println(ty.P("Herald supports rich typography" + ty.FootnoteRef(1) + " with multiple elements" + ty.FootnoteRef(2)))
	fmt.Println(ty.FootnoteSection([]string{
		"Built on lipgloss v2",
		"Headings, lists, alerts, and more",
	}))
}
