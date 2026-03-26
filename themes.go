package herald

import (
	"charm.land/lipgloss/v2"
)

// DraculaTheme returns a Theme based on the Dracula color palette.
// Colors match huh's ThemeDracula. Dracula is a dark-only palette;
// adaptive colors are used for API consistency.
func DraculaTheme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	theme := ThemeFromPalette(ColorPalette{
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

	theme.Alerts = DefaultAlertConfigs(AlertPalette{
		Note:      lightDark(lipgloss.Color("#1e6e99"), lipgloss.Color("#8be9fd")), // cyan
		Tip:       lightDark(lipgloss.Color("#1e9651"), lipgloss.Color("#50fa7b")), // green
		Important: lightDark(lipgloss.Color("#6d3fc0"), lipgloss.Color("#bd93f9")), // purple
		Warning:   lightDark(lipgloss.Color("#b07d2b"), lipgloss.Color("#f1fa8c")), // yellow
		Caution:   lightDark(lipgloss.Color("#c44040"), lipgloss.Color("#ff5555")), // red
	})

	// Semantic palette: Dracula green/yellow/red/cyan
	draculaSP := SemanticPalette{
		Success: lightDark(lipgloss.Color("#1e9651"), lipgloss.Color("#50fa7b")), // green
		Warning: lightDark(lipgloss.Color("#b07d2b"), lipgloss.Color("#f1fa8c")), // yellow
		Error:   lightDark(lipgloss.Color("#c44040"), lipgloss.Color("#ff5555")), // red
		Info:    lightDark(lipgloss.Color("#1e6e99"), lipgloss.Color("#8be9fd")), // cyan
	}
	base := theme.Badge.GetForeground()
	surface := theme.Tag.GetBackground()
	theme.SuccessBadge, theme.WarningBadge, theme.ErrorBadge, theme.InfoBadge = defaultSemanticBadgeStyles(draculaSP, base)
	theme.SuccessTag, theme.WarningTag, theme.ErrorTag, theme.InfoTag = defaultSemanticTagStyles(draculaSP, surface)

	return theme
}

// CatppuccinTheme returns a Theme based on the Catppuccin color palette.
// Adapts to the terminal background: Mocha on dark, Latte on light.
// Colors match huh's ThemeCatppuccin.
func CatppuccinTheme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	theme := ThemeFromPalette(ColorPalette{
		Primary:   lightDark(lipgloss.Color("#4c4f69"), lipgloss.Color("#cdd6f4")),
		Secondary: lightDark(lipgloss.Color("#8839ef"), lipgloss.Color("#cba6f7")),
		Tertiary:  lightDark(lipgloss.Color("#179299"), lipgloss.Color("#94e2d5")),
		Accent:    lightDark(lipgloss.Color("#ea76cb"), lipgloss.Color("#f5c2e7")),
		Highlight: lightDark(lipgloss.Color("#e64553"), lipgloss.Color("#eba0ac")),
		Muted:     lightDark(lipgloss.Color("#9ca0b0"), lipgloss.Color("#6c7086")),
		Text:      lightDark(lipgloss.Color("#4c4f69"), lipgloss.Color("#cdd6f4")),
		Surface:   lightDark(lipgloss.Color("#ccd0da"), lipgloss.Color("#313244")),
		Base:      lightDark(lipgloss.Color("#eff1f5"), lipgloss.Color("#1e1e2e")),
	})

	theme.Alerts = DefaultAlertConfigs(AlertPalette{
		Note:      lightDark(lipgloss.Color("#1e66f5"), lipgloss.Color("#89b4fa")), // blue
		Tip:       lightDark(lipgloss.Color("#40a02b"), lipgloss.Color("#a6e3a1")), // green
		Important: lightDark(lipgloss.Color("#8839ef"), lipgloss.Color("#cba6f7")), // mauve
		Warning:   lightDark(lipgloss.Color("#fe640b"), lipgloss.Color("#fab387")), // peach
		Caution:   lightDark(lipgloss.Color("#d20f39"), lipgloss.Color("#f38ba8")), // red
	})

	// Semantic palette: Catppuccin green/peach/red/blue
	catppuccinSP := SemanticPalette{
		Success: lightDark(lipgloss.Color("#40a02b"), lipgloss.Color("#a6e3a1")), // green
		Warning: lightDark(lipgloss.Color("#fe640b"), lipgloss.Color("#fab387")), // peach
		Error:   lightDark(lipgloss.Color("#d20f39"), lipgloss.Color("#f38ba8")), // red
		Info:    lightDark(lipgloss.Color("#1e66f5"), lipgloss.Color("#89b4fa")), // blue
	}
	base := theme.Badge.GetForeground()
	surface := theme.Tag.GetBackground()
	theme.SuccessBadge, theme.WarningBadge, theme.ErrorBadge, theme.InfoBadge = defaultSemanticBadgeStyles(catppuccinSP, base)
	theme.SuccessTag, theme.WarningTag, theme.ErrorTag, theme.InfoTag = defaultSemanticTagStyles(catppuccinSP, surface)

	return theme
}

// Base16Theme returns a Theme based on ANSI base16 terminal colors.
// Colors match huh's ThemeBase16. Base16 uses standard ANSI color indices;
// adaptive colors are used for API consistency.
func Base16Theme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	theme := ThemeFromPalette(ColorPalette{
		Primary:   lightDark(lipgloss.Color("0"), lipgloss.Color("7")), // black / white
		Secondary: lightDark(lipgloss.Color("4"), lipgloss.Color("6")), // blue / cyan
		Tertiary:  lightDark(lipgloss.Color("2"), lipgloss.Color("2")), // green
		Accent:    lightDark(lipgloss.Color("5"), lipgloss.Color("3")), // magenta / yellow
		Highlight: lightDark(lipgloss.Color("1"), lipgloss.Color("1")), // red
		Muted:     lightDark(lipgloss.Color("8"), lipgloss.Color("8")), // bright black
		Text:      lightDark(lipgloss.Color("0"), lipgloss.Color("7")), // black / white
		Surface:   lightDark(lipgloss.Color("7"), lipgloss.Color("8")), // white / bright black
		Base:      lightDark(lipgloss.Color("7"), lipgloss.Color("0")), // white / black
	})

	theme.Alerts = DefaultAlertConfigs(AlertPalette{
		Note:      lightDark(lipgloss.Color("4"), lipgloss.Color("4")), // blue
		Tip:       lightDark(lipgloss.Color("2"), lipgloss.Color("2")), // green
		Important: lightDark(lipgloss.Color("5"), lipgloss.Color("5")), // magenta
		Warning:   lightDark(lipgloss.Color("3"), lipgloss.Color("3")), // yellow
		Caution:   lightDark(lipgloss.Color("1"), lipgloss.Color("1")), // red
	})

	// Semantic palette: ANSI green/yellow/red/blue
	base16SP := SemanticPalette{
		Success: lightDark(lipgloss.Color("2"), lipgloss.Color("2")), // green
		Warning: lightDark(lipgloss.Color("3"), lipgloss.Color("3")), // yellow
		Error:   lightDark(lipgloss.Color("1"), lipgloss.Color("1")), // red
		Info:    lightDark(lipgloss.Color("4"), lipgloss.Color("4")), // blue
	}
	base := theme.Badge.GetForeground()
	surface := theme.Tag.GetBackground()
	theme.SuccessBadge, theme.WarningBadge, theme.ErrorBadge, theme.InfoBadge = defaultSemanticBadgeStyles(base16SP, base)
	theme.SuccessTag, theme.WarningTag, theme.ErrorTag, theme.InfoTag = defaultSemanticTagStyles(base16SP, surface)

	return theme
}

// CharmTheme returns a Theme based on Charm's brand color palette.
// Auto-detects terminal background for light/dark variants.
// Colors match huh's ThemeCharm.
func CharmTheme() Theme {
	lightDark := lipgloss.LightDark(hasDarkBG())

	theme := ThemeFromPalette(ColorPalette{
		Primary:   lightDark(lipgloss.Color("235"), lipgloss.Color("#FFFDF5")),
		Secondary: lightDark(lipgloss.Color("#5A56E0"), lipgloss.Color("#7571F9")),
		Tertiary:  lightDark(lipgloss.Color("#02BA84"), lipgloss.Color("#02BF87")),
		Accent:    lightDark(lipgloss.Color("#C740B0"), lipgloss.Color("#F780E2")),
		Highlight: lightDark(lipgloss.Color("#C7304E"), lipgloss.Color("#ED567A")),
		Muted:     lightDark(lipgloss.Color("243"), lipgloss.Color("243")),
		Text:      lightDark(lipgloss.Color("235"), lipgloss.Color("#FFFDF5")),
		Surface:   lightDark(lipgloss.Color("254"), lipgloss.Color("238")),
		Base:      lightDark(lipgloss.Color("255"), lipgloss.Color("236")),
	})

	theme.Alerts = DefaultAlertConfigs(AlertPalette{
		Note:      lightDark(lipgloss.Color("#2563EB"), lipgloss.Color("#60A5FA")), // charm blue
		Tip:       lightDark(lipgloss.Color("#02BA84"), lipgloss.Color("#02BF87")), // charm green
		Important: lightDark(lipgloss.Color("#5A56E0"), lipgloss.Color("#7571F9")), // charm purple
		Warning:   lightDark(lipgloss.Color("#C740B0"), lipgloss.Color("#F780E2")), // charm pink/warm
		Caution:   lightDark(lipgloss.Color("#C7304E"), lipgloss.Color("#ED567A")), // charm red
	})

	// Semantic palette: Charm green/pink/red/blue
	charmSP := SemanticPalette{
		Success: lightDark(lipgloss.Color("#02BA84"), lipgloss.Color("#02BF87")), // charm green
		Warning: lightDark(lipgloss.Color("#C740B0"), lipgloss.Color("#F780E2")), // charm pink/warm
		Error:   lightDark(lipgloss.Color("#C7304E"), lipgloss.Color("#ED567A")), // charm red
		Info:    lightDark(lipgloss.Color("#2563EB"), lipgloss.Color("#60A5FA")), // charm blue
	}
	base := theme.Badge.GetForeground()
	surface := theme.Tag.GetBackground()
	theme.SuccessBadge, theme.WarningBadge, theme.ErrorBadge, theme.InfoBadge = defaultSemanticBadgeStyles(charmSP, base)
	theme.SuccessTag, theme.WarningTag, theme.ErrorTag, theme.InfoTag = defaultSemanticTagStyles(charmSP, surface)

	return theme
}
