// `mapsloop`: explicit loops that copy, collect keys, or collect values from a
// map are replaced by `maps.Copy`, `maps.Keys`, and `maps.Values`.
package mapsloop

func Copy(dst, src map[string]int) {
	for k, v := range src {
		dst[k] = v
	}
}

func Keys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values(m map[string]int) []int {
	vals := make([]int, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}
