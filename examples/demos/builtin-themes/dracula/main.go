// Dracula theme showcase.
// Run: go run ./examples/demos/builtin-themes/dracula/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New(herald.WithTheme(herald.DraculaTheme()))

	fmt.Println(ty.H1("Dracula Theme"))
	fmt.Println(ty.H4("Render rich text elements in your terminal"))

	fmt.Println(ty.P("A Go library for HTML-analogous TUI typography, built on lipgloss v2."))
	fmt.Println(ty.Blockquote("Good design is as little design as possible.\n— Dieter Rams"))
	fmt.Println()

	fmt.Println(ty.CodeBlock("ty := herald.New()\nfmt.Println(ty.H1(\"Hello, World!\"))"))

	tyLN := herald.New(herald.WithTheme(herald.DraculaTheme()), herald.WithCodeLineNumbers(true))
	fmt.Println(tyLN.CodeBlock("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"))

	fmt.Println(ty.HR())
	fmt.Println()

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

	fmt.Println(
		ty.Bold("Bold") + "  " +
			ty.Italic("Italic") + "  " +
			ty.Mark("Highlight") + "  " +
			ty.Kbd("Ctrl") + "+" + ty.Kbd("C") + "  " +
			ty.Code("inline code"),
	)
	fmt.Println()

	fmt.Println(ty.Tip("Herald supports Note, Tip, Important, Warning, and Caution alerts."))
}
