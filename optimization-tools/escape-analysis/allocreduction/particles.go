// Package allocreduction shows a small, realistic code change that cuts heap
// allocations dramatically — and how escape analysis makes it visible.
//
// Scenario: building a batch of records in a hot loop (think: parsing rows,
// collecting events, simulating particles). The idiomatic-but-costly habit is
// to collect a slice of POINTERS. Switching to a slice of VALUES lets the
// compiler keep every element in one backing array instead of escaping each
// one to the heap.
//
// See the difference yourself:
//
//	go build -gcflags='-m=2' ./allocreduction      # watch &Particle{...} escapes to heap (slow only)
//	go test -bench=. -benchmem ./allocreduction     # allocs/op: ~N (slow) vs ~1 (fast)
package allocreduction

// Particle is a plain value type — no pointers inside, cheap to copy
type Particle struct {
	X, Y, Z float64
}

// BuildSlow collects []*Particle. Each `&Particle{...}` outlives the loop
// iteration (it is stored in a slice that is returned), so escape analysis is
// forced to move every element to the heap: ~N allocations for N elements.
//
// -gcflags=-m reports:  &Particle{...} escapes to heap
func BuildSlow(n int) []*Particle {
	out := make([]*Particle, 0, n)
	for i := range n {
		out = append(out, &Particle{X: float64(i), Y: float64(i) * 2, Z: float64(i) * 3})
	}
	return out
}

// BuildFast collects []Particle. The values are copied straight into the
// slice's backing array, so only that one array is allocated (plus growth
// reallocations if the capacity is not pre-sized): ~1 allocation total.
//
// The change is mechanical: []*Particle -> []Particle and &Particle{...} ->
// Particle{...}. Callers that only read fields need no other change.
func BuildFast(n int) []Particle {
	out := make([]Particle, 0, n)
	for i := range n {
		out = append(out, Particle{X: float64(i), Y: float64(i) * 2, Z: float64(i) * 3})
	}
	return out
}

// sumSlow / sumFast exist so the benchmarks do real work with the result and
// the compiler cannot optimize the construction away.

func sumSlow(ps []*Particle) float64 {
	var s float64
	for _, p := range ps {
		s += p.X + p.Y + p.Z
	}
	return s
}

func sumFast(ps []Particle) float64 {
	var s float64
	for i := range ps {
		s += ps[i].X + ps[i].Y + ps[i].Z
	}
	return s
}
