// Flat and nested lists with mixed nesting.
// Run: go run ./examples/demos/lists/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	// Flat lists
	fmt.Println(ty.H3("Unordered List"))
	fmt.Println(ty.UL("Apples", "Bananas", "Cherries"))
	fmt.Println()

	fmt.Println(ty.H3("Ordered List"))
	fmt.Println(ty.OL("First item", "Second item", "Third item"))
	fmt.Println()

	// Nested unordered list
	fmt.Println(ty.H3("Nested Unordered List"))
	fmt.Println(ty.NestUL(
		herald.Item("Fruits"),
		herald.ItemWithChildren("Vegetables",
			herald.Item("Carrots"),
			herald.Item("Peas"),
			herald.ItemWithChildren("Leafy Greens",
				herald.Item("Spinach"),
				herald.Item("Kale"),
			),
		),
		herald.Item("Grains"),
	))
	fmt.Println()

	// Nested ordered list
	fmt.Println(ty.H3("Nested Ordered List"))
	fmt.Println(ty.NestOL(
		herald.Item("Introduction"),
		herald.ItemWithOLChildren("Main Topics",
			herald.Item("Architecture"),
			herald.Item("Design"),
		),
		herald.Item("Conclusion"),
	))
	fmt.Println()

	// Mixed nesting
	fmt.Println(ty.H3("Mixed — Ordered inside Unordered"))
	fmt.Println(ty.NestUL(
		herald.Item("Overview"),
		herald.ItemWithOLChildren("Ranked Desserts",
			herald.Item("Ice cream"),
			herald.Item("Cake"),
			herald.Item("Pie"),
		),
		herald.Item("Summary"),
	))
}
