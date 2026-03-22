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

func TestBoxBorderSet(t *testing.T) {
	bs := BoxBorderSet()

	fields := []struct {
		name string
		val  string
		want string
	}{
		{"Top", bs.Top, "─"},
		{"Bottom", bs.Bottom, "─"},
		{"Left", bs.Left, "│"},
		{"Right", bs.Right, "│"},
		{"Header", bs.Header, "─"},
		{"Row", bs.Row, "─"},
		{"TopLeft", bs.TopLeft, "┌"},
		{"TopRight", bs.TopRight, "┐"},
		{"BottomLeft", bs.BottomLeft, "└"},
		{"BottomRight", bs.BottomRight, "┘"},
		{"TopJunction", bs.TopJunction, "┬"},
		{"BottomJunction", bs.BottomJunction, "┴"},
		{"LeftJunction", bs.LeftJunction, "├"},
		{"RightJunction", bs.RightJunction, "┤"},
		{"Cross", bs.Cross, "┼"},
		{"HeaderLeft", bs.HeaderLeft, "├"},
		{"HeaderRight", bs.HeaderRight, "┤"},
		{"HeaderCross", bs.HeaderCross, "┼"},
		{"FooterLeft", bs.FooterLeft, "├"},
		{"FooterRight", bs.FooterRight, "┤"},
		{"FooterCross", bs.FooterCross, "┼"},
	}

	for _, f := range fields {
		t.Run(f.name, func(t *testing.T) {
			if f.val != f.want {
				t.Errorf("BoxBorderSet().%s = %q, want %q", f.name, f.val, f.want)
			}
		})
	}
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
