package herald

import (
	"image/color"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestDefaultSemanticPalette(t *testing.T) {
	p := ColorPalette{
		Primary:   color.Black,
		Secondary: lipgloss.Color("#0000FF"),
		Tertiary:  lipgloss.Color("#00FF00"),
		Accent:    lipgloss.Color("#FFFF00"),
		Highlight: lipgloss.Color("#FF0000"),
		Muted:     lipgloss.Color("#888888"),
		Text:      color.Black,
		Surface:   lipgloss.Color("#333333"),
		Base:      lipgloss.Color("#FFFFFF"),
	}

	sp := DefaultSemanticPalette(p)

	// Success should derive from Tertiary
	if sp.Success != p.Tertiary {
		t.Errorf("Success: expected Tertiary, got different color")
	}
	// Warning should derive from Accent
	if sp.Warning != p.Accent {
		t.Errorf("Warning: expected Accent, got different color")
	}
	// Error should derive from Highlight
	if sp.Error != p.Highlight {
		t.Errorf("Error: expected Highlight, got different color")
	}
	// Info should derive from Secondary
	if sp.Info != p.Secondary {
		t.Errorf("Info: expected Secondary, got different color")
	}
}

func TestSemanticPaletteFromThemeFromPalette(t *testing.T) {
	p := ColorPalette{
		Primary:   color.Black,
		Secondary: lipgloss.Color("#0000FF"),
		Tertiary:  lipgloss.Color("#00FF00"),
		Accent:    lipgloss.Color("#FFFF00"),
		Highlight: lipgloss.Color("#FF0000"),
		Muted:     lipgloss.Color("#888888"),
		Text:      color.Black,
		Surface:   lipgloss.Color("#333333"),
		Base:      lipgloss.Color("#FFFFFF"),
	}

	theme := ThemeFromPalette(p)

	// All 8 semantic styles should render non-empty output
	styles := []struct {
		name  string
		style lipgloss.Style
	}{
		{"SuccessBadge", theme.SuccessBadge},
		{"WarningBadge", theme.WarningBadge},
		{"ErrorBadge", theme.ErrorBadge},
		{"InfoBadge", theme.InfoBadge},
		{"SuccessTag", theme.SuccessTag},
		{"WarningTag", theme.WarningTag},
		{"ErrorTag", theme.ErrorTag},
		{"InfoTag", theme.InfoTag},
	}

	for _, tc := range styles {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.style.Render("test")
			if result == "" {
				t.Errorf("%s rendered empty string", tc.name)
			}
		})
	}
}

func TestSemanticBadgeStylesHaveBoldAndPadding(t *testing.T) {
	sp := SemanticPalette{
		Success: lipgloss.Color("#00FF00"),
		Warning: lipgloss.Color("#FFFF00"),
		Error:   lipgloss.Color("#FF0000"),
		Info:    lipgloss.Color("#0000FF"),
	}
	base := lipgloss.Color("#FFFFFF")

	success, warning, errStyle, info := defaultSemanticBadgeStyles(sp, base)

	for _, tc := range []struct {
		name  string
		style lipgloss.Style
	}{
		{"Success", success},
		{"Warning", warning},
		{"Error", errStyle},
		{"Info", info},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.style.GetBold() {
				t.Errorf("%s badge should be bold", tc.name)
			}
			top, right, bottom, left := tc.style.GetPadding()
			if top != 0 || bottom != 0 || left != 1 || right != 1 {
				t.Errorf("%s badge padding: got (%d,%d,%d,%d), want (0,1,0,1)", tc.name, top, right, bottom, left)
			}
		})
	}
}

func TestSemanticTagStylesHavePadding(t *testing.T) {
	sp := SemanticPalette{
		Success: lipgloss.Color("#00FF00"),
		Warning: lipgloss.Color("#FFFF00"),
		Error:   lipgloss.Color("#FF0000"),
		Info:    lipgloss.Color("#0000FF"),
	}
	surface := lipgloss.Color("#333333")

	success, warning, errStyle, info := defaultSemanticTagStyles(sp, surface)

	for _, tc := range []struct {
		name  string
		style lipgloss.Style
	}{
		{"Success", success},
		{"Warning", warning},
		{"Error", errStyle},
		{"Info", info},
	} {
		t.Run(tc.name, func(t *testing.T) {
			top, right, bottom, left := tc.style.GetPadding()
			if top != 0 || bottom != 0 || left != 1 || right != 1 {
				t.Errorf("%s tag padding: got (%d,%d,%d,%d), want (0,1,0,1)", tc.name, top, right, bottom, left)
			}
		})
	}
}

func TestSemanticStylesInAllThemes(t *testing.T) {
	themes := map[string]func() Theme{
		"Default":    DefaultTheme,
		"Dracula":    DraculaTheme,
		"Catppuccin": CatppuccinTheme,
		"Base16":     Base16Theme,
		"Charm":      CharmTheme,
	}

	for name, themeFn := range themes {
		t.Run(name, func(t *testing.T) {
			theme := themeFn()
			styles := []struct {
				label string
				style lipgloss.Style
			}{
				{"SuccessBadge", theme.SuccessBadge},
				{"WarningBadge", theme.WarningBadge},
				{"ErrorBadge", theme.ErrorBadge},
				{"InfoBadge", theme.InfoBadge},
				{"SuccessTag", theme.SuccessTag},
				{"WarningTag", theme.WarningTag},
				{"ErrorTag", theme.ErrorTag},
				{"InfoTag", theme.InfoTag},
			}

			for _, s := range styles {
				result := s.style.Render("test")
				if result == "" {
					t.Errorf("%s: %s rendered empty", name, s.label)
				}
			}
		})
	}
}
