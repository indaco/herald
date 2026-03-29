// Project overview explorer using herald with bubbletea.
//
// A sidebar list (bubbles/list) lets you browse topics while a scrollable
// content pane (bubbles/viewport) shows herald-rendered output. Left/right
// arrows switch focus between panels, up/down scrolls the content pane or
// navigates the sidebar list depending on focus. Tab also toggles focus.
// Press q to quit.
//
// This example is a separate Go module with its own go.mod to keep
// charm.land/bubbletea/v2 out of herald's core dependencies.
//
// Run from the repository root:
//
//	cd examples/206_bubbletea-explorer && go run .
package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/indaco/herald"
)

// ---------------------------------------------------------------------------
// Styles
// ---------------------------------------------------------------------------

const sidebarWidth = 30

var (
	sidebarStyle = lipgloss.NewStyle().
			Width(sidebarWidth).
			Padding(1, 1)

	viewportStyle = lipgloss.NewStyle().
			Padding(1, 2)

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555"))

	focusedBorder = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7C3AED"))

	unfocusedBorder = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#555555"))

	footerKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A3A3A3"))

	footerValStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5"))
)

// ---------------------------------------------------------------------------
// Topic items
// ---------------------------------------------------------------------------

type topic struct {
	name string
	desc string
}

func (t topic) Title() string       { return t.name }
func (t topic) Description() string { return t.desc }
func (t topic) FilterValue() string { return t.name }

func topics() []list.Item {
	return []list.Item{
		topic{"Overview", "Introduction & highlights"},
		topic{"Installation", "Quick start guide"},
		topic{"Typography", "Headings, paragraphs & rules"},
		topic{"Lists", "Flat, nested & mixed lists"},
		topic{"Tables & Data", "Tables, badges & key-values"},
		topic{"Themes", "Built-in & custom themes"},
		topic{"Alerts & Inline", "Callouts & inline styles"},
	}
}

// ---------------------------------------------------------------------------
// Content builders - one per topic
// ---------------------------------------------------------------------------

// contentBuilder is a function that renders a topic at a given width.
type contentBuilder func(ty *herald.Typography, width int) string

// wrapText soft-wraps a string to the given width using lipgloss, which is
// ANSI-aware and preserves escape sequences across line breaks.
func wrapText(s string, width int) string {
	return lipgloss.NewStyle().Width(width).Render(s)
}

func buildOverview(ty *herald.Typography, width int) string {
	return ty.Compose(
		ty.H1("Herald"),
		wrapText(ty.P(
			"HTML-inspired typography for terminal UIs in Go, built on "+
				ty.Bold("lipgloss v2")+". Herald maps familiar element names "+
				"(H1\u2013H6, P, Blockquote, Code, HR, lists) to styled terminal output.",
		), width),
		ty.HR(),
		ty.H2("Highlights"),
		ty.UL(
			"Familiar HTML element names",
			"Built-in themes: Dracula, Catppuccin, Base16, Charm",
			"Pairs with huh, bubbletea, tview, and other TUI libraries",
			"Pluggable syntax highlighting via "+ty.Code("WithCodeFormatter"),
			"Custom palettes from just 9 colors",
		),
		ty.HR(),
		ty.Blockquote(
			"The integration point with bubbletea is "+ty.Code("viewport.SetContent()")+
				".\nHerald renders the styled text, bubbletea handles the interactivity.",
		),
	)
}

func buildInstallation(ty *herald.Typography, width int) string {
	return ty.Compose(
		ty.H1("Installation"),
		wrapText(ty.P("Requires Go 1.25 or later."), width),
		ty.CodeBlock("go get github.com/indaco/herald@latest", "sh"),
		ty.H2("Quick Start"),
		ty.CodeBlock(`package main

import (
    "fmt"
    "github.com/indaco/herald"
)

func main() {
    ty := herald.New()
    fmt.Println(ty.H1("Hello, Herald!"))
    fmt.Println(ty.P("Rich terminal typography, simply."))
    fmt.Println(ty.UL("Headings", "Lists", "Inline styles"))
}`, "go"),
		wrapText(ty.Tip("Press "+ty.Kbd("\u2190")+"/"+ty.Kbd("\u2192")+" to switch focus. Use "+ty.Kbd("\u2191")+"/"+ty.Kbd("\u2193")+" to scroll or navigate."), width),
	)
}

func buildTypography(ty *herald.Typography, width int) string {
	return ty.Compose(
		ty.H1("Typography Elements"),
		ty.H2("Headings"),
		wrapText(ty.P("H1\u2013H3 render with an underline; H4\u2013H6 with a left-bar prefix."), width),
		ty.H3("This is H3"),
		ty.H4("This is H4"),
		ty.H5("This is H5"),
		ty.H6("This is H6"),
		ty.HR(),
		ty.H2("Block Elements"),
		ty.Blockquote("Blockquotes indent with a left bar.\nThey support multi-line input."),
		wrapText(ty.P("Inline: "+ty.Code("os.Exit(1)")+" vs a fenced block below."), width),
		ty.CodeBlock(`func greet(name string) string {
    return "Hello, " + name
}`, "go"),
		ty.HR(),
		ty.H2("Inline Styles"),
		wrapText(ty.P(
			ty.Bold("Bold")+" \u00b7 "+
				ty.Italic("Italic")+" \u00b7 "+
				ty.Underline("Underline")+" \u00b7 "+
				ty.Strikethrough("Strikethrough")+" \u00b7 "+
				ty.Mark("Marked")+" \u00b7 "+
				ty.Small("Small"),
		), width),
		wrapText(ty.P(
			"Keyboard: "+ty.Kbd("Ctrl")+" + "+ty.Kbd("C")+"  |  "+
				"Link: "+ty.Link("pkg.go.dev", "https://pkg.go.dev/github.com/indaco/herald"),
		), width),
		ty.HRWithLabel("Horizontal Rules"),
		wrapText(ty.P("Plain "+ty.Code("HR()")+" and labeled "+ty.Code("HRWithLabel()")+":"), width),
		ty.HR(),
		ty.HRWithLabel("Section Break"),
		ty.Blockquote("The beginning is the most important part of the work.\n\u2014 Plato"),
	)
}

func buildLists(ty *herald.Typography, width int) string {
	tyH := herald.New(
		herald.WithTheme(herald.CatppuccinTheme()),
		herald.WithHierarchicalNumbers(true),
	)
	return ty.Compose(
		ty.H1("Lists"),
		ty.H2("Unordered"),
		ty.UL("Design", "Implement", "Test"),
		ty.H2("Ordered"),
		ty.OL("Plan", "Build", "Ship", "Iterate"),
		ty.H2("Nested \u2014 mixed sub-lists"),
		ty.NestUL(
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
		),
		ty.H2("Hierarchical Ordered"),
		tyH.NestOL(
			herald.Item("Introduction"),
			herald.ItemWithOLChildren("Core Concepts",
				herald.Item("Typography"),
				herald.Item("Themes"),
			),
			herald.Item("Conclusion"),
		),
		ty.HRWithLabel("Definition List"),
		ty.DL([][2]string{
			{"Typography", "The art of arranging type to make written language readable and appealing."},
			{"Theme", "A collection of styles that define the visual appearance of all elements."},
			{"Functional Option", "A pattern where configuration is passed as variadic function arguments."},
		}),
	)
}

func buildTablesAndData(ty *herald.Typography, width int) string {
	return ty.Compose(
		ty.H2("Tables & Data"),
		wrapText(ty.P("Herald renders bordered tables with header styling, optional striped rows, column alignment, and captions."), width),
		ty.TableWithOpts(
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
		),
		ty.HRWithLabel("Badges & Tags"),
		ty.P(
			ty.Badge("STABLE")+" "+ty.Tag("v1.0.0")+" "+ty.Badge("GO")+" "+ty.Tag("lipgloss-v2"),
		),
		wrapText(ty.P("Use "+ty.Code("Badge()")+" for bold status pills and "+ty.Code("Tag()")+" for subtle category labels."), width),
		ty.HRWithLabel("Key-Value Pairs"),
		ty.KVGroup([][2]string{
			{"Package", "github.com/indaco/herald"},
			{"Go version", "1.25+"},
			{"Dependencies", "lipgloss v2 (single dep)"},
			{"License", "MIT"},
			{"Test coverage", "95%+"},
		}),
	)
}

func buildThemes(ty *herald.Typography, width int) string {
	return ty.Compose(
		ty.H1("Themes"),
		wrapText(ty.P(
			"Herald ships with named themes that match "+
				ty.Bold("huh")+"'s built-in palettes. "+
				"Colors auto-adapt to light/dark terminal backgrounds.",
		), width),
		ty.H2("Built-in Themes"),
		ty.DL([][2]string{
			{"Rose Pine", "Default \u2014 warm, muted palette with dark and light variants"},
			{"Dracula", "Matches huh.ThemeDracula()"},
			{"Catppuccin", "Mocha (dark) / Latte (light)"},
			{"Base16", "ANSI base16 terminal colors"},
			{"Charm", "Charm brand colors"},
		}),
		ty.H2("Applying a Theme"),
		ty.CodeBlock(
			`ty := herald.New(herald.WithTheme(herald.DraculaTheme()))`, "go",
		),
		ty.H2("Custom Palette"),
		wrapText(ty.P(
			"Derive a full theme from 9 colors with "+ty.Code("ColorPalette")+". "+
				"Use "+ty.Code("lipgloss.LightDark")+" for adaptive colors that adjust "+
				"to the terminal background:",
		), width),
		ty.CodeBlock(`lightDark := lipgloss.LightDark(lipgloss.HasDarkBackground(os.Stdin, os.Stdout))

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
ty := herald.New(herald.WithPalette(palette))`, "go"),
		ty.Tip("All built-in themes are themselves built with "+ty.Code("ThemeFromPalette()")+"."),
	)
}

func buildAlertsAndInline(ty *herald.Typography, width int) string {
	return ty.Compose(
		ty.H2("Alerts & Inline Styles"),
		wrapText(ty.P("Herald provides five GitHub-style alert types:"), width),
		wrapText(ty.Note("Useful information that users should know, even when skimming."), width),
		wrapText(ty.Tip("Helpful advice for doing things better or more easily."), width),
		wrapText(ty.Important("Key information users need to know to achieve their goal."), width),
		wrapText(ty.Warning("Urgent info that needs immediate user attention to avoid problems."), width),
		wrapText(ty.Caution("Advises about risks or negative outcomes of certain actions."), width),
		ty.HRWithLabel("More Inline Styles"),
		wrapText(ty.P(
			"Inline code: "+ty.Code("fmt.Println()")+
				"  Keyboard: "+ty.Kbd("Ctrl")+"+"+ty.Kbd("C")+
				"  Highlight: "+ty.Mark("important"),
		), width),
		ty.P("Links: "+ty.Link("Herald", "https://github.com/indaco/herald")),
		ty.P(
			"Abbreviations: "+ty.Abbr("TUI", "Terminal User Interface")+
				", "+ty.Abbr("CLI", "Command-Line Interface"),
		),
		ty.P(
			"Subscript: H"+ty.Sub("2")+"O  "+
				"Superscript: E=mc"+ty.Sup("2"),
		),
		ty.HRWithLabel("Diff Markers"),
		ty.Del("typography.InlineCode(text)")+"\n"+ty.Ins("typography.Code(text, lang)"),
	)
}

// allBuilders returns the ordered list of content builder functions.
var allBuilders = []contentBuilder{
	buildOverview,
	buildInstallation,
	buildTypography,
	buildLists,
	buildTablesAndData,
	buildThemes,
	buildAlertsAndInline,
}

// ---------------------------------------------------------------------------
// Model
// ---------------------------------------------------------------------------

type pane int

const (
	sidebarPane pane = iota
	contentPane
)

type model struct {
	ty        *herald.Typography
	list      list.Model
	viewport  viewport.Model
	contents  []string
	focus     pane
	ready     bool
	lastIndex int
	lastWidth int
}

func initialModel(ty *herald.Typography) model {
	delegate := list.NewDefaultDelegate()
	delegate.SetSpacing(0)
	delegate.ShowDescription = true

	l := list.New(topics(), delegate, sidebarWidth-4, 0)
	l.Title = "Herald Explorer"
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return model{
		ty:        ty,
		list:      l,
		focus:     sidebarPane,
		lastIndex: -1,
	}
}

// renderContents rebuilds all topic content at the given width.
func (m *model) renderContents(width int) {
	m.contents = make([]string, len(allBuilders))
	for i, build := range allBuilders {
		m.contents[i] = build(m.ty, width)
	}
	m.lastWidth = width
}

// contentWidth returns the usable text width inside the viewport, accounting
// for borders, padding, and the divider column.
func contentWidth(termWidth int) int {
	// border (1+1) + viewportStyle padding (2+2) = 6
	return termWidth - sidebarWidth - 3 - 6
}

func (m model) Init() tea.Cmd {
	return nil
}

// handleKeyPress processes keyboard input and returns true with an optional
// command when the key was fully handled (no further processing needed).
func (m *model) handleKeyPress(key string) (bool, tea.Cmd) {
	switch key {
	case "q", "ctrl+c":
		return true, tea.Quit
	case "tab", "right":
		if m.focus == sidebarPane {
			m.focus = contentPane
		}
		return true, nil
	case "left":
		if m.focus == contentPane {
			m.focus = sidebarPane
		}
		return true, nil
	default:
		return false, nil
	}
}

// handleWindowSize updates layout dimensions and re-renders content when the
// terminal is resized.
func (m *model) handleWindowSize(msg tea.WindowSizeMsg) {
	vpW := contentWidth(msg.Width)
	vpHeight := msg.Height - 6 // padding (2) + border (2) + newline (1) + footer (1)

	if !m.ready {
		m.viewport = viewport.New(
			viewport.WithWidth(vpW+6), // add padding back for viewport frame
			viewport.WithHeight(vpHeight),
		)
		m.viewport.MouseWheelEnabled = true
		m.list.SetHeight(vpHeight)
		m.renderContents(vpW)
		m.ready = true
		return
	}

	m.viewport.SetWidth(vpW + 6)
	m.viewport.SetHeight(vpHeight)
	m.list.SetHeight(vpHeight)
	if vpW != m.lastWidth {
		m.renderContents(vpW)
		// Re-apply current content at new width.
		idx := m.list.Index()
		if idx >= 0 && idx < len(m.contents) {
			m.viewport.SetContent(m.contents[idx])
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if handled, cmd := m.handleKeyPress(msg.String()); handled {
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.handleWindowSize(msg)
	}

	// Route input to the focused pane.
	if m.focus == sidebarPane {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	// Update viewport content when selection changes.
	idx := m.list.Index()
	if idx != m.lastIndex && idx >= 0 && idx < len(m.contents) {
		m.viewport.SetContent(m.contents[idx])
		m.viewport.GotoTop()
		m.lastIndex = idx
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() tea.View {
	if !m.ready {
		v := tea.NewView("\n  Initializing...")
		v.AltScreen = true
		return v
	}

	// Apply focus-aware borders.
	var sidebarBorder, contentBorder lipgloss.Style
	if m.focus == sidebarPane {
		sidebarBorder = focusedBorder
		contentBorder = unfocusedBorder
	} else {
		sidebarBorder = unfocusedBorder
		contentBorder = focusedBorder
	}

	sidebar := sidebarBorder.Render(sidebarStyle.Render(m.list.View()))
	content := contentBorder.Render(viewportStyle.Render(m.viewport.View()))
	divider := dividerStyle.Render(strings.Repeat("│\n", max(lipgloss.Height(sidebar), 1)))

	body := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, divider, content)

	// Footer.
	pct := fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100)
	footer := footerKeyStyle.Render("  \u2190/\u2192") +
		footerValStyle.Render(" switch  ") +
		footerKeyStyle.Render("\u2191/\u2193") +
		footerValStyle.Render(" navigate  ") +
		footerKeyStyle.Render("q") +
		footerValStyle.Render(" quit  ") +
		footerKeyStyle.Render(pct)

	v := tea.NewView(body + "\n" + footer)
	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion
	return v
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	ty := herald.New(herald.WithTheme(herald.CatppuccinTheme()))
	p := tea.NewProgram(initialModel(ty))

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
