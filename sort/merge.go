package sort

type T = int // NOTE: generic type placeholder

// Merge performs a merge sort of vals. The returned slice is sorted in
// ascending order as defined by the standard <, <=, >, >= operators.  It is a
// stable sorting algorithm, meaning that equal values maintain their original
// order.
//
// It runs in O(n log n) time complexity and O(n) space complexity.
func Merge /*[T algo.Ordered]*/ (vals []T) []T {
	return splitSort(vals)
}

func splitSort /*[T algo.Ordered]*/ (vals []T) []T {
	n := len(vals)
	if n < 2 {
		// unsplittable, this is the recursion's stop condition
		return vals
	}

	// split vals in half recursively ("divide and conquer" strategy)
	half := n / 2
	return merge(splitSort(vals[:half]), splitSort(vals[half:]))
}

func merge /*[T algo.Ordered]*/ (v1, v2 []T) []T {
	n1, n2 := len(v1), len(v2)
	dst := make([]T, n1+n2)

	// merge back v1 and v2, selecting the smallest value from either slice:
	// [3, 6, 8] + [2, 7, 9] => [2, 3, 6, 7, 8, 9]
	// Equal values are taken from the v1 (left) slice first, to maintain
	// sort stability.
	for i, j1, j2 := 0, 0, 0; i < n1+n2; i++ {
		if j2 >= n2 || (j1 < n1 && v1[j1] <= v2[j2]) {
			dst[i] = v1[j1]
			j1++
		} else {
			dst[i] = v2[j2]
			j2++
		}
	}
	return dst
}

// MergeFunc performs a merge sort of vals. The returned slice is sorted in
// ascending order as defined by the cmp function. It calls cmp to check
// ordering of pairs of values, and it should return -1 if the first value is
// smaller, 1 if it is larger, and 0 if they are equal. It is a stable
// sorting algorithm, meaning that equal values maintain their original
// order.
//
// It runs in O(n log n) time complexity and O(n) space complexity.
func MergeFunc /*[T algo.Ordered]*/ (vals []T, cmp func(T, T) int) []T {
	return splitSortFunc(vals, cmp)
}

func splitSortFunc /*[T algo.Ordered]*/ (vals []T, cmp func(T, T) int) []T {
	n := len(vals)
	if n < 2 {
		// unsplittable, this is the recursion's stop condition
		return vals
	}

	// split vals in half recursively ("divide and conquer" strategy)
	half := n / 2
	return mergeFunc(splitSortFunc(vals[:half], cmp),
		splitSortFunc(vals[half:], cmp), cmp)
}

func mergeFunc /*[T algo.Ordered]*/ (v1, v2 []T, cmp func(T, T) int) []T {
	n1, n2 := len(v1), len(v2)
	dst := make([]T, n1+n2)

	// merge back v1 and v2, selecting the smallest value from either slice:
	// [3, 6, 8] + [2, 7, 9] => [2, 3, 6, 7, 8, 9]
	// Equal values are taken from the v1 (left) slice first, to maintain
	// sort stability.
	for i, j1, j2 := 0, 0, 0; i < n1+n2; i++ {
		if j2 >= n2 || (j1 < n1 && cmp(v1[j1], v2[j2]) <= 0) {
			dst[i] = v1[j1]
			j1++
		} else {
			dst[i] = v2[j2]
			j2++
		}
	}
	return dst
}
