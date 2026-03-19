// Using herald's built-in named themes that match huh's color palettes.
// Colors auto-adapt to light/dark terminal backgrounds.
// Run: go run ./examples/05_builtin-themes/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	themes := []struct {
		name  string
		theme herald.Theme
	}{
		{"Dracula", herald.DraculaTheme()},
		{"Catppuccin", herald.CatppuccinTheme()},
		{"Base16", herald.Base16Theme()},
		{"Charm", herald.CharmTheme()},
	}

	for i := range themes {
		th := &themes[i]
		ty := herald.New(herald.WithTheme(th.theme))

		fmt.Println(ty.H1(th.name + " Theme"))
		fmt.Println(ty.H2("Subheading"))
		fmt.Println(ty.H3("Section"))
		fmt.Println(ty.H4("Subsection"))
		fmt.Println()

		fmt.Println(ty.P("Body text styled with the " + th.name + " palette."))
		fmt.Println(ty.Blockquote("A blockquote in muted tones."))
		fmt.Println()

		fmt.Println(ty.Code("inline") + " and block:")
		fmt.Println()
		fmt.Println(ty.CodeBlock("fmt.Println(\"hello\")"))

		fmt.Println(ty.UL("First item", "Second item", "Third item"))
		fmt.Println()

		fmt.Println(ty.Bold("Bold") + " " + ty.Italic("Italic") + " " + ty.Mark("Highlight"))
		fmt.Println(ty.Kbd("Ctrl") + "+" + ty.Kbd("C"))
		fmt.Println(ty.Link("Link", "https://example.com"))
		fmt.Println()
		fmt.Println(ty.HR())
		fmt.Println()
	}
}
