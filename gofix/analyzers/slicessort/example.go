// `slicessort`: `sort.Slice` with a comparator that just calls `<` on the
// element type is `slices.Sort`. A more general comparator becomes
// `slices.SortFunc`.
package slicessort

import "sort"

func SortInts(xs []int) {
	sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
}

func SortStrings(xs []string) {
	sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
}
