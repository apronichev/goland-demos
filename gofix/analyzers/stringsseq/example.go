// `stringsseq`: ranging over `strings.Split` / `strings.Fields` allocates a
// slice that's immediately thrown away. The `SplitSeq` / `FieldsSeq` iterators
// (Go 1.24+) yield each piece without that intermediate slice.
package stringsseq

import "strings"

func CountCSV(s string) int {
	n := 0
	for _, p := range strings.Split(s, ",") {
		if p != "" {
			n++
		}
	}
	return n
}

func CountWords(s string) int {
	n := 0
	for _, w := range strings.Fields(s) {
		_ = w
		n++
	}
	return n
}
