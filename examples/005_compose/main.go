// Compose multiple rendered blocks into a single output.
// Run: go run ./examples/005_compose/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	// Compose joins rendered blocks with double newlines,
	// skipping any empty blocks automatically.
	page := ty.Compose(
		ty.H1("Release Notes"),
		ty.P("Herald v0.9.0 brings the Compose method for building multi-block layouts with less boilerplate."),
		ty.H2("What's New"),
		ty.UL(
			"Compose method for joining rendered blocks",
			"Reduced verbosity when building full pages",
			"Empty blocks are automatically skipped",
		),
		ty.H2("Example"),
		ty.CodeBlock("page := ty.Compose(\n\tty.H1(\"Title\"),\n\tty.P(\"Body text.\"),\n\tty.UL(\"one\", \"two\", \"three\"),\n)"),
		ty.HR(),
		ty.KVGroup([][2]string{
			{"Version", "0.9.0"},
			{"License", "MIT"},
			{"Author", "indaco"},
		}),
	)

	fmt.Println(page)
}
