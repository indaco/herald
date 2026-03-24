// Multi-step project scaffolder wizard using herald with huh.
//
// This example is a separate Go module with its own go.mod to keep
// charm.land/huh/v2 out of herald's core dependencies.
//
// Run from the repository root:
//
//	cd examples/204_huh-wizard && go run .
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"charm.land/huh/v2"
	"github.com/indaco/herald"
)

// languageFeatures maps each language to its available feature set.
var languageFeatures = map[string][]huh.Option[string]{
	"go": {
		huh.NewOption("CI/CD (GitHub Actions)", "ci"),
		huh.NewOption("Docker", "docker"),
		huh.NewOption("Tests", "tests"),
		huh.NewOption("Linting (golangci-lint)", "lint"),
		huh.NewOption("Documentation (GoDoc)", "docs"),
		huh.NewOption("Makefile", "makefile"),
		huh.NewOption("Pre-commit hooks", "hooks"),
	},
	"rust": {
		huh.NewOption("CI/CD (GitHub Actions)", "ci"),
		huh.NewOption("Docker", "docker"),
		huh.NewOption("Tests", "tests"),
		huh.NewOption("Linting (clippy)", "lint"),
		huh.NewOption("Documentation (rustdoc)", "docs"),
		huh.NewOption("Cargo workspace", "workspace"),
		huh.NewOption("Cross-compilation", "cross"),
	},
	"python": {
		huh.NewOption("CI/CD (GitHub Actions)", "ci"),
		huh.NewOption("Docker", "docker"),
		huh.NewOption("Tests (pytest)", "tests"),
		huh.NewOption("Linting (ruff)", "lint"),
		huh.NewOption("Documentation (Sphinx)", "docs"),
		huh.NewOption("Virtual environment", "venv"),
		huh.NewOption("Type checking (mypy)", "types"),
	},
	"typescript": {
		huh.NewOption("CI/CD (GitHub Actions)", "ci"),
		huh.NewOption("Docker", "docker"),
		huh.NewOption("Tests (vitest)", "tests"),
		huh.NewOption("Linting (ESLint + Prettier)", "lint"),
		huh.NewOption("Documentation (TypeDoc)", "docs"),
		huh.NewOption("Monorepo (turborepo)", "monorepo"),
		huh.NewOption("Bundler (esbuild)", "bundler"),
	},
}

// languageTips provides a language-specific tip for each supported language.
var languageTips = map[string]string{
	"go":         "Go projects benefit from a clear module path and minimal dependencies.\nConsider enabling tests and linting from the start.",
	"rust":       "Rust's ownership model catches bugs at compile time.\nEnable clippy and tests early to keep your codebase clean.",
	"python":     "Python projects thrive with virtual environments and type hints.\nAdding ruff and mypy from the start prevents technical debt.",
	"typescript": "TypeScript gives you type safety with JavaScript's ecosystem.\nPair ESLint with Prettier for consistent, error-free code.",
}

// languageDisplayName returns a human-friendly name for a language key.
func languageDisplayName(lang string) string {
	names := map[string]string{
		"go":         "Go",
		"rust":       "Rust",
		"python":     "Python",
		"typescript": "TypeScript",
	}
	if n, ok := names[lang]; ok {
		return n
	}
	return lang
}

// initCommands returns the shell commands to initialize a project for the given language.
func initCommands(name, lang string) string {
	switch lang {
	case "go":
		return fmt.Sprintf("mkdir %s && cd %s\ngo mod init github.com/yourname/%s\ngo mod tidy", name, name, name)
	case "rust":
		return fmt.Sprintf("cargo new %s\ncd %s\ncargo build", name, name)
	case "python":
		return fmt.Sprintf("mkdir %s && cd %s\npython -m venv .venv\nsource .venv/bin/activate\npip install -e .", name, name)
	case "typescript":
		return fmt.Sprintf("mkdir %s && cd %s\nnpm init -y\nnpm install typescript --save-dev\nnpx tsc --init", name, name)
	default:
		return fmt.Sprintf("mkdir %s && cd %s", name, name)
	}
}

func main() {
	ty := herald.New(herald.WithTheme(herald.DraculaTheme()))

	// -----------------------------------------------------------------------
	// Step 1: Welcome screen + collect project name and language
	// -----------------------------------------------------------------------

	fmt.Println(ty.H1("Project Scaffolder"))
	fmt.Println(ty.P("Create a new project with best practices baked in."))
	fmt.Println(ty.Blockquote("The beginning is the most important part of the work.\n-- Plato"))
	fmt.Println(ty.HR())

	var (
		name     string
		language string
	)

	form1 := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project name").
				Description("A short, lowercase identifier for your project.").
				Placeholder("my-awesome-app").
				Value(&name).
				Validate(func(s string) error {
					s = strings.TrimSpace(s)
					if s == "" {
						return fmt.Errorf("project name cannot be empty")
					}
					if strings.Contains(s, " ") {
						return fmt.Errorf("project name cannot contain spaces")
					}
					return nil
				}),
			huh.NewSelect[string]().
				Title("Primary language").
				Description("Choose the main language for your project.").
				Options(
					huh.NewOption("Go", "go"),
					huh.NewOption("Rust", "rust"),
					huh.NewOption("Python", "python"),
					huh.NewOption("TypeScript", "typescript"),
				).
				Value(&language),
		),
	).WithTheme(huh.ThemeFunc(huh.ThemeDracula))

	if err := form1.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	name = strings.TrimSpace(name)

	// -----------------------------------------------------------------------
	// Step 2: Language-specific features
	// -----------------------------------------------------------------------

	fmt.Println()
	fmt.Println(ty.H2(fmt.Sprintf("Configure %s Project", languageDisplayName(language))))
	fmt.Println(ty.Tip(languageTips[language]))
	fmt.Println()

	featureOpts := languageFeatures[language]

	featureLabels := make([]string, len(featureOpts))
	for i, opt := range featureOpts {
		featureLabels[i] = opt.Key
	}
	fmt.Println(ty.P("Available features for " + ty.Bold(languageDisplayName(language)) + " projects:"))
	fmt.Println(ty.UL(featureLabels...))
	fmt.Println(ty.HR())

	var features []string

	form2 := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select features to include").
				Description("Use space to toggle, enter to confirm.").
				Options(featureOpts...).
				Value(&features),
		),
	).WithTheme(huh.ThemeFunc(huh.ThemeDracula))

	if err := form2.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// -----------------------------------------------------------------------
	// Step 3: Confirmation
	// -----------------------------------------------------------------------

	var confirm bool

	form3 := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Ready to scaffold this project?").
				Value(&confirm),
		),
	).WithTheme(huh.ThemeFunc(huh.ThemeDracula))

	if err := form3.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if !confirm {
		fmt.Println()
		fmt.Println(ty.P("Aborted. No files were created."))
		os.Exit(0)
	}

	// -----------------------------------------------------------------------
	// Step 4: Summary
	// -----------------------------------------------------------------------

	fmt.Println()
	fmt.Println(ty.H1(name))
	fmt.Println()

	fmt.Println(
		ty.Badge(languageDisplayName(language)) + "  " +
			strings.Join(featureTags(ty, features), "  "),
	)
	fmt.Println()

	fmt.Println(ty.KVGroup([][2]string{
		{"Project", name},
		{"Language", languageDisplayName(language)},
		{"Features", fmt.Sprintf("%d selected", len(features))},
		{"Created", time.Now().Format("2006-01-02 15:04")},
	}))
	fmt.Println()

	fmt.Println(ty.HRWithLabel("Selected Features"))
	fmt.Println()

	if len(features) > 0 {
		featureNames := resolveFeatureNames(language, features)
		fmt.Println(ty.UL(featureNames...))
	} else {
		fmt.Println(ty.P(ty.Italic("No features selected.")))
	}

	fmt.Println()
	fmt.Println(ty.HRWithLabel("Next Steps"))
	fmt.Println()

	fmt.Println(ty.OL(
		"Initialize the project with the commands below",
		"Review the generated configuration files",
		"Install dependencies and run the test suite",
		"Set up your IDE with the recommended extensions",
		"Make your first commit and push to a remote repository",
	))
	fmt.Println()

	fmt.Println(ty.CodeBlock(initCommands(name, language), language))
	fmt.Println()

	fmt.Println(ty.Note(
		"Run " + ty.Code("git init", "") + " inside your project directory to start version control. " +
			"Then add a " + ty.Code(".gitignore", "") + " appropriate for " + languageDisplayName(language) + " projects.",
	))
	fmt.Println()

	fmt.Println(ty.HR())
}

// featureTags converts selected feature keys into herald Tag elements.
func featureTags(ty *herald.Typography, features []string) []string {
	tags := make([]string, len(features))
	for i, f := range features {
		tags[i] = ty.Tag(f)
	}
	return tags
}

// resolveFeatureNames maps feature keys back to their human-friendly labels
// for the selected language.
func resolveFeatureNames(lang string, keys []string) []string {
	opts := languageFeatures[lang]
	lookup := make(map[string]string, len(opts))
	for _, opt := range opts {
		lookup[opt.Value] = opt.Key
	}
	names := make([]string, 0, len(keys))
	for _, k := range keys {
		if label, ok := lookup[k]; ok {
			names = append(names, label)
		} else {
			names = append(names, k)
		}
	}
	return names
}
