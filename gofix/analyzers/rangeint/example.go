// `rangeint`: a three-clause `for i := 0; i < n; i++` whose body doesn't
// otherwise need the upper bound is just `for i := range n` (Go 1.22+).
package rangeint

func Sum(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}

func Fill(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i * i
	}
	return out
}
