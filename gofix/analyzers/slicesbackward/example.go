// `slicesbackward`: a manual reverse loop on a slice index becomes a range
// over `slices.Backward(s)` (Go 1.23+), which is easier to read and harder to
// off-by-one.
package slicesbackward

func ReverseInPlace(s []int) []int {
	out := make([]int, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return out
}
