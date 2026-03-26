package herald

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// SemanticPalette defines four colors for common status semantics. It is
// separate from ColorPalette (which maps to typographic elements) and from
// AlertPalette (which maps to GitHub-style alert callouts). A SemanticPalette
// is used to derive themed badge and tag styles for status indicators.
type SemanticPalette struct {
	Success color.Color // green  - running, passed, healthy
	Warning color.Color // yellow/amber - expiring, degraded
	Error   color.Color // red - failed, critical, down
	Info    color.Color // blue - informational, neutral status
}

// DefaultSemanticPalette derives a SemanticPalette from the given ColorPalette
// using sensible defaults:
//
//   - Success -> Tertiary  (green in most themes)
//   - Warning -> Accent    (yellow/amber in most themes)
//   - Error   -> Highlight (red in most themes)
//   - Info    -> Secondary (blue/purple in most themes)
func DefaultSemanticPalette(p ColorPalette) SemanticPalette {
	return SemanticPalette{
		Success: p.Tertiary,
		Warning: p.Accent,
		Error:   p.Highlight,
		Info:    p.Secondary,
	}
}

// defaultSemanticBadgeStyles builds the four semantic badge styles from a
// SemanticPalette. Badge styles use a bold pill with the semantic color as
// background and the given base color as foreground.
func defaultSemanticBadgeStyles(sp SemanticPalette, base color.Color) (success, warning, errStyle, info lipgloss.Style) {
	badge := func(bg color.Color) lipgloss.Style {
		return lipgloss.NewStyle().
			Background(bg).
			Foreground(base).
			Bold(true).
			Padding(0, 1)
	}
	return badge(sp.Success), badge(sp.Warning), badge(sp.Error), badge(sp.Info)
}

// defaultSemanticTagStyles builds the four semantic tag styles from a
// SemanticPalette. Tag styles use the semantic color as foreground and the
// given surface color as background.
func defaultSemanticTagStyles(sp SemanticPalette, surface color.Color) (success, warning, errStyle, info lipgloss.Style) {
	tag := func(fg color.Color) lipgloss.Style {
		return lipgloss.NewStyle().
			Foreground(fg).
			Background(surface).
			Padding(0, 1)
	}
	return tag(sp.Success), tag(sp.Warning), tag(sp.Error), tag(sp.Info)
}
