// Basic usage of herald with the default Rose Pine theme.
// Run: go run ./examples/basic/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	// Headings
	fmt.Println(ty.H1("Heading 1 — Double Underline"))
	fmt.Println(ty.H2("Heading 2 — Single Underline"))
	fmt.Println(ty.H3("Heading 3 — Dotted Underline"))
	fmt.Println(ty.H4("Heading 4 — Bar Prefix"))
	fmt.Println(ty.H5("Heading 5 — Bar Prefix"))
	fmt.Println(ty.H6("Heading 6 — Bar Prefix"))

	// Block elements
	fmt.Println(ty.P("This is a paragraph. It wraps text with the paragraph style."))
	fmt.Println(ty.Blockquote("A wise person once said something profound.\nAnd then said more on a second line."))
	fmt.Println()
	fmt.Println(ty.Code("fmt.Println()"))
	fmt.Println()
	fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"Hello, World!\")\n}"))
	fmt.Println(ty.HR())
	fmt.Println()

	// Lists
	fmt.Println(ty.H3("Unordered List"))
	fmt.Println(ty.UL("Apples", "Bananas", "Cherries"))
	fmt.Println()

	fmt.Println(ty.H3("Ordered List"))
	fmt.Println(ty.OL("First item", "Second item", "Third item"))
	fmt.Println()

	// Nested lists
	fmt.Println(ty.H3("Nested Unordered List"))
	fmt.Println(ty.NestUL(
		herald.Item("Fruits"),
		herald.ItemWithChildren("Vegetables",
			herald.Item("Carrots"),
			herald.Item("Peas"),
		),
		herald.ItemWithOLChildren("Ranked Desserts",
			herald.Item("Ice cream"),
			herald.Item("Cake"),
		),
	))
	fmt.Println()

	fmt.Println(ty.H3("Nested Ordered List"))
	fmt.Println(ty.NestOL(
		herald.Item("Introduction"),
		herald.ItemWithChildren("Main Topics",
			herald.Item("Architecture"),
			herald.Item("Design"),
		),
		herald.Item("Conclusion"),
	))
	fmt.Println()

	// Inline styles
	fmt.Println(ty.H3("Inline Styles"))
	fmt.Println(ty.Bold("Bold text"))
	fmt.Println(ty.Italic("Italic text"))
	fmt.Println(ty.Underline("Underlined text"))
	fmt.Println(ty.Strikethrough("Strikethrough text"))
	fmt.Println(ty.Small("Small/faint text"))
	fmt.Println(ty.Mark("Highlighted text"))
	fmt.Println(ty.Kbd("Ctrl") + " + " + ty.Kbd("C"))
	fmt.Println(ty.Sub("subscript") + " and " + ty.Sup("superscript"))
	fmt.Println()

	// Links & Abbreviations
	fmt.Println(ty.H3("Links & Abbreviations"))
	fmt.Println(ty.Link("https://go.dev"))
	fmt.Println(ty.Link("Go website", "https://go.dev"))
	fmt.Println(ty.Abbr("HTML"))
	fmt.Println(ty.Abbr("CSS", "Cascading Style Sheets"))
	fmt.Println()

	// Definition list
	fmt.Println(ty.H3("Definition List"))
	fmt.Println(ty.DL([][2]string{
		{"Go", "A statically typed, compiled language"},
		{"Rust", "A systems programming language"},
		{"Python", "A dynamic, interpreted language"},
	}))
	fmt.Println()

	fmt.Println(ty.DT("Manual Term"))
	fmt.Println(ty.DD("Manual description using DT/DD directly"))
}
