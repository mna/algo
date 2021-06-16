package slices

// TODO: implement this? https://github.com/golang/go/issues/45955
// And anything missing from this? https://github.com/golang/go/wiki/SliceTricks

type T = int // NOTE: generic type placeholder

// Prepend inserts elements at the beginning of a slice. If it has sufficient
// capacity, the destination is resliced to accommodate the new elements,
// otherwise a new underlying array will be allocated. Prepend returns the
// updated slice, it is therefore necessary to store the result of Prepend.
//
// It is very similar to the append builtin, except it adds new elements to
// the front of the slice, so the time complexity is O(n).
func Prepend /*[T algo.Any]*/ (vals []T, v ...T) []T {
	newVals := growSlice(vals, len(vals)+len(v))
	copy(newVals[len(v):], vals)
	copy(newVals, v)
	return newVals
}

// returns a (possibly new) slice with at least minCap capacity and len ==
// minCap. No elements are copied or moved, it just adjusts capacity and len.
// This is similar logic used internally for the append builtin (see
// https://github.com/golang/go/blob/master/src/runtime/slice.go#L144).
func growSlice /* [T algo.Any]*/ (vals []T, minCap int) []T {
	oldCap := cap(vals)
	if oldCap >= minCap {
		// slice capacity is already big enough, just adjust len
		vals = vals[:minCap]
		return vals
	}

	newCap := oldCap
	doubleCap := newCap + newCap
	if minCap > doubleCap {
		newCap = minCap
	} else {
		if oldCap < 1024 {
			newCap = doubleCap
		} else {
			// Check 0 < newCap to detect overflow
			// and prevent an infinite loop.
			for 0 < newCap && newCap < minCap {
				newCap += newCap / 4
			}
			// Set newCap to the requested minCap when
			// the newCap calculation overflowed.
			if newCap <= 0 {
				newCap = minCap
			}
		}
	}
	return make([]T, minCap, newCap)
}
