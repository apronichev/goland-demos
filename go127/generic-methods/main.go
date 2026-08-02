package main

import (
	"fmt"
	"slices"
	"strings"
)

// Set is a generic set type.
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, it := range items {
		s[it] = struct{}{}
	}
	return s
}

// Go 1.27: a method may declare its own type parameters.
//
// Before 1.27 this had to be a package-level generic helper, because the
// receiver could not introduce new type parameters:
//
//	func Map[T comparable, U any](s Set[T], f func(T) U) []U { ... }
//	names := Map(users, func(u User) string { return u.Name })
//
// Now the operation lives in the type's own namespace and reads naturally.

func (s Set[T]) Map[U any](f func(T) U) []U {
	out := make([]U, 0, len(s))
	for v := range s {
		out = append(out, f(v))
	}
	return out
}

type User struct {
	ID   int
	Name string
}

func main() {
	users := NewSet(
		User{1, "Alice"},
		User{2, "Bob"},
		User{3, "Carol"},
	)

	// Call the generic method directly on the value, inferring U per call.
	names := users.Map(func(u User) string { return u.Name })
	ids := users.Map(func(u User) int { return u.ID })

	// Sort for stable output (map iteration order is randomized).
	slices.Sort(names)
	slices.Sort(ids)

	fmt.Println("names:", strings.Join(names, ", "))
	fmt.Println("ids:  ", ids)
}
