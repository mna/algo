package rings

type T = int // NOTE: generic type placeholder

// WriteMode determines how the ring buffer behaves when a write is made and
// the buffer is full.
type WriteMode int

const (
	// AllowOverwrite means that a write when the buffer is full will overwrite
	// the oldest values so that all values to write are written.
	AllowOverwrite WriteMode = iota
	// PreventOverwrite means that a write when the buffer is full will only
	// be made if the buffer has sufficient space for all values to write.
	PreventOverwrite
)

// Buffer is a ring buffer. It can store up to a pre-determined capacity
// number of elements, and elements are removed on read and are read in
// a first-in-first-out (FIFO) order. The write behaviour when full can
// be specified on write, to either overwrite the oldest element or fail.
type Buffer /*[T algo.Any]*/ struct {
	items      []T
	start, end int
}

// MakeCap returns a buffer of some element type with the specified capacity.
func MakeCap /*[T algo.Any]*/ (capacity int) *Buffer /*[T]*/ {
	return &Buffer{
		items: make([]T, capacity),
	}
}

// MakeFrom returns a buffer of some element type initialized with the
// provided values so that the first value will be the first to be removed.
// The number of values provided determine the capacity of the buffer or, in
// other words, the returned buffer is full.
func MakeFrom /*[T algo.Any]*/ (vs ...T) *Buffer /*[T]*/ {
	b := MakeCap /*[T]*/ (len(vs))
	b.Write(AllowOverwrite, vs...)
	return b
}

// Len reports the number of elements in b.
func (b *Buffer /*[T]*/) Len() int {
	panic("unimplemented")
}

// Read returns the oldest value in the ring buffer, freeing its space for a
// new value. If the buffer is empty, it returns the zero value of T. Len can
// be used to check if there are values to read.
//
// It runs in O(1) time and space complexity. It does not allocate.
func (b *Buffer /*[T]*/) Read() T {
	panic("unimplemented")
}

// Write stores the values vs in the buffer according the the specified write
// mode. It returns the number of older values that were overwritten by the
// call, or -1 if the mode is PreventOverwrite and the buffer does not have
// sufficient free space for all values.
//
// It runs in O(1) time and space complexity (O(n) with respect to the number
// of values to write). It does not allocate.
func (b *Buffer /*[T]*/) Write(mode WriteMode, vs ...T) int {
	_, _ = b.start, b.end
	_, _ = mode, vs
	panic("unimplemented")
}

// Peek returns the oldest value in the ring buffer without removing it from
// the buffer. If the buffer is empty, it returns the zero value of T. Len can
// be used to check if there are values to peek.
//
// It runs in O(1) time and space complexity. It does not allocate.
func (b *Buffer /*[T]*/) Peek() T {
	panic("unimplemented")
}
