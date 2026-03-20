// Table elements: bordered, minimal, alignment, striped rows, caption, footer,
// and auto-truncation.
// Run: go run ./examples/demos/tables/
package main

import (
	"fmt"

	"github.com/indaco/herald"
)

func main() {
	ty := herald.New()

	fmt.Println(ty.Table([][]string{
		{"Name", "Role", "Status"},
		{"Alice", "Admin", "Active"},
		{"Bob", "Editor", "Idle"},
		{"Charlie", "Viewer", "Active"},
	}))
	fmt.Println()

	tyMin := herald.New(herald.WithTableBorderSet(herald.MinimalBorderSet()))
	fmt.Println(tyMin.Table([][]string{
		{"Language", "Typing", "Compiled"},
		{"Go", "Static", "Yes"},
		{"Python", "Dynamic", "No"},
		{"Rust", "Static", "Yes"},
	}))
	fmt.Println()

	// Striped rows with right-aligned IDs.
	fmt.Println(ty.TableWithOpts([][]string{
		{"ID", "Name", "Department"},
		{"1", "Alice", "Engineering"},
		{"2", "Bob", "Marketing"},
		{"3", "Charlie", "Design"},
		{"4", "Diana", "Sales"},
	},
		herald.WithStripedRows(true),
		herald.WithColumnAlign(0, herald.AlignRight),
	))
	fmt.Println()

	// Footer row with alignment and caption.
	fmt.Println(ty.TableWithOpts([][]string{
		{"Item", "Qty", "Price"},
		{"Widget", "10", "$9.99"},
		{"Gadget", "5", "$24.50"},
		{"Doohickey", "20", "$1.25"},
		{"Total", "35", "$35.74"},
	},
		herald.WithCaption("Table: Order Summary"),
		herald.WithFooterRow(true),
		herald.WithColumnAlign(1, herald.AlignRight),
		herald.WithColumnAlign(2, herald.AlignRight),
	))
	fmt.Println()

	// Auto-truncation on a specific column.
	fmt.Println(ty.TableWithOpts([][]string{
		{"Name", "Email", "Role"},
		{"Alice Johnson", "alice.johnson@example.com", "Admin"},
		{"Bob Smith", "bob.smith@example.com", "Editor"},
	},
		herald.WithColumnMaxWidth(1, 15),
	))
}
