package main

import (
	"fmt"
	"testing"
)

// Pizza order for our examples
type Order struct {
	ID       int
	Customer string
	Item     string
}

func TestPizzaOrder(t *testing.T) {
	// NEW in Go 1.25: Test Attributes (metadata tags)
	t.Attr("test_id", "ORD-001")
	t.Attr("component", "ordering-system")
	t.Attr("severity", "critical")

	order := Order{ID: 1, Customer: "Bob", Item: "Pepperoni Pizza"}

	// Regular log (includes file:line)
	t.Logf("Processing order #%d", order.ID)

	// NEW: Clean output without file:line clutter
	out := t.Output()
	fmt.Fprintln(out, "\n🍕 Order Receipt (Clean Output):")
	fmt.Fprintf(out, "   Order: #%d\n", order.ID)
	fmt.Fprintf(out, "   Customer: %s\n", order.Customer)
	fmt.Fprintf(out, "   Item: %s\n", order.Item)
	fmt.Fprintln(out, "   Status: ✅ Confirmed")
}

func BenchmarkOrderProcessing(b *testing.B) {
	// Attributes work in benchmarks too
	b.Attr("operation", "order-creation")
	b.Attr("database", "postgresql")

	for i := 0; i < b.N; i++ {
		_ = Order{ID: i, Customer: "Test", Item: "Pizza"}
	}
}

// This test is SAFE - runs without parallel
func TestSafeAllocations(t *testing.T) {
	testing.AllocsPerRun(10, func() {
		orders := make([]Order, 0, 5)
		orders = append(orders, Order{ID: 1, Customer: "Alice", Item: "Pizza"})
		// Force the compiler to keep the allocation
		if len(orders) == 0 {
			panic("unexpected")
		}
	})
}

// This test would PANIC in Go 1.25 - demonstrates the new safety feature
func TestUnsafeAllocations(t *testing.T) {
	// Uncomment the next line to see the panic in Go 1.25!
	// t.Parallel() // ⚠️ DANGER: This causes AllocsPerRun to panic!

	t.Log("⚠️  WARNING: If t.Parallel() was enabled, AllocsPerRun would panic!")
	t.Log("   Go 1.25 prevents flaky allocation measurements")

	// This would panic if t.Parallel() was uncommented:
	/*
		testing.AllocsPerRun(10, func() {
			_ = make([]byte, 1024)
		})
	*/
}
