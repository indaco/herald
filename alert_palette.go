package herald

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// AlertPalette defines the five colors used to derive alert styles. It is
// separate from ColorPalette because the existing 9-color palette does not
// cleanly map to 5 alert semantics.
type AlertPalette struct {
	Note      color.Color // blue
	Tip       color.Color // green
	Important color.Color // purple
	Warning   color.Color // yellow/amber
	Caution   color.Color // red
}

// DefaultAlertConfigs builds a full map[AlertType]AlertConfig from an
// AlertPalette, using default icons and labels and creating a lipgloss.Style
// with the palette color as foreground.
func DefaultAlertConfigs(ap AlertPalette) map[AlertType]AlertConfig {
	return map[AlertType]AlertConfig{
		AlertNote: {
			Icon:  DefaultAlertNoteIcon,
			Label: DefaultAlertNoteLabel,
			Style: lipgloss.NewStyle().Foreground(ap.Note),
		},
		AlertTip: {
			Icon:  DefaultAlertTipIcon,
			Label: DefaultAlertTipLabel,
			Style: lipgloss.NewStyle().Foreground(ap.Tip),
		},
		AlertImportant: {
			Icon:  DefaultAlertImportantIcon,
			Label: DefaultAlertImportantLabel,
			Style: lipgloss.NewStyle().Foreground(ap.Important),
		},
		AlertWarning: {
			Icon:  DefaultAlertWarningIcon,
			Label: DefaultAlertWarningLabel,
			Style: lipgloss.NewStyle().Foreground(ap.Warning),
		},
		AlertCaution: {
			Icon:  DefaultAlertCautionIcon,
			Label: DefaultAlertCautionLabel,
			Style: lipgloss.NewStyle().Foreground(ap.Caution),
		},
	}
}
