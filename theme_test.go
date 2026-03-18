package herald

import (
	"testing"
)

func TestDefaultTheme(t *testing.T) {
	theme := DefaultTheme()
	assertDefaultTokens(t, theme)
}

func assertDefaultTokens(t *testing.T, theme Theme) {
	t.Helper()

	if theme.BulletChar == "" {
		t.Error("BulletChar should not be empty")
	}
	if theme.HRChar == "" {
		t.Error("HRChar should not be empty")
	}
	if theme.HRWidth <= 0 {
		t.Errorf("HRWidth should be positive, got %d", theme.HRWidth)
	}
	if theme.BlockquoteBar == "" {
		t.Error("BlockquoteBar should not be empty")
	}
	if theme.BulletChar != DefaultBulletChar {
		t.Errorf("expected bullet %q, got %q", DefaultBulletChar, theme.BulletChar)
	}
	if theme.HRWidth != DefaultHRWidth {
		t.Errorf("expected HRWidth %d, got %d", DefaultHRWidth, theme.HRWidth)
	}
}

func TestHasDarkBGEnvOverride(t *testing.T) {
	t.Run("force dark", func(t *testing.T) {
		t.Setenv("HERALD_FORCE_DARK", "1")
		if !hasDarkBG() {
			t.Error("expected dark when HERALD_FORCE_DARK=1")
		}
	})

	t.Run("force light", func(t *testing.T) {
		t.Setenv("HERALD_FORCE_DARK", "0")
		if hasDarkBG() {
			t.Error("expected light when HERALD_FORCE_DARK=0")
		}
	})

	t.Run("force dark with true", func(t *testing.T) {
		t.Setenv("HERALD_FORCE_DARK", "true")
		if !hasDarkBG() {
			t.Error("expected dark when HERALD_FORCE_DARK=true")
		}
	})

	t.Run("force light with false", func(t *testing.T) {
		t.Setenv("HERALD_FORCE_DARK", "false")
		if hasDarkBG() {
			t.Error("expected light when HERALD_FORCE_DARK=false")
		}
	})
}

func TestDefaultThemeIdempotent(t *testing.T) {
	t1 := DefaultTheme()
	t2 := DefaultTheme()

	if t1.HRWidth != t2.HRWidth {
		t.Error("DefaultTheme should return consistent values")
	}
	if t1.BulletChar != t2.BulletChar {
		t.Error("DefaultTheme should return consistent values")
	}
}
