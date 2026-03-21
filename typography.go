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

// NestUL renders a nested unordered list from the provided ListItems.
func (t *Typography) NestUL(items ...ListItem) string {
	return t.renderNestedList(items, Unordered, 0, "")
}

// NestOL renders a nested ordered list from the provided ListItems.
func (t *Typography) NestOL(items ...ListItem) string {
	return t.renderNestedList(items, Ordered, 0, "")
}

// renderNestedList recursively renders a list at the given depth.
// The prefix parameter carries the parent number for hierarchical numbering
// (e.g. "2" so children become "2.1", "2.2").
func (t *Typography) renderNestedList(items []ListItem, kind ListKind, depth int, prefix string) string {
	if len(items) == 0 {
		return ""
	}

	indent := strings.Repeat(" ", depth*t.theme.ListIndent)
	lines := make([]string, 0, len(items))

	for i, item := range items {
		var marker string
		var childPrefix string
		if kind == Ordered {
			num := fmt.Sprintf("%d", i+1)
			if t.theme.HierarchicalNumbers && prefix != "" {
				num = prefix + "." + num
			}
			marker = t.theme.ListBullet.Render(num + ".")
			childPrefix = num
		} else {
			chars := t.theme.NestedBulletChars
			bullet := chars[depth%len(chars)]
			marker = t.theme.ListBullet.Render(bullet)
			childPrefix = ""
		}
		lines = append(lines, indent+marker+" "+t.theme.ListItem.Render(item.Text))

		if len(item.Children) > 0 {
			child := t.renderNestedList(item.Children, item.Kind, depth+1, childPrefix)
			lines = append(lines, child)
		}
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
// in the style. When ShowLineNumbers is true, line numbers are prepended to
// each line after formatting.
func (t *Typography) CodeBlock(text string, lang ...string) string {
	content := text
	if t.theme.CodeFormatter != nil && len(lang) > 0 && lang[0] != "" {
		content = t.theme.CodeFormatter(text, lang[0])
	}
	if t.theme.ShowLineNumbers {
		return t.codeBlockWithLineNumbers(content)
	}
	return t.theme.CodeBlock.Render(content)
}

// codeBlockWithLineNumbers renders a code block with a line number gutter.
// It builds two separate columns (numbers+separator and code) and joins them
// horizontally so that each column is independently styled, avoiding nested
// Render calls that break background propagation.
func (t *Typography) codeBlockWithLineNumbers(content string) string {
	lines := strings.Split(content, "\n")
	width := len(fmt.Sprintf("%d", len(lines)))

	bg := t.theme.CodeBlock.GetBackground()

	// Build the gutter column: right-aligned numbers + separator.
	gutter := make([]string, len(lines))
	for i := range lines {
		gutter[i] = fmt.Sprintf("%*d", width, i+1) + t.theme.CodeLineNumberSep
	}
	gutterStyle := t.theme.CodeLineNumber.
		Background(bg).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(2).
		MarginBottom(t.theme.CodeBlock.GetMarginBottom())

	// Build the code column with matching background and right padding.
	codeStyle := lipgloss.NewStyle().
		Foreground(t.theme.CodeBlock.GetForeground()).
		Background(bg).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingRight(2).
		PaddingLeft(1).
		MarginBottom(t.theme.CodeBlock.GetMarginBottom())

	return lipgloss.JoinHorizontal(lipgloss.Top,
		gutterStyle.Render(strings.Join(gutter, "\n")),
		codeStyle.Render(content),
	)
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

// Ins renders inserted (added) text with a prefix marker.
func (t *Typography) Ins(text string) string {
	return t.theme.Ins.Render(t.theme.InsPrefix + text)
}

// Del renders deleted (removed) text with a prefix marker.
func (t *Typography) Del(text string) string {
	return t.theme.Del.Render(t.theme.DelPrefix + text)
}

// ---------------------------------------------------------------------------
// Alerts
// ---------------------------------------------------------------------------

// Alert renders a GitHub-style alert callout. The header line (bar + icon +
// label) is bold + colored; content lines have a colored bar but unstyled
// text, matching GitHub's visual style. If the alert type is not configured,
// it falls back to blockquote rendering.
func (t *Typography) Alert(at AlertType, text string) string {
	cfg, ok := t.theme.Alerts[at]
	if !ok {
		return t.Blockquote(text)
	}

	bar := t.theme.AlertBar
	style := cfg.Style

	// Header: bar + icon + label, rendered bold + colored
	headerStyle := style.Bold(true)
	header := headerStyle.Render(bar + " " + cfg.Icon + " " + cfg.Label)

	// Content: colored bar + unstyled text (matching GitHub alerts)
	lines := strings.Split(text, "\n")
	content := make([]string, len(lines))
	for i, line := range lines {
		content[i] = style.Render(bar) + " " + line
	}

	return header + "\n" + strings.Join(content, "\n")
}

// Note renders a blue informational alert.
func (t *Typography) Note(text string) string { return t.Alert(AlertNote, text) }

// Tip renders a green helpful-hint alert.
func (t *Typography) Tip(text string) string { return t.Alert(AlertTip, text) }

// Important renders a purple important-information alert.
func (t *Typography) Important(text string) string { return t.Alert(AlertImportant, text) }

// Warning renders a yellow/amber warning alert.
func (t *Typography) Warning(text string) string { return t.Alert(AlertWarning, text) }

// Caution renders a red caution/danger alert.
func (t *Typography) Caution(text string) string { return t.Alert(AlertCaution, text) }

// ---------------------------------------------------------------------------
// Table
// ---------------------------------------------------------------------------

// Alignment represents the horizontal alignment of text within a table cell.
type Alignment int

const (
	// AlignLeft aligns cell content to the left (default).
	AlignLeft Alignment = iota
	// AlignCenter centers cell content horizontally.
	AlignCenter
	// AlignRight aligns cell content to the right.
	AlignRight
)

// TableOption is a functional option for per-table configuration.
// It is distinct from Option, which configures the Typography instance.
type TableOption func(*tableConfig)

// CaptionPosition specifies where a table caption is rendered.
type CaptionPosition int

const (
	// CaptionTop renders the caption above the table.
	CaptionTop CaptionPosition = iota
	// CaptionBottom renders the caption below the table.
	CaptionBottom
)

// tableConfig holds per-table settings applied via TableOption.
type tableConfig struct {
	alignments      map[int]Alignment // column index -> alignment
	rowSeparators   bool              // draw horizontal lines between body rows
	stripedRows     bool              // alternate row styles for readability
	caption         string            // optional caption text
	captionPosition CaptionPosition   // top or bottom
	footerRow       bool              // treat last row as a footer
	maxColWidth     int               // global max column width (0 = unlimited)
	maxColWidths    map[int]int       // per-column max widths (overrides global)
}

// WithColumnAlign sets the alignment for a specific column (0-indexed).
// Columns without an explicit alignment default to AlignLeft.
func WithColumnAlign(col int, align Alignment) TableOption {
	return func(cfg *tableConfig) {
		if cfg.alignments == nil {
			cfg.alignments = make(map[int]Alignment)
		}
		cfg.alignments[col] = align
	}
}

// WithRowSeparators enables horizontal separator lines between every body row.
func WithRowSeparators(enabled bool) TableOption {
	return func(cfg *tableConfig) {
		cfg.rowSeparators = enabled
	}
}

// WithStripedRows enables alternating row background styles for readability.
// Odd body rows use the TableStripedCell style instead of TableCell.
func WithStripedRows(enabled bool) TableOption {
	return func(cfg *tableConfig) {
		cfg.stripedRows = enabled
	}
}

// WithCaption adds a caption above the table.
func WithCaption(text string) TableOption {
	return func(cfg *tableConfig) {
		cfg.caption = text
		cfg.captionPosition = CaptionTop
	}
}

// WithCaptionBottom adds a caption below the table.
func WithCaptionBottom(text string) TableOption {
	return func(cfg *tableConfig) {
		cfg.caption = text
		cfg.captionPosition = CaptionBottom
	}
}

// WithFooterRow treats the last row as a footer with a distinct separator
// and the TableFooter style.
func WithFooterRow(enabled bool) TableOption {
	return func(cfg *tableConfig) {
		cfg.footerRow = enabled
	}
}

// WithMaxColumnWidth sets a global maximum visible width for all columns.
// Cells exceeding this width are truncated with "…". A value of 0 disables
// truncation. Per-column limits set via WithColumnMaxWidth take precedence.
func WithMaxColumnWidth(n int) TableOption {
	return func(cfg *tableConfig) {
		cfg.maxColWidth = n
	}
}

// WithColumnMaxWidth sets the maximum visible width for a specific column
// (0-indexed). Cells exceeding this width are truncated with "…".
// This overrides the global WithMaxColumnWidth for the given column.
func WithColumnMaxWidth(col, n int) TableOption {
	return func(cfg *tableConfig) {
		if cfg.maxColWidths == nil {
			cfg.maxColWidths = make(map[int]int)
		}
		cfg.maxColWidths[col] = n
	}
}

// WithColumnAligns sets the alignment for columns 0, 1, 2, ... in order.
// Columns beyond the length of the slice default to AlignLeft.
func WithColumnAligns(aligns ...Alignment) TableOption {
	return func(cfg *tableConfig) {
		if cfg.alignments == nil {
			cfg.alignments = make(map[int]Alignment)
		}
		for i, a := range aligns {
			cfg.alignments[i] = a
		}
	}
}

// truncateCell truncates a string to maxWidth visible characters, appending "…"
// if truncation occurs. Returns the original string if it fits.
func truncateCell(s string, maxWidth int) string {
	if maxWidth <= 0 || lipgloss.Width(s) <= maxWidth {
		return s
	}
	runes := []rune(s)
	// Reserve 1 character for the ellipsis.
	for i := len(runes); i > 0; i-- {
		candidate := string(runes[:i]) + "…"
		if lipgloss.Width(candidate) <= maxWidth {
			return candidate
		}
	}
	return "…"
}

// truncateRows returns a copy of rows with cells truncated according to the
// config's max column width settings. The original rows are not modified.
func truncateRows(rows [][]string, cfg *tableConfig) [][]string {
	if cfg.maxColWidth <= 0 && len(cfg.maxColWidths) == 0 {
		return rows
	}
	out := make([][]string, len(rows))
	for i, row := range rows {
		newRow := make([]string, len(row))
		for c, cell := range row {
			maxW := cfg.maxColWidth
			if w, ok := cfg.maxColWidths[c]; ok {
				maxW = w
			}
			if maxW > 0 {
				newRow[c] = truncateCell(cell, maxW)
			} else {
				newRow[c] = cell
			}
		}
		out[i] = newRow
	}
	return out
}

// tableColumnWidths computes the maximum cell width per column across all rows.
func tableColumnWidths(rows [][]string, cols int) []int {
	widths := make([]int, cols)
	for _, row := range rows {
		for c := range cols {
			cell := ""
			if c < len(row) {
				cell = row[c]
			}
			if w := lipgloss.Width(cell); w > widths[c] {
				widths[c] = w
			}
		}
	}
	return widths
}

// tableHLine builds a horizontal separator line for a table.
func (t *Typography) tableHLine(colWidths []int, pad int, left, fill, junction, right string) string {
	segments := make([]string, len(colWidths))
	for c, w := range colWidths {
		segments[c] = strings.Repeat(fill, w+pad*2)
	}
	return t.theme.TableBorder.Render(left + strings.Join(segments, junction) + right)
}

// alignCell pads a rendered cell string to the given total width according to
// the specified alignment. The returned string has exactly totalWidth visible
// characters (excluding the surrounding padStr added by the caller).
func alignCell(rendered string, cellWidth, totalWidth int, align Alignment) string {
	gap := totalWidth - cellWidth
	if gap <= 0 {
		return rendered
	}
	switch align {
	case AlignRight:
		return strings.Repeat(" ", gap) + rendered
	case AlignCenter:
		left := gap / 2
		right := gap - left
		return strings.Repeat(" ", left) + rendered + strings.Repeat(" ", right)
	default: // AlignLeft
		return rendered + strings.Repeat(" ", gap)
	}
}

// tableRow renders a single data row with cell content and vertical separators.
func (t *Typography) tableRow(row []string, style lipgloss.Style, colWidths []int, padStr string, bordered bool, aligns map[int]Alignment) string {
	bs := t.theme.TableBorderSet
	cols := len(colWidths)
	cells := make([]string, cols)
	for c := range cols {
		cell := ""
		if c < len(row) {
			cell = row[c]
		}
		rendered := style.Render(cell)
		cellWidth := lipgloss.Width(cell)
		align := AlignLeft
		if a, ok := aligns[c]; ok {
			align = a
		}
		cells[c] = padStr + alignCell(rendered, cellWidth, colWidths[c], align) + padStr
	}

	sep, end := "", ""
	if bordered {
		sep = t.theme.TableBorder.Render(bs.Left)
		end = t.theme.TableBorder.Render(bs.Right)
	}
	inner := t.theme.TableBorder.Render("│")
	return sep + strings.Join(cells, inner) + end
}

// Table renders a table from a slice of rows. The first row is treated as the
// header. Each row is a slice of cell strings. Rows may have different lengths;
// shorter rows are padded with empty cells. Returns an empty string if rows is
// nil or empty.
func (t *Typography) Table(rows [][]string) string {
	return t.renderTable(rows, &tableConfig{})
}

// TableWithOpts renders a table like Table but accepts per-table options such
// as column alignment. The first row is treated as the header.
//
//	t.TableWithOpts(rows,
//	    herald.WithColumnAlign(0, herald.AlignCenter),
//	    herald.WithColumnAlign(2, herald.AlignRight),
//	)
func (t *Typography) TableWithOpts(rows [][]string, opts ...TableOption) string {
	cfg := &tableConfig{}
	for _, o := range opts {
		o(cfg)
	}
	return t.renderTable(rows, cfg)
}

// tableMaxCols returns the maximum number of columns across all rows.
func tableMaxCols(rows [][]string) int {
	cols := 0
	for _, row := range rows {
		if len(row) > cols {
			cols = len(row)
		}
	}
	return cols
}

// renderTable is the shared implementation for Table and TableWithOpts.
func (t *Typography) renderTable(rows [][]string, cfg *tableConfig) string {
	if len(rows) == 0 {
		return ""
	}

	cols := tableMaxCols(rows)
	if cols == 0 {
		return ""
	}

	// Apply auto-truncation before computing widths.
	rows = truncateRows(rows, cfg)

	aligns := cfg.alignments
	if aligns == nil {
		aligns = make(map[int]Alignment)
	}

	bs := t.theme.TableBorderSet
	pad := t.theme.TableCellPad
	padStr := strings.Repeat(" ", pad)
	colWidths := tableColumnWidths(rows, cols)
	bordered := bs.TopLeft != ""

	// Determine body range and footer row.
	bodyEnd := len(rows)
	hasFooter := cfg.footerRow && len(rows) > 2
	if hasFooter {
		bodyEnd = len(rows) - 1
	}

	var sb strings.Builder

	// Caption (top).
	t.writeCaption(&sb, cfg, CaptionTop, false)

	// Top border.
	if bordered {
		sb.WriteString(t.tableHLine(colWidths, pad, bs.TopLeft, bs.Top, bs.TopJunction, bs.TopRight))
		sb.WriteByte('\n')
	}

	// Header row.
	sb.WriteString(t.tableRow(rows[0], t.theme.TableHeader, colWidths, padStr, bordered, aligns))
	sb.WriteByte('\n')
	sb.WriteString(t.tableHLine(colWidths, pad, bs.HeaderLeft, bs.Header, bs.HeaderCross, bs.HeaderRight))

	// Body rows.
	hasRowSep := cfg.rowSeparators && bs.Row != ""
	t.renderTableBody(rows[1:bodyEnd], &sb, colWidths, padStr, bordered, aligns, hasRowSep, cfg.stripedRows)

	// Footer row.
	if hasFooter {
		sb.WriteByte('\n')
		sb.WriteString(t.tableHLine(colWidths, pad, bs.FooterLeft, bs.Header, bs.FooterCross, bs.FooterRight))
		sb.WriteByte('\n')
		sb.WriteString(t.tableRow(rows[len(rows)-1], t.theme.TableFooter, colWidths, padStr, bordered, aligns))
	}

	// Bottom border.
	if bordered {
		sb.WriteByte('\n')
		sb.WriteString(t.tableHLine(colWidths, pad, bs.BottomLeft, bs.Bottom, bs.BottomJunction, bs.BottomRight))
	}

	// Caption (bottom).
	t.writeCaption(&sb, cfg, CaptionBottom, true)

	return sb.String()
}

// writeCaption writes the table caption to the builder if it matches the
// given position. When newlineBefore is true, a newline is prepended.
func (t *Typography) writeCaption(sb *strings.Builder, cfg *tableConfig, pos CaptionPosition, newlineBefore bool) {
	if cfg.caption == "" || cfg.captionPosition != pos {
		return
	}
	if newlineBefore {
		sb.WriteByte('\n')
	}
	sb.WriteString(t.theme.TableCaption.Render(cfg.caption))
	if !newlineBefore {
		sb.WriteByte('\n')
	}
}

// renderTableBody writes body rows to the builder, handling row separators and
// striped row styling.
func (t *Typography) renderTableBody(bodyRows [][]string, sb *strings.Builder, colWidths []int, padStr string, bordered bool, aligns map[int]Alignment, hasRowSep, striped bool) {
	bs := t.theme.TableBorderSet
	for i, row := range bodyRows {
		sb.WriteByte('\n')
		if hasRowSep && i > 0 {
			left, right := bs.LeftJunction, bs.RightJunction
			if !bordered {
				left, right = "", ""
			}
			sb.WriteString(t.tableHLine(colWidths, t.theme.TableCellPad, left, bs.Row, bs.Cross, right))
			sb.WriteByte('\n')
		}
		style := t.theme.TableCell
		if striped && i%2 == 1 {
			style = t.theme.TableStripedCell
		}
		sb.WriteString(t.tableRow(row, style, colWidths, padStr, bordered, aligns))
	}
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

// ---------------------------------------------------------------------------
// Address
// ---------------------------------------------------------------------------

// Address renders a styled contact/author information block.
func (t *Typography) Address(text string) string {
	return t.theme.Address.Render(text)
}

// AddressCard renders a contact/author block inside a bordered card.
// The border color is taken from the AddressCardBorder theme style.
func (t *Typography) AddressCard(text string) string {
	borderColor := t.theme.AddressCardBorder.GetForeground()
	style := t.theme.AddressCard.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Padding(0, 1)
	return style.Render(text)
}

// ---------------------------------------------------------------------------
// Badge
// ---------------------------------------------------------------------------

// Badge renders text as a styled pill/tag label.
func (t *Typography) Badge(text string) string {
	return t.theme.Badge.Render(text)
}

// BadgeWithStyle renders a badge using a one-off style override, useful for
// semantic variants (success, warning, error) without changing the theme.
func (t *Typography) BadgeWithStyle(text string, style lipgloss.Style) string {
	return style.Render(text)
}

// Tag renders text as a subtle pill/category label.
func (t *Typography) Tag(text string) string {
	return t.theme.Tag.Render(text)
}

// TagWithStyle renders a tag using a one-off style override.
func (t *Typography) TagWithStyle(text string, style lipgloss.Style) string {
	return style.Render(text)
}
