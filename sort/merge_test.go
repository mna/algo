package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestMerge(t *testing.T) {
	src := time.Now().UnixNano()
	r := rand.New(rand.NewSource(src))
	t.Logf("random seed: %d", src)

	cases := []struct {
		in  []int
		out []int
	}{
		{nil, nil},
		{sortedSlice(1, 1), sortedSlice(1, 1)},
		{shuffledSlice(r, sortedSlice(2, 1)), sortedSlice(2, 1)},
		{shuffledSlice(r, sortedSlice(3, 1)), sortedSlice(3, 1)},
		{shuffledSlice(r, sortedSlice(4, 1)), sortedSlice(4, 1)},
		{shuffledSlice(r, sortedSlice(5, 1)), sortedSlice(5, 1)},
		{shuffledSlice(r, sortedSlice(6, 1)), sortedSlice(6, 1)},
		{shuffledSlice(r, sortedSlice(7, 1)), sortedSlice(7, 1)},
		{shuffledSlice(r, sortedSlice(8, 1)), sortedSlice(8, 1)},
		{shuffledSlice(r, sortedSlice(9, 1)), sortedSlice(9, 1)},
		{shuffledSlice(r, sortedSlice(10, 1)), sortedSlice(10, 1)},
		{shuffledSlice(r, sortedSlice(10, 10)), sortedSlice(10, 10)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.in), func(t *testing.T) {
			got := Merge(c.in)
			if !cmp.Equal(c.out, got) {
				t.Fatalf("want %v, got %v", c.out, got)
			}
		})
	}
}

func shuffledSlice(r *rand.Rand, vals []int) []int {
	r.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})
	return vals
}

func sortedSlice(n, mul int) []int {
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = (i + 1) * mul
	}
	return vals
}
