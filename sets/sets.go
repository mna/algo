package sets

import (
	"sort"
)

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
	var vals []T

	if len(s) > 0 {
		vals = make([]T, 0, len(s))
		for k := range s {
			vals = append(vals, k)
		}
	}
	return vals
}

// IsDisjoint returns true if the intersection of s and other is an empty
// set, false otherwise.
//
// It runs in O(n) time complexity where n is the number of values in the
// smallest set between s and other.
func (s Set /*[T]*/) IsDisjoint(other Set /*[T]*/) bool {
	iter, cmp := s, other
	if other.Len() < s.Len() {
		iter, cmp = other, s
	}
	for k := range iter {
		if cmp.Contains(k) {
			return false
		}
	}
	return true
}

// IsSubset returns true if every value of s is in other, false otherwise.  If
// strict is true and s contains the same values as other, then it returns
// false.
//
// It runs in O(n) time complexity where n is the number of values in s.
func (s Set /*[T]*/) IsSubset(other Set /*[T]*/, strict bool) bool {
	if (!strict && s.Len() <= other.Len()) || (strict && s.Len() < other.Len()) {
		for k := range s {
			if !other.Contains(k) {
				return false
			}
		}
		return true
	}
	return false
}

// IsSuperset returns true if every value of other is in s, false otherwise.
// If strict is true and other contains the same values as s, then it returns
// false.
//
// It runs in O(n) time complexity where n is the number of values in other.
func (s Set /*[T]*/) IsSuperset(other Set /*[T]*/, strict bool) bool {
	return other.IsSubset(s, strict)
}

// IsEqual returns true if s contains the same values as other, false
// otherwise.
//
// It runs in O(n) time complexity where n is the number of values in s.
func (s Set /*[T]*/) IsEqual(other Set) bool {
	if s.Len() != other.Len() {
		return false
	}
	return s.IsSubset(other, false)
}

// Intersect returns a new Set that contains the intersection of all sets.
// If no set is provided, it returns a nil Set. If a single Set is provided,
// it returns a copy of that Set (that is, it always creates a new Set if
// at least one set is provided).
//
// It runs in O(n*m) time complexity where n is the smallest number of values
// in any set and m is the number of sets to intersect minus one (i.e. for
// all practical purposes where a handful of sets are provided, it runs in O(n)).
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
//
// Its time complexity is the same as for Intersect.
func IntersectInto /*[T any.Comparable]*/ (dst Set /*[T]*/, sets ...Set /*[T]*/) {
	if len(sets) == 0 {
		return
	}
	if len(sets) == 1 {
		for k := range sets[0] {
			dst.Add(k)
		}
		return
	}

	// start with the set that has the fewest values, as this is the maximum
	// number of elements that the intersection may generate.
	sort.Slice(sets, func(i, j int) bool {
		return len(sets[i]) < len(sets[j])
	})

	// a temporary set is required as we look for values present in all sets.
	// Create one if dst is not empty, otherwise dst can be safely used as tmp.
	// As a special-case, if there are only two sets to intersect, no need for a
	// temporary set as we can add the intersection values directly in dst even
	// if it is not empty - we do not need temporary storage.
	tmp := dst
	tmpIsDst := true
	if len(dst) > 0 && len(sets) > 2 {
		tmpIsDst = false
		tmp = Make /*[T]*/ ()
	}

	// first iteration is over the values of the (smallest) first set, and if the
	// second set contains the value, add it to tmp.
	first, second := sets[0], sets[1]
	for k := range first {
		if second.Contains(k) {
			tmp.Add(k)
		}
	}

	// next, for all subsequent sets, loop over the values of tmp and if the set
	// does *not* contain the value, remove it from tmp.
	for i, set := range sets[2:] {
		last := i+2 == len(sets)-1
		for k := range tmp {
			if !set.Contains(k) {
				tmp.Delete(k)
			} else if last && !tmpIsDst {
				// this condition avoids a final loop over the tmp set's values: if
				// this is the last set to process and it contains the value, then
				// it will need to be added to dst, so do it right away.
				dst.Add(k)
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
	var s Set /*[T]*/
	if len(sets) > 0 {
		s = MakeCap(sets[0].Len())
	}
	UnionInto(s, sets...)
	return s
}

// UnionInto is like Union, but the union of the sets is stored in dst.  If no
// set is provided for the union, then dst is untouched.
//
// Its time complexity is the same as Union.
func UnionInto /*[T any.Comparable]*/ (dst Set /*[T]*/, sets ...Set /*[T]*/) {
	for _, set := range sets {
		for k := range set {
			dst.Add(k)
		}
	}
}

// Diff returns a new Set that is the difference of all sets, that is, the
// values in the first set that are not in any of the other sets. If no set is
// provided, it returns a nil Set. If a single Set is provided, it returns a
// copy of that Set (it always creates a new Set if at least one set is
// provided).
//
// It runs in O(n*m) time complexity where n is the number of values of the
// first set and m is the number of sets to diff minus one (i.e. for all
// practical purposes where a handful of sets are provided, it runs in O(n)).
func Diff /*[T any.Comparable]*/ (sets ...Set /*[T]*/) Set /*[T]*/ {
	if len(sets) == 0 {
		return nil
	}
	s := Make /*[T]*/ ()
	DiffInto(s, sets...)
	return s
}

// DiffInto is like Diff, but the difference of the sets is stored in dst. The
// dst set's values are not used to find the difference of values, only as
// destination storage. If no set is provided for the difference, then dst is
// untouched. If a single set is provided, then all its values are added to
// dst.
//
// Its time complexity is the same as Diff.
func DiffInto /*[T any.Comparable]*/ (dst Set /*[T]*/, sets ...Set /*[T]*/) {
	if len(sets) == 0 {
		return
	}
	if len(sets) == 1 {
		for k := range sets[0] {
			dst.Add(k)
		}
		return
	}

	// a temporary set is required as we look for the difference.  Create one if
	// dst is not empty, otherwise dst can be safely used as tmp.  As a
	// special-case, if there are only two sets to diff, no need for a temporary
	// set as we can add the diff values directly in dst even if it is not empty
	// - we do not need temporary storage.
	tmp := dst
	tmpIsDst := true
	if len(dst) > 0 && len(sets) > 2 {
		tmpIsDst = false
		tmp = Make /*[T]*/ ()
	}

	// first iteration is over the values of the first set, and if the second set
	// does not contain the value, add it to tmp.
	first, second := sets[0], sets[1]
	for k := range first {
		if !second.Contains(k) {
			tmp.Add(k)
		}
	}

	// next, for all subsequent sets, loop over the values of tmp and if the set
	// *does* contain the value, remove it from tmp.
	for i, set := range sets[2:] {
		last := i+2 == len(sets)-1
		for k := range tmp {
			if set.Contains(k) {
				tmp.Delete(k)
			} else if last && !tmpIsDst {
				// this condition avoids a final loop over the tmp set's values: if
				// this is the last set to process and it does not contain the value,
				// then it will need to be added to dst, so do it right away.
				dst.Add(k)
			}
		}
	}
}

// SymmetricDiff returns a new Set that contains values that are in either of
// the sets but not in any other. If no set is provided, it returns a nil Set.
// If a single Set is provided, it returns a copy of that Set (it always
// creates a new Set if at least one set is provided).
//
// It runs in O(n*m) time complexity where n is the number of values in each
// set and m is the number of sets (i.e. for all practical purposes where a
// handful of sets are provided, it runs in O(n)).
func SymmetricDiff /*[T any.Comparable]*/ (sets ...Set /*[T]*/) Set /*[T]*/ {
	if len(sets) == 0 {
		return nil
	}
	s := Make /*[T]*/ ()
	SymmetricDiffInto(s, sets...)
	return s
}

// SymmetricDiffInto is like SymmetricDiff, but the resulting values are
// stored in dst. The dst set's values are not used to find the symmetric
// difference, only as destination storage. If no set is provided for the
// symmetric difference, then dst is untouched. If a single set is provided,
// then all its values are added to dst.
//
// Its time complexity is the same as SymmetricDiff.
func SymmetricDiffInto /*[T any.Comparable]*/ (dst Set /*[T]*/, sets ...Set /*[T]*/) {
	if len(sets) == 0 {
		return
	}
	if len(sets) == 1 {
		for k := range sets[0] {
			dst.Add(k)
		}
		return
	}

	// a temporary map is required as we look for the symmetric difference.  It
	// will keep track of which values are still in the final set (value=true)
	// and which values have been seen in more than one set, so should not be in
	// the final set (value=false).
	tmp := make(map[T]bool, len(sets[0]))

	// loop over the values of each set in sequence and if the value is not in
	// tmp, add it, and if it is in tmp, remove it.
	for _, set := range sets {
		for k := range set {
			contains, exists := tmp[k]
			if contains {
				tmp[k] = false
			} else if !exists {
				tmp[k] = true
			}
		}
	}

	// copy all values from tmp to dst
	for k, ok := range tmp {
		if ok {
			dst.Add(k)
		}
	}
}
