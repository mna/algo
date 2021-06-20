package stacks

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mna/algo/sort"
)

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

	t.Run("ZeroValue", func(t *testing.T) {
		var s Stack
		if s.Len() != 0 {
			t.Fatalf("want len %d, got %d", 0, s.Len())
		}
		if v := s.Pop(); v != 0 {
			t.Fatalf("want empty pop %d, got %d", 0, v)
		}
		s.Push(1)
		if got := s.Pop(); got != 1 {
			t.Fatalf("want %d, got %d", 1, got)
		}
	})

	t.Run("PushPop", func(t *testing.T) {
		cases := [][]T{
			sortedSlice(1),
			sortedSlice(2),
			sortedSlice(3),
			sortedSlice(4),
			sortedSlice(5),
			sortedSlice(6),
			sortedSlice(7),
			sortedSlice(8),
			sortedSlice(9),
			sortedSlice(10),
			sortedSlice(100),
			sortedSlice(1000),
			sortedSlice(10000),
		}
		for _, c := range cases {
			t.Run(fmt.Sprintf("%d", len(c)), func(t *testing.T) {
				got := make([]T, 0, len(c))
				s := MakeFrom(c...)
				for s.Len() > 0 {
					got = append(got, s.Pop())
				}

				sort.Reverse(c)
				if !cmp.Equal(c, got) {
					t.Fatalf("want %v, got %v", c, got)
				}
			})
		}
	})

	t.Run("MixPushPop", func(t *testing.T) {
		cases := []struct {
			pushPop [][2]int // number of push & pop per iteration
			want    []T      // expected pop values in order
		}{
			{[][2]int{{1, 1}}, []T{1}},
			{[][2]int{{2, 1}}, []T{2}},
			{[][2]int{{3, 2}, {1, 2}}, []T{3, 2, 4, 1}},
			{[][2]int{{0, 1}, {1, 0}, {3, 1}, {0, 2}}, []T{0, 4, 3, 2}},
		}
		for _, c := range cases {
			t.Run(fmt.Sprintf("%v", c.pushPop), func(t *testing.T) {
				s := Make /*[T]*/ ()

				var v int
				var got []T
				for _, pp := range c.pushPop {
					for i := 0; i < pp[0]; i++ {
						v++
						s.Push(v)
					}
					for i := 0; i < pp[1]; i++ {
						got = append(got, s.Pop())
					}
				}

				if !cmp.Equal(c.want, got) {
					t.Fatalf("want %v, got %v", c.want, got)
				}
			})
		}
	})
}

// ns control the slice's creation:
// ns[0] = how many items, default 0
// ns[1] = multiplier, default 1
// ns[2] = value starts at (before multiplier is applied), default 1
func sortedSlice(ns ...int) []int {
	n, mul, start := 0, 1, 1
	if len(ns) > 0 {
		n = ns[0]
	}
	if len(ns) > 1 {
		mul = ns[1]
	}
	if len(ns) > 2 {
		start = ns[2]
	}

	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = start * mul
		start++
	}
	return vals
}
