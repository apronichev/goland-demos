package allocreduction

import "testing"

const n = 1000

// BenchmarkBuildSlow allocates one heap object per element (~N allocs/op).
func BenchmarkBuildSlow(b *testing.B) {
	var sink float64
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink += sumSlow(BuildSlow(n))
	}
	_ = sink
}

// BenchmarkBuildFast allocates only the backing array (~1 alloc/op).
func BenchmarkBuildFast(b *testing.B) {
	var sink float64
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink += sumFast(BuildFast(n))
	}
	_ = sink
}

// TestBuildEquivalent confirms both versions produce the same data, so the
// optimization is behavior-preserving for read-only consumers.
func TestBuildEquivalent(t *testing.T) {
	slow := BuildSlow(n)
	fast := BuildFast(n)
	if len(slow) != len(fast) {
		t.Fatalf("length mismatch: slow=%d fast=%d", len(slow), len(fast))
	}
	if got, want := sumFast(fast), sumSlow(slow); got != want {
		t.Fatalf("sum mismatch: fast=%v slow=%v", got, want)
	}
}

// TestAllocationsReduced compares the two versions by allocation count and
// guards the optimization against regressions. BuildSlow allocates one heap
// object per element (~n); BuildFast keeps elements in one backing array.
func TestAllocationsReduced(t *testing.T) {
	slowAllocs := testing.AllocsPerRun(100, func() { _ = sumSlow(BuildSlow(n)) })
	fastAllocs := testing.AllocsPerRun(100, func() { _ = sumFast(BuildFast(n)) })

	t.Logf("allocs/op: slow=%.0f fast=%.0f", slowAllocs, fastAllocs)

	// Slow must allocate per element; if this drops, the example is no longer
	// illustrating the problem.
	if slowAllocs < float64(n) {
		t.Errorf("BuildSlow expected to allocate at least %d times, got %.0f", n, slowAllocs)
	}
	// Fast must allocate dramatically less — the whole point of the change.
	if fastAllocs >= slowAllocs/10 {
		t.Errorf("BuildFast should allocate far less than BuildSlow: fast=%.0f slow=%.0f", fastAllocs, slowAllocs)
	}
}
