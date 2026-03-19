package herald

import (
	"os"

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

	// Callbacks
	CodeFormatter func(code, language string) string // optional syntax highlighter

	// Heading decoration
	H1UnderlineChar string // character repeated under H1 (e.g. "═")
	H2UnderlineChar string // character repeated under H2 (e.g. "─")
	H3UnderlineChar string // character repeated under H3 (e.g. "·")
	HeadingBarChar  string // left-bar prefix for H4-H6 (e.g. "▎")

	// Configurable tokens
	BulletChar          string   // character used for unordered list bullets
	NestedBulletChars   []string // bullet chars cycling per depth for nested lists
	ListIndent          int      // spaces per nesting level for nested lists
	HierarchicalNumbers bool     // use hierarchical numbering (e.g. 2.1, 2.2) for nested OL
	HRChar              string   // character repeated for horizontal rules
	HRWidth             int      // width of horizontal rule in characters
	BlockquoteBar       string   // left-bar character for blockquotes

	// Alert elements
	Alerts   map[AlertType]AlertConfig // per-type icon, label, and style
	AlertBar string                    // left-bar character for alerts
}

// Default token values used by DefaultTheme and ThemeFromPalette.
const (
	DefaultH1UnderlineChar = "═"
	DefaultH2UnderlineChar = "─"
	DefaultH3UnderlineChar = "·"
	DefaultHeadingBarChar  = "▎"
	DefaultBulletChar      = "•"
	DefaultListIndent      = 2
	DefaultHRChar          = "─"
	DefaultHRWidth         = 40
	DefaultBlockquoteBar   = "│"
	DefaultAlertBar        = "│"
)

// DefaultNestedBulletChars is the default set of bullet characters that
// cycle through nesting levels in nested unordered lists.
var DefaultNestedBulletChars = []string{"•", "◦", "▪", "▹"}

// hasDarkBG returns whether the terminal has a dark background.
// It respects the HERALD_FORCE_DARK env var for tooling (e.g. screenshots).
func hasDarkBG() bool {
	if v := os.Getenv("HERALD_FORCE_DARK"); v != "" {
		return v == "1" || v == "true"
	}
	return lipgloss.HasDarkBackground(os.Stdin, os.Stdout)
}

// DefaultTheme returns a Theme based on the Rose Pine color palette.
func DefaultTheme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	theme := ThemeFromPalette(ColorPalette{
		Primary:   lightDark(lipgloss.Color("#286983"), lipgloss.Color("#E0DEF4")), // pine / text
		Secondary: lightDark(lipgloss.Color("#7c6f93"), lipgloss.Color("#C4A7E7")), // iris (deeper)
		Tertiary:  lightDark(lipgloss.Color("#3e8fb0"), lipgloss.Color("#9CCFD8")), // foam (deeper)
		Accent:    lightDark(lipgloss.Color("#D7827E"), lipgloss.Color("#F6C177")), // rose / gold
		Highlight: lightDark(lipgloss.Color("#B4637A"), lipgloss.Color("#EA9A97")), // love
		Muted:     lightDark(lipgloss.Color("#797593"), lipgloss.Color("#6E6A86")), // subtle
		Text:      lightDark(lipgloss.Color("#575279"), lipgloss.Color("#E0DEF4")), // text
		Surface:   lightDark(lipgloss.Color("#DFDAD9"), lipgloss.Color("#393552")), // overlay (darker)
		Base:      lightDark(lipgloss.Color("#FAF4ED"), lipgloss.Color("#191724")), // base
	})

	// Override alerts with Rose Pine-specific colors:
	// Note=foam/blue, Tip=pine/green, Important=iris/purple, Warning=gold/amber, Caution=love/red
	theme.Alerts = DefaultAlertConfigs(AlertPalette{
		Note:      lightDark(lipgloss.Color("#3e8fb0"), lipgloss.Color("#9CCFD8")), // foam
		Tip:       lightDark(lipgloss.Color("#286983"), lipgloss.Color("#31748F")), // pine
		Important: lightDark(lipgloss.Color("#7c6f93"), lipgloss.Color("#C4A7E7")), // iris
		Warning:   lightDark(lipgloss.Color("#D7827E"), lipgloss.Color("#F6C177")), // gold
		Caution:   lightDark(lipgloss.Color("#B4637A"), lipgloss.Color("#EB6F92")), // love
	})

	return theme
}
