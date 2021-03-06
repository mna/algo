package search

import (
	"fmt"
	"testing"
)

func BenchmarkBinary(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				got := Binary(vals, n+1)
				if got != -1 {
					b.Fatalf("want -1, got %d", got)
				}
			}
		})
	}
}

func BenchmarkBinaryFunc(b *testing.B) {
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

	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				got := BinaryFunc(vals, n+1, cmp)
				if got != -1 {
					b.Fatalf("want -1, got %d", got)
				}
			}
		})
	}
}
