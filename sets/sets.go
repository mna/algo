package sets

import "sort"

// NOTE: Set type, Make function, Add, Delete, Contains and Len methods
// are adapted from the example sets package in the Type Parameters proposal
// See https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md.

type T = int // NOTE: generic type placeholder

// Set is a set of values.
type Set /*[T algo.Comparable]*/ map[T]struct{}

// Make returns a set of some element type.
func Make /*[T algo.Comparable]*/ () Set /*[T]*/ {
	return make(Set /*[T]*/)
}

// MakeCap returns a set of some element type with an initial capacity.
func MakeCap /*[T algo.Comparable]*/ (capacity int) Set /*[T]*/ {
	return make(Set /*[T]*/, capacity)
}

// MakeFrom returns a set of some element type initialized with the
// provided values.
func MakeFrom /*[T algo.Comparable]*/ (vs ...T) Set /*[T]*/ {
	s := make(Set /*[T]*/, len(vs))
	s.Add(vs...)
	return s
}

// Add adds value(s) to the set s. If v is already in s this has no effect.
//
// It runs in O(1) (amortized) time complexity. It is the "amortized"
// complexity because the Set is backed by a hash map, and an O(n) rehash
// operation may happen when the storage buckets need to grow. Of course
// it is O(n) with respect to the number of values to add.
func (s Set /*[T]*/) Add(vs ...T) {
	for _, v := range vs {
		s[v] = struct{}{}
	}
}

// Delete removes v from the set s. If v is not in s this has no effect.
//
// It runs in O(1) time complexity (O(n) with respect to the number of values
// to delete).
func (s Set /*[T]*/) Delete(vs ...T) {
	for _, v := range vs {
		delete(s, v)
	}
}

// Contains reports whether v is in s.
//
// It runs in O(1) time complexity.
func (s Set /*[T]*/) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Len reports the number of elements in s.
func (s Set /*[T]*/) Len() int {
	return len(s)
}

// Values returns a slice of all values present in the Set. The order is
// undefined.
//
// It runs in O(n) time complexity where n is the number of values in the set.
func (s Set /*[T]*/) Values() []T {
	vals := make([]T, 0, len(s))
	for k := range s {
		vals = append(vals, k)
	}
	return vals
}

// Intersect returns a new Set that contains the intersection of all sets.
// If no set is provided, it returns a nil Set. If a single Set is provided,
// it returns a copy of that Set (that is, it always creates a new Set if
// at least one set is provided).
func Intersect /*[T any.Comparable]*/ (sets ...Set /*[T]*/) Set /*[T]*/ {
	if len(sets) == 0 {
		return nil
	}
	s := Make /*[T]*/ ()
	IntersectInto(s, sets...)
	return s
}

// IntersectInto is like Intersect, but the intersection of the sets is stored
// in dst. The dst set's values are not used to find the intersection of values,
// only as destination storage.
//
// If no set is provided for the intersection, then dst is untouched. If a single
// set is provided, then all its values are added to dst.
func IntersectInto /*[T any.Comparable]*/ (dst Set /*[T]*/, sets ...Set /*[T]*/) {
	if len(sets) == 0 {
		return
	}

	// start with the set that has the fewest values, as this is the maximum
	// number of elements that the intersection may generate.
	sort.Slice(sets, func(i, j int) bool {
		return len(sets[i]) < len(sets[j])
	})

	// a temporary set is required as we look for values present in all sets.
	// Create one if dst is not empty.
	tmp := dst
	tmpIsDst := true
	if len(tmp) > 0 {
		tmpIsDst = false
		tmp = Make /*[T]*/ ()
	}

	for i, set := range sets {
		first, last := i == 0, i == len(sets)-1
		if first {
			// first set, add all its values to the tmp set
			for k := range set {
				// optimization - if first == last, add immediately to dst
				if last {
					dst.Add(k)
				} else {
					tmp.Add(k)
				}
			}
		} else {
			// subsequent sets, remove from tmp if it doesn't hold tmp's value
			for k := range tmp {
				if !set.Contains(k) {
					tmp.Delete(k)
				} else if last && !tmpIsDst {
					// if this is the last set and it contains k, then insert it into dst
					dst.Add(k)
				}
			}
		}
	}
}

// Union returns a new Set that is the union of all sets. If no set is
// provided, it returns a nil Set. If a single Set is provided, it returns a
// copy of that Set (that is, it always creates a new Set if at least one set
// is provided).
//
// It runs in O(n * m) time complexity where n is the number of values per set
// and m is the number of sets. A more useful way to think about it may be to
// say it runs in O(n) where n is the total number of values in all sets.
func Union /*[T any.Comparable]*/ (sets ...Set /*[T]*/) Set /*[T]*/ {
	if len(sets) == 0 {
		return nil
	}
	s := Make /*[T]*/ ()
	UnionInto(s, sets...)
	return s
}

// UnionInto is like Union, but the union of the sets is stored in dst.  If no
// set is provided for the union, then dst is untouched.
//
// It runs in O(n * m) time complexity where n is the number of values per set
// and m is the number of sets. A more useful way to think about it may be to
// say it runs in O(n) where n is the total number of values in all sets.
func UnionInto /*[T any.Comparable]*/ (dst Set /*[T]*/, sets ...Set /*[T]*/) {
	for _, set := range sets {
		for k := range set {
			dst.Add(k)
		}
	}
}
