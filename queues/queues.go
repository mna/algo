package queues

type T = int // NOTE: generic type placeholder

// TODO: currently, it wastes the capacity of the slice by slicing forward,
// keeping a potentially large backing array alive even if just one value
// remains in the queue. Use a current index pointer instead of the slice's
// [0:len) range to improve memory usage.

// Queue is a queue data structure, implementing a first-in-first-out (FIFO)
// insertion and retrieval mechanism. Its zero-value is ready to use.
type Queue /*[T algo.Any]*/ struct {
	items []T
}

// Make returns a queue of some element type.
func Make /*[T algo.Any]*/ () *Queue /*[T]*/ {
	return new(Queue)
}

// MakeCap returns a queue of some element type with an initial capacity.
func MakeCap /*[T algo.Any]*/ (capacity int) *Queue /*[T]*/ {
	return &Queue{
		items: make([]T, 0, capacity),
	}
}

// MakeFrom returns a queue of some element type initialized with the provided
// values so that the first value will be the first to be removed.
func MakeFrom /*[T algo.Any]*/ (vs ...T) *Queue /*[T]*/ {
	q := MakeCap /*[T]*/ (len(vs))
	q.Enqueue(vs...)
	return q
}

// Len reports the number of elements in q.
func (q *Queue /*[T]*/) Len() int {
	return len(q.items)
}

// Enqueue adds the provided values vs in order to the queue, so that the
// first value provided would be the first to be removed.
//
// It runs in O(1) (amortized) time complexity (O(n) with respect to the
// number of values to add).
func (q *Queue /*[T]*/) Enqueue(vs ...T) {
	q.items = append(q.items, vs...)
}

// Dequeue removes the first item from the queue (the first one added). If the
// queue is empty, it returns the zero value of T. Len can be used to check if
// there are values to dequeue.
//
// It runs in O(1) time and space complexity. It does not allocate.
func (q *Queue /*[T]*/) Dequeue() T {
	var v T
	if len(q.items) > 0 {
		v = q.items[0]
		q.items = q.items[1:]
	}
	return v
}

// Peek returns the item at the front of the queue (the first one added)
// without removing it from the queue. If the queue is empty, it returns the
// zero value of T. Len can be used to check if there are values to peek.
//
// It runs in O(1) time and space complexity. It does not allocate.
func (q *Queue /*[T]*/) Peek() T {
	var v T
	if len(q.items) > 0 {
		v = q.items[0]
	}
	return v
}
