package herald

import (
	"charm.land/lipgloss/v2"
)

// DraculaTheme returns a Theme based on the Dracula color palette.
// Colors match huh's ThemeDracula. Dracula is a dark-only palette;
// adaptive colors are used for API consistency.
func DraculaTheme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	return ThemeFromPalette(ColorPalette{
		Primary:   lightDark(lipgloss.Color("#282a36"), lipgloss.Color("#f8f8f2")),
		Secondary: lightDark(lipgloss.Color("#6d3fc0"), lipgloss.Color("#bd93f9")),
		Tertiary:  lightDark(lipgloss.Color("#1e9651"), lipgloss.Color("#50fa7b")),
		Accent:    lightDark(lipgloss.Color("#b07d2b"), lipgloss.Color("#f1fa8c")),
		Highlight: lightDark(lipgloss.Color("#d63384"), lipgloss.Color("#ff79c6")),
		Muted:     lightDark(lipgloss.Color("#6272a4"), lipgloss.Color("#6272a4")),
		Text:      lightDark(lipgloss.Color("#282a36"), lipgloss.Color("#f8f8f2")),
		Surface:   lightDark(lipgloss.Color("#e8e6ef"), lipgloss.Color("#44475a")),
		Base:      lightDark(lipgloss.Color("#f8f8f2"), lipgloss.Color("#282a36")),
	})
}

// CatppuccinTheme returns a Theme based on the Catppuccin color palette.
// Adapts to the terminal background: Mocha on dark, Latte on light.
// Colors match huh's ThemeCatppuccin.
func CatppuccinTheme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	return ThemeFromPalette(ColorPalette{
		Primary:   lightDark(lipgloss.Color("#4c4f69"), lipgloss.Color("#cdd6f4")),
		Secondary: lightDark(lipgloss.Color("#8839ef"), lipgloss.Color("#cba6f7")),
		Tertiary:  lightDark(lipgloss.Color("#179299"), lipgloss.Color("#94e2d5")),
		Accent:    lightDark(lipgloss.Color("#ea76cb"), lipgloss.Color("#f5c2e7")),
		Highlight: lightDark(lipgloss.Color("#e64553"), lipgloss.Color("#eba0ac")),
		Muted:     lightDark(lipgloss.Color("#7c7f93"), lipgloss.Color("#6c7086")),
		Text:      lightDark(lipgloss.Color("#4c4f69"), lipgloss.Color("#cdd6f4")),
		Surface:   lightDark(lipgloss.Color("#ccd0da"), lipgloss.Color("#313244")),
		Base:      lightDark(lipgloss.Color("#eff1f5"), lipgloss.Color("#1e1e2e")),
	})
}

// Base16Theme returns a Theme based on ANSI base16 terminal colors.
// Colors match huh's ThemeBase16. Base16 uses standard ANSI color indices;
// adaptive colors are used for API consistency.
func Base16Theme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	return ThemeFromPalette(ColorPalette{
		Primary:   lightDark(lipgloss.Color("0"), lipgloss.Color("7")),  // black / white
		Secondary: lightDark(lipgloss.Color("4"), lipgloss.Color("6")),  // blue / cyan
		Tertiary:  lightDark(lipgloss.Color("2"), lipgloss.Color("2")),  // green
		Accent:    lightDark(lipgloss.Color("5"), lipgloss.Color("3")),  // magenta / yellow
		Highlight: lightDark(lipgloss.Color("1"), lipgloss.Color("1")),  // red
		Muted:     lightDark(lipgloss.Color("8"), lipgloss.Color("8")),  // bright black
		Text:      lightDark(lipgloss.Color("0"), lipgloss.Color("7")),  // black / white
		Surface:   lightDark(lipgloss.Color("15"), lipgloss.Color("8")), // bright white / bright black
		Base:      lightDark(lipgloss.Color("7"), lipgloss.Color("0")),  // white / black
	})
}

// CharmTheme returns a Theme based on Charm's brand color palette.
// Auto-detects terminal background for light/dark variants.
// Colors match huh's ThemeCharm.
func CharmTheme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	return ThemeFromPalette(ColorPalette{
		Primary:   lightDark(lipgloss.Color("235"), lipgloss.Color("#FFFDF5")),
		Secondary: lightDark(lipgloss.Color("#5A56E0"), lipgloss.Color("#7571F9")),
		Tertiary:  lightDark(lipgloss.Color("#02BA84"), lipgloss.Color("#02BF87")),
		Accent:    lightDark(lipgloss.Color("#C740B0"), lipgloss.Color("#F780E2")),
		Highlight: lightDark(lipgloss.Color("#C7304E"), lipgloss.Color("#ED567A")),
		Muted:     lightDark(lipgloss.Color("243"), lipgloss.Color("243")),
		Text:      lightDark(lipgloss.Color("235"), lipgloss.Color("#FFFDF5")),
		Surface:   lightDark(lipgloss.Color("254"), lipgloss.Color("238")),
		Base:      lightDark(lipgloss.Color("252"), lipgloss.Color("236")),
	})
}
