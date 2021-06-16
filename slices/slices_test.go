package slices

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrepend(t *testing.T) {
	cases := []struct {
		in  []int
		add []int
		out []int
	}{
		{nil, nil, nil},
		{[]int{1}, nil, []int{1}},
		{[]int{1}, []int{2}, []int{2, 1}},
		{[]int{1, 2}, []int{3}, []int{3, 1, 2}},
		{[]int{1, 2, 3}, []int{4}, []int{4, 1, 2, 3}},
		{[]int{1}, []int{2, 3, 4}, []int{2, 3, 4, 1}},
		{sortedSlice(5, 1), sortedSlice(3, 10), []int{10, 20, 30, 1, 2, 3, 4, 5}},
		{sortedSlice(3, 1), sortedSlice(6, 10), []int{10, 20, 30, 40, 50, 60, 1, 2, 3}},
		{sortedSlice(1023, 1), []int{1}, append([]int{1}, sortedSlice(1023, 1)...)},
		{sortedSlice(1024, 1), []int{1}, append([]int{1}, sortedSlice(1024, 1)...)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v <= %v", c.in, c.add), func(t *testing.T) {
			got := Prepend(c.in, c.add...)
			if !cmp.Equal(c.out, got) {
				t.Fatalf("want %d, got %d", c.out, got)
			}
			t.Logf("after call: len=%d, cap=%d", len(got), cap(got))
		})
	}
}

func sortedSlice(n, mul int) []int {
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = (i + 1) * mul
	}
	return vals
}
