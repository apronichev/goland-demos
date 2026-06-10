// `waitgroupgo`: Go 1.25 added `(*WaitGroup).Go(func())`, which couples the
// `Add(1)` and the `defer Done()` to the goroutine launch — so the classic
// three-line idiom collapses into one call.
package waitgroupgo

import "sync"

func Fanout(tasks []func()) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			task()
		}()
	}
	wg.Wait()
}
