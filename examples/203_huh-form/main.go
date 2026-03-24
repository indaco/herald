// Using herald with huh for interactive TUI forms.
//
// This example is a separate Go module with its own go.mod to keep
// charm.land/huh/v2 out of herald's core dependencies.
//
// Run from the repository root:
//
// cd examples/203_huh-form && go run .
package main

import (
	"fmt"
	"os"

	"charm.land/huh/v2"
	"github.com/indaco/herald"
)

func main() {
	ty := herald.New(herald.WithTheme(herald.DraculaTheme()))

	// --- Herald: introduction block ---
	fmt.Println(ty.H1("Project Scaffolder"))
	fmt.Println(ty.P("This tool will scaffold a new project for you."))
	fmt.Println(ty.Blockquote("All fields are required. You can change them later in your config file."))
	fmt.Println(ty.HR())

	// --- huh: collect input ---
	var (
		name     string
		language string
		confirm  bool
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project name").
				Value(&name),
			huh.NewSelect[string]().
				Title("Primary language").
				Options(
					huh.NewOption("Go", "go"),
					huh.NewOption("TypeScript", "typescript"),
					huh.NewOption("Python", "python"),
				).
				Value(&language),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Ready to scaffold?").
				Value(&confirm),
		),
	).WithAccessible(true).WithTheme(huh.ThemeFunc(huh.ThemeDracula))

	if err := form.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if !confirm {
		fmt.Println(ty.P("Aborted. No files were created."))
		os.Exit(0)
	}

	// --- Herald: results block ---
	fmt.Println(ty.HR())
	fmt.Println(ty.H2("Project Created"))
	fmt.Println(ty.P("Your project has been scaffolded with the following settings:"))
	fmt.Println(ty.DL([][2]string{
		{"Name", name},
		{"Language", language},
	}))
	fmt.Println(ty.P(
		"Next: " + ty.Bold("cd "+name) + " and run " + ty.Code("go mod tidy", ""),
	))
}
