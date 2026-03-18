package herald

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

// Typography is the central renderer. It holds a Theme and exposes methods
// for every supported typographic element.
type Typography struct {
	theme Theme
}

// New creates a new Typography instance with the default theme, then applies
// any provided functional options.
func New(opts ...Option) *Typography {
	t := &Typography{
		theme: DefaultTheme(),
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

// Theme returns a copy of the current theme.
func (t *Typography) Theme() Theme {
	return t.theme
}

// ---------------------------------------------------------------------------
// Headings
// ---------------------------------------------------------------------------

// headingWithUnderline renders a heading with a repeated underline character
// beneath the text. The underline matches the visible width of the text.
func (t *Typography) headingWithUnderline(text string, style lipgloss.Style, char string) string {
	rendered := style.UnsetMarginBottom().Render(text)
	underline := style.UnsetMarginBottom().Render(strings.Repeat(char, len([]rune(text))))
	return rendered + "\n" + underline + style.Render("")
}

// headingWithBar renders a heading prefixed by a vertical bar character.
func (t *Typography) headingWithBar(text string, style lipgloss.Style, bar string) string {
	return style.Render(bar + " " + text)
}

// H1 renders text as a level-1 heading with a double-line underline.
func (t *Typography) H1(text string) string {
	return t.headingWithUnderline(text, t.theme.H1, t.theme.H1UnderlineChar)
}

// H2 renders text as a level-2 heading with a single-line underline.
func (t *Typography) H2(text string) string {
	return t.headingWithUnderline(text, t.theme.H2, t.theme.H2UnderlineChar)
}

// H3 renders text as a level-3 heading with a dotted underline.
func (t *Typography) H3(text string) string {
	return t.headingWithUnderline(text, t.theme.H3, t.theme.H3UnderlineChar)
}

// H4 renders text as a level-4 heading with a bar prefix.
func (t *Typography) H4(text string) string {
	return t.headingWithBar(text, t.theme.H4, t.theme.HeadingBarChar)
}

// H5 renders text as a level-5 heading with a bar prefix.
func (t *Typography) H5(text string) string {
	return t.headingWithBar(text, t.theme.H5, t.theme.HeadingBarChar)
}

// H6 renders text as a level-6 heading with a bar prefix.
func (t *Typography) H6(text string) string {
	return t.headingWithBar(text, t.theme.H6, t.theme.HeadingBarChar)
}

// ---------------------------------------------------------------------------
// Block elements
// ---------------------------------------------------------------------------

// P renders a paragraph.
func (t *Typography) P(text string) string {
	return t.theme.Paragraph.Render(text)
}

// Blockquote renders a blockquote with a left border bar. Multi-line text
// is handled by prepending the bar to every line.
func (t *Typography) Blockquote(text string) string {
	bar := t.theme.BlockquoteBar
	lines := strings.Split(text, "\n")
	quoted := make([]string, len(lines))
	for i, line := range lines {
		quoted[i] = bar + " " + line
	}
	return t.theme.Blockquote.Render(strings.Join(quoted, "\n"))
}

// UL renders an unordered (bulleted) list from the provided items.
func (t *Typography) UL(items ...string) string {
	if len(items) == 0 {
		return ""
	}
	bullet := t.theme.BulletChar
	lines := make([]string, len(items))
	for i, item := range items {
		marker := t.theme.ListBullet.Render(bullet)
		lines[i] = marker + " " + t.theme.ListItem.Render(item)
	}
	return strings.Join(lines, "\n")
}

// OL renders an ordered (numbered) list from the provided items.
func (t *Typography) OL(items ...string) string {
	if len(items) == 0 {
		return ""
	}
	lines := make([]string, len(items))
	for i, item := range items {
		num := fmt.Sprintf("%d.", i+1)
		marker := t.theme.ListBullet.Render(num)
		lines[i] = marker + " " + t.theme.ListItem.Render(item)
	}
	return strings.Join(lines, "\n")
}

// Code renders inline code. If a language is provided and a CodeFormatter is
// set on the theme, the formatter is applied before wrapping in the style.
func (t *Typography) Code(text string, lang ...string) string {
	content := text
	if t.theme.CodeFormatter != nil && len(lang) > 0 && lang[0] != "" {
		content = t.theme.CodeFormatter(text, lang[0])
	}
	return t.theme.CodeInline.Render(content)
}

// CodeBlock renders a fenced code block. If a language is provided and a
// CodeFormatter is set on the theme, the formatter is applied before wrapping
// in the style.
func (t *Typography) CodeBlock(text string, lang ...string) string {
	content := text
	if t.theme.CodeFormatter != nil && len(lang) > 0 && lang[0] != "" {
		content = t.theme.CodeFormatter(text, lang[0])
	}
	return t.theme.CodeBlock.Render(content)
}

// HR renders a horizontal rule.
func (t *Typography) HR() string {
	line := strings.Repeat(t.theme.HRChar, t.theme.HRWidth)
	return t.theme.HR.Render(line)
}

// ---------------------------------------------------------------------------
// Inline styles
// ---------------------------------------------------------------------------

// Bold renders bold text.
func (t *Typography) Bold(text string) string {
	return t.theme.Bold.Render(text)
}

// Italic renders italic text.
func (t *Typography) Italic(text string) string {
	return t.theme.Italic.Render(text)
}

// Underline renders underlined text.
func (t *Typography) Underline(text string) string {
	return t.theme.Underline.Render(text)
}

// Strikethrough renders strikethrough text.
func (t *Typography) Strikethrough(text string) string {
	return t.theme.Strikethrough.Render(text)
}

// Small renders small/faint text.
func (t *Typography) Small(text string) string {
	return t.theme.Small.Render(text)
}

// Mark renders highlighted text.
func (t *Typography) Mark(text string) string {
	return t.theme.Mark.Render(text)
}

// Link renders a styled URL or link text. If both label and url are provided,
// it renders as "label (url)". If only one argument is given, it is treated
// as both the label and the URL.
func (t *Typography) Link(label string, url ...string) string {
	if len(url) > 0 && url[0] != "" && url[0] != label {
		return t.theme.Link.Render(label) + " (" + t.theme.Small.Render(url[0]) + ")"
	}
	return t.theme.Link.Render(label)
}

// Kbd renders a keyboard key indicator.
func (t *Typography) Kbd(text string) string {
	return t.theme.Kbd.Render(text)
}

// Abbr renders an abbreviation. If a description is provided it is shown in
// parentheses after the abbreviation.
func (t *Typography) Abbr(abbr string, desc ...string) string {
	styled := t.theme.Abbr.Render(abbr)
	if len(desc) > 0 && desc[0] != "" {
		styled += " (" + desc[0] + ")"
	}
	return styled
}

// Sub renders a subscript marker. In a terminal we prefix with an underscore
// to visually indicate subscript.
func (t *Typography) Sub(text string) string {
	return t.theme.Sub.Render("_" + text)
}

// Sup renders a superscript marker. In a terminal we prefix with a caret
// to visually indicate superscript.
func (t *Typography) Sup(text string) string {
	return t.theme.Sup.Render("^" + text)
}

// ---------------------------------------------------------------------------
// Definition list
// ---------------------------------------------------------------------------

// DL renders a definition list from term-description pairs. Each pair is a
// two-element array: [term, description]. Odd-length slices ignore the last
// unpaired element.
func (t *Typography) DL(pairs [][2]string) string {
	if len(pairs) == 0 {
		return ""
	}
	lines := make([]string, 0, len(pairs)*2)
	for _, pair := range pairs {
		lines = append(lines, t.theme.DT.Render(pair[0]), t.theme.DD.Render(pair[1]))
	}
	return strings.Join(lines, "\n")
}

// DT renders a single definition term.
func (t *Typography) DT(text string) string {
	return t.theme.DT.Render(text)
}

// DD renders a single definition description.
func (t *Typography) DD(text string) string {
	return t.theme.DD.Render(text)
}
