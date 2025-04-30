package inspections

/*
Simplified example from real project.

Explanation:
The code has a variable shadowing issue. Inside the `if conf == nil block`, a new local variable conf is declared,
which shadows the outer conf. As a result, any changes to conf inside the if block do not affect the outer one.
After the block, the original (still nil) conf is used, potentially causing a nil pointer dereference
when accessing conf.Root.
*/
func _() {
	conf := LoadConfig()

	if conf == nil {
		conf, err := LoadConfigFromPath() // `conf` shadows outer `conf` variable
		if err != nil || conf == nil {
			return
		}
	}
	process(conf.Root)
}

/* Interprocedural analysis (will be available in 252).
Interprocedural analysis tracks data and control flow across function boundaries to detect nil dereference bugs.
It models how nil values are passed, returned, and propagated between functions, allowing the analyzer to identify
unsafe dereferences that aren’t visible within a single function’s scope.
*/

/*
Explanation:
The problem is that NewUser() may return nil (check definition), but the code immediately accesses `user.Age`
and `user.Name` without checking.

Possible fix:
* Make NewUser() return (User, error)
* Check `user` for nil before accessing it
*/
func _() {
	user := NewUser("123")
	println(user.Age, user.Name)
}

/*
Explanation:
The error handling in the code is correct, but there is an issue with the contract of the (result, error) return value
in the LoadUser function. The function LoadUser can return (nil, nil) meaning there's no error, but the user is still nil.
In such a case even `err` is checked for nil, `user` can still be `nil` too.

Possible fix:
Don't return (nil, nil) in LoadUser => return (nil, errors.New("..."))
*/
func _() {
	user, err := LoadUser()
	if err != nil {
		return
	}
	println(user.Age, user.Name)
}
