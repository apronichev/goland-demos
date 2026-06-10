// `stringscut`: `strings.Index` followed by manual slicing is the textbook
// case for `strings.Cut`, which returns the prefix, suffix, and a found bool
// in a single call.
package stringscut

import "strings"

func SplitKV(s string) (key, value string, ok bool) {
	i := strings.Index(s, "=")
	if i < 0 {
		return s, "", false
	}
	return s[:i], s[i+1:], true
}
