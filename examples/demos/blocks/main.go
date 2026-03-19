// Block-level elements: paragraph, blockquote, code, code block, HR, and definition list.
// Run: go run ./examples/demos/blocks/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	fmt.Println(ty.P("This is a paragraph. It wraps text with the paragraph style."))
	fmt.Println(ty.Blockquote("A wise person once said something profound.\nAnd then said more on a second line."))
	fmt.Println()

	fmt.Println(ty.Code("fmt.Println()"))
	fmt.Println()
	fmt.Println(ty.CodeBlock("func main() {\n\tfmt.Println(\"Hello, World!\")\n}"))

	fmt.Println(ty.HR())
	fmt.Println()

	fmt.Println(ty.DL([][2]string{
		{"Go", "A statically typed, compiled language"},
		{"Rust", "A systems programming language"},
		{"Python", "A dynamic, interpreted language"},
	}))
	fmt.Println()

	fmt.Println(ty.DT("Manual Term"))
	fmt.Println(ty.DD("Manual description using DT/DD directly"))
}
