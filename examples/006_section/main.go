// Section groups a heading with its content so Compose does not insert
// an extra blank line between them.
// Run: go run ./examples/006_section/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	// Without Section, Compose would insert a double newline between
	// each H4 and its list, creating unwanted vertical space.
	// Section joins them with a single newline instead.
	page := ty.Compose(
		ty.H2("Shopping List"),
		ty.Section(
			ty.H4("Fruits"),
			ty.UL("Apples", "Bananas", "Cherries"),
		),
		ty.Section(
			ty.H4("Vegetables"),
			ty.UL("Carrots", "Spinach", "Peppers"),
		),
		ty.HR(),
		ty.P("BR inserts a simple line break:"),
		ty.P("Line one"+ty.BR()+"Line two"),
	)

	fmt.Println(page)
}
