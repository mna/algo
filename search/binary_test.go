package search

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	cases := []struct {
		in     []int
		search int
		out    int
	}{
		{nil, 1, -1},
		{sortedSlice(1, 1), 1, 0},
		{sortedSlice(1, 1), -1, -1},
		{sortedSlice(1, 1), 2, -1},
		{sortedSlice(2, 1), 0, -1},
		{sortedSlice(2, 1), 1, 0},
		{sortedSlice(2, 1), 2, 1},
		{sortedSlice(2, 1), 3, -1},
		{sortedSlice(3, 1), 0, -1},
		{sortedSlice(3, 1), 1, 0},
		{sortedSlice(3, 1), 2, 1},
		{sortedSlice(3, 1), 3, 2},
		{sortedSlice(3, 1), 4, -1},
		{sortedSlice(4, 1), 0, -1},
		{sortedSlice(4, 1), 1, 0},
		{sortedSlice(4, 1), 2, 1},
		{sortedSlice(4, 1), 3, 2},
		{sortedSlice(4, 1), 4, 3},
		{sortedSlice(4, 1), 5, -1},
		{sortedSlice(5, 1), 0, -1},
		{sortedSlice(5, 1), 1, 0},
		{sortedSlice(5, 1), 2, 1},
		{sortedSlice(5, 1), 3, 2},
		{sortedSlice(5, 1), 4, 3},
		{sortedSlice(5, 1), 5, 4},
		{sortedSlice(5, 1), 6, -1},
		{sortedSlice(6, 1), 0, -1},
		{sortedSlice(6, 1), 1, 0},
		{sortedSlice(6, 1), 2, 1},
		{sortedSlice(6, 1), 3, 2},
		{sortedSlice(6, 1), 4, 3},
		{sortedSlice(6, 1), 5, 4},
		{sortedSlice(6, 1), 6, 5},
		{sortedSlice(6, 1), 7, -1},
		{sortedSlice(7, 1), 0, -1},
		{sortedSlice(7, 1), 1, 0},
		{sortedSlice(7, 1), 2, 1},
		{sortedSlice(7, 1), 3, 2},
		{sortedSlice(7, 1), 4, 3},
		{sortedSlice(7, 1), 5, 4},
		{sortedSlice(7, 1), 6, 5},
		{sortedSlice(7, 1), 7, 6},
		{sortedSlice(7, 1), 8, -1},
		{sortedSlice(8, 1), 0, -1},
		{sortedSlice(8, 1), 1, 0},
		{sortedSlice(8, 1), 2, 1},
		{sortedSlice(8, 1), 3, 2},
		{sortedSlice(8, 1), 4, 3},
		{sortedSlice(8, 1), 5, 4},
		{sortedSlice(8, 1), 6, 5},
		{sortedSlice(8, 1), 7, 6},
		{sortedSlice(8, 1), 8, 7},
		{sortedSlice(8, 1), 9, -1},
		{sortedSlice(9, 1), 0, -1},
		{sortedSlice(9, 1), 1, 0},
		{sortedSlice(9, 1), 2, 1},
		{sortedSlice(9, 1), 3, 2},
		{sortedSlice(9, 1), 4, 3},
		{sortedSlice(9, 1), 5, 4},
		{sortedSlice(9, 1), 6, 5},
		{sortedSlice(9, 1), 7, 6},
		{sortedSlice(9, 1), 8, 7},
		{sortedSlice(9, 1), 9, 8},
		{sortedSlice(9, 1), 10, -1},
		{sortedSlice(10, 1), 0, -1},
		{sortedSlice(10, 1), 1, 0},
		{sortedSlice(10, 1), 2, 1},
		{sortedSlice(10, 1), 3, 2},
		{sortedSlice(10, 1), 4, 3},
		{sortedSlice(10, 1), 5, 4},
		{sortedSlice(10, 1), 6, 5},
		{sortedSlice(10, 1), 7, 6},
		{sortedSlice(10, 1), 8, 7},
		{sortedSlice(10, 1), 9, 8},
		{sortedSlice(10, 1), 10, 9},
		{sortedSlice(10, 1), 11, -1},
		{sortedSlice(10, 10), 11, -1},
		{sortedSlice(10, 10), 82, -1},
		{sortedSlice(10, 10), 99, -1},
		{sortedSlice(10, 10), 100, 9},
		{sortedSlice(10, 10), 10, 0},
		{sortedSlice(10, 10), 1, -1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d in %v", c.search, c.in), func(t *testing.T) {
			got := Binary(c.in, c.search)
			if got != c.out {
				t.Fatalf("want %d, got %d", c.out, got)
			}
		})
	}
}

func TestBinaryFunc(t *testing.T) {
	cmp := func(v1, v2 int) int {
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
		in     []int
		search int
		out    int
	}{
		{nil, 1, -1},
		{sortedSlice(1, 1), 1, 0},
		{sortedSlice(1, 1), -1, -1},
		{sortedSlice(1, 1), 2, -1},
		{sortedSlice(2, 1), 0, -1},
		{sortedSlice(2, 1), 1, 0},
		{sortedSlice(2, 1), 2, 1},
		{sortedSlice(2, 1), 3, -1},
		{sortedSlice(3, 1), 0, -1},
		{sortedSlice(3, 1), 1, 0},
		{sortedSlice(3, 1), 2, 1},
		{sortedSlice(3, 1), 3, 2},
		{sortedSlice(3, 1), 4, -1},
		{sortedSlice(4, 1), 0, -1},
		{sortedSlice(4, 1), 1, 0},
		{sortedSlice(4, 1), 2, 1},
		{sortedSlice(4, 1), 3, 2},
		{sortedSlice(4, 1), 4, 3},
		{sortedSlice(4, 1), 5, -1},
		{sortedSlice(5, 1), 0, -1},
		{sortedSlice(5, 1), 1, 0},
		{sortedSlice(5, 1), 2, 1},
		{sortedSlice(5, 1), 3, 2},
		{sortedSlice(5, 1), 4, 3},
		{sortedSlice(5, 1), 5, 4},
		{sortedSlice(5, 1), 6, -1},
		{sortedSlice(6, 1), 0, -1},
		{sortedSlice(6, 1), 1, 0},
		{sortedSlice(6, 1), 2, 1},
		{sortedSlice(6, 1), 3, 2},
		{sortedSlice(6, 1), 4, 3},
		{sortedSlice(6, 1), 5, 4},
		{sortedSlice(6, 1), 6, 5},
		{sortedSlice(6, 1), 7, -1},
		{sortedSlice(7, 1), 0, -1},
		{sortedSlice(7, 1), 1, 0},
		{sortedSlice(7, 1), 2, 1},
		{sortedSlice(7, 1), 3, 2},
		{sortedSlice(7, 1), 4, 3},
		{sortedSlice(7, 1), 5, 4},
		{sortedSlice(7, 1), 6, 5},
		{sortedSlice(7, 1), 7, 6},
		{sortedSlice(7, 1), 8, -1},
		{sortedSlice(8, 1), 0, -1},
		{sortedSlice(8, 1), 1, 0},
		{sortedSlice(8, 1), 2, 1},
		{sortedSlice(8, 1), 3, 2},
		{sortedSlice(8, 1), 4, 3},
		{sortedSlice(8, 1), 5, 4},
		{sortedSlice(8, 1), 6, 5},
		{sortedSlice(8, 1), 7, 6},
		{sortedSlice(8, 1), 8, 7},
		{sortedSlice(8, 1), 9, -1},
		{sortedSlice(9, 1), 0, -1},
		{sortedSlice(9, 1), 1, 0},
		{sortedSlice(9, 1), 2, 1},
		{sortedSlice(9, 1), 3, 2},
		{sortedSlice(9, 1), 4, 3},
		{sortedSlice(9, 1), 5, 4},
		{sortedSlice(9, 1), 6, 5},
		{sortedSlice(9, 1), 7, 6},
		{sortedSlice(9, 1), 8, 7},
		{sortedSlice(9, 1), 9, 8},
		{sortedSlice(9, 1), 10, -1},
		{sortedSlice(10, 1), 0, -1},
		{sortedSlice(10, 1), 1, 0},
		{sortedSlice(10, 1), 2, 1},
		{sortedSlice(10, 1), 3, 2},
		{sortedSlice(10, 1), 4, 3},
		{sortedSlice(10, 1), 5, 4},
		{sortedSlice(10, 1), 6, 5},
		{sortedSlice(10, 1), 7, 6},
		{sortedSlice(10, 1), 8, 7},
		{sortedSlice(10, 1), 9, 8},
		{sortedSlice(10, 1), 10, 9},
		{sortedSlice(10, 1), 11, -1},
		{sortedSlice(10, 10), 11, -1},
		{sortedSlice(10, 10), 82, -1},
		{sortedSlice(10, 10), 99, -1},
		{sortedSlice(10, 10), 100, 9},
		{sortedSlice(10, 10), 10, 0},
		{sortedSlice(10, 10), 1, -1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d in %v", c.search, c.in), func(t *testing.T) {
			got := BinaryFunc(c.in, c.search, cmp)
			if got != c.out {
				t.Fatalf("want %d, got %d", c.out, got)
			}
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
