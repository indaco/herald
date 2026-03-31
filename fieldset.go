package herald

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// Fieldset renders content inside a bordered box with an optional legend
// embedded in the top border. If legend is empty, a plain border box is
// rendered. An optional width override can be passed; 0 or omitted means
// auto-fit to the content. The rendering uses manual string building
// because lipgloss.Border cannot embed a legend in the top line.
func (t *Typography) Fieldset(legend, content string, width ...int) string {
	w := t.theme.FieldsetWidth
	if len(width) > 0 && width[0] > 0 {
		w = width[0]
	}

	innerWidth := t.fieldsetInnerWidth(legend, content, w)
	bs := t.theme.FieldsetBorder

	var sb strings.Builder
	t.fieldsetTopBorder(&sb, legend, innerWidth, bs)
	sb.WriteByte('\n')
	t.fieldsetContent(&sb, content, innerWidth, bs)
	sb.WriteByte('\n')
	t.fieldsetBottomBorder(&sb, innerWidth, bs)
	return sb.String()
}

// fieldsetInnerWidth computes the inner width (the space between the left
// and right border characters). When w == 0, auto-fit to the wider of
// legend+4 or maxContentLine+2.
func (t *Typography) fieldsetInnerWidth(legend, content string, w int) int {
	if w > 0 {
		// Explicit width: innerWidth = w - 2 (for left+right border chars)
		return w - 2
	}

	contentLines := strings.Split(content, "\n")
	maxCW := 0
	for _, line := range contentLines {
		if lw := lipgloss.Width(line); lw > maxCW {
			maxCW = lw
		}
	}

	legendWidth := lipgloss.Width(legend)
	inner := maxCW + 2 // +2 for left/right padding inside border
	if legend != "" {
		if lw := legendWidth + 4; lw > inner { // +4 for dash+space around legend
			inner = lw
		}
	}
	return inner
}

// fieldsetTopBorder writes the top border line. With a legend:
//
//	╭─ Legend ──────╮
//
// Without a legend:
//
//	╭──────────────╮
func (t *Typography) fieldsetTopBorder(sb *strings.Builder, legend string, innerWidth int, bs lipgloss.Style) {
	if legend == "" {
		sb.WriteString(bs.Render("╭" + strings.Repeat("─", innerWidth) + "╮"))
		return
	}

	legendRendered := t.theme.FieldsetLegend.Render(legend)
	legendWidth := lipgloss.Width(legend)
	rightDashes := max(
		// 3 = "─ " before legend + " " after legend
		innerWidth-legendWidth-3, 1)

	// Three segments: "╭─ " + legend + " ─...─╮"
	// Spaces are rendered inside the border-styled segments to avoid visual gaps.
	sb.WriteString(bs.Render("╭─ "))
	sb.WriteString(legendRendered)
	sb.WriteString(bs.Render(" " + strings.Repeat("─", rightDashes) + "╮"))
}

// fieldsetContent writes the content lines, each padded and wrapped with
// border characters.
func (t *Typography) fieldsetContent(sb *strings.Builder, content string, innerWidth int, bs lipgloss.Style) {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		styledLine := t.theme.Fieldset.Render(line)
		lineWidth := lipgloss.Width(styledLine)
		pad := max(
			// -2 for left/right space padding
			innerWidth-2-lineWidth, 0)
		sb.WriteString(bs.Render("│"))
		sb.WriteString(" " + styledLine + strings.Repeat(" ", pad) + " ")
		sb.WriteString(bs.Render("│"))
		if i < len(lines)-1 {
			sb.WriteByte('\n')
		}
	}
}

// fieldsetBottomBorder writes the bottom border line: ╰───...───╯
func (t *Typography) fieldsetBottomBorder(sb *strings.Builder, innerWidth int, bs lipgloss.Style) {
	sb.WriteString(bs.Render("╰" + strings.Repeat("─", innerWidth) + "╯"))
}
