// `errorsastype`: `errors.As` requires a separately declared sentinel variable
// and a `&` to bind into it. `errors.AsType[T]` (Go 1.26+) is generic, returns
// the typed error directly, and is harder to misuse.
package errorsastype

import (
	"errors"
	"fmt"
)

type PathError struct {
	Op  string
	Err error
}

func (e *PathError) Error() string { return e.Op + ": " + e.Err.Error() }

func Describe(err error) string {
	var pe *PathError
	if errors.As(err, &pe) {
		return fmt.Sprintf("path error during %s", pe.Op)
	}
	return "other"
}
