// Customizing herald with functional options: override individual styles,
// decoration characters, and tokens without replacing the entire theme.
// Run: go run ./examples/100_custom-options/
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/indaco/herald"
)

func main() {
	ty := herald.New(
		// Override heading styles
		herald.WithH1Style(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF6600")).
			MarginBottom(1)),

		// Change decoration characters
		herald.WithH1UnderlineChar("="),
		herald.WithH2UnderlineChar("-"),
		herald.WithH3UnderlineChar("~"),
		herald.WithHeadingBarChar("|"),

		// List customization
		herald.WithBulletChar("-"),
		herald.WithListBulletStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6600"))),

		// Horizontal rule
		herald.WithHRWidth(60),
		herald.WithHRChar("*"),

		// Blockquote
		herald.WithBlockquoteBar(">"),
		herald.WithBlockquoteStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Italic(true).
			PaddingLeft(2)),

		// Inline styles
		herald.WithMarkStyle(lipgloss.NewStyle().
			Background(lipgloss.Color("#FF6600")).
			Foreground(lipgloss.Color("#FFFFFF"))),
		herald.WithKbdStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#444444")).
			Bold(true).
			Padding(0, 1)),
	)

	fmt.Println(ty.H1("Custom Options Demo"))
	fmt.Println(ty.H2("Subheading with dashes"))
	fmt.Println(ty.H3("Section with tildes"))
	fmt.Println(ty.H4("Minor heading with pipe bar"))
	fmt.Println()

	fmt.Println(ty.P("Paragraph text using the default paragraph style."))
	fmt.Println(ty.Blockquote("Custom blockquote with > prefix."))
	fmt.Println()

	fmt.Println(ty.UL("Custom dash bullets", "Instead of dots", "More readable"))
	fmt.Println()

	fmt.Println(ty.HR())
	fmt.Println()

	fmt.Println(ty.Mark("Custom highlighted text"))
	fmt.Println(ty.Kbd("Ctrl") + " + " + ty.Kbd("C"))
}
