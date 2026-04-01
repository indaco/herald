// herald with the default Rose Pine theme.
// Run: go run ./examples/000_default-theme/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	// Headings
	fmt.Println(ty.H1("Heading 1 - Double Underline"))
	fmt.Println(ty.H2("Heading 2 - Single Underline"))
	fmt.Println(ty.H3("Heading 3 - Dotted Underline"))
	fmt.Println(ty.H4("Heading 4 - Bar Prefix"))
	fmt.Println(ty.H5("Heading 5 - Bar Prefix"))
	fmt.Println(ty.H6("Heading 6 - Bar Prefix"))

	// Block elements
	fmt.Println(ty.P("This is a paragraph. It wraps text with the paragraph style."))
	fmt.Println(ty.Blockquote("A wise person once said something profound.\nAnd then said more on a second line."))
	fmt.Println()
	fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"Hello, World!\")\n}"))
	fmt.Println(ty.HR())
	fmt.Println(ty.HRWithLabel("Section"))
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
	fmt.Println(ty.Code("fmt.Println()"))
	fmt.Println(ty.Bold("Bold text"))
	fmt.Println(ty.Italic("Italic text"))
	fmt.Println(ty.Underline("Underlined text"))
	fmt.Println(ty.Strikethrough("Strikethrough text"))
	fmt.Println(ty.Small("Small/faint text"))
	fmt.Println(ty.Mark("Highlighted text"))
	fmt.Println(ty.Kbd("Ctrl") + " + " + ty.Kbd("C"))
	fmt.Println(ty.Sub("subscript") + " and " + ty.Sup("superscript"))
	fmt.Println(ty.Ins("added line"))
	fmt.Println(ty.Del("removed line"))
	fmt.Println()

	// Quotations & citations
	fmt.Println(ty.H3("Quotations & Citations"))
	fmt.Println(ty.Q("To be, or not to be"))
	fmt.Println(ty.Cite("The Go Programming Language"))
	fmt.Println("Output: " + ty.Samp("Hello, World!"))
	fmt.Println("Set " + ty.Var("PORT") + " to configure the server")
	fmt.Println()

	// Links & Abbreviations
	fmt.Println(ty.H3("Links & Abbreviations"))
	fmt.Println(ty.Link("https://go.dev"))
	fmt.Println(ty.Link("Go website", "https://go.dev"))
	fmt.Println(ty.Abbr("HTML"))
	fmt.Println(ty.Abbr("CSS", "Cascading Style Sheets"))
	fmt.Println()

	// Alerts
	fmt.Println(ty.H3("Alerts"))
	fmt.Println(ty.Note("Useful information that users should know, even when skimming content."))
	fmt.Println()
	fmt.Println(ty.Tip("Helpful advice for doing things better or more easily."))
	fmt.Println()
	fmt.Println(ty.Important("Key information users need to know to achieve their goal."))
	fmt.Println()
	fmt.Println(ty.Warning("Urgent info that needs immediate user attention to avoid problems."))
	fmt.Println()
	fmt.Println(ty.Caution("Advises about risks or negative outcomes of certain actions."))
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
	fmt.Println()

	// Key-Value
	fmt.Println(ty.H3("Key-Value"))
	fmt.Println(ty.KV("Name", "Alice"))
	fmt.Println()
	fmt.Println(ty.KVGroup([][2]string{
		{"Name", "Alice"},
		{"Role", "Admin"},
		{"Status", "Active"},
	}))
	fmt.Println()

	// KVGroupWithOpts: no separator, pre-styled keys, indented
	fmt.Println(ty.H3("KVGroupWithOpts"))
	fmt.Println(ty.KVGroupWithOpts([][2]string{
		{ty.Var("--output") + " " + ty.Small("string"), "Output destination " + ty.Small("(default: stdout)")},
		{ty.Var("--verbose"), "Enable verbose output"},
		{ty.Var("--port") + " " + ty.Small("int"), "Port number " + ty.Small("(default: 8080)")},
	}, herald.WithKVGroupSeparator(""), herald.WithKVRawKeys(true), herald.WithKVRawValues(true), herald.WithKVIndent(2)))
	fmt.Println()

	// Address
	fmt.Println(ty.H3("Address"))
	fmt.Println(ty.Address("Jane Doe\njane@example.com\nSan Francisco, CA"))
	fmt.Println()
	fmt.Println(ty.AddressCard("Jane Doe\njane@example.com\nSan Francisco, CA"))
	fmt.Println()

	// Badge / Tag
	fmt.Println(ty.H3("Badge / Tag"))
	fmt.Println(ty.Badge("SUCCESS") + " " + ty.Badge("BETA") + " " + ty.Tag("v2.0") + " " + ty.Tag("go"))
	fmt.Println()

	// Fieldset
	fmt.Println(ty.H3("Fieldset"))
	fmt.Println(ty.Fieldset("Server Config", "Host:  localhost\nPort:  8080\nTLS:   enabled"))
	fmt.Println()
	fmt.Println(ty.Fieldset("", "Plain bordered box without a legend"))
	fmt.Println()

	// Figure
	fmt.Println(ty.H3("Figure"))
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

	// Footnote
	fmt.Println(ty.H3("Footnote"))
	fmt.Println(ty.P("Herald supports rich typography" + ty.FootnoteRef(1) + " with multiple elements" + ty.FootnoteRef(2)))
	fmt.Println(ty.FootnoteSection([]string{
		"Built on lipgloss v2",
		"Headings, lists, alerts, and more",
	}))
	fmt.Println()
}
