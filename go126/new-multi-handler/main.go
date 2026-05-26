package main

import "fmt"

// -----------------------------
// Domain type
// -----------------------------

type Version int

func (v Version) Less(o Version) bool {
	return v < o
}

// -----------------------------
// BEFORE Go 1.26
// Comparator-based approach
// -----------------------------

func MinWithComparator[T any](a, b T, less func(T, T) bool) T {
	if less(a, b) {
		return a
	}
	return b
}

// -----------------------------
// Go 1.26
// Recursive constraint approach
// -----------------------------

type Ordered[T Ordered[T]] interface {
	Less(T) bool
}

func Min[T Ordered[T]](a, b T) T {
	if a.Less(b) {
		return a
	}
	return b
}

// -----------------------------
// Demo
// -----------------------------

func main() {
	v1 := Version(1)
	v2 := Version(2)

	// Old approach: must pass comparator every time
	min1 := MinWithComparator(v1, v2, func(a, b Version) bool {
		return a.Less(b)
	})

	fmt.Println("min (old approach):", min1)

	// New Go 1.26 approach: no comparator needed
	min2 := Min(v1, v2)
	fmt.Println("min (new approach):", min2)

}
