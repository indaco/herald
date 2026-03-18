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

// --- Callback options ---

// WithCodeFormatter sets a callback that receives raw code and a language hint,
// returning syntax-highlighted text. The highlighted text is then wrapped in
// the CodeInline or CodeBlock style for consistent padding/margins.
func WithCodeFormatter(fn func(code, language string) string) Option {
	return func(ty *Typography) { ty.theme.CodeFormatter = fn }
}
