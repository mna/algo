package stacks

type T = int // NOTE: generic type placeholder

// Stack is a stack data structure, implementing a last-in-first-out (LIFO)
// insertion and retrieval mechanism. Its zero-value is ready to use.
type Stack /*[T algo.Any]*/ struct {
	items []T
}

// Make returns a stack of some element type.
func Make /*[T algo.Any]*/ () *Stack /*[T]*/ {
	return new(Stack)
}

// MakeCap returns a stack of some element type with an initial capacity.
func MakeCap /*[T algo.Any]*/ (capacity int) *Stack /*[T]*/ {
	return &Stack{
		items: make([]T, 0, capacity),
	}
}

// MakeFrom returns a stack of some element type initialized with the
// provided values so that the last value will be the first to be removed.
func MakeFrom /*[T algo.Any]*/ (vs ...T) *Stack /*[T]*/ {
	s := MakeCap /*[T]*/ (len(vs))
	s.Push(vs...)
	return s
}

// Len reports the number of elements in s.
func (s *Stack /*[T]*/) Len() int {
	return len(s.items)
}

// Push adds the provided values vs in order to the stack, so that the
// last value provided would be the first to be removed.
//
// It runs in O(1) (amortized) time complexity (O(n) with respect to the
// number of values to add).
func (s *Stack /*[T]*/) Push(vs ...T) {
	s.items = append(s.items, vs...)
}

// Pop removes the item at the top of the stack (the last one added). If
// the stack is empty, it returns the zero value of T. Len can be used
// to check if there are values to pop.
//
// It runs in O(1) time and space complexity. It does not allocate.
func (s *Stack /*[T]*/) Pop() T {
	var v T
	if n := len(s.items); n > 0 {
		v = s.items[n-1]
		s.items = s.items[:n-1]
	}
	return v
}
