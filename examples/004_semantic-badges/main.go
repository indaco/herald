// Semantic badges and tags demo with the default Rose Pine theme.
// Run: go run ./examples/004_semantic-badges/
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	fmt.Println(ty.H2("Semantic Badges"))
	fmt.Println(
		ty.SuccessBadge("running"),
		ty.ErrorBadge("failed"),
		ty.WarningBadge("expiring"),
		ty.InfoBadge("pending"),
	)
	fmt.Println()

	fmt.Println(ty.H2("Semantic Tags"))
	fmt.Println(
		ty.SuccessTag("healthy"),
		ty.ErrorTag("critical"),
		ty.WarningTag("degraded"),
		ty.InfoTag("maintenance"),
	)
	fmt.Println()

	// Custom semantic palette
	fmt.Println(ty.H2("Custom Semantic Palette"))
	custom := herald.New(
		herald.WithSemanticPalette(herald.SemanticPalette{
			Success: lipgloss.Color("#22c55e"),
			Warning: lipgloss.Color("#f59e0b"),
			Error:   lipgloss.Color("#ef4444"),
			Info:    lipgloss.Color("#3b82f6"),
		}),
	)
	fmt.Println(
		custom.SuccessBadge("passed"),
		custom.ErrorBadge("error"),
		custom.WarningBadge("warning"),
		custom.InfoBadge("info"),
	)
}
