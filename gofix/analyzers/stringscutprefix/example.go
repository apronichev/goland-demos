// `stringscutprefix`: `HasPrefix` followed by `TrimPrefix` walks the prefix
// twice. `strings.CutPrefix` does both in one pass and returns whether the
// prefix was present.
package stringscutprefix

import "strings"

func Strip(s string) (string, bool) {
	if strings.HasPrefix(s, "v") {
		return strings.TrimPrefix(s, "v"), true
	}
	return s, false
}
