package stacks

import "testing"

func TestStack(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		sts := []*Stack /*[T]*/ {
			Make(),
			MakeCap(10),
			MakeFrom(),
		}
		for _, s := range sts {
			if s.Len() != 0 {
				t.Fatalf("want len %d, got %d", 0, s.Len())
			}
			if v := s.Pop(); v != 0 {
				t.Fatalf("want empty pop %d, got %d", 0, v)
			}
		}
	})
}
