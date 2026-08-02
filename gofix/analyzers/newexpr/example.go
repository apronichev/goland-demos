// `newexpr`: Go 1.26 extends the `new` builtin to accept an expression, so
// `new(123)` returns a `*int` pointing to a fresh `123`. Wrapper functions
// like `IntPtr` (common when populating JSON/protobuf pointer fields) become
// inlinable thin shells, and existing call sites get rewritten too.
package newexpr

func IntPtr(x int) *int { return &x }

func StringPtr(s string) *string { return &s }

type Config struct {
	Retries *int
	Label   *string
}

func Default() Config {
	return Config{
		Retries: IntPtr(3),
		Label:   StringPtr("default"),
	}
}
