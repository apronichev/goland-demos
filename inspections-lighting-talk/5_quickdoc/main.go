package main

import "fmt"

/*
This demo show inference of nilability in quick documentation.

This feature displays information about the nilability of function/method parameters and return values
in the quick documentation. Nilability is a property that determines whether a parameter or return value
can be nil or not:

* Return is Nilable (yellow): The function can return nil.
* Return is NotNil (green): The function cannot return nil.
* Parameter is NilSafe (green): It is safe to pass nil as the parameter value.
* Parameter is NilUnsafe (yellow): Passing nil as the parameter may lead to a panic.

The nil safe annotations are highlighted in green, and nil unsafe annotations are highlighted in yellow.
Programmers should pay attention to yellow annotation.

This feature helps developers quickly understand how nil behaves in the context of function calls,
improving code safety and reducing runtime errors.
*/

type User struct {
	Id   int
	Name string
}

/*
This function is nil unsafe in terms of its parameters:

* ctx has an explicit dereference.
* With user, it's more interesting — a dereference may occur in the Validate function.

Call to action: Try to make the code safe by using comments.
*/

func CopyUser(user *User, ctx *Ctx) *User {
	if ctx.DebugEnabled {
		println("copy user...")
	}
	if !user.Validate() || user == nil {
		return nil
	}
	return &User{
		Id:   user.Id,
		Name: user.Name,
	}
}

func main() {
	u1 := &User{Name: "Gopher", Id: 42}
	u2 := CopyUser(u1, nil)
	fmt.Printf("Name: %s, id: %d", u2.Name, u2.Id)
}
