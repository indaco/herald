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
	Ins           lipgloss.Style
	Del           lipgloss.Style

	// Definition list
	DT lipgloss.Style // definition term
	DD lipgloss.Style // definition description

	// Address element
	Address           lipgloss.Style // style for address/contact blocks
	AddressCard       lipgloss.Style // content style for bordered address card
	AddressCardBorder lipgloss.Style // border color/style for address card (same pattern as TableBorder)

	// Badge / Tag elements
	Badge lipgloss.Style // style for bold pill/status labels
	Tag   lipgloss.Style // style for subtle pill/category labels

	// Footnote elements
	FootnoteRef     lipgloss.Style // style for inline reference markers (e.g. "[1]")
	FootnoteItem    lipgloss.Style // style for each footnote entry in the section
	FootnoteDivider lipgloss.Style // style for the divider line above the footnote section

	// Callbacks
	CodeFormatter func(code, language string) string // optional syntax highlighter

	// Heading decoration
	H1UnderlineChar string // character repeated under H1 (e.g. "═")
	H2UnderlineChar string // character repeated under H2 (e.g. "─")
	H3UnderlineChar string // character repeated under H3 (e.g. "·")
	HeadingBarChar  string // left-bar prefix for H4-H6 (e.g. "▎")

	// Code block line numbers
	CodeLineNumber    lipgloss.Style // style for line number text
	CodeLineNumberSep string         // separator between line number and code (e.g. "│")
	ShowLineNumbers   bool           // whether to render line numbers in CodeBlock

	// Configurable tokens
	BulletChar           string   // character used for unordered list bullets
	NestedBulletChars    []string // bullet chars cycling per depth for nested lists
	ListIndent           int      // spaces per nesting level for nested lists
	HierarchicalNumbers  bool     // use hierarchical numbering (e.g. 2.1, 2.2) for nested OL
	HRChar               string   // character repeated for horizontal rules
	HRWidth              int      // width of horizontal rule in characters
	BlockquoteBar        string   // left-bar character for blockquotes
	InsPrefix            string   // prefix for inserted text (e.g. "+")
	DelPrefix            string   // prefix for deleted text (e.g. "-")
	FootnoteDividerChar  string   // character for footnote section divider (default "─")
	FootnoteDividerWidth int      // width of footnote divider (default 20)

	// Table elements
	TableHeader      lipgloss.Style // style for header cell text (e.g. bold + primary color)
	TableCell        lipgloss.Style // style for body cell text
	TableStripedCell lipgloss.Style // style for alternating (odd) body rows when striping is enabled
	TableFooter      lipgloss.Style // style for footer row cells
	TableCaption     lipgloss.Style // style for table caption text
	TableBorder      lipgloss.Style // style (color) applied to border characters
	TableBorderSet   TableBorderSet // box-drawing character set
	TableCellPad     int            // spaces of padding inside each cell (default 1)

	// Alert elements
	Alerts   map[AlertType]AlertConfig // per-type icon, label, and style
	AlertBar string                    // left-bar character for alerts
}

// TableBorderSet holds all box-drawing characters needed to render a table.
type TableBorderSet struct {
	Top            string // horizontal line for top border
	Bottom         string // horizontal line for bottom border
	Left           string // vertical line for left border
	Right          string // vertical line for right border
	Header         string // horizontal line separating header from body
	Row            string // horizontal line between rows (optional, empty = no row separators)
	TopLeft        string // top-left corner
	TopRight       string // top-right corner
	BottomLeft     string // bottom-left corner
	BottomRight    string // bottom-right corner
	TopJunction    string // top edge junction (┬)
	BottomJunction string // bottom edge junction (┴)
	LeftJunction   string // left edge junction (├)
	RightJunction  string // right edge junction (┤)
	Cross          string // interior junction (┼)
	HeaderLeft     string // header-row left junction
	HeaderRight    string // header-row right junction
	HeaderCross    string // header-row interior junction
	FooterLeft     string // footer-row left junction
	FooterRight    string // footer-row right junction
	FooterCross    string // footer-row interior junction
}

// BoxBorderSet returns a TableBorderSet using full Unicode box-drawing characters.
func BoxBorderSet() TableBorderSet {
	return TableBorderSet{
		Top:            "─",
		Bottom:         "─",
		Left:           "│",
		Right:          "│",
		Header:         "─",
		Row:            "─",
		TopLeft:        "┌",
		TopRight:       "┐",
		BottomLeft:     "└",
		BottomRight:    "┘",
		TopJunction:    "┬",
		BottomJunction: "┴",
		LeftJunction:   "├",
		RightJunction:  "┤",
		Cross:          "┼",
		HeaderLeft:     "├",
		HeaderRight:    "┤",
		HeaderCross:    "┼",
		FooterLeft:     "├",
		FooterRight:    "┤",
		FooterCross:    "┼",
	}
}

// MinimalBorderSet returns a TableBorderSet with no outer borders — only column
// separators and a header underline.
func MinimalBorderSet() TableBorderSet {
	return TableBorderSet{
		Left:        "",
		Right:       "",
		Header:      "─",
		Row:         "─",
		HeaderLeft:  "",
		HeaderRight: "",
		HeaderCross: "┼",
		FooterLeft:  "",
		FooterRight: "",
		FooterCross: "┼",
		Cross:       "┼",
	}
}

// Default token values used by DefaultTheme and ThemeFromPalette.
const (
	DefaultH1UnderlineChar      = "═"
	DefaultH2UnderlineChar      = "─"
	DefaultH3UnderlineChar      = "·"
	DefaultHeadingBarChar       = "▎"
	DefaultBulletChar           = "•"
	DefaultListIndent           = 2
	DefaultHRChar               = "─"
	DefaultHRWidth              = 40
	DefaultBlockquoteBar        = "│"
	DefaultAlertBar             = "│"
	DefaultCodeLineNumberSep    = "│"
	DefaultTableCellPad         = 1
	DefaultInsPrefix            = "+"
	DefaultDelPrefix            = "-"
	DefaultFootnoteDividerChar  = "─"
	DefaultFootnoteDividerWidth = 20
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
		Muted:     lightDark(lipgloss.Color("#9893A5"), lipgloss.Color("#6E6A86")), // subtle
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
