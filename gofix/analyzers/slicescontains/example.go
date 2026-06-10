// `slicescontains`: a hand-rolled membership loop is `slices.Contains` (for
// `==`-comparable elements) or `slices.ContainsFunc` (for a predicate).
package slicescontains

import "strings"

func HasInt(xs []int, target int) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}

func HasPrefix(xs []string, prefix string) bool {
	for _, x := range xs {
		if strings.HasPrefix(x, prefix) {
			return true
		}
	}
	return false
}
