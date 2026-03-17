// Package herald provides a reusable TUI typography library inspired by CSS
// frameworks such as Tailwind Typography, Bootstrap, and Pico CSS.
//
// It offers HTML-analogous rendering functions (H1-H6, P, Blockquote, lists,
// inline styles, etc.) that output styled strings via lipgloss v2.
//
// Quick start:
//
//	ty := herald.New()                    // default theme
//	fmt.Println(ty.H1("Hello, World!"))
//	fmt.Println(ty.P("Some body text."))
//	fmt.Println(ty.UL("Apples", "Bananas", "Cherries"))
//
// Customization via functional options:
//
//	ty := herald.New(
//	    herald.WithHRWidth(60),
//	    herald.WithBulletChar("-"),
//	    herald.WithH1Style(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))),
//	)
package herald
