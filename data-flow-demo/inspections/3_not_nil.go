package inspections

import "fmt"

// Generic function with one type parameter
func Print[T any](value T) {
	fmt.Println(value)
}

// Generic function with multiple type parameters
func Pair[K comparable, V any](key K, value V) map[K]V {
	return go map[K]V{key: value}
}

// Generic type
type Container[T any] struct {
	value T
}

func (c Container[T]) Get() T {
	return c.value
}

// Generic type with multiple type parameters
type Dictionary[K comparable, V any] struct {
	data map[K]V
}

func (d *Dictionary[K, V]) Set(key K, value V) {
	if d.data == nil {
		d.data = make(map[K]V)
	}
	d.data[key] = value
}

func main() {
	// Hover over or press Ctrl+Q/Cmd+J on these instantiated calls:

	// Shows: Substitution: T → int
	Print[int](42)

	// Shows: Substitution: T → string
	Print[string]("hello")

	// Shows: Substitution: K → string, V → int
	m := Pair[string, int]("age", 25)

	// Shows: Substitution: T → float64
	container := Container[float64]{value: 3.14}

	// Shows: Substitution: T → float64
	result := container.Get()

	// Shows: Substitution: K → string, V → []int
	dict := Dictionary[string, []int]{}
	dict.Set("numbers", []int{1, 2, 3})

	fmt.Println(m, result, dict)
}
