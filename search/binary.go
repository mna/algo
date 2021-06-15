package search

type T = int // TODO: generic type placeholder

// Binary performs a binary search on vals and returns the index at which v
// was found or -1 if it is not in vals. The vals slice must already be sorted
// in ascending order as defined by the standard <, == and > operators.
//
// It runs in O(log n) time complexity and O(1) space complexity. It does not
// allocate.
func Binary /*[T algo.OrderedComparable]*/ (vals []T, v T) int {
	start, end := 0, len(vals)
	for start < end {
		// uint conversion is to avoid overflow for very big slices - after the
		// division, it necessarily fits in an int. Note that the compiler
		// will automatically optimize the division by 2 to a right shift,
		// so no need to make the code more clever.
		half := int(uint(start+end) / 2)
		cur := vals[half]

		if cur == v {
			return half
		}
		if cur < v {
			start = half + 1
		} else {
			end = half
		}
	}
	return -1
}

// BinaryFunc performs a binary search on vals and returns the index at which
// v was found or -1 if it is not in vals. The vals slice must already be
// sorted using the same ordering as the one reported by the cmp function. It
// calls cmp to check ordering of pairs of values, and it should return -1 if
// the first value is smaller, 1 if it is larger, and 0 if they are equal.
//
// It runs in O(log n) time complexity and O(1) space complexity. It does not
// allocate.
func BinaryFunc /*[T algo.Any]*/ (vals []T, v T, cmp func(T, T) int) int {
	start, end := 0, len(vals)
	for start < end {
		// uint conversion is to avoid overflow for very big slices - after the
		// division, it necessarily fits in an int. Note that the compiler
		// will automatically optimize the division by 2 to a right shift,
		// so no need to make the code more clever.
		half := int(uint(start+end) / 2)
		cur := vals[half]

		order := cmp(cur, v)
		if order == 0 {
			return half
		}
		if order < 0 {
			start = half + 1
		} else {
			end = half
		}
	}
	return -1
}
