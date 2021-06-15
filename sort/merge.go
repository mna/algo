package sort

type T = int // NOTE: generic type placeholder

// Merge performs a merge sort of vals. The returned slice is sorted in
// ascending order as defined by the standard <, <=, >, >= operators.
//
// It runs in O(n log n) time complexity and O(n) space complexity.
func Merge /*[T algo.Ordered]*/ (vals []T) []T {
	return splitSort(vals)
}

func splitSort /*[T algo.Ordered]*/ (vals []T) []T {
	n := len(vals)
	if n < 2 {
		return vals
	}

	half := n / 2
	return merge(splitSort(vals[:half]), splitSort(vals[half:]))
}

func merge /*[T algo.Ordered]*/ (v1, v2 []T) []T {
	n1, n2 := len(v1), len(v2)
	dst := make([]T, n1+n2)

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
