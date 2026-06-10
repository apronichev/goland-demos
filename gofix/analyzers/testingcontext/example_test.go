// `testingcontext`: inside a test, `t.Context()` (Go 1.24+) already returns a
// context that's cancelled when the test ends, so the manual
// `context.WithCancel` + `defer cancel()` pair is just noise.
package testingcontext

import (
	"context"
	"testing"
)

func TestSomething(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := work(ctx); err != nil {
		t.Fatal(err)
	}
}

func work(ctx context.Context) error {
	<-ctx.Done()
	return nil
}
