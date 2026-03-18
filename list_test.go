package herald

import (
	"strings"
	"testing"
)

func TestItem(t *testing.T) {
	li := Item("hello")
	if li.Text != "hello" {
		t.Errorf("expected %q, got %q", "hello", li.Text)
	}
	if len(li.Children) != 0 {
		t.Errorf("expected no children, got %d", len(li.Children))
	}
}

func TestItems(t *testing.T) {
	items := Items("a", "b", "c")
	if len(items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(items))
	}
	for i, text := range []string{"a", "b", "c"} {
		if items[i].Text != text {
			t.Errorf("items[%d]: expected %q, got %q", i, text, items[i].Text)
		}
	}
}

func TestItemWithChildren(t *testing.T) {
	li := ItemWithChildren("parent", Item("child1"), Item("child2"))
	if li.Text != "parent" {
		t.Errorf("expected %q, got %q", "parent", li.Text)
	}
	if len(li.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(li.Children))
	}
	if li.Kind != Unordered {
		t.Errorf("expected Unordered, got %d", li.Kind)
	}
}

func TestItemWithOLChildren(t *testing.T) {
	li := ItemWithOLChildren("parent", Item("child1"))
	if li.Kind != Ordered {
		t.Errorf("expected Ordered, got %d", li.Kind)
	}
}

func TestNestUL(t *testing.T) {
	ty := New()

	t.Run("empty", func(t *testing.T) {
		result := ty.NestUL()
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})

	t.Run("flat items", func(t *testing.T) {
		result := stripANSI(ty.NestUL(Item("a"), Item("b")))
		lines := strings.Split(result, "\n")
		if len(lines) != 2 {
			t.Fatalf("expected 2 lines, got %d: %q", len(lines), result)
		}
		// First level uses first bullet char
		if !strings.Contains(lines[0], "a") {
			t.Errorf("expected line to contain %q, got %q", "a", lines[0])
		}
	})

	t.Run("nested items", func(t *testing.T) {
		result := stripANSI(ty.NestUL(
			Item("top"),
			ItemWithChildren("parent",
				Item("child1"),
				Item("child2"),
			),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 4 {
			t.Fatalf("expected 4 lines, got %d: %q", len(lines), result)
		}
		// Children should be indented
		if !strings.HasPrefix(lines[2], "  ") {
			t.Errorf("expected child line to be indented, got %q", lines[2])
		}
	})

	t.Run("bullet cycling", func(t *testing.T) {
		// Build 3 levels deep
		result := stripANSI(ty.NestUL(
			ItemWithChildren("l0",
				ItemWithChildren("l1",
					Item("l2"),
				),
			),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 3 {
			t.Fatalf("expected 3 lines, got %d: %q", len(lines), result)
		}
		// Default bullets: "•", "◦", "▪", "▹"
		bullets := DefaultNestedBulletChars
		if !strings.Contains(lines[0], bullets[0]) {
			t.Errorf("depth 0: expected bullet %q in %q", bullets[0], lines[0])
		}
		if !strings.Contains(lines[1], bullets[1]) {
			t.Errorf("depth 1: expected bullet %q in %q", bullets[1], lines[1])
		}
		if !strings.Contains(lines[2], bullets[2]) {
			t.Errorf("depth 2: expected bullet %q in %q", bullets[2], lines[2])
		}
	})
}

func TestNestOL(t *testing.T) {
	ty := New()

	t.Run("empty", func(t *testing.T) {
		result := ty.NestOL()
		if result != "" {
			t.Errorf("expected empty string, got %q", result)
		}
	})

	t.Run("flat items", func(t *testing.T) {
		result := stripANSI(ty.NestOL(Item("a"), Item("b"), Item("c")))
		lines := strings.Split(result, "\n")
		if len(lines) != 3 {
			t.Fatalf("expected 3 lines, got %d", len(lines))
		}
		if !strings.Contains(lines[0], "1.") {
			t.Errorf("expected %q in %q", "1.", lines[0])
		}
		if !strings.Contains(lines[2], "3.") {
			t.Errorf("expected %q in %q", "3.", lines[2])
		}
	})

	t.Run("nested numbering resets", func(t *testing.T) {
		result := stripANSI(ty.NestOL(
			ItemWithOLChildren("parent",
				Item("child1"),
				Item("child2"),
			),
			Item("second"),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 4 {
			t.Fatalf("expected 4 lines, got %d: %q", len(lines), result)
		}
		// Parent is 1., second is 2.
		if !strings.Contains(lines[0], "1.") {
			t.Errorf("expected parent to be 1., got %q", lines[0])
		}
		if !strings.Contains(lines[3], "2.") {
			t.Errorf("expected second to be 2., got %q", lines[3])
		}
		// Children restart at 1.
		if !strings.Contains(lines[1], "1.") {
			t.Errorf("expected child1 to be 1., got %q", lines[1])
		}
		if !strings.Contains(lines[2], "2.") {
			t.Errorf("expected child2 to be 2., got %q", lines[2])
		}
	})
}

func TestMixedNesting(t *testing.T) {
	ty := New()

	t.Run("OL children inside UL", func(t *testing.T) {
		result := stripANSI(ty.NestUL(
			ItemWithOLChildren("parent",
				Item("first"),
				Item("second"),
			),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 3 {
			t.Fatalf("expected 3 lines, got %d: %q", len(lines), result)
		}
		// Children should have numbered markers
		if !strings.Contains(lines[1], "1.") {
			t.Errorf("expected numbered child, got %q", lines[1])
		}
	})

	t.Run("UL children inside OL", func(t *testing.T) {
		result := stripANSI(ty.NestOL(
			ItemWithChildren("parent",
				Item("bullet1"),
				Item("bullet2"),
			),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 3 {
			t.Fatalf("expected 3 lines, got %d: %q", len(lines), result)
		}
		// Parent should be numbered
		if !strings.Contains(lines[0], "1.") {
			t.Errorf("expected numbered parent, got %q", lines[0])
		}
		// Children should have bullet chars
		if !strings.Contains(lines[1], DefaultNestedBulletChars[1]) {
			t.Errorf("expected bullet child, got %q", lines[1])
		}
	})
}

func TestNestULCustomIndent(t *testing.T) {
	ty := New(WithListIndent(4))

	result := stripANSI(ty.NestUL(
		ItemWithChildren("parent",
			Item("child"),
		),
	))
	lines := strings.Split(result, "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	// Child should be indented by 4 spaces
	if !strings.HasPrefix(lines[1], "    ") {
		t.Errorf("expected 4-space indent, got %q", lines[1])
	}
}

func TestNestULCustomBullets(t *testing.T) {
	ty := New(WithNestedBulletChars([]string{"*", "o", "-", ">"}))

	result := stripANSI(ty.NestUL(
		ItemWithChildren("l0",
			ItemWithChildren("l1",
				Item("l2"),
			),
		),
	))
	lines := strings.Split(result, "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if !strings.Contains(lines[0], "*") {
		t.Errorf("depth 0: expected *, got %q", lines[0])
	}
	if !strings.Contains(lines[1], "o") {
		t.Errorf("depth 1: expected o, got %q", lines[1])
	}
	if !strings.Contains(lines[2], "-") {
		t.Errorf("depth 2: expected -, got %q", lines[2])
	}
}

func TestHierarchicalNumbers(t *testing.T) {
	ty := New(WithHierarchicalNumbers(true))

	t.Run("nested OL uses parent prefix", func(t *testing.T) {
		result := stripANSI(ty.NestOL(
			Item("Introduction"),
			ItemWithOLChildren("Main Topics",
				Item("Architecture"),
				Item("Design"),
			),
			Item("Conclusion"),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 5 {
			t.Fatalf("expected 5 lines, got %d: %q", len(lines), result)
		}
		// Parent items: 1., 2., 3.
		if !strings.Contains(lines[0], "1.") {
			t.Errorf("expected 1. in %q", lines[0])
		}
		if !strings.Contains(lines[1], "2.") {
			t.Errorf("expected 2. in %q", lines[1])
		}
		// Children: 2.1., 2.2.
		if !strings.Contains(lines[2], "2.1.") {
			t.Errorf("expected 2.1. in %q", lines[2])
		}
		if !strings.Contains(lines[3], "2.2.") {
			t.Errorf("expected 2.2. in %q", lines[3])
		}
		if !strings.Contains(lines[4], "3.") {
			t.Errorf("expected 3. in %q", lines[4])
		}
	})

	t.Run("three levels deep", func(t *testing.T) {
		result := stripANSI(ty.NestOL(
			ItemWithOLChildren("A",
				ItemWithOLChildren("B",
					Item("C"),
				),
			),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 3 {
			t.Fatalf("expected 3 lines, got %d: %q", len(lines), result)
		}
		if !strings.Contains(lines[0], "1.") {
			t.Errorf("expected 1. in %q", lines[0])
		}
		if !strings.Contains(lines[1], "1.1.") {
			t.Errorf("expected 1.1. in %q", lines[1])
		}
		if !strings.Contains(lines[2], "1.1.1.") {
			t.Errorf("expected 1.1.1. in %q", lines[2])
		}
	})

	t.Run("disabled by default", func(t *testing.T) {
		tyDefault := New()
		result := stripANSI(tyDefault.NestOL(
			ItemWithOLChildren("A",
				Item("B"),
			),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 2 {
			t.Fatalf("expected 2 lines, got %d", len(lines))
		}
		// Child should be "1." not "1.1."
		if strings.Contains(lines[1], "1.1.") {
			t.Errorf("hierarchical numbering should be off by default, got %q", lines[1])
		}
	})

	t.Run("mixed UL children skip prefix", func(t *testing.T) {
		result := stripANSI(ty.NestOL(
			ItemWithChildren("A",
				Item("bullet"),
			),
		))
		lines := strings.Split(result, "\n")
		if len(lines) != 2 {
			t.Fatalf("expected 2 lines, got %d", len(lines))
		}
		// UL child should use a bullet, not a number
		if strings.Contains(lines[1], "1.1.") {
			t.Errorf("UL children should not use hierarchical numbers, got %q", lines[1])
		}
	})
}
