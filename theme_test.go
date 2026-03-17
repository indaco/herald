package herald

import (
	"testing"
)

func TestDefaultTheme(t *testing.T) {
	theme := DefaultTheme()

	t.Run("bullet char", func(t *testing.T) {
		if theme.BulletChar == "" {
			t.Error("BulletChar should not be empty")
		}
	})

	t.Run("HR char", func(t *testing.T) {
		if theme.HRChar == "" {
			t.Error("HRChar should not be empty")
		}
	})

	t.Run("HR width", func(t *testing.T) {
		if theme.HRWidth <= 0 {
			t.Errorf("HRWidth should be positive, got %d", theme.HRWidth)
		}
	})

	t.Run("blockquote bar", func(t *testing.T) {
		if theme.BlockquoteBar == "" {
			t.Error("BlockquoteBar should not be empty")
		}
	})

	t.Run("default values", func(t *testing.T) {
		if theme.BulletChar != "\u2022" {
			t.Errorf("expected bullet %q, got %q", "\u2022", theme.BulletChar)
		}
		if theme.HRWidth != 40 {
			t.Errorf("expected HRWidth 40, got %d", theme.HRWidth)
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
