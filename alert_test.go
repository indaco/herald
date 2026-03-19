package herald

import (
	"image/color"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestAlertRendering(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name      string
		alertType AlertType
		text      string
		icon      string
		label     string
	}{
		{"Note", AlertNote, "Useful information.", DefaultAlertNoteIcon, DefaultAlertNoteLabel},
		{"Tip", AlertTip, "A helpful hint.", DefaultAlertTipIcon, DefaultAlertTipLabel},
		{"Important", AlertImportant, "Key information.", DefaultAlertImportantIcon, DefaultAlertImportantLabel},
		{"Warning", AlertWarning, "Be careful.", DefaultAlertWarningIcon, DefaultAlertWarningLabel},
		{"Caution", AlertCaution, "Danger ahead.", DefaultAlertCautionIcon, DefaultAlertCautionLabel},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := stripANSI(ty.Alert(tc.alertType, tc.text))

			if !strings.Contains(result, tc.icon) {
				t.Errorf("expected icon %q in %q", tc.icon, result)
			}
			if !strings.Contains(result, tc.label) {
				t.Errorf("expected label %q in %q", tc.label, result)
			}
			if !strings.Contains(result, tc.text) {
				t.Errorf("expected text %q in %q", tc.text, result)
			}
			// Every line should have the bar
			bar := ty.theme.AlertBar
			for line := range strings.SplitSeq(result, "\n") {
				if !strings.HasPrefix(strings.TrimSpace(line), bar) {
					t.Errorf("expected line to start with bar %q, got %q", bar, line)
				}
			}
		})
	}
}

func TestAlertMultiline(t *testing.T) {
	ty := newTestTypography()
	text := "First line.\nSecond line.\nThird line."
	result := stripANSI(ty.Warning(text))

	bar := ty.theme.AlertBar
	lines := strings.Split(result, "\n")
	// Header line + 3 content lines = 4 total
	if len(lines) != 4 {
		t.Errorf("expected 4 lines, got %d: %q", len(lines), result)
	}
	for i, line := range lines {
		if !strings.Contains(line, bar) {
			t.Errorf("line %d missing bar: %q", i, line)
		}
	}
}

func TestAlertEmptyText(t *testing.T) {
	ty := newTestTypography()
	result := stripANSI(ty.Note(""))

	// Should still have header with icon and label
	if !strings.Contains(result, DefaultAlertNoteIcon) {
		t.Errorf("empty alert should still have icon, got %q", result)
	}
	if !strings.Contains(result, DefaultAlertNoteLabel) {
		t.Errorf("empty alert should still have label, got %q", result)
	}
}

func TestAlertConvenienceMethods(t *testing.T) {
	ty := newTestTypography()

	tests := []struct {
		name      string
		fn        func(string) string
		alertType AlertType
	}{
		{"Note", ty.Note, AlertNote},
		{"Tip", ty.Tip, AlertTip},
		{"Important", ty.Important, AlertImportant},
		{"Warning", ty.Warning, AlertWarning},
		{"Caution", ty.Caution, AlertCaution},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			text := "Test text for " + tc.name
			convenience := stripANSI(tc.fn(text))
			direct := stripANSI(ty.Alert(tc.alertType, text))
			if convenience != direct {
				t.Errorf("%s() != Alert(%d, ...)\n  got:  %q\n  want: %q", tc.name, tc.alertType, convenience, direct)
			}
		})
	}
}

func TestAlertCustomIcon(t *testing.T) {
	ty := New(WithAlertIcon(AlertNote, ">>"))
	result := stripANSI(ty.Note("Hello"))

	if !strings.Contains(result, ">>") {
		t.Errorf("expected custom icon '>>' in %q", result)
	}
	if strings.Contains(result, DefaultAlertNoteIcon) {
		t.Errorf("should not contain default icon %q in %q", DefaultAlertNoteIcon, result)
	}
}

func TestAlertCustomLabel(t *testing.T) {
	ty := New(WithAlertLabel(AlertTip, "Hint"))
	result := stripANSI(ty.Tip("Hello"))

	if !strings.Contains(result, "Hint") {
		t.Errorf("expected custom label 'Hint' in %q", result)
	}
	if strings.Contains(result, DefaultAlertTipLabel) {
		t.Errorf("should not contain default label %q in %q", DefaultAlertTipLabel, result)
	}
}

func TestAlertCustomStyle(t *testing.T) {
	customStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	ty := New(WithAlertStyle(AlertWarning, customStyle))

	// Should not panic and should render
	result := ty.Warning("styled")
	if result == "" {
		t.Error("WithAlertStyle should produce non-empty output")
	}
}

func TestAlertBar(t *testing.T) {
	ty := New(WithAlertBar("║"))
	result := stripANSI(ty.Note("Hello"))

	if !strings.Contains(result, "║") {
		t.Errorf("expected custom bar '║' in %q", result)
	}
}

func TestAlertPalette(t *testing.T) {
	red := lipgloss.Color("#FF0000")
	green := lipgloss.Color("#00FF00")
	blue := lipgloss.Color("#0000FF")
	yellow := lipgloss.Color("#FFFF00")
	purple := lipgloss.Color("#800080")

	ty := New(WithAlertPalette(AlertPalette{
		Note:      blue,
		Tip:       green,
		Important: purple,
		Warning:   yellow,
		Caution:   red,
	}))

	// All 5 types should work and produce output
	for _, at := range []AlertType{AlertNote, AlertTip, AlertImportant, AlertWarning, AlertCaution} {
		result := ty.Alert(at, "test")
		if result == "" {
			t.Errorf("WithAlertPalette: AlertType %d produced empty output", at)
		}
	}
}

func TestAlertInAllThemes(t *testing.T) {
	themes := map[string]func() Theme{
		"Default":    DefaultTheme,
		"Dracula":    DraculaTheme,
		"Catppuccin": CatppuccinTheme,
		"Base16":     Base16Theme,
		"Charm":      CharmTheme,
	}

	alertTypes := []AlertType{AlertNote, AlertTip, AlertImportant, AlertWarning, AlertCaution}

	for name, themeFn := range themes {
		t.Run(name, func(t *testing.T) {
			theme := themeFn()
			if theme.Alerts == nil {
				t.Fatalf("%s theme has nil Alerts map", name)
			}
			for _, at := range alertTypes {
				cfg, ok := theme.Alerts[at]
				if !ok {
					t.Errorf("%s theme missing AlertType %d", name, at)
					continue
				}
				if cfg.Icon == "" {
					t.Errorf("%s theme AlertType %d has empty icon", name, at)
				}
				if cfg.Label == "" {
					t.Errorf("%s theme AlertType %d has empty label", name, at)
				}
			}
			if theme.AlertBar == "" {
				t.Errorf("%s theme has empty AlertBar", name)
			}
		})
	}
}

func TestAlertFallbackToBlockquote(t *testing.T) {
	// Create a typography with an empty Alerts map to trigger fallback
	ty := New(WithTheme(Theme{
		Alerts:        map[AlertType]AlertConfig{},
		AlertBar:      DefaultAlertBar,
		BlockquoteBar: DefaultBlockquoteBar,
		Blockquote:    lipgloss.NewStyle(),
	}))

	// An unknown/unconfigured alert type should fall back to blockquote
	result := stripANSI(ty.Alert(AlertNote, "fallback text"))
	if !strings.Contains(result, "fallback text") {
		t.Errorf("fallback should contain text, got %q", result)
	}
	if !strings.Contains(result, DefaultBlockquoteBar) {
		t.Errorf("fallback should use blockquote bar, got %q", result)
	}
}

func TestAlertNilMapGuard(t *testing.T) {
	nilTheme := Theme{
		BlockquoteBar: DefaultBlockquoteBar,
		Blockquote:    lipgloss.NewStyle(),
		AlertBar:      DefaultAlertBar,
	}

	t.Run("WithAlertIcon on nil map", func(t *testing.T) {
		ty := New(WithTheme(nilTheme), WithAlertIcon(AlertNote, "!"))
		result := stripANSI(ty.Note("test"))
		if !strings.Contains(result, "!") {
			t.Errorf("expected custom icon '!' in %q", result)
		}
	})

	t.Run("WithAlertLabel on nil map", func(t *testing.T) {
		ty := New(WithTheme(nilTheme), WithAlertLabel(AlertNote, "Info"))
		result := stripANSI(ty.Note("test"))
		if !strings.Contains(result, "Info") {
			t.Errorf("expected custom label 'Info' in %q", result)
		}
	})

	t.Run("WithAlertStyle on nil map", func(t *testing.T) {
		ty := New(WithTheme(nilTheme), WithAlertStyle(AlertNote, lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))))
		result := ty.Note("test")
		if result == "" {
			t.Error("expected non-empty output")
		}
	})
}

func TestAlertPaletteFromColorPalette(t *testing.T) {
	// Verify that ThemeFromPalette produces non-nil Alerts
	theme := ThemeFromPalette(ColorPalette{
		Primary:   color.Black,
		Secondary: color.Black,
		Tertiary:  color.Black,
		Accent:    color.Black,
		Highlight: color.Black,
		Muted:     color.Black,
		Text:      color.Black,
		Surface:   color.Black,
		Base:      color.Black,
	})

	if theme.Alerts == nil {
		t.Fatal("ThemeFromPalette should produce non-nil Alerts")
	}
	if len(theme.Alerts) != 5 {
		t.Errorf("expected 5 alert configs, got %d", len(theme.Alerts))
	}
	if theme.AlertBar == "" {
		t.Error("ThemeFromPalette should set AlertBar")
	}
}
