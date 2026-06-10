// `forvar`: Go 1.22 made each range iteration have a fresh loop variable, so
// the classic `x := x` shadowing trick before launching a goroutine is now
// redundant and can be removed.
package forvar

func Run(xs []int) {
	done := make(chan int, len(xs))
	for _, x := range xs {
		x := x // ← redundant since Go 1.22
		go func() { done <- x }()
	}
	for range xs {
		<-done
	}
}
