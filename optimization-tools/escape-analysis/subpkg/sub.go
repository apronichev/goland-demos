package subpkg

func SpillInt() *int {
	x := 42
	return &x
}

func NoEscape(x int) int {
	return x + 1
}
