package herald

import "charm.land/lipgloss/v2"

// AlertType identifies a GitHub-style alert category.
type AlertType int

const (
	// AlertNote renders a blue informational alert.
	AlertNote AlertType = iota
	// AlertTip renders a green helpful-hint alert.
	AlertTip
	// AlertImportant renders a purple important-information alert.
	AlertImportant
	// AlertWarning renders a yellow/amber warning alert.
	AlertWarning
	// AlertCaution renders a red caution/danger alert.
	AlertCaution
)

// AlertConfig holds the display properties for a single alert type.
type AlertConfig struct {
	Icon  string         // Unicode icon for the header line
	Label string         // Text label (e.g. "Note")
	Style lipgloss.Style // Foreground color applied to bar, icon, label, and content
}

// Default icon constants for each alert type. All are plain Unicode (no emoji)
// for broad terminal compatibility.
const (
	DefaultAlertNoteIcon      = "ℹ"
	DefaultAlertTipIcon       = "✦"
	DefaultAlertImportantIcon = "‼"
	DefaultAlertWarningIcon   = "⚠"
	DefaultAlertCautionIcon   = "◇"
)

// Default label constants for each alert type.
const (
	DefaultAlertNoteLabel      = "Note"
	DefaultAlertTipLabel       = "Tip"
	DefaultAlertImportantLabel = "Important"
	DefaultAlertWarningLabel   = "Warning"
	DefaultAlertCautionLabel   = "Caution"
)
