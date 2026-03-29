// Demonstrates flat lists, nested lists, mixed nesting, and hierarchical numbering.
// Run: go run ./examples/001_lists/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	// --- Flat lists (UL / OL) ---

	fmt.Println(ty.H2("Flat Lists"))

	fmt.Println(ty.H3("Unordered List"))
	fmt.Println(ty.UL("Apples", "Bananas", "Cherries"))
	fmt.Println()

	fmt.Println(ty.H3("Ordered List"))
	fmt.Println(ty.OL("First item", "Second item", "Third item"))
	fmt.Println()

	// --- Nested unordered list ---

	fmt.Println(ty.H2("Nested Lists"))

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

	// --- Nested ordered list ---

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

	// --- Mixed nesting (OL children inside UL) ---

	fmt.Println(ty.H3("Mixed Nesting - Ordered inside Unordered"))
	fmt.Println(ty.NestUL(
		herald.Item("Overview"),
		herald.ItemWithOLChildren("Ranked Desserts",
			herald.Item("Ice cream"),
			herald.Item("Cake"),
			herald.Item("Pie"),
		),
		herald.Item("Summary"),
	))
	fmt.Println()

	// --- Mixed nesting (UL children inside OL) ---

	fmt.Println(ty.H3("Mixed Nesting - Unordered inside Ordered"))
	fmt.Println(ty.NestOL(
		herald.Item("Planning"),
		herald.ItemWithChildren("Resources",
			herald.Item("Books"),
			herald.Item("Articles"),
		),
		herald.Item("Execution"),
	))
	fmt.Println()

	// --- Batch items with Items() ---

	fmt.Println(ty.H3("Batch Items with Items()"))
	fmt.Println(ty.NestUL(herald.Items("Alpha", "Beta", "Gamma")...))
	fmt.Println()

	// --- Hierarchical numbering ---

	fmt.Println(ty.H2("Hierarchical Numbering"))

	tyHier := herald.New(herald.WithHierarchicalNumbers(true))

	fmt.Println(tyHier.H3("Outline-style Ordered List"))
	fmt.Println(tyHier.NestOL(
		herald.Item("Introduction"),
		herald.ItemWithOLChildren("Main Topics",
			herald.Item("Architecture"),
			herald.ItemWithOLChildren("Design",
				herald.Item("UI"),
				herald.Item("API"),
			),
		),
		herald.Item("Conclusion"),
	))
	fmt.Println()

	// --- Custom options ---

	fmt.Println(ty.H2("Custom Options"))

	tyCustom := herald.New(
		herald.WithNestedBulletChars([]string{"*", "o", "-", ">"}),
		herald.WithListIndent(4),
	)

	fmt.Println(tyCustom.H3("Custom Bullets and Indent"))
	fmt.Println(tyCustom.NestUL(
		herald.Item("Level 0"),
		herald.ItemWithChildren("Level 0 with children",
			herald.Item("Level 1"),
			herald.ItemWithChildren("Level 1 with children",
				herald.Item("Level 2"),
				herald.ItemWithChildren("Level 2 with children",
					herald.Item("Level 3"),
				),
			),
		),
	))
}
