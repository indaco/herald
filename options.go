package herald

import "charm.land/lipgloss/v2"

// Option is a functional option for configuring a Typography instance.
type Option func(*Typography)

// WithTheme sets the entire theme at once.
func WithTheme(t Theme) Option {
	return func(ty *Typography) {
		ty.theme = t
	}
}

// WithPalette derives a complete theme from a ColorPalette and sets it.
// It is a convenience shortcut for WithTheme(ThemeFromPalette(p)).
func WithPalette(p ColorPalette) Option {
	return func(ty *Typography) {
		ty.theme = ThemeFromPalette(p)
	}
}

// --- Heading style options ---

// WithH1Style overrides the H1 heading style.
func WithH1Style(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.H1 = s }
}

// WithH2Style overrides the H2 heading style.
func WithH2Style(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.H2 = s }
}

// WithH3Style overrides the H3 heading style.
func WithH3Style(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.H3 = s }
}

// WithH4Style overrides the H4 heading style.
func WithH4Style(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.H4 = s }
}

// WithH5Style overrides the H5 heading style.
func WithH5Style(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.H5 = s }
}

// WithH6Style overrides the H6 heading style.
func WithH6Style(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.H6 = s }
}

// --- Block element style options ---

// WithParagraphStyle overrides the paragraph style.
func WithParagraphStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Paragraph = s }
}

// WithBlockquoteStyle overrides the blockquote style.
func WithBlockquoteStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Blockquote = s }
}

// WithBlockquoteBarStyle overrides the blockquote bar character style.
func WithBlockquoteBarStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.BlockquoteBarStyle = s }
}

// WithCodeInlineStyle overrides the inline code style.
func WithCodeInlineStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.CodeInline = s }
}

// WithCodeBlockStyle overrides the code block style.
func WithCodeBlockStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.CodeBlock = s }
}

// WithHRStyle overrides the horizontal rule style.
func WithHRStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.HR = s }
}

// WithHRLabelStyle overrides the label style for labeled horizontal rules.
func WithHRLabelStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.HRLabel = s }
}

// --- Inline style options ---

// WithBoldStyle overrides the bold style.
func WithBoldStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Bold = s }
}

// WithItalicStyle overrides the italic style.
func WithItalicStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Italic = s }
}

// WithUnderlineStyle overrides the underline style.
func WithUnderlineStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Underline = s }
}

// WithStrikethroughStyle overrides the strikethrough style.
func WithStrikethroughStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Strikethrough = s }
}

// WithSmallStyle overrides the small/faint style.
func WithSmallStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Small = s }
}

// WithMarkStyle overrides the highlight/mark style.
func WithMarkStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Mark = s }
}

// WithLinkStyle overrides the link style.
func WithLinkStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Link = s }
}

// WithKbdStyle overrides the keyboard key style.
func WithKbdStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Kbd = s }
}

// WithAbbrStyle overrides the abbreviation style.
func WithAbbrStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Abbr = s }
}

// WithInsStyle overrides the inserted text style.
func WithInsStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Ins = s }
}

// WithDelStyle overrides the deleted text style.
func WithDelStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Del = s }
}

// --- List style options ---

// WithListBulletStyle overrides the bullet/number marker style.
func WithListBulletStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.ListBullet = s }
}

// WithListItemStyle overrides the list item text style.
func WithListItemStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.ListItem = s }
}

// --- Definition list style options ---

// WithDTStyle overrides the definition term style.
func WithDTStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.DT = s }
}

// WithDDStyle overrides the definition description style.
func WithDDStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.DD = s }
}

// --- Address style option ---

// WithAddressStyle overrides the address/contact block style.
func WithAddressStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Address = s }
}

// WithAddressCardStyle overrides the address card content style.
func WithAddressCardStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.AddressCard = s }
}

// WithAddressCardBorderStyle overrides the address card border style.
func WithAddressCardBorderStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.AddressCardBorder = s }
}

// --- Badge / Tag style options ---

// WithBadgeStyle overrides the badge/tag pill style.
func WithBadgeStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Badge = s }
}

// WithTagStyle overrides the tag/category label style.
func WithTagStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.Tag = s }
}

// --- Footnote style options ---

// WithFootnoteRefStyle overrides the inline footnote reference marker style.
func WithFootnoteRefStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.FootnoteRef = s }
}

// WithFootnoteItemStyle overrides the footnote item style.
func WithFootnoteItemStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.FootnoteItem = s }
}

// WithFootnoteDividerStyle overrides the footnote divider style.
func WithFootnoteDividerStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.FootnoteDivider = s }
}

// WithFootnoteDividerChar sets the character used for the footnote section divider.
func WithFootnoteDividerChar(c string) Option {
	return func(ty *Typography) { ty.theme.FootnoteDividerChar = c }
}

// WithFootnoteDividerWidth sets the width of the footnote section divider.
func WithFootnoteDividerWidth(w int) Option {
	return func(ty *Typography) {
		if w > 0 {
			ty.theme.FootnoteDividerWidth = w
		}
	}
}

// --- Heading decoration options ---

// WithH1UnderlineChar sets the character used for the H1 underline.
func WithH1UnderlineChar(c string) Option {
	return func(ty *Typography) { ty.theme.H1UnderlineChar = c }
}

// WithH2UnderlineChar sets the character used for the H2 underline.
func WithH2UnderlineChar(c string) Option {
	return func(ty *Typography) { ty.theme.H2UnderlineChar = c }
}

// WithH3UnderlineChar sets the character used for the H3 underline.
func WithH3UnderlineChar(c string) Option {
	return func(ty *Typography) { ty.theme.H3UnderlineChar = c }
}

// WithHeadingBarChar sets the bar prefix character for H4-H6.
func WithHeadingBarChar(c string) Option {
	return func(ty *Typography) { ty.theme.HeadingBarChar = c }
}

// --- Token options ---

// WithBulletChar sets the bullet character for unordered lists.
func WithBulletChar(c string) Option {
	return func(ty *Typography) { ty.theme.BulletChar = c }
}

// WithHRChar sets the character used for horizontal rules.
func WithHRChar(c string) Option {
	return func(ty *Typography) { ty.theme.HRChar = c }
}

// WithHRWidth sets the width of horizontal rules in characters.
func WithHRWidth(w int) Option {
	return func(ty *Typography) {
		if w > 0 {
			ty.theme.HRWidth = w
		}
	}
}

// WithBlockquoteBar sets the left-bar character for blockquotes.
func WithBlockquoteBar(c string) Option {
	return func(ty *Typography) { ty.theme.BlockquoteBar = c }
}

// WithInsPrefix sets the prefix for inserted text.
func WithInsPrefix(c string) Option {
	return func(ty *Typography) { ty.theme.InsPrefix = c }
}

// WithDelPrefix sets the prefix for deleted text.
func WithDelPrefix(c string) Option {
	return func(ty *Typography) { ty.theme.DelPrefix = c }
}

// --- Nested list options ---

// WithListIndent sets the number of spaces per nesting level for nested lists.
func WithListIndent(n int) Option {
	return func(ty *Typography) {
		if n > 0 {
			ty.theme.ListIndent = n
		}
	}
}

// WithNestedBulletChars sets the bullet characters that cycle through
// nesting levels in nested unordered lists.
func WithNestedBulletChars(chars []string) Option {
	return func(ty *Typography) {
		if len(chars) > 0 {
			ty.theme.NestedBulletChars = chars
		}
	}
}

// WithHierarchicalNumbers enables hierarchical numbering for nested ordered
// lists (e.g. 1., 1.1, 1.2, 2., 2.1). When false (default), each nested
// sub-list restarts numbering at 1.
func WithHierarchicalNumbers(enabled bool) Option {
	return func(ty *Typography) { ty.theme.HierarchicalNumbers = enabled }
}

// --- Table options ---

// WithTableHeaderStyle overrides the table header cell style.
func WithTableHeaderStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.TableHeader = s }
}

// WithTableCellStyle overrides the table body cell style.
func WithTableCellStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.TableCell = s }
}

// WithTableStripedCellStyle overrides the style for alternating body rows
// when striped rows are enabled.
func WithTableStripedCellStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.TableStripedCell = s }
}

// WithTableFooterStyle overrides the table footer row style.
func WithTableFooterStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.TableFooter = s }
}

// WithTableCaptionStyle overrides the table caption style.
func WithTableCaptionStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.TableCaption = s }
}

// WithTableBorderStyle overrides the style applied to table border characters.
func WithTableBorderStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.TableBorder = s }
}

// WithTableBorderSet sets the box-drawing character set for tables.
func WithTableBorderSet(bs TableBorderSet) Option {
	return func(ty *Typography) { ty.theme.TableBorderSet = bs }
}

// WithTableCellPad sets the number of spaces of padding inside each table cell.
func WithTableCellPad(n int) Option {
	return func(ty *Typography) {
		if n >= 0 {
			ty.theme.TableCellPad = n
		}
	}
}

// --- Alert options ---

// WithAlertStyle overrides the style for a specific alert type.
func WithAlertStyle(at AlertType, s lipgloss.Style) Option {
	return func(ty *Typography) {
		if ty.theme.Alerts == nil {
			ty.theme.Alerts = make(map[AlertType]AlertConfig)
		}
		cfg := ty.theme.Alerts[at]
		cfg.Style = s
		ty.theme.Alerts[at] = cfg
	}
}

// WithAlertIcon overrides the icon for a specific alert type.
func WithAlertIcon(at AlertType, icon string) Option {
	return func(ty *Typography) {
		if ty.theme.Alerts == nil {
			ty.theme.Alerts = make(map[AlertType]AlertConfig)
		}
		cfg := ty.theme.Alerts[at]
		cfg.Icon = icon
		ty.theme.Alerts[at] = cfg
	}
}

// WithAlertLabel overrides the label for a specific alert type.
func WithAlertLabel(at AlertType, label string) Option {
	return func(ty *Typography) {
		if ty.theme.Alerts == nil {
			ty.theme.Alerts = make(map[AlertType]AlertConfig)
		}
		cfg := ty.theme.Alerts[at]
		cfg.Label = label
		ty.theme.Alerts[at] = cfg
	}
}

// WithAlertBar sets the left-bar character for alerts.
func WithAlertBar(c string) Option {
	return func(ty *Typography) { ty.theme.AlertBar = c }
}

// WithAlertPalette rebuilds all alert configs from the given AlertPalette,
// using default icons and labels.
func WithAlertPalette(ap AlertPalette) Option {
	return func(ty *Typography) {
		ty.theme.Alerts = DefaultAlertConfigs(ap)
	}
}

// --- Code block line number options ---

// WithCodeLineNumbers enables or disables line numbers in code blocks.
func WithCodeLineNumbers(enabled bool) Option {
	return func(ty *Typography) { ty.theme.ShowLineNumbers = enabled }
}

// WithCodeLineNumberStyle overrides the style for code block line numbers.
func WithCodeLineNumberStyle(s lipgloss.Style) Option {
	return func(ty *Typography) { ty.theme.CodeLineNumber = s }
}

// WithCodeLineNumberSep sets the separator between line numbers and code content.
func WithCodeLineNumberSep(sep string) Option {
	return func(ty *Typography) { ty.theme.CodeLineNumberSep = sep }
}

// --- Callback options ---

// WithCodeFormatter sets a callback that receives raw code and a language hint,
// returning syntax-highlighted text. The highlighted text is then wrapped in
// the CodeInline or CodeBlock style for consistent padding/margins.
func WithCodeFormatter(fn func(code, language string) string) Option {
	return func(ty *Typography) { ty.theme.CodeFormatter = fn }
}
