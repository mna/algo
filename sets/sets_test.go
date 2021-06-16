package sets

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestSet(t *testing.T) {
	t.Run("NilEmpty", func(t *testing.T) {
		var s Set /*[int]*/
		if s.Len() != 0 {
			t.Fatalf("want %d, got %d", 0, s.Len())
		}
		vals := s.Values()
		if len(vals) != 0 {
			t.Fatalf("want %d, got %d", 0, len(vals))
		}

		s = Make()
		if s.Len() != 0 {
			t.Fatalf("want %d, got %d", 0, s.Len())
		}
		vals = s.Values()
		if len(vals) != 0 {
			t.Fatalf("want %d, got %d", 0, len(vals))
		}
	})

	t.Run("MakeAddLen", func(t *testing.T) {
		s := Make()
		s.Add(1)
		if s.Len() != 1 {
			t.Fatalf("want %d, got %d", 1, s.Len())
		}
		s.Add(1, 2, 3)
		if s.Len() != 3 {
			t.Fatalf("want %d, got %d", 3, s.Len())
		}
	})

	t.Run("MakeCapAddValues", func(t *testing.T) {
		s := MakeCap(5)
		s.Add(1, 2, 3, 2, 3)
		vals := s.Values()
		if len(vals) != 3 {
			t.Fatalf("want %d, got %d", 3, len(vals))
		}
		if !cmp.Equal(vals, []int{3, 2, 1}, cmpopts.SortSlices(sortCmpSlice)) {
			t.Fatalf("want %v in any order, got %v", []int{1, 2, 3}, vals)
		}
	})

	t.Run("MakeFromDeleteContains", func(t *testing.T) {
		vals := []int{1, 2, 3}
		s := MakeFrom(vals...)
		if !s.Contains(1) {
			t.Fatalf("Contains(1): want %t, got %t", true, s.Contains(1))
		}
		s.Delete(1, 4)
		if s.Contains(1) {
			t.Fatalf("Contains(1): want %t, got %t", false, s.Contains(1))
		}
		if !s.Contains(2) {
			t.Fatalf("Contains(2): want %t, got %t", true, s.Contains(1))
		}
	})

	t.Run("Union", func(t *testing.T) {
		s0 := Union()
		if s0 != nil {
			t.Fatalf("want nil, got %v", s0)
		}

		vals1 := sortedSlice(5, 1)
		s1 := Union(MakeFrom(vals1...))
		if s1vals, want := s1.Values(), sortedSlice(5, 1); !cmp.Equal(s1vals, want, cmpopts.SortSlices(sortCmpSlice)) {
			t.Fatalf("want %v, got %v", want, s1vals)
		}

		vals2 := sortedSlice(3, 10)
		s2 := Union(MakeFrom(vals1...), MakeFrom(vals2...))
		if s2vals, want := s2.Values(), append(vals1, vals2...); !cmp.Equal(s2vals, want, cmpopts.SortSlices(sortCmpSlice)) {
			t.Fatalf("want %v, got %v", want, s2vals)
		}

		vals3 := sortedSlice(5, 10)
		s3 := Union(MakeFrom(vals1...), MakeFrom(vals2...), MakeFrom(vals3...))
		if s3vals, want := s3.Values(), append(vals1, vals3...); !cmp.Equal(s3vals, want, cmpopts.SortSlices(sortCmpSlice)) {
			t.Fatalf("want %v, got %v", want, s3vals)
		}
	})

	t.Run("UnionInto", func(t *testing.T) {
		dst := Make()
		UnionInto(dst)
		if dst.Len() != 0 {
			t.Fatalf("want %d, got %d", 0, dst.Len())
		}

		dst.Add(55)
		UnionInto(dst, MakeFrom(sortedSlice(2, 1)...))
		if vals, want := dst.Values(), append([]int{55}, sortedSlice(2, 1)...); !cmp.Equal(vals, want, cmpopts.SortSlices(sortCmpSlice)) {
			t.Fatalf("want %v, got %v", want, vals)
		}
	})

	t.Run("Intersect", func(t *testing.T) {
		s0 := Intersect()
		if s0 != nil {
			t.Fatalf("want nil, got %v", s0)
		}

		s1 := Intersect(Make())
		if s1.Len() != 0 {
			t.Fatalf("want %d, got %d", 0, s1.Len())
		}

		s2 := Intersect(MakeFrom(sortedSlice(3, 1)...))
		if vals, want := s2.Values(), sortedSlice(3, 1); !cmp.Equal(vals, want, cmpopts.SortSlices(sortCmpSlice)) {
			t.Fatalf("want %v, got %v", want, vals)
		}

		s3 := Intersect(MakeFrom(sortedSlice(3, 1)...), MakeFrom(sortedSlice(5, 1)...))
		if vals, want := s3.Values(), sortedSlice(3, 1); !cmp.Equal(vals, want, cmpopts.SortSlices(sortCmpSlice)) {
			t.Fatalf("want %v, got %v", want, vals)
		}

		s4 := Intersect(MakeFrom(sortedSlice(3, 1)...), MakeFrom(sortedSlice(2, 1)...), MakeFrom(sortedSlice(1, 1)...))
		if vals, want := s4.Values(), []int{1}; !cmp.Equal(vals, want, cmpopts.SortSlices(sortCmpSlice)) {
			t.Fatalf("want %v, got %v", want, vals)
		}

		s5 := Intersect(MakeFrom(sortedSlice(3, 1)...), MakeFrom(sortedSlice(2, 10)...))
		if s5.Len() != 0 {
			t.Fatalf("want %d, got %d", 0, s5.Len())
		}
	})
}

func sortCmpSlice(i, j int) bool {
	return i < j
}

func sortedSlice(n, mul int) []int {
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = (i + 1) * mul
	}
	return vals
}
