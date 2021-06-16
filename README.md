[![Build Status](https://github.com/mna/algo/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/mna/algo/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/mna/algo.svg)](https://pkg.go.dev/github.com/mna/algo)

# Algo: fundamental algorithms and data structures in Go

With the Go programming language soon gaining support for [type parameters
(_generics_)](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md),
it becomes interesting to implement some of the fundamental algorithms and data
structures not previously found in the standard library (pre Go1.18 at least)
in a generic way - where before those would've been implemented repeatedly for
different data types, or implemented with the empty interface (`interface{}`),
losing type safety.

Eventually, I assume many of those algorithms will end up in the standard
library, but in the meantime it is a good excuse to revisit them and get a feel
for the generics proposal. It can also serve as a reference for implementation
of those fundamental algorithms and data structures - I will take great care to
make the code clear and well documented (the goal is not to make it as fast as
possible at the expense of code clarity, e.g. reaching for the `unsafe` package
is likely out of scope for this package), and will mention the Big O
complexities, with benchmarks to back those up.

That being said, please do rise an issue or send a pull request if there is an
obvious optimization missing or if an implementation is not quite right.

Once generics land officially, I will update the code to build with the proper
constraints and type arguments. For now, it is a mix of inline comments and
type aliases to bring the syntax close to the generics one, and the generic
`T` is aliased to `int`, e.g.:

```go
// in package search, file binary.go

type T = int
func Binary /*[T algo.OrderedComparable]*/ (vals []T, v T) int {
  ...
}


// in package algo, file constraints.go

type Comparable interface{}
type Ordered interface {
	/*
		type int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr,
			float32, float64,
			string
	*/
}
type OrderedComparable interface {
	Ordered
	Comparable
}
```

## Installation

    $ go get [-u] [-t] github.com/mna/algo

## License

The [BSD 3-Clause license](http://opensource.org/licenses/BSD-3-Clause).
