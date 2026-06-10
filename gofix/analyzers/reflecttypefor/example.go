// `reflecttypefor`: `reflect.TypeOf((*T)(nil)).Elem()` and
// `reflect.TypeOf(T{})` both ask for a `reflect.Type` known at compile time.
// `reflect.TypeFor[T]()` (Go 1.22+) does the same without the runtime value.
package reflecttypefor

import "reflect"

type User struct {
	Name string
	Age  int
}

var userType = reflect.TypeOf(User{})

func StringType() reflect.Type {
	return reflect.TypeOf("")
}
