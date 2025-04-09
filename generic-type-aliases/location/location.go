package pkgA

type Location[T any] struct {
	Object T
	X, Y   float64
}
