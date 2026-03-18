// Using ColorPalette to create a consistent theme from just 8 colors.
// Run: go run ./examples/palette/
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/indaco/herald"
)

func main() {
	// Dracula-inspired palette
	palette := herald.ColorPalette{
		Primary:   lipgloss.Color("#bd93f9"), // purple
		Secondary: lipgloss.Color("#ff79c6"), // pink
		Tertiary:  lipgloss.Color("#8be9fd"), // cyan
		Accent:    lipgloss.Color("#ffb86c"), // orange
		Highlight: lipgloss.Color("#ff5555"), // red
		Muted:     lipgloss.Color("#6272a4"), // comment gray
		Text:      lipgloss.Color("#f8f8f2"), // foreground
		Surface:   lipgloss.Color("#44475a"), // current line
		Base:      lipgloss.Color("#282a36"), // background
	}

	ty := herald.New(herald.WithPalette(palette))

	fmt.Println(ty.H1("Dracula Palette Demo"))
	fmt.Println(ty.H2("Subheading"))
	fmt.Println(ty.H3("Section"))
	fmt.Println(ty.H4("Subsection"))
	fmt.Println(ty.H5("Minor heading"))
	fmt.Println(ty.H6("Smallest heading"))

	fmt.Println(ty.P("A full theme generated from just 8 colors."))
	fmt.Println(ty.Blockquote("Muted text in a blockquote."))
	fmt.Println()

	fmt.Println(ty.Code("inline code") + " within text")
	fmt.Println()
	fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"dracula\")\n}"))

	fmt.Println(ty.HR())
	fmt.Println()

	fmt.Println(ty.UL("Primary", "Secondary", "Tertiary"))
	fmt.Println()
	fmt.Println(ty.OL("First", "Second", "Third"))
	fmt.Println()

	fmt.Println(ty.Bold("Bold") + ", " + ty.Italic("italic") + ", " + ty.Mark("highlighted"))
	fmt.Println(ty.Kbd("Ctrl") + " + " + ty.Kbd("C"))
	fmt.Println(ty.Link("Dracula Theme", "https://draculatheme.com"))
	fmt.Println(ty.Abbr("TUI", "Terminal User Interface"))
	fmt.Println()

	fmt.Println(ty.DL([][2]string{
		{"Primary", "Main accent — headings, emphasis"},
		{"Muted", "Subdued elements — comments, borders"},
	}))
}
