// `stringsbuilder`: repeated `s += …` inside a loop is quadratic. A
// `strings.Builder` (or `bytes.Buffer`) builds the result in linear time.
package stringsbuilder

func Join(parts []string, sep string) string {
	var out string
	for i, p := range parts {
		if i > 0 {
			out += sep
		}
		out += p
	}
	return out
}
