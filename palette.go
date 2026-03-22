package herald

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// ColorPalette defines a minimal set of colors from which a full Theme can
// be derived. This allows users to share a single color palette across
// herald and other Charm ecosystem libraries (e.g. huh).
type ColorPalette struct {
	Primary   color.Color // main text, H1 headings
	Secondary color.Color // H2, list bullets, accents
	Tertiary  color.Color // H3, links
	Accent    color.Color // H4, marks/highlights
	Highlight color.Color // H5, abbreviations
	Muted     color.Color // H6, comments, sub/sup, blockquote, DD, HR
	Text      color.Color // default body text, paragraphs, list items, DT
	Surface   color.Color // background for kbd
	Base      color.Color // background for inline code, code blocks, mark foreground
}

// ThemeFromPalette constructs a complete Theme by mapping palette colors to
// all style fields. Configurable tokens use the same defaults as DefaultTheme.
func ThemeFromPalette(p ColorPalette) Theme {
	return Theme{
		H1: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Primary).
			MarginBottom(1),

		H2: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Secondary).
			MarginBottom(1),

		H3: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Tertiary).
			MarginBottom(1),

		H4: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Accent).
			MarginBottom(1),

		H5: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Highlight).
			MarginBottom(1),

		H6: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Muted).
			MarginBottom(1),

		Paragraph: lipgloss.NewStyle().
			MarginBottom(1),

		Blockquote: lipgloss.NewStyle().
			Foreground(p.Muted).
			Italic(true),

		BlockquoteBarStyle: lipgloss.NewStyle().
			Foreground(p.Muted).
			PaddingLeft(1),

		CodeInline: lipgloss.NewStyle().
			Foreground(p.Text).
			Background(p.Base),

		CodeBlock: lipgloss.NewStyle().
			Foreground(p.Text).
			Background(p.Base).
			Padding(1, 2).
			MarginBottom(1),

		HR: lipgloss.NewStyle().
			Foreground(p.Muted),

		HRLabel: lipgloss.NewStyle().
			Foreground(p.Muted),

		ListBullet: lipgloss.NewStyle().
			Foreground(p.Secondary),

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
			Background(p.Accent).
			Foreground(p.Base),

		Link: lipgloss.NewStyle().
			Foreground(p.Tertiary).
			Underline(true),

		Kbd: lipgloss.NewStyle().
			Foreground(p.Text).
			Background(p.Surface).
			Bold(true).
			Padding(0, 1),

		Abbr: lipgloss.NewStyle().
			Underline(true).
			Foreground(p.Highlight),

		Sub: lipgloss.NewStyle().
			Foreground(p.Muted),

		Sup: lipgloss.NewStyle().
			Foreground(p.Muted),

		Ins: lipgloss.NewStyle().
			Foreground(p.Tertiary),

		Del: lipgloss.NewStyle().
			Foreground(p.Highlight).
			Strikethrough(true),

		DT: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Text),

		DD: lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(p.Muted),

		Address: lipgloss.NewStyle().
			Foreground(p.Muted).
			Italic(true).
			PaddingLeft(2),

		AddressCard: lipgloss.NewStyle().
			Foreground(p.Muted).
			Italic(true),
		AddressCardBorder: lipgloss.NewStyle().
			Foreground(p.Muted),

		Badge: lipgloss.NewStyle().
			Background(p.Secondary).
			Foreground(p.Base).
			Bold(true).
			Padding(0, 1),

		Tag: lipgloss.NewStyle().
			Foreground(p.Secondary).
			Background(p.Surface).
			Padding(0, 1),

		FootnoteRef:     lipgloss.NewStyle().Foreground(p.Tertiary),
		FootnoteItem:    lipgloss.NewStyle().Foreground(p.Muted),
		FootnoteDivider: lipgloss.NewStyle().Foreground(p.Muted),

		H1UnderlineChar: DefaultH1UnderlineChar,
		H2UnderlineChar: DefaultH2UnderlineChar,
		H3UnderlineChar: DefaultH3UnderlineChar,
		HeadingBarChar:  DefaultHeadingBarChar,

		BulletChar:           DefaultBulletChar,
		NestedBulletChars:    DefaultNestedBulletChars,
		ListIndent:           DefaultListIndent,
		HRChar:               DefaultHRChar,
		HRWidth:              DefaultHRWidth,
		BlockquoteBar:        DefaultBlockquoteBar,
		InsPrefix:            DefaultInsPrefix,
		DelPrefix:            DefaultDelPrefix,
		FootnoteDividerChar:  DefaultFootnoteDividerChar,
		FootnoteDividerWidth: DefaultFootnoteDividerWidth,

		CodeLineNumber:    lipgloss.NewStyle().Foreground(p.Muted).Background(p.Base),
		CodeLineNumberSep: DefaultCodeLineNumberSep,

		TableHeader: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Primary),
		TableCell:        lipgloss.NewStyle().Foreground(p.Text),
		TableStripedCell: lipgloss.NewStyle().Foreground(p.Text).Background(p.Surface),
		TableFooter: lipgloss.NewStyle().
			Bold(true).
			Foreground(p.Text),
		TableCaption:   lipgloss.NewStyle().Foreground(p.Muted).Italic(true),
		TableBorder:    lipgloss.NewStyle().Foreground(p.Muted),
		TableBorderSet: BoxBorderSet(),
		TableCellPad:   DefaultTableCellPad,

		AlertBar: DefaultAlertBar,
		Alerts: DefaultAlertConfigs(AlertPalette{
			Note:      p.Tertiary,  // blue/cyan
			Tip:       p.Tertiary,  // green in some palettes
			Important: p.Secondary, // purple
			Warning:   p.Accent,    // yellow/amber
			Caution:   p.Highlight, // red
		}),
	}
}
