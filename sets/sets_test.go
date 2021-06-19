package sets

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type setOpCase struct {
	desc       string
	dstValues  []T   // if non-nil, use the Into variation with those initial values in dst
	setsValues [][]T // values for each sets to create and pass to the operation
	want       []T   // expected values in the resulting set
}

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
		cases := []setOpCase{
			{"no set", nil, nil, nil},
			{"empty set", nil, [][]T{{}}, nil},
			{"single set", nil, [][]T{sortedSlice(5, 1)}, []T{1, 2, 3, 4, 5}},
			{"two sets", nil, [][]T{sortedSlice(5, 1), sortedSlice(3, 10)}, []T{1, 2, 3, 4, 5, 10, 20, 30}},
			{"three sets", nil, [][]T{sortedSlice(5, 1), sortedSlice(3, 10), sortedSlice(5, 10)}, []T{1, 2, 3, 4, 5, 10, 20, 30, 40, 50}},
			{"empty dst no set", []T{}, nil, nil},
			{"non-empty dst no set", []T{55}, nil, []T{55}},
			{"non-empty dst, single set", []T{55}, [][]T{sortedSlice(3, 1)}, []T{1, 2, 3, 55}},
			{"non-empty dst, two sets", []T{55}, [][]T{sortedSlice(3, 1), sortedSlice(4, 1)}, []T{1, 2, 3, 4, 55}},
			{"non-empty dst, three sets", []T{1, 55}, [][]T{sortedSlice(3, 1), sortedSlice(4, 1), sortedSlice(1, 10)}, []T{1, 2, 3, 4, 10, 55}},
		}
		for _, c := range cases {
			t.Run(c.desc, func(t *testing.T) {
				sets := make([]Set /*[T]*/, len(c.setsValues))
				for i, vals := range c.setsValues {
					sets[i] = MakeFrom(vals...)
				}

				var got Set /*[T]*/
				if c.dstValues != nil {
					got = MakeFrom(c.dstValues...)
					UnionInto(got, sets...)
				} else {
					got = Union(sets...)
				}

				if vals := got.Values(); !cmp.Equal(vals, c.want, cmpopts.SortSlices(sortCmpSlice)) {
					t.Fatalf("want %v, got %v", c.want, vals)
				}
			})
		}
	})

	t.Run("Intersect", func(t *testing.T) {
		cases := []setOpCase{
			{"no set", nil, nil, nil},
			{"empty set", nil, [][]T{{}}, nil},
			{"single set", nil, [][]T{sortedSlice(3, 1)}, []T{1, 2, 3}},
			{"two sets", nil, [][]T{sortedSlice(3, 1), sortedSlice(5, 1)}, []T{1, 2, 3}},
			{"two sets no overlap", nil, [][]T{sortedSlice(3, 1), sortedSlice(2, 10)}, nil},
			{"three sets", nil, [][]T{sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{1}},
			{"four sets", nil, [][]T{sortedSlice(4, 1), sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{1}},
			{"four sets varied content", nil, [][]T{{1, 3, 4}, {3, 4}, {3, 4, 5}, {1, 2, 3, 5}}, []T{3}},
			{"empty dst no set", []T{}, nil, nil},
			{"empty dst single set", []T{}, [][]T{sortedSlice(2, 1)}, []T{1, 2}},
			{"empty dst two sets", []T{}, [][]T{sortedSlice(2, 1), sortedSlice(1, 1)}, []T{1}},
			{"empty dst three sets", []T{}, [][]T{sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{1}},
			{"empty dst four sets", []T{}, [][]T{sortedSlice(4, 1), sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{1}},
			{"non-empty dst no set", []T{55}, nil, []T{55}},
			{"non-empty dst single set", []T{55}, [][]T{sortedSlice(3, 1)}, []T{55, 1, 2, 3}},
			{"non-empty dst two sets", []T{55}, [][]T{sortedSlice(3, 1), sortedSlice(2, 1)}, []T{55, 1, 2}},
			{"non-empty dst three sets", []T{55}, [][]T{sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{1, 55}},
			{"non-empty dst four sets", []T{55}, [][]T{sortedSlice(4, 1), sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{1, 55}},
			{"non-empty dst overlaps four sets", []T{4, 55}, [][]T{sortedSlice(4, 1), sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{1, 4, 55}},
			{"non-empty dst four sets varied content", []T{55}, [][]T{{1, 3, 4}, {3, 4}, {3, 4, 5}, {1, 2, 3, 5}}, []T{3, 55}},
		}
		for _, c := range cases {
			t.Run(c.desc, func(t *testing.T) {
				sets := make([]Set /*[T]*/, len(c.setsValues))
				for i, vals := range c.setsValues {
					sets[i] = MakeFrom(vals...)
				}

				var got Set /*[T]*/
				if c.dstValues != nil {
					got = MakeFrom(c.dstValues...)
					IntersectInto(got, sets...)
				} else {
					got = Intersect(sets...)
				}

				if vals := got.Values(); !cmp.Equal(vals, c.want, cmpopts.SortSlices(sortCmpSlice)) {
					t.Fatalf("want %v, got %v", c.want, vals)
				}
			})
		}
	})

	t.Run("Diff", func(t *testing.T) {
		cases := []setOpCase{
			{"no set", nil, nil, nil},
			{"empty set", nil, [][]T{{}}, nil},
			{"single set", nil, [][]T{sortedSlice(3, 1)}, []T{1, 2, 3}},
			{"two sets none", nil, [][]T{sortedSlice(3, 1), sortedSlice(5, 1)}, nil},
			{"two sets some", nil, [][]T{sortedSlice(5, 1), sortedSlice(3, 1)}, []T{4, 5}},
			{"two sets no overlap", nil, [][]T{sortedSlice(3, 1), sortedSlice(2, 10)}, []T{1, 2, 3}},
			{"three sets", nil, [][]T{sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{3}},
			{"four sets", nil, [][]T{sortedSlice(4, 1), sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{4}},
			{"four sets varied content", nil, [][]T{{1, 3, 4, 6}, {3, 4}, {3, 4, 5}, {2, 3, 6}}, []T{1}},
			{"empty dst no set", []T{}, nil, nil},
			{"empty dst single set", []T{}, [][]T{sortedSlice(2, 1)}, []T{1, 2}},
			{"empty dst two sets", []T{}, [][]T{sortedSlice(2, 1), sortedSlice(1, 1)}, []T{2}},
			{"empty dst three sets", []T{}, [][]T{sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{3}},
			{"empty dst four sets", []T{}, [][]T{sortedSlice(4, 1), sortedSlice(1, 1), sortedSlice(2, 1), sortedSlice(3, 1)}, []T{4}},
			{"non-empty dst no set", []T{55}, nil, []T{55}},
			{"non-empty dst single set", []T{55}, [][]T{sortedSlice(3, 1)}, []T{55, 1, 2, 3}},
			{"non-empty dst two sets", []T{55}, [][]T{sortedSlice(3, 1), sortedSlice(2, 1)}, []T{55, 3}},
			{"non-empty dst three sets", []T{55}, [][]T{sortedSlice(3, 1), sortedSlice(1, 1), sortedSlice(2, 1)}, []T{3, 55}},
			{"non-empty dst four sets", []T{55}, [][]T{sortedSlice(4, 1), sortedSlice(2, 1), sortedSlice(1, 1), sortedSlice(3, 1)}, []T{4, 55}},
			{"non-empty dst overlaps four sets", []T{4, 55}, [][]T{sortedSlice(4, 1), sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{4, 55}},
			{"non-empty dst four sets varied content", []T{55}, [][]T{{1, 3, 4}, {3, 5}, {3, 4, 5}, {2, 3, 5}}, []T{1, 55}},
		}
		for _, c := range cases {
			t.Run(c.desc, func(t *testing.T) {
				sets := make([]Set /*[T]*/, len(c.setsValues))
				for i, vals := range c.setsValues {
					sets[i] = MakeFrom(vals...)
				}

				var got Set /*[T]*/
				if c.dstValues != nil {
					got = MakeFrom(c.dstValues...)
					DiffInto(got, sets...)
				} else {
					got = Diff(sets...)
				}

				if vals := got.Values(); !cmp.Equal(vals, c.want, cmpopts.SortSlices(sortCmpSlice)) {
					t.Fatalf("want %v, got %v", c.want, vals)
				}
			})
		}
	})

	t.Run("SymmetricDiff", func(t *testing.T) {
		cases := []setOpCase{
			{"no set", nil, nil, nil},
			{"empty set", nil, [][]T{{}}, nil},
			{"single set", nil, [][]T{sortedSlice(3, 1)}, []T{1, 2, 3}},
			{"two sets", nil, [][]T{sortedSlice(3, 1), sortedSlice(5, 1)}, []T{4, 5}},
			{"two sets some in both", nil, [][]T{sortedSlice(5, 1), sortedSlice(3, 2)}, []T{1, 3, 5, 6}},
			{"two sets no overlap", nil, [][]T{sortedSlice(3, 1), sortedSlice(2, 10)}, []T{1, 2, 3, 10, 20}},
			{"three sets", nil, [][]T{sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{3}},
			{"four sets", nil, [][]T{sortedSlice(4, 1), sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{4}},
			{"four sets varied content", nil, [][]T{{1, 3, 4, 6}, {3, 4}, {3, 4, 5}, {2, 3, 6}}, []T{1, 2, 5}},
			{"empty dst no set", []T{}, nil, nil},
			{"empty dst single set", []T{}, [][]T{sortedSlice(2, 1)}, []T{1, 2}},
			{"empty dst two sets", []T{}, [][]T{sortedSlice(2, 1), sortedSlice(1, 1)}, []T{2}},
			{"empty dst three sets", []T{}, [][]T{sortedSlice(3, 1), sortedSlice(2, 1), {4}}, []T{3, 4}},
			{"empty dst four sets", []T{}, [][]T{sortedSlice(4, 1), sortedSlice(1, 1), sortedSlice(2, 1), sortedSlice(3, 1)}, []T{4}},
			{"non-empty dst no set", []T{55}, nil, []T{55}},
			{"non-empty dst single set", []T{55}, [][]T{sortedSlice(3, 1)}, []T{55, 1, 2, 3}},
			{"non-empty dst two sets", []T{55}, [][]T{sortedSlice(3, 1), sortedSlice(2, 1)}, []T{55, 3}},
			{"non-empty dst three sets", []T{55}, [][]T{sortedSlice(2, 1), sortedSlice(1, 1), sortedSlice(3, 1)}, []T{3, 55}},
			{"non-empty dst four sets", []T{55}, [][]T{sortedSlice(4, 1), sortedSlice(2, 1), sortedSlice(1, 1), sortedSlice(3, 1)}, []T{4, 55}},
			{"non-empty dst overlaps four sets", []T{4, 55}, [][]T{sortedSlice(4, 1), sortedSlice(3, 1), sortedSlice(2, 1), sortedSlice(1, 1)}, []T{4, 55}},
			{"non-empty dst four sets varied content", []T{55}, [][]T{{1, 3, 4}, {3, 5}, {3, 4, 5}, {2, 3, 5}}, []T{1, 2, 55}},
		}
		for _, c := range cases {
			t.Run(c.desc, func(t *testing.T) {
				sets := make([]Set /*[T]*/, len(c.setsValues))
				for i, vals := range c.setsValues {
					sets[i] = MakeFrom(vals...)
				}

				var got Set /*[T]*/
				if c.dstValues != nil {
					got = MakeFrom(c.dstValues...)
					SymmetricDiffInto(got, sets...)
				} else {
					got = SymmetricDiff(sets...)
				}

				if vals := got.Values(); !cmp.Equal(vals, c.want, cmpopts.SortSlices(sortCmpSlice)) {
					t.Fatalf("want %v, got %v", c.want, vals)
				}
			})
		}
	})
}

func sortCmpSlice(i, j int) bool {
	return i < j
}

func TestSortedSlice(t *testing.T) {
	cases := []struct {
		ns   []int
		want []int
	}{
		{nil, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{1, 2}},
		{[]int{3}, []int{1, 2, 3}},
		{[]int{1, 10}, []int{10}},
		{[]int{2, 10}, []int{10, 20}},
		{[]int{3, 10}, []int{10, 20, 30}},
		{[]int{1, 10, 2}, []int{20}},
		{[]int{2, 10, 2}, []int{20, 30}},
		{[]int{3, 10, 2}, []int{20, 30, 40}},
		{[]int{1, 1, 4}, []int{4}},
		{[]int{2, 1, 4}, []int{4, 5}},
		{[]int{3, 1, 4}, []int{4, 5, 6}},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.ns), func(t *testing.T) {
			got := sortedSlice(c.ns...)
			if !cmp.Equal(got, c.want, cmpopts.SortSlices(sortCmpSlice)) {
				t.Fatalf("want %v, got %v", c.want, got)
			}
		})
	}
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
