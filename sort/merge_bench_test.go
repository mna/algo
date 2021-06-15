package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMerge(b *testing.B) {
	src := time.Now().UnixNano()
	r := rand.New(rand.NewSource(src))
	b.Logf("random seed: %d", src)

	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := shuffledSlice(r, sortedSlice(n, 1))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				Merge(vals)
			}
		})
	}
}
