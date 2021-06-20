package stacks

import (
	"fmt"
	"testing"
)

func BenchmarkStack_Push(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		// with the right cap on initial creation, Push is O(1)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			s := MakeCap /*[int]*/ (n + b.N)
			s.Push(vals...)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				s.Push(i)
			}
		})

		// with no explicit cap provided, Push is amortized O(1)
		b.Run(fmt.Sprintf("amortized n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			s := Make /*[int]*/ ()
			s.Push(vals...)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				s.Push(i)
			}
		})
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		// hard to do proper benchmarking here, as b.N will be larger than any n
		// so it pops from an empty stack. I could stop the timer after Pop and
		// always Push one more value, but I'm afraid that's not ideal to play
		// with the timer inside the hot loop either.
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			s := MakeFrom(vals...)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				s.Pop()
			}
		})
	}
}
