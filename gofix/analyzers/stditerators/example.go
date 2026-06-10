// `stditerators`: well-known stdlib types now expose iterator methods. An
// index loop like `for i := 0; i < x.Len(); i++ { use(x.At(i)) }` is rewritten
// to range over the matching iterator (e.g. `x.Variables()` on
// `go/types.Tuple`).
package stditerators

import (
	"go/types"
	"strings"
)

func TupleNames(t *types.Tuple) string {
	var names []string
	for i := 0; i < t.Len(); i++ {
		names = append(names, t.At(i).Name())
	}
	return strings.Join(names, ",")
}
