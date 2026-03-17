package herald

import (
	"charm.land/lipgloss/v2"
)

// Theme holds all style definitions used by Typography to render elements.
// Each field corresponds to a visual element and is a lipgloss.Style.
type Theme struct {
	// Headings H1-H6
	H1 lipgloss.Style
	H2 lipgloss.Style
	H3 lipgloss.Style
	H4 lipgloss.Style
	H5 lipgloss.Style
	H6 lipgloss.Style

	// Text block elements
	Paragraph  lipgloss.Style
	Blockquote lipgloss.Style
	CodeInline lipgloss.Style
	CodeBlock  lipgloss.Style
	HR         lipgloss.Style

	// List elements
	ListBullet lipgloss.Style // style for the bullet/number marker
	ListItem   lipgloss.Style // style for list item text

	// Inline styles
	Bold          lipgloss.Style
	Italic        lipgloss.Style
	Underline     lipgloss.Style
	Strikethrough lipgloss.Style
	Small         lipgloss.Style
	Mark          lipgloss.Style
	Link          lipgloss.Style
	Kbd           lipgloss.Style
	Abbr          lipgloss.Style
	Sub           lipgloss.Style
	Sup           lipgloss.Style

	// Definition list
	DT lipgloss.Style // definition term
	DD lipgloss.Style // definition description

	// Heading decoration
	H1UnderlineChar string // character repeated under H1 (e.g. "═")
	H2UnderlineChar string // character repeated under H2 (e.g. "─")
	H3UnderlineChar string // character repeated under H3 (e.g. "·")
	HeadingBarChar  string // left-bar prefix for H4-H6 (e.g. "▎")

	// Configurable tokens
	BulletChar    string // character used for unordered list bullets
	HRChar        string // character repeated for horizontal rules
	HRWidth       int    // width of horizontal rule in characters
	BlockquoteBar string // left-bar character for blockquotes
}

// DefaultTheme returns a Theme with sensible default styles that look great
// in most terminal environments.
func DefaultTheme() Theme {
	return Theme{
		H1: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#E0DEF4")).
			MarginBottom(1),

		H2: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#C4A7E7")).
			MarginBottom(1),

		H3: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#9CCFD8")).
			MarginBottom(1),

		H4: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F6C177")).
			MarginBottom(1),

		H5: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#EA9A97")).
			MarginBottom(1),

		H6: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#908CAA")).
			MarginBottom(1),

		Paragraph: lipgloss.NewStyle().
			MarginBottom(1),

		Blockquote: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#908CAA")).
			Italic(true).
			PaddingLeft(2),

		CodeInline: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E0DEF4")).
			Background(lipgloss.Color("#393552")),

		CodeBlock: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E0DEF4")).
			Background(lipgloss.Color("#2A273F")).
			Padding(1, 2).
			MarginBottom(1),

		HR: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6E6A86")),

		ListBullet: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#C4A7E7")),

		ListItem: lipgloss.NewStyle(),

		Bold: lipgloss.NewStyle().
			Bold(true),

		Italic: lipgloss.NewStyle().
			Italic(true),

		Underline: lipgloss.NewStyle().
			Underline(true),

		Strikethrough: lipgloss.NewStyle().
			Strikethrough(true),

		Small: lipgloss.NewStyle().
			Faint(true),

		Mark: lipgloss.NewStyle().
			Background(lipgloss.Color("#F6C177")).
			Foreground(lipgloss.Color("#191724")),

		Link: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9CCFD8")).
			Underline(true),

		Kbd: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E0DEF4")).
			Background(lipgloss.Color("#393552")).
			Bold(true).
			Padding(0, 1),

		Abbr: lipgloss.NewStyle().
			Underline(true).
			Foreground(lipgloss.Color("#EA9A97")),

		Sub: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#908CAA")),

		Sup: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#908CAA")),

		DT: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#E0DEF4")),

		DD: lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color("#908CAA")),

		H1UnderlineChar: "═",
		H2UnderlineChar: "─",
		H3UnderlineChar: "·",
		HeadingBarChar:  "▎",

		BulletChar:    "\u2022",
		HRChar:        "\u2500",
		HRWidth:       40,
		BlockquoteBar: "\u2502",
	}
}
