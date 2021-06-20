package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.in), func(t *testing.T) {
			Reverse(c.in)
			if !cmp.Equal(c.in, c.out) {
				t.Fatalf("want %v, got %v", c.out, c.in)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	cases := [][]int{
		nil,
		{1},
		{1, 2},
		{3, 2, 1},
	}

	seed := time.Now().UnixNano()
	t.Logf("random seed: %d", seed)
	r := rand.New(rand.NewSource(seed))

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c), func(t *testing.T) {
			n := len(c)
			Shuffle(r, c)
			if len(c) != n {
				t.Fatalf("want %d elements, got %d", n, len(c))
			}
		})
	}
}

func TestReverseCmpFunc(t *testing.T) {
	cmpFn := func(v1, v2 int) int {
		switch {
		case v1 > v2:
			return 1
		case v1 < v2:
			return -1
		default:
			return 0
		}
	}
	cases := []struct {
		in  [2]int
		out int
	}{
		{[2]int{0, 0}, 0},
		{[2]int{1, 0}, -1},
		{[2]int{1, 2}, 1},
	}

	revFn := ReverseCmpFunc(cmpFn)
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.in), func(t *testing.T) {
			got := revFn(c.in[0], c.in[1])
			if got != c.out {
				t.Fatalf("want %v, got %v", c.out, got)
			}
		})
	}
}
