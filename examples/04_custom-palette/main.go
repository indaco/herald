// Using ColorPalette with adaptive light/dark colors to create a custom theme.
// Run: go run ./examples/04_custom-palette/
package main

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/indaco/herald"
)

func main() {
	// lipgloss.LightDark returns a function that picks the right color
	// based on the terminal background — define your palette once.
	lightDark := lipgloss.LightDark(lipgloss.HasDarkBackground(os.Stdin, os.Stdout))

	// Nord-inspired palette — cool blues, muted frost tones
	palette := herald.ColorPalette{
		Primary:   lightDark(lipgloss.Color("#5E81AC"), lipgloss.Color("#88C0D0")), // frost blue
		Secondary: lightDark(lipgloss.Color("#81A1C1"), lipgloss.Color("#81A1C1")), // frost lighter
		Tertiary:  lightDark(lipgloss.Color("#8FBCBB"), lipgloss.Color("#8FBCBB")), // frost teal
		Accent:    lightDark(lipgloss.Color("#EBCB8B"), lipgloss.Color("#EBCB8B")), // aurora yellow
		Highlight: lightDark(lipgloss.Color("#BF616A"), lipgloss.Color("#BF616A")), // aurora red
		Muted:     lightDark(lipgloss.Color("#7B88A1"), lipgloss.Color("#4C566A")), // polar night muted
		Text:      lightDark(lipgloss.Color("#2E3440"), lipgloss.Color("#ECEFF4")), // polar / snow
		Surface:   lightDark(lipgloss.Color("#D8DEE9"), lipgloss.Color("#3B4252")), // snow / polar
		Base:      lightDark(lipgloss.Color("#ECEFF4"), lipgloss.Color("#2E3440")), // snow / polar night
	}

	ty := herald.New(herald.WithPalette(palette))

	fmt.Println(ty.H1("Nord Palette Demo"))
	fmt.Println(ty.H2("Subheading"))
	fmt.Println(ty.H3("Section"))
	fmt.Println(ty.H4("Subsection"))
	fmt.Println(ty.H5("Minor heading"))
	fmt.Println(ty.H6("Smallest heading"))

	fmt.Println(ty.P("A full theme generated from 9 adaptive colors."))
	fmt.Println(ty.Blockquote("Muted text in a blockquote."))
	fmt.Println()

	fmt.Println(ty.Code("inline code") + " within text")
	fmt.Println()
	fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"hello\")\n}"))

	fmt.Println(ty.HR())
	fmt.Println()

	fmt.Println(ty.UL("Primary", "Secondary", "Tertiary"))
	fmt.Println()
	fmt.Println(ty.OL("First", "Second", "Third"))
	fmt.Println()

	fmt.Println(ty.Bold("Bold") + ", " + ty.Italic("italic") + ", " + ty.Mark("highlighted"))
	fmt.Println(ty.Kbd("Ctrl") + " + " + ty.Kbd("C"))
	fmt.Println(ty.Link("Go website", "https://go.dev"))
	fmt.Println(ty.Abbr("TUI", "Terminal User Interface"))
	fmt.Println()

	fmt.Println(ty.DL([][2]string{
		{"Primary", "Main accent — headings, emphasis"},
		{"Muted", "Subdued elements — comments, borders"},
	}))
}
