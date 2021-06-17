package sets

import (
	"fmt"
	"testing"
)

func BenchmarkSet_Add(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		// with the right cap on initial creation, Add is O(1)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			s := MakeCap /*[int]*/ (n + b.N)
			s.Add(vals...)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				s.Add(n + 1 + i)
			}
		})

		// with no explicit cap provided, Add is amortized O(1)
		b.Run(fmt.Sprintf("amortized n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			s := Make /*[int]*/ ()
			s.Add(vals...)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				s.Add(n + 1 + i)
			}
		})
	}
}

func BenchmarkSet_Delete(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			s := MakeFrom(vals...)
			indices := indicesSlice(vals, b.N)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				s.Delete(vals[indices[i]])
			}
		})
	}
}

func BenchmarkSet_Contains(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vals := sortedSlice(n, 1)
			s := MakeFrom(vals...)
			indices := indicesSlice(vals, b.N)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				if !s.Contains(vals[indices[i]]) {
					b.Fatal("Contains returned false")
				}
			}
		})
	}
}

// for comparison, the built-in map returns times similar to the Set.Contains.
func BenchmarkBuiltinMap_Contains(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			m := make(map[int]bool, n)
			vals := sortedSlice(n, 1)
			for _, v := range vals {
				m[v] = true
			}
			indices := indicesSlice(vals, b.N)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, ok := m[vals[indices[i]]]
				if !ok {
					b.Fatal("Contains returned false")
				}
			}
		})
	}
}

// returns a slice of N valid indices into vals, to be used for benchmarks
// where N is b.N. WARNING: that may create huge slices.
func indicesSlice(vals []int, N int) []int {
	indices := make([]int, N)
	for i := 0; i < N; i++ {
		indices[i] = i % len(vals)
	}
	return indices
}
