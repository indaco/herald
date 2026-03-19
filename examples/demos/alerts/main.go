// GitHub-style alerts: Note, Tip, Important, Warning, Caution.
// Run: go run ./examples/demos/alerts/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	fmt.Println(ty.Note("Useful information that users should know, even when skimming content."))
	fmt.Println()
	fmt.Println(ty.Tip("Helpful advice for doing things better or more easily."))
	fmt.Println()
	fmt.Println(ty.Important("Key information users need to know to achieve their goal."))
	fmt.Println()
	fmt.Println(ty.Warning("Urgent info that needs immediate user attention to avoid problems."))
	fmt.Println()
	fmt.Println(ty.Caution("Advises about risks or negative outcomes of certain actions."))
}
