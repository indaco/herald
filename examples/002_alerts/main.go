// GitHub-style alerts demo with the default Rose Pine theme.
// Run: go run ./examples/002_alerts/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	fmt.Println(ty.H2("GitHub-Style Alerts"))
	fmt.Println()

	// All 5 alert types with default settings
	fmt.Println(ty.Note("Useful information that users should know, even when skimming content."))
	fmt.Println()

	fmt.Println(ty.Tip("Helpful advice for doing things better or more easily."))
	fmt.Println()

	fmt.Println(ty.Important("Key information users need to know to achieve their goal."))
	fmt.Println()

	fmt.Println(ty.Warning("Urgent info that needs immediate user attention to avoid problems."))
	fmt.Println()

	fmt.Println(ty.Caution("Advises about risks or negative outcomes of certain actions."))
	fmt.Println()

	// Multi-line alert
	fmt.Println(ty.H3("Multi-line Alert"))
	fmt.Println(ty.Warning("This is the first line of a warning.\nThis is the second line.\nAnd a third for good measure."))
	fmt.Println()

	// Custom icon example
	fmt.Println(ty.H3("Custom Icon"))
	custom := herald.New(herald.WithAlertIcon(herald.AlertNote, ">>"))
	fmt.Println(custom.Note("This note uses a custom icon."))
}
