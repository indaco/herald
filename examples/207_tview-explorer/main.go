// Project overview explorer using herald with tview.
//
// A sidebar list (tview.List) lets you browse topics while a scrollable
// content pane (tview.TextView) shows herald-rendered output. The bridge
// between the two libraries is tview.ANSIWriter, which converts lipgloss
// ANSI escape sequences into tview's native color tags.
// Left/right arrows switch focus, up/down navigates or scrolls, q quits.
//
// This example is a separate Go module with its own go.mod to keep
// github.com/rivo/tview out of herald's core dependencies.
//
// Run from the repository root:
//
//	cd examples/207_tview-explorer && go run .
package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/indaco/herald"
	"github.com/rivo/tview"
)

// ---------------------------------------------------------------------------
// Topic definitions
// ---------------------------------------------------------------------------

// topic holds a sidebar label and a render function that writes
// herald-styled content into a tview.TextView via tview.ANSIWriter.
type topic struct {
	title  string
	render func(ty *herald.Typography, tv *tview.TextView)
}

var topicList = []topic{
	{
		title: "Overview",
		render: func(ty *herald.Typography, tv *tview.TextView) {
			w := tview.ANSIWriter(tv)
			fmt.Fprintln(w, ty.H1("Herald"))
			fmt.Fprintln(w, ty.P(
				"HTML-inspired typography for terminal UIs in Go, built on "+
					ty.Bold("lipgloss v2")+". Herald maps familiar element names "+
					"(H1\u2013H6, P, Blockquote, Code, HR, lists) to styled terminal output.",
			))
			fmt.Fprintln(w, ty.HR())
			fmt.Fprintln(w, ty.H2("Highlights"))
			fmt.Fprintln(w, ty.UL(
				"Familiar HTML element names",
				"Built-in themes: Dracula, Catppuccin, Base16, Charm",
				"Pairs with huh, bubbletea, tview, and other TUI libraries",
				"Pluggable syntax highlighting via "+ty.Code("WithCodeFormatter"),
				"Custom palettes from just 9 colors",
			))
			fmt.Fprintln(w, ty.HR())
			fmt.Fprintln(w, ty.Blockquote(
				"The integration point with tview is "+ty.Code("tview.ANSIWriter")+".\n"+
					"It translates lipgloss ANSI output into tview's internal color tags.",
			))
		},
	},
	{
		title: "Installation",
		render: func(ty *herald.Typography, tv *tview.TextView) {
			w := tview.ANSIWriter(tv)
			fmt.Fprintln(w, ty.H1("Installation"))
			fmt.Fprintln(w, ty.P("Requires Go 1.25 or later."))
			fmt.Fprintln(w, ty.CodeBlock("go get github.com/indaco/herald@latest", "sh"))
			fmt.Fprintln(w, ty.H2("Quick Start"))
			fmt.Fprintln(w, ty.CodeBlock(`package main

import (
    "fmt"
    "github.com/indaco/herald"
)

func main() {
    ty := herald.New()
    fmt.Println(ty.H1("Hello, Herald!"))
    fmt.Println(ty.P("Rich terminal typography, simply."))
    fmt.Println(ty.UL("Headings", "Lists", "Inline styles"))
}`, "go"))
		},
	},
	{
		title: "Typography",
		render: func(ty *herald.Typography, tv *tview.TextView) {
			w := tview.ANSIWriter(tv)
			fmt.Fprintln(w, ty.H1("Typography Elements"))

			fmt.Fprintln(w, ty.H2("Headings"))
			fmt.Fprintln(w, ty.P("H1\u2013H3 render with an underline; H4\u2013H6 with a left-bar prefix."))
			fmt.Fprintln(w, ty.H3("This is H3"))
			fmt.Fprintln(w, ty.H4("This is H4"))
			fmt.Fprintln(w, ty.H5("This is H5"))
			fmt.Fprintln(w, ty.H6("This is H6"))

			fmt.Fprintln(w, ty.HR())

			fmt.Fprintln(w, ty.H2("Block Elements"))
			fmt.Fprintln(w, ty.Blockquote(
				"Blockquotes indent with a left bar.\nThey support multi-line input.",
			))
			fmt.Fprintln(w, ty.P(
				"Inline: "+ty.Code("os.Exit(1)")+" vs a fenced block below.",
			))
			fmt.Fprintln(w, ty.CodeBlock(`func greet(name string) string {
    return "Hello, " + name
}`, "go"))

			fmt.Fprintln(w, ty.HR())

			fmt.Fprintln(w, ty.H2("Inline Styles"))
			fmt.Fprintln(w, ty.P(
				ty.Bold("Bold")+" \u00b7 "+
					ty.Italic("Italic")+" \u00b7 "+
					ty.Underline("Underline")+" \u00b7 "+
					ty.Strikethrough("Strikethrough")+" \u00b7 "+
					ty.Mark("Marked")+" \u00b7 "+
					ty.Small("Small"),
			))
			fmt.Fprintln(w, ty.P(
				"Keyboard: "+ty.Kbd("Ctrl")+" + "+ty.Kbd("C")+"  |  "+
					"Link: "+ty.Link("pkg.go.dev", "https://pkg.go.dev/github.com/indaco/herald"),
			))

			fmt.Fprintln(w, ty.HRWithLabel("Horizontal Rules"))
			fmt.Fprintln(w, ty.P("Plain "+ty.Code("HR()")+" and labeled "+ty.Code("HRWithLabel()")+":"))
			fmt.Fprintln(w, ty.HR())
			fmt.Fprintln(w, ty.HRWithLabel("Section Break"))
			fmt.Fprintln(w, ty.Blockquote("The beginning is the most important part of the work.\n\u2014 Plato"))
		},
	},
	{
		title: "Lists",
		render: func(ty *herald.Typography, tv *tview.TextView) {
			w := tview.ANSIWriter(tv)
			fmt.Fprintln(w, ty.H1("Lists"))

			fmt.Fprintln(w, ty.H2("Unordered"))
			fmt.Fprintln(w, ty.UL("Design", "Implement", "Test"))

			fmt.Fprintln(w, ty.H2("Ordered"))
			fmt.Fprintln(w, ty.OL("Plan", "Build", "Ship", "Iterate"))

			fmt.Fprintln(w, ty.H2("Nested \u2014 mixed sub-lists"))
			fmt.Fprintln(w, ty.NestUL(
				herald.Item("Frontend"),
				herald.ItemWithChildren("Backend",
					herald.Item("REST API"),
					herald.Item("gRPC"),
				),
				herald.ItemWithOLChildren("DevOps",
					herald.Item("CI/CD"),
					herald.Item("Monitoring"),
					herald.Item("Alerting"),
				),
			))

			fmt.Fprintln(w, ty.H2("Hierarchical Ordered"))
			tyH := herald.New(herald.WithHierarchicalNumbers(true))
			fmt.Fprintln(w, tyH.NestOL(
				herald.Item("Introduction"),
				herald.ItemWithOLChildren("Core Concepts",
					herald.Item("Typography"),
					herald.Item("Themes"),
				),
				herald.Item("Conclusion"),
			))

			fmt.Fprintln(w, ty.HRWithLabel("Definition List"))
			fmt.Fprintln(w, ty.DL([][2]string{
				{"Typography", "The art of arranging type to make written language readable and appealing."},
				{"Theme", "A collection of styles that define the visual appearance of all elements."},
				{"Functional Option", "A pattern where configuration is passed as variadic function arguments."},
			}))
		},
	},
	{
		title: "Tables & Data",
		render: func(ty *herald.Typography, tv *tview.TextView) {
			w := tview.ANSIWriter(tv)
			fmt.Fprintln(w, ty.H2("Tables & Data"))
			fmt.Fprintln(w, ty.P("Herald renders bordered tables with header styling, optional striped rows, column alignment, and captions."))
			fmt.Fprintln(w, ty.TableWithOpts(
				[][]string{
					{"Method", "Element", "Category"},
					{"H1-H6", "Headings", "Block"},
					{"P", "Paragraph", "Block"},
					{"UL / OL", "Lists", "Block"},
					{"CodeBlock", "Fenced code", "Block"},
					{"Table", "Data table", "Block"},
					{"Bold", "Strong text", "Inline"},
					{"Code", "Inline code", "Inline"},
					{"Kbd", "Key indicator", "Inline"},
				},
				herald.WithCaption("Herald API Elements"),
				herald.WithStripedRows(true),
			))

			fmt.Fprintln(w, ty.HRWithLabel("Badges & Tags"))
			fmt.Fprintln(w, ty.P(
				ty.Badge("STABLE")+" "+ty.Tag("v1.0.0")+" "+ty.Badge("GO")+" "+ty.Tag("lipgloss-v2"),
			))
			fmt.Fprintln(w, ty.P("Use "+ty.Code("Badge()")+" for bold status pills and "+ty.Code("Tag()")+" for subtle category labels."))

			fmt.Fprintln(w, ty.HRWithLabel("Key-Value Pairs"))
			fmt.Fprintln(w, ty.KVGroup([][2]string{
				{"Package", "github.com/indaco/herald"},
				{"Go version", "1.25+"},
				{"Dependencies", "lipgloss v2 (single dep)"},
				{"License", "MIT"},
				{"Test coverage", "95%+"},
			}))
		},
	},
	{
		title: "Themes",
		render: func(ty *herald.Typography, tv *tview.TextView) {
			w := tview.ANSIWriter(tv)
			fmt.Fprintln(w, ty.H1("Themes"))
			fmt.Fprintln(w, ty.P(
				"Herald ships with named themes that match "+
					ty.Bold("huh")+"'s built-in palettes. "+
					"Colors auto-adapt to light/dark terminal backgrounds.",
			))

			fmt.Fprintln(w, ty.H2("Built-in Themes"))
			fmt.Fprintln(w, ty.DL([][2]string{
				{"Rose Pine", "Default \u2014 warm, muted palette with dark and light variants"},
				{"Dracula", "Matches huh.ThemeDracula()"},
				{"Catppuccin", "Mocha (dark) / Latte (light)"},
				{"Base16", "ANSI base16 terminal colors"},
				{"Charm", "Charm brand colors"},
			}))

			fmt.Fprintln(w, ty.H2("Applying a Theme"))
			fmt.Fprintln(w, ty.CodeBlock(
				`ty := herald.New(herald.WithTheme(herald.DraculaTheme()))`, "go",
			))

			fmt.Fprintln(w, ty.H2("Custom Palette"))
			fmt.Fprintln(w, ty.P(
				"Derive a full theme from 9 colors with "+ty.Code("ColorPalette")+". "+
					"Use "+ty.Code("lipgloss.LightDark")+" for adaptive colors that adjust "+
					"to the terminal background:",
			))
			fmt.Fprintln(w, ty.CodeBlock(`lightDark := lipgloss.LightDark(lipgloss.HasDarkBackground(os.Stdin, os.Stdout))

// Nord-inspired palette
palette := herald.ColorPalette{
    Primary:   lightDark(lipgloss.Color("#5E81AC"), lipgloss.Color("#88C0D0")),
    Secondary: lightDark(lipgloss.Color("#81A1C1"), lipgloss.Color("#81A1C1")),
    Tertiary:  lightDark(lipgloss.Color("#8FBCBB"), lipgloss.Color("#8FBCBB")),
    Accent:    lightDark(lipgloss.Color("#EBCB8B"), lipgloss.Color("#EBCB8B")),
    Highlight: lightDark(lipgloss.Color("#BF616A"), lipgloss.Color("#BF616A")),
    Muted:     lightDark(lipgloss.Color("#7B88A1"), lipgloss.Color("#4C566A")),
    Text:      lightDark(lipgloss.Color("#2E3440"), lipgloss.Color("#ECEFF4")),
    Surface:   lightDark(lipgloss.Color("#D8DEE9"), lipgloss.Color("#3B4252")),
    Base:      lightDark(lipgloss.Color("#ECEFF4"), lipgloss.Color("#2E3440")),
}
ty := herald.New(herald.WithPalette(palette))`, "go"))

			fmt.Fprintln(w, ty.Tip("All built-in themes are themselves built with "+ty.Code("ThemeFromPalette()")+"."))
		},
	},
	{
		title: "Alerts & Inline",
		render: func(ty *herald.Typography, tv *tview.TextView) {
			w := tview.ANSIWriter(tv)
			fmt.Fprintln(w, ty.H2("Alerts & Inline Styles"))
			fmt.Fprintln(w, ty.P("Herald provides five GitHub-style alert types:"))

			fmt.Fprintln(w, ty.Note("Useful information that users should know, even when skimming."))
			fmt.Fprintln(w, ty.Tip("Helpful advice for doing things better or more easily."))
			fmt.Fprintln(w, ty.Important("Key information users need to know to achieve their goal."))
			fmt.Fprintln(w, ty.Warning("Urgent info that needs immediate user attention to avoid problems."))
			fmt.Fprintln(w, ty.Caution("Advises about risks or negative outcomes of certain actions."))

			fmt.Fprintln(w, ty.HRWithLabel("More Inline Styles"))

			fmt.Fprintln(w, ty.P(
				"Inline code: "+ty.Code("fmt.Println()")+
					"  Keyboard: "+ty.Kbd("Ctrl")+"+"+ty.Kbd("C")+
					"  Highlight: "+ty.Mark("important"),
			))
			fmt.Fprintln(w, ty.P(
				"Links: "+ty.Link("Herald", "https://github.com/indaco/herald"),
			))
			fmt.Fprintln(w, ty.P(
				"Abbreviations: "+ty.Abbr("TUI", "Terminal User Interface")+
					", "+ty.Abbr("CLI", "Command-Line Interface"),
			))
			fmt.Fprintln(w, ty.P(
				"Subscript: H"+ty.Sub("2")+"O  "+
					"Superscript: E=mc"+ty.Sup("2"),
			))

			fmt.Fprintln(w, ty.HRWithLabel("Diff Markers"))
			fmt.Fprintln(w, ty.Del("typography.InlineCode(text)"))
			fmt.Fprintln(w, ty.Ins("typography.Code(text, lang)"))
		},
	},
}

// ---------------------------------------------------------------------------
// Application
// ---------------------------------------------------------------------------

func main() {
	app := tview.NewApplication()
	ty := herald.New(herald.WithTheme(herald.CatppuccinTheme()))

	// Right pane: scrollable content.
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWordWrap(true)
	textView.
		SetBorder(true).
		SetTitle(" Content ").
		SetBorderColor(tcell.ColorDimGray)

	// Render a topic into the content pane via ANSIWriter.
	renderTopic := func(idx int) {
		textView.Clear()
		topicList[idx].render(ty, textView)
		textView.ScrollToBeginning()
	}

	// Left pane: topic list.
	list := tview.NewList().
		SetHighlightFullLine(true).
		SetSecondaryTextColor(tcell.ColorDimGray).
		ShowSecondaryText(false)
	list.
		SetBorder(true).
		SetTitle(" Herald Explorer ").
		SetBorderColor(tcell.ColorDimGray)

	for i, t := range topicList {
		idx := i
		list.AddItem(t.title, "", 0, func() { renderTopic(idx) })
	}

	// Update content when the selected topic changes.
	list.SetChangedFunc(func(idx int, _ string, _ string, _ rune) {
		renderTopic(idx)
	})

	// Show first topic on start.
	renderTopic(0)

	// Footer bar with navigation hints.
	footer := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)
	footer.SetText("  [gray]\u2190/\u2192[-] [white]switch[-]  [gray]\u2191/\u2193[-] [white]navigate[-]  [gray]q[-] [white]quit[-]")

	// focusSidebar / focusContent update border colors and app focus.
	focusSidebar := func() {
		app.SetFocus(list)
		list.SetBorderColor(tcell.ColorMediumPurple)
		textView.SetBorderColor(tcell.ColorDimGray)
	}
	focusContent := func() {
		app.SetFocus(textView)
		list.SetBorderColor(tcell.ColorDimGray)
		textView.SetBorderColor(tcell.ColorMediumPurple)
	}

	// Layout: sidebar (fixed 24 cols) | content (flexible), footer at bottom.
	columns := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(list, 24, 0, true).
		AddItem(textView, 0, 1, false)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(columns, 0, 1, true).
		AddItem(footer, 1, 0, false)

	// Tab and left/right switch focus; q quits.
	layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab, tcell.KeyRight:
			if list.HasFocus() {
				focusContent()
			} else if event.Key() == tcell.KeyTab {
				focusSidebar()
			}
			return nil
		case tcell.KeyLeft:
			if textView.HasFocus() {
				focusSidebar()
			}
			return nil
		case tcell.KeyRune:
			if event.Rune() == 'q' {
				app.Stop()
				return nil
			}
		}
		return event
	})

	// Highlight the initially focused pane.
	list.SetBorderColor(tcell.ColorMediumPurple)

	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
