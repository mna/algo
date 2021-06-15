package algo

// Any allows any type.
type Any interface{} // TODO: predeclared `any` constraint in Go2

// Comparable allows any type that can be compared using == and !=.
type Comparable interface{} // TODO: predeclared `comparable` constraint in Go2

// Ordered allows any type that supports <, <=, >, >= operators.
type Ordered interface { // TODO: not clear if that will be a predeclared constraint in Go2, declare it if needed
	/*
		type int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr,
			float32, float64,
			string
	*/
}

// OrderedComparable is any type that supports all operators of Ordered
// and Comparable constraints.
type OrderedComparable interface {
	Ordered
	Comparable
}
