// `atomictypes`: the primitive `sync/atomic` functions (AddInt32, LoadUint64,
// …) take a `*T` and rely on the caller to never touch the variable
// non-atomically. The typed wrappers (`atomic.Int32`, `atomic.Uint64`, …) make
// the atomicity part of the type and also fix the 32-bit alignment trap.
package atomictypes

import "sync/atomic"

type Counter struct {
	hits int64
}

func (c *Counter) Bump() {
	atomic.AddInt64(&c.hits, 1)
}

func (c *Counter) Read() int64 {
	return atomic.LoadInt64(&c.hits)
}
