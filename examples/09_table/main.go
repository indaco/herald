// Table rendering: bordered (default), minimal, custom padding, custom styles,
// column alignment, row separators, striped rows, captions, footer rows,
// and auto-truncation.
// Run: go run ./examples/09_table/
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/indaco/herald"
)

func main() {
	// --- Bordered table (default) ---
	ty := herald.New()

	fmt.Println(ty.H3("Bordered Table (default)"))
	fmt.Println(ty.Table([][]string{
		{"Name", "Role", "Status"},
		{"Alice", "Admin", "Active"},
		{"Bob", "Editor", "Idle"},
		{"Charlie", "Viewer", "Active"},
	}))
	fmt.Println()

	// --- Minimal table ---
	tyMin := herald.New(herald.WithTableBorderSet(herald.MinimalBorderSet()))

	fmt.Println(ty.H3("Minimal Table"))
	fmt.Println(tyMin.Table([][]string{
		{"Language", "Typing", "Compiled"},
		{"Go", "Static", "Yes"},
		{"Python", "Dynamic", "No"},
		{"Rust", "Static", "Yes"},
	}))
	fmt.Println()

	// --- Custom cell padding ---
	tyPad := herald.New(herald.WithTableCellPad(2))

	fmt.Println(ty.H3("Custom Cell Padding (2)"))
	fmt.Println(tyPad.Table([][]string{
		{"Feature", "Herald", "Plain"},
		{"Headings", "H1–H6", "fmt.Println"},
		{"Lists", "UL/OL/Nested", "Manual"},
		{"Tables", "Bordered/Minimal", "N/A"},
	}))
	fmt.Println()

	// --- Custom styles ---
	tyCustom := herald.New(
		herald.WithTableHeaderStyle(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF6600"))),
		herald.WithTableBorderStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555"))),
		herald.WithTableCellStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC"))),
	)

	fmt.Println(ty.H3("Custom Styles"))
	fmt.Println(tyCustom.Table([][]string{
		{"Command", "Description"},
		{"go build", "Compile packages"},
		{"go test", "Run tests"},
		{"go fmt", "Format source code"},
	}))
	fmt.Println()

	// --- Column alignment ---
	fmt.Println(ty.H3("Column Alignment"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"Language", "Stars", "License"},
		{"Go", "125000", "BSD-3"},
		{"Rust", "98000", "MIT/Apache"},
		{"Python", "65000", "PSF"},
	},
		herald.WithColumnAlign(1, herald.AlignRight),
		herald.WithColumnAlign(2, herald.AlignCenter),
	))
	fmt.Println()

	// --- All columns aligned at once ---
	fmt.Println(ty.H3("Bulk Column Alignment"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"Item", "Qty", "Price"},
		{"Widget", "42", "$9.99"},
		{"Gadget", "7", "$49.50"},
		{"Doohickey", "123", "$1.25"},
	},
		herald.WithColumnAligns(herald.AlignLeft, herald.AlignCenter, herald.AlignRight),
	))
	fmt.Println()

	// --- Row separators ---
	fmt.Println(ty.H3("Row Separators"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"Name", "Score", "Grade"},
		{"Alice", "95", "A"},
		{"Bob", "82", "B"},
		{"Charlie", "71", "C"},
	},
		herald.WithRowSeparators(true),
		herald.WithColumnAlign(1, herald.AlignRight),
	))
	fmt.Println()

	// --- Striped rows ---
	fmt.Println(ty.H3("Striped Rows"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"ID", "Name", "Department"},
		{"1", "Alice", "Engineering"},
		{"2", "Bob", "Marketing"},
		{"3", "Charlie", "Design"},
		{"4", "Diana", "Engineering"},
		{"5", "Eve", "Sales"},
	},
		herald.WithStripedRows(true),
		herald.WithColumnAlign(0, herald.AlignRight),
	))
	fmt.Println()

	// --- Caption ---
	fmt.Println(ty.H3("Table with Caption"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"Language", "Year", "Creator"},
		{"Go", "2009", "Google"},
		{"Rust", "2010", "Mozilla"},
		{"Zig", "2016", "A. Kelley"},
	},
		herald.WithCaption("Table 1: Programming Languages"),
	))
	fmt.Println()

	// --- Caption bottom ---
	fmt.Println(ty.H3("Table with Bottom Caption"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"Metric", "Value"},
		{"Latency", "12ms"},
		{"Throughput", "1.2k rps"},
	},
		herald.WithCaptionBottom("Source: internal benchmarks"),
	))
	fmt.Println()

	// --- Footer row ---
	fmt.Println(ty.H3("Footer Row"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"Item", "Qty", "Price"},
		{"Widget", "10", "$9.99"},
		{"Gadget", "5", "$24.50"},
		{"Doohickey", "20", "$1.25"},
		{"Total", "35", "$35.74"},
	},
		herald.WithFooterRow(true),
		herald.WithColumnAlign(1, herald.AlignRight),
		herald.WithColumnAlign(2, herald.AlignRight),
	))
	fmt.Println()

	// --- Auto-truncation (global) ---
	fmt.Println(ty.H3("Auto-Truncation (max 12)"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"Package", "Description"},
		{"herald", "HTML-inspired typography for terminal UIs"},
		{"lipgloss", "Style definitions for nice terminal layouts"},
		{"huh", "Build interactive forms in the terminal"},
	},
		herald.WithMaxColumnWidth(12),
	))
	fmt.Println()

	// --- Per-column truncation ---
	fmt.Println(ty.H3("Per-Column Truncation"))
	fmt.Println(ty.TableWithOpts([][]string{
		{"Name", "Email", "Role"},
		{"Alice Johnson", "alice.johnson@example.com", "Administrator"},
		{"Bob Smith", "bob.smith@example.com", "Editor"},
	},
		herald.WithColumnMaxWidth(1, 15),
	))
}
