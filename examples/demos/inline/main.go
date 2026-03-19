// Inline styles, links, and abbreviations.
// Run: go run ./examples/demos/inline/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	fmt.Println(ty.Bold("Bold text"))
	fmt.Println(ty.Italic("Italic text"))
	fmt.Println(ty.Underline("Underlined text"))
	fmt.Println(ty.Strikethrough("Strikethrough text"))
	fmt.Println(ty.Small("Small/faint text"))
	fmt.Println(ty.Mark("Highlighted text"))
	fmt.Println(ty.Code("inline code"))
	fmt.Println(ty.Kbd("Ctrl") + " + " + ty.Kbd("C"))
	fmt.Println(ty.Sub("subscript") + " and " + ty.Sup("superscript"))
	fmt.Println()

	fmt.Println(ty.Link("https://go.dev"))
	fmt.Println(ty.Link("Go website", "https://go.dev"))
	fmt.Println(ty.Abbr("HTML"))
	fmt.Println(ty.Abbr("CSS", "Cascading Style Sheets"))
}
