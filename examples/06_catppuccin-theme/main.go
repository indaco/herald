// Building a complete herald Theme from the Catppuccin color palette.
// This example uses the Mocha flavor; swap with Frappe, Macchiato, or Latte.
//
// This example is a separate Go module with its own go.mod to keep
// github.com/catppuccin/go out of herald's core dependencies.
//
// Run from the repository root:
//
//	cd examples/06_catppuccin-theme && go run .
package main

import (
	"fmt"

	catppuccin "github.com/catppuccin/go"

	"charm.land/lipgloss/v2"
	"github.com/indaco/herald"
)

// catppuccinTheme builds a herald.Theme from a Catppuccin flavor.
func catppuccinTheme(f catppuccin.Flavor) herald.Theme {
	text := lipgloss.Color(f.Text().Hex)
	subtext := lipgloss.Color(f.Subtext0().Hex)
	overlay := lipgloss.Color(f.Overlay0().Hex)
	surface := lipgloss.Color(f.Surface0().Hex)
	base := lipgloss.Color(f.Base().Hex)
	mauve := lipgloss.Color(f.Mauve().Hex)
	lavender := lipgloss.Color(f.Lavender().Hex)
	teal := lipgloss.Color(f.Teal().Hex)
	peach := lipgloss.Color(f.Peach().Hex)
	pink := lipgloss.Color(f.Pink().Hex)
	rosewater := lipgloss.Color(f.Rosewater().Hex)
	yellow := lipgloss.Color(f.Yellow().Hex)
	sky := lipgloss.Color(f.Sky().Hex)
	mantle := lipgloss.Color(f.Mantle().Hex)

	return herald.Theme{
		H1: lipgloss.NewStyle().Bold(true).Foreground(text).MarginBottom(1),
		H2: lipgloss.NewStyle().Bold(true).Foreground(mauve).MarginBottom(1),
		H3: lipgloss.NewStyle().Bold(true).Foreground(teal).MarginBottom(1),
		H4: lipgloss.NewStyle().Bold(true).Foreground(peach).MarginBottom(1),
		H5: lipgloss.NewStyle().Bold(true).Foreground(pink).MarginBottom(1),
		H6: lipgloss.NewStyle().Bold(true).Foreground(subtext).MarginBottom(1),

		Paragraph:  lipgloss.NewStyle().MarginBottom(1),
		Blockquote: lipgloss.NewStyle().Foreground(overlay).Italic(true).PaddingLeft(2),
		CodeInline: lipgloss.NewStyle().Foreground(text).Background(surface),
		CodeBlock:  lipgloss.NewStyle().Foreground(text).Background(mantle).Padding(1, 2).MarginBottom(1),
		HR:         lipgloss.NewStyle().Foreground(overlay),

		ListBullet: lipgloss.NewStyle().Foreground(mauve),
		ListItem:   lipgloss.NewStyle(),

		Bold:          lipgloss.NewStyle().Bold(true),
		Italic:        lipgloss.NewStyle().Italic(true),
		Underline:     lipgloss.NewStyle().Underline(true),
		Strikethrough: lipgloss.NewStyle().Strikethrough(true),
		Small:         lipgloss.NewStyle().Faint(true),
		Mark:          lipgloss.NewStyle().Background(yellow).Foreground(base),
		Link:          lipgloss.NewStyle().Foreground(sky).Underline(true),
		Kbd:           lipgloss.NewStyle().Foreground(text).Background(surface).Bold(true).Padding(0, 1),
		Abbr:          lipgloss.NewStyle().Underline(true).Foreground(rosewater),
		Sub:           lipgloss.NewStyle().Foreground(overlay),
		Sup:           lipgloss.NewStyle().Foreground(overlay),

		DT: lipgloss.NewStyle().Bold(true).Foreground(lavender),
		DD: lipgloss.NewStyle().PaddingLeft(4).Foreground(subtext),

		H1UnderlineChar: "═",
		H2UnderlineChar: "─",
		H3UnderlineChar: "·",
		HeadingBarChar:  "▎",
		BulletChar:      "•",
		HRChar:          "─",
		HRWidth:         40,
		BlockquoteBar:   "│",
	}
}

func main() {
	ty := herald.New(herald.WithTheme(catppuccinTheme(catppuccin.Mocha)))

	fmt.Println(ty.H1("Catppuccin Mocha Theme"))
	fmt.Println(ty.H2("Subheading"))
	fmt.Println(ty.H3("Section"))
	fmt.Println(ty.H4("Subsection"))
	fmt.Println()

	fmt.Println(ty.P("This example demonstrates herald styled with the Catppuccin Mocha palette."))
	fmt.Println(ty.Blockquote("A blockquote using overlay colors."))
	fmt.Println()

	fmt.Println(ty.Code("inline code") + " within text")
	fmt.Println()
	fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"catppuccin\")\n}"))

	fmt.Println(ty.HR())
	fmt.Println()

	fmt.Println(ty.UL("Rosewater", "Flamingo", "Mauve"))
	fmt.Println()
	fmt.Println(ty.OL("Latte", "Frappe", "Macchiato", "Mocha"))
	fmt.Println()

	fmt.Println(ty.Bold("Bold") + ", " + ty.Italic("italic") + ", " + ty.Mark("highlighted"))
	fmt.Println(ty.Kbd("Ctrl") + " + " + ty.Kbd("C"))
	fmt.Println(ty.Link("Catppuccin", "https://catppuccin.com"))
	fmt.Println(ty.Abbr("TUI", "Terminal User Interface"))
	fmt.Println()

	fmt.Println(ty.DL([][2]string{
		{"Mocha", "Dark, warm flavor"},
		{"Latte", "Light, soft flavor"},
	}))
}
