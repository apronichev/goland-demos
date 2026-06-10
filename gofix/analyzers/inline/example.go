// `inline`: a `//go:fix inline` directive on a function or constant tells the
// analyzer that callers/references should be rewritten to use the body. Use
// this to migrate a deprecated wrapper to its replacement, or to bump callers
// to a new package version, without a manual sweep.
//
// Run `go fix -inline ./analyzers/inline/...` to rewrite uses of `Square`
// (below) into `Pow(_, 2)`, and uses of `Ptr` into `Pointer`.
package inline

// Pow is the "new" API.
func Pow(x, y int) int {
	r := 1
	for range y {
		r *= x
	}
	return r
}

// Deprecated: prefer Pow(x, 2).
//
//go:fix inline
func Square(x int) int { return Pow(x, 2) }

func Areas(sides []int) []int {
	out := make([]int, len(sides))
	for i, s := range sides {
		out[i] = Pow(s, 2)
	}
	return out
}

const Pointer = 1

//go:fix inline
const Ptr = Pointer

var _ = Ptr
