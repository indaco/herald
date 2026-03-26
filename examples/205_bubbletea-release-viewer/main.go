// Scrollable release notes viewer using herald with bubbletea.
//
// This example is a separate Go module with its own go.mod to keep
// charm.land/bubbletea/v2 out of herald's core dependencies.
//
// Run from the repository root:
//
//	cd examples/205_bubbletea-release-viewer && go run .
package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/indaco/herald"
)

// Styles for header and footer bars.
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#7C3AED")).
			Padding(0, 1)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#4C1D95")).
			Padding(0, 1)

	footerKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A3A3A3"))

	footerValStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5"))
)

type model struct {
	viewport viewport.Model
	content  string
	ready    bool
}

func initialModel(content string) model {
	return model{content: content}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := 1
		footerHeight := 1
		verticalMargin := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(
				viewport.WithWidth(msg.Width),
				viewport.WithHeight(msg.Height-verticalMargin),
			)
			m.viewport.MouseWheelEnabled = true
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.SetWidth(msg.Width)
			m.viewport.SetHeight(msg.Height - verticalMargin)
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() tea.View {
	if !m.ready {
		v := tea.NewView("\n  Initializing...")
		v.AltScreen = true
		return v
	}

	header := lipgloss.JoinHorizontal(lipgloss.Center,
		titleStyle.Render("Herald"),
		infoStyle.Render("Release Notes Viewer"),
	)

	pct := fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100)
	footer := footerKeyStyle.Render("  q") +
		footerValStyle.Render(" quit  ") +
		footerKeyStyle.Render("\u2191/\u2193") +
		footerValStyle.Render(" scroll  ") +
		footerKeyStyle.Render("pgup/pgdn") +
		footerValStyle.Render(" page  ") +
		footerKeyStyle.Render(pct)

	// Pad header and footer to viewport width.
	headerGap := strings.Repeat(" ", max(0, m.viewport.Width()-lipgloss.Width(header)))
	footerGap := strings.Repeat(" ", max(0, m.viewport.Width()-lipgloss.Width(footer)))

	content := header + headerGap + "\n" +
		m.viewport.View() + "\n" +
		footer + footerGap

	v := tea.NewView(content)
	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion
	return v
}

// buildReleaseNotes assembles realistic release notes using herald typography.
func buildReleaseNotes(ty *herald.Typography) string {
	var b strings.Builder

	// Project heading and version badge.
	b.WriteString(ty.H1("Herald v1.0.0"))
	b.WriteString("\n\n")

	b.WriteString(ty.Tag("v1.0.0") + "  " + ty.Badge("STABLE") + "  " + ty.Tag("go1.25+"))
	b.WriteString("\n\n")

	// Release metadata.
	b.WriteString(ty.KVGroup([][2]string{
		{"Release Date", "2026-03-23"},
		{"Author", "indaco"},
		{"Commits", "47"},
		{"License", "MIT"},
	}))
	b.WriteString("\n\n")

	// Opening quote.
	b.WriteString(ty.Blockquote("Typography matters. Herald brings HTML-grade text styling to\nyour terminal, so every heading, list, and callout looks intentional."))
	b.WriteString("\n\n")

	// --- Features section ---
	b.WriteString(ty.HRWithLabel("Features"))
	b.WriteString("\n\n")

	b.WriteString(ty.H3("Typography Elements"))
	b.WriteString("\n\n")

	b.WriteString(ty.P("This release introduces the full suite of typographic primitives:"))
	b.WriteString("\n\n")

	b.WriteString(ty.UL(
		"Six heading levels ("+ty.Code("H1", "")+" through "+ty.Code("H6", "")+") with configurable underlines and bar prefixes",
		"Paragraphs, blockquotes, and horizontal rules with labeled variants",
		"Ordered and unordered lists with nested sub-items and hierarchical numbering",
		"Fenced code blocks with optional line numbers and syntax highlighting callbacks",
		"Inline styles: "+ty.Bold("bold")+", "+ty.Italic("italic")+", "+ty.Code("code", "")+", and "+ty.Kbd("Kbd")+" indicators",
		"Definition lists, key-value pairs, and address cards",
		"GitHub-style alerts: Note, Tip, Important, Warning, and Caution",
	))
	b.WriteString("\n\n")

	b.WriteString(ty.H3("Themes"))
	b.WriteString("\n\n")

	b.WriteString(ty.P("Herald ships with several built-in themes that adapt to your terminal background:"))
	b.WriteString("\n\n")

	b.WriteString(ty.OL(
		"Rose Pine (default) -- a warm, muted palette inspired by the Rose Pine project",
		"Catppuccin -- the pastel theme used in this very demo"+ty.FootnoteRef(1),
		"Dracula -- a dark-mode staple with vibrant accent colors",
		"Nord -- cool, arctic-inspired blue tones",
		"Tokyo Night -- soft purples and blues from the Tokyo Night family",
	))
	b.WriteString("\n\n")

	b.WriteString(ty.P("Custom themes are supported via " + ty.Code("ThemeFromPalette()", "") + " or full " + ty.Code("WithTheme()", "") + " overrides."))
	b.WriteString("\n\n")

	b.WriteString(ty.H3("Badge and Tag System"))
	b.WriteString("\n\n")

	b.WriteString(ty.P(
		"New pill-style labels for status indicators and categories. " +
			"Use " + ty.Code("Badge()", "") + " for bold status pills and " + ty.Code("Tag()", "") + " for subtle category markers. " +
			"Both support one-off style overrides via " + ty.Code("BadgeWithStyle()", "") + " and " + ty.Code("TagWithStyle()", "") + "."),
	)
	b.WriteString("\n\n")

	b.WriteString(ty.H3("Footnote System"))
	b.WriteString("\n\n")

	b.WriteString(ty.P(
		"Inline footnote references" + ty.FootnoteRef(2) + " paired with a collected footnote section " +
			"at the bottom of your document. References are rendered as bracketed numbers " +
			"and the section includes a configurable divider line."),
	)
	b.WriteString("\n\n")

	// --- Breaking Changes section ---
	b.WriteString(ty.HRWithLabel("Breaking Changes"))
	b.WriteString("\n\n")

	b.WriteString(ty.Warning(
		"This release includes API changes that require updates to existing code.\nPlease review the migration guide below before upgrading."))
	b.WriteString("\n\n")

	b.WriteString(ty.P("The following functions have been renamed for consistency:"))
	b.WriteString("\n\n")

	b.WriteString(ty.Del("typography.InlineCode(text)"))
	b.WriteString("\n")
	b.WriteString(ty.Ins("typography.Code(text, lang)"))
	b.WriteString("\n\n")

	b.WriteString(ty.Del("typography.Separator()"))
	b.WriteString("\n")
	b.WriteString(ty.Ins("typography.HR()"))
	b.WriteString("\n\n")

	b.WriteString(ty.Del("typography.LabeledSeparator(label)"))
	b.WriteString("\n")
	b.WriteString(ty.Ins("typography.HRWithLabel(label)"))
	b.WriteString("\n\n")

	b.WriteString(ty.P(
		"The " + ty.Code("Table()", "") + " method now accepts " + ty.Code("[][]string", "") + " instead of " +
			"a custom struct. Use " + ty.Code("TableWithOpts()", "") + " for column alignment, " +
			"striped rows, and captions."))
	b.WriteString("\n\n")

	// --- Migration section ---
	b.WriteString(ty.HRWithLabel("Migration"))
	b.WriteString("\n\n")

	b.WriteString(ty.P("Update your import path and adjust renamed methods:"))
	b.WriteString("\n\n")

	b.WriteString(ty.CodeBlock(`// Before
ty := herald.New()
fmt.Println(ty.InlineCode("go build"))
fmt.Println(ty.Separator())

// After
ty := herald.New()
fmt.Println(ty.Code("go build", ""))
fmt.Println(ty.HR())`, "go"))
	b.WriteString("\n\n")

	b.WriteString(ty.Tip("Run " + ty.Code("go vet ./...", "") + " after migrating to catch any missed renames at compile time."))
	b.WriteString("\n\n")

	// --- Performance section ---
	b.WriteString(ty.HRWithLabel("Performance"))
	b.WriteString("\n\n")

	b.WriteString(ty.P("Key performance characteristics of this release:"))
	b.WriteString("\n\n")

	b.WriteString(ty.KVGroup([][2]string{
		{"Render overhead", "< 2ms for a full-page document"},
		{"Allocations", "Zero allocations for cached theme lookups"},
		{"Binary size", "No CGo, single dependency (lipgloss v2)"},
		{"Startup time", "< 1ms to construct a Typography instance"},
	}))
	b.WriteString("\n\n")

	b.WriteString(ty.Note("Herald has no runtime dependencies beyond " + ty.Code("lipgloss/v2", "") + ". " +
		"All rendering is pure Go string manipulation."))
	b.WriteString("\n\n")

	// --- Acknowledgments section ---
	b.WriteString(ty.HRWithLabel("Acknowledgments"))
	b.WriteString("\n\n")

	b.WriteString(ty.Blockquote("A heartfelt thank you to the Charm community for lipgloss,\nbubbletea, and the entire ecosystem that makes TUI development a joy."))
	b.WriteString("\n\n")

	b.WriteString(ty.P(
		"Special thanks to all contributors who reported issues, " +
			"submitted pull requests, and tested pre-release builds. " +
			"Press " + ty.Kbd("q") + " to exit this viewer."))
	b.WriteString("\n\n")

	// --- Footnotes ---
	b.WriteString(ty.FootnoteSection([]string{
		"Catppuccin theme: https://github.com/catppuccin",
		"The footnote system was inspired by academic paper conventions for terminal documentation.",
	}))
	b.WriteString("\n\n")

	b.WriteString(ty.HR())
	b.WriteString("\n")

	return b.String()
}

func main() {
	ty := herald.New(herald.WithTheme(herald.CatppuccinTheme()))
	content := buildReleaseNotes(ty)

	p := tea.NewProgram(initialModel(content))

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
