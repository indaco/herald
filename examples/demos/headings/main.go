// All six heading levels with the default Rose Pine theme.
// Run: go run ./examples/demos/headings/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	fmt.Println(ty.H1("Heading 1 - Double Underline"))
	fmt.Println(ty.H2("Heading 2 - Single Underline"))
	fmt.Println(ty.H3("Heading 3 - Dotted Underline"))
	fmt.Println(ty.H4("Heading 4 - Bar Prefix"))
	fmt.Println(ty.H5("Heading 5 - Bar Prefix"))
	fmt.Println(ty.H6("Heading 6 - Bar Prefix"))
}
