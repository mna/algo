package sort

import "math/rand"

// TODO: maybe move this into slices package?

// Reverse reverses the position of the elements of the slice in-place.
//
// It runs in O(n) time complexity and O(1) space complexity. It does not
// allocate.
func Reverse /*[T algo.Any]*/ (vals []T) {
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
}

// Shuffle shuffles the positions of the elements of the slice in-place
// using the provided *rand.Rand.
func Shuffle /*[T algo.Any]*/ (r *rand.Rand, vals []T) {
	r.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})
}

// ReverseCmpFunc takes an ordering comparison function cmp and returns a new
// ordering comparison function that generates the reverse order of cmp.  That
// is, it returns -1 where cmp returns 1 and 1 where it returns -1.
func ReverseCmpFunc /*[T algo.Any]*/ (cmp func(T, T) int) func(T, T) int {
	return func(v1, v2 T) int {
		v := cmp(v1, v2)
		return -v
	}
}
