package sort

import (
	"fmt"
	"testing"
)

func BenchmarkReverse(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				Reverse(vals)
			}
		})
	}
}
