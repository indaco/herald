package herald

import (
	"os"
	"sync"

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
	Paragraph          lipgloss.Style
	Blockquote         lipgloss.Style
	BlockquoteBarStyle lipgloss.Style // style for the blockquote bar character(s)
	CodeInline         lipgloss.Style
	CodeBlock          lipgloss.Style
	HR                 lipgloss.Style
	HRLabel            lipgloss.Style // style for the label text within a labeled separator

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
	Q             lipgloss.Style // quotation mark + text styling
	Cite          lipgloss.Style // citation (italic + muted)
	Samp          lipgloss.Style // sample output (monospace, distinct from CodeInline)
	Var           lipgloss.Style // variable name (italic monospace)

	// Definition list
	DT lipgloss.Style // definition term
	DD lipgloss.Style // definition description

	// Key-value elements
	KVKey       lipgloss.Style // style for key text in key-value pairs
	KVValue     lipgloss.Style // style for value text in key-value pairs
	KVSeparator string         // separator between key and value (default ":")

	// Address element
	Address           lipgloss.Style // style for address/contact blocks
	AddressCard       lipgloss.Style // content style for bordered address card
	AddressCardBorder lipgloss.Style // border color/style for address card (same pattern as TableBorder)

	// Badge / Tag elements
	Badge lipgloss.Style // style for bold pill/status labels
	Tag   lipgloss.Style // style for subtle pill/category labels

	// Semantic badge/tag styles (derived from SemanticPalette)
	SuccessBadge lipgloss.Style
	WarningBadge lipgloss.Style
	ErrorBadge   lipgloss.Style
	InfoBadge    lipgloss.Style

	SuccessTag lipgloss.Style
	WarningTag lipgloss.Style
	ErrorTag   lipgloss.Style
	InfoTag    lipgloss.Style

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
	CodeLineNumber       lipgloss.Style // style for line number text
	CodeLineNumberSep    string         // separator between line number and code (e.g. "│")
	ShowLineNumbers      bool           // whether to render line numbers in CodeBlock
	CodeLineNumberOffset int            // starting line number for code blocks (default 1)

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
	QuoteOpen            string   // opening quotation mark for Q (default "\u201C")
	QuoteClose           string   // closing quotation mark for Q (default "\u201D")
	FootnoteDividerChar  string   // character for footnote section divider (default "─")
	FootnoteDividerWidth int      // width of footnote divider (default 20)

	// Figure elements
	FigureCaption         lipgloss.Style  // style for figure caption text
	FigureCaptionPosition CaptionPosition // default CaptionBottom

	// Fieldset elements
	Fieldset       lipgloss.Style // content text style
	FieldsetBorder lipgloss.Style // border characters style
	FieldsetLegend lipgloss.Style // legend text style
	FieldsetWidth  int            // default 0 = auto-fit

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
	ColumnSep      string // vertical separator between columns (defaults to "│")
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
		ColumnSep:      "│",
	}
}

// MinimalBorderSet returns a TableBorderSet with no outer borders - only column
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
		ColumnSep:   "│",
	}
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

	DefaultAlertBar             = "│"
	DefaultCodeLineNumberSep    = "│"
	DefaultCodeLineNumberOffset = 1
	DefaultTableCellPad         = 1
	DefaultInsPrefix            = "+"
	DefaultDelPrefix            = "-"
	DefaultFootnoteDividerChar  = "─"
	DefaultFootnoteDividerWidth = 20
	DefaultKVSeparator          = ":"
	DefaultQuoteOpen            = "\u201C" // left curly double quote
	DefaultQuoteClose           = "\u201D" // right curly double quote
	MaxWidthChars               = 500
)

// DefaultNestedBulletChars returns the default set of bullet characters that
// cycle through nesting levels in nested unordered lists.
// A fresh slice is returned on each call to prevent mutation of shared state.
func DefaultNestedBulletChars() []string {
	return []string{"•", "◦", "▪", "▹"}
}

var (
	termBGOnce sync.Once
	termBGDark bool
)

// hasDarkBG returns whether the terminal has a dark background.
// The terminal query is cached to avoid repeated I/O, but the
// HERALD_FORCE_DARK env var is always checked first so it can be
// changed between calls (e.g. in tests).
func hasDarkBG() bool {
	if v := os.Getenv("HERALD_FORCE_DARK"); v != "" {
		return v == "1" || v == "true"
	}
	termBGOnce.Do(func() {
		termBGDark = lipgloss.HasDarkBackground(os.Stdin, os.Stdout)
	})
	return termBGDark
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

	// Semantic palette: Rose Pine foam/gold/love/iris
	rosePineSP := SemanticPalette{
		Success: lightDark(lipgloss.Color("#286983"), lipgloss.Color("#9CCFD8")), // foam/pine (green)
		Warning: lightDark(lipgloss.Color("#D7827E"), lipgloss.Color("#F6C177")), // gold (amber)
		Error:   lightDark(lipgloss.Color("#B4637A"), lipgloss.Color("#EB6F92")), // love (red)
		Info:    lightDark(lipgloss.Color("#3e8fb0"), lipgloss.Color("#9CCFD8")), // foam (blue)
	}
	base := theme.Badge.GetForeground()
	surface := theme.Tag.GetBackground()
	theme.SuccessBadge, theme.WarningBadge, theme.ErrorBadge, theme.InfoBadge = defaultSemanticBadgeStyles(rosePineSP, base)
	theme.SuccessTag, theme.WarningTag, theme.ErrorTag, theme.InfoTag = defaultSemanticTagStyles(rosePineSP, surface)

	return theme
}
