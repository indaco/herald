// Compact highlight reel for the README hero image.
// Run: go run ./examples/demos/hero/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	// Headings (just two to show the range)
	fmt.Println(ty.H1("Herald - TUI Typography"))
	fmt.Println(ty.H4("Render rich text elements in your terminal"))

	// Paragraph & blockquote
	fmt.Println(ty.P("A Go library for HTML-analogous TUI typography, built on lipgloss v2."))
	fmt.Println(ty.Blockquote("Good design is as little design as possible.\n- Dieter Rams"))
	fmt.Println()

	// Code
	fmt.Println(ty.CodeBlock("ty := herald.New()\nfmt.Println(ty.H1(\"Hello, World!\"))"))

	fmt.Println(ty.HR())
	fmt.Println()

	// Lists (one flat, one nested)
	fmt.Println(ty.UL("Headings H1–H6", "Lists & nested lists", "Inline styles & alerts"))
	fmt.Println()
	fmt.Println(ty.NestOL(
		herald.Item("Getting Started"),
		herald.ItemWithChildren("Features",
			herald.Item("Typography"),
			herald.Item("Theming"),
		),
	))
	fmt.Println()

	// Table
	fmt.Println(ty.Table([][]string{
		{"Name", "Role", "Status"},
		{"Alice", "Admin", "Active"},
		{"Bob", "Editor", "Idle"},
		{"Charlie", "Viewer", "Active"},
	}))
	fmt.Println()

	// Inline styles (single line showcase)
	fmt.Println(
		ty.Bold("Bold") + "  " +
			ty.Italic("Italic") + "  " +
			ty.Mark("Highlight") + "  " +
			ty.Kbd("Ctrl") + "+" + ty.Kbd("C") + "  " +
			ty.Code("inline code"),
	)
	fmt.Println(ty.Q("quoted") + "  " + ty.Cite("citation") + "  " + ty.Samp("output") + "  " + ty.Var("x"))
	fmt.Println(ty.Ins("added") + "  " + ty.Del("removed") + "  " + ty.Badge("NEW") + "  " + ty.Tag("go"))
	fmt.Println(ty.SuccessBadge("running") + "  " + ty.ErrorBadge("failed") + "  " + ty.WarningBadge("expiring") + "  " + ty.InfoBadge("pending"))
	fmt.Println()

	// Fieldset
	fmt.Println(ty.Fieldset("Config", "Host: localhost\nPort: 8080"))
	fmt.Println()

	// Key-value
	fmt.Println(ty.KVGroup([][2]string{
		{"Version", "0.7.0"},
		{"License", "MIT"},
		{"Go", "1.25+"},
	}))
	fmt.Println()

	// Alert (just one)
	fmt.Println(ty.Tip("Herald supports Note, Tip, Important, Warning, and Caution alerts."))
}
