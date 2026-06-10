// `minmax`: `if a > b { x = a } else { x = b }` and its variants collapse to

// the builtin `max(a, b)` (Go 1.21+); the `<` shape collapses to `min`.

package minmax

func MaxInt(a, b int) int {
	var m int
	if a > b {
		m = a
	} else {
		m = b
	}
	return m
}

func ClampHigh(x, hi int) int {
	var out int
	if x > hi {
		out = hi
	} else {
		out = x
	}
	return out
}
