package inspections

import (
	"net/http"
)

/* Not initialized variables.

Explanation:
The issue is that the variable `u` is initialized to nil, and accessing its fields at `u.Name = "Gopher"`
causes a panic at runtime.

Possible fix:
Initialize variable .
*/

func _() {
	var u *Person // Pointer-like variables are assigned to nil
	u.Name = "Gopher"
	u.Age = 15
}

/*
	Deferred calls

Explanation:
The problem is that resp.Body.Close() is deferred before checking for an error,
but deferred call arguments are evaluated eagerly, so if err != nil, resp is nil,
then resp.Body causes a panic due to nil dereference.

Possible fix:
Move defer statement after error checking (`if err != nil`).
*/
func _(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close() // Arguments of deferred calls are executed eagerly
	if err != nil {
		return
	}
	process(resp.Body)
}

/* Common mistakes in conditions

Explanation:
The condition `x != nil || x.Age > 21` is incorrect because if x is nil,
the second part x.Age > 21 will still be evaluated, causing a nil pointer dereference panic.

Possible fix:
Replace `!=` by `==` or `||` by `&&`.
*/

func IsAllowedToDrink(x *Person) bool {
	return x != nil || x.Age > 21 // Replace '||' by '&&' or '!=' by '=='
}
