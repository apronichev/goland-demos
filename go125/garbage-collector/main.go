package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

type CacheEntry struct {
	Key   string
	Value string
	TTL   time.Time
}

func main() {
	if isGreenTeaGC() {
		fmt.Println("Running with GreenTea" + "GC")
	} else {
		fmt.Println("Running with standard " + "GC")
	}

	// Test 1: Many small objects performance
	fmt.Println("\nTest 1: Small Objects Handling")
	testSmallObjects()

	// Test 2: CPU scalability test
	fmt.Println("\nTest 2: CPU Scalability")
	testCPUScalability()

	// Show memory stats
	showMemoryStats()
}

func testSmallObjects() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	initialAlloc := m.Alloc

	start := time.Now()

	// Create millions of small objects
	cache := make(map[string]*CacheEntry)
	for i := 0; i < 1_000_000; i++ {
		key := fmt.Sprintf("key-%d", i)
		cache[key] = &CacheEntry{
			Key:   key,
			Value: fmt.Sprintf("value-%d", i),
			TTL:   time.Now().Add(time.Hour),
		}

		// Delete old entries to trigger GC
		if i > 100_000 {
			delete(cache, fmt.Sprintf("key-%d", i-100_000))
		}
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m)

	fmt.Printf("  • Objects created: 1,000,000\n")
	fmt.Printf("  • Time taken: %v\n", duration)
	fmt.Printf("  • Memory allocated: %.2f MB\n", float64(m.Alloc-initialAlloc)/1024/1024)
	fmt.Printf("  • GC runs: %d\n", m.NumGC)
	fmt.Printf("  • Avg pause per GC: %.2f ms\n", float64(m.PauseTotalNs)/float64(m.NumGC)/1_000_000)
}

func testCPUScalability() {
	// Test with different CPU counts
	maxCPUs := runtime.NumCPU()

	for cpus := 1; cpus <= maxCPUs; cpus *= 2 {
		runtime.GOMAXPROCS(cpus)

		start := time.Now()
		var wg sync.WaitGroup

		// Run parallel workload
		for i := 0; i < cpus; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()

				// Each goroutine creates many small objects
				local := make([]*CacheEntry, 0, 10000)
				for j := 0; j < 10000; j++ {
					local = append(local, &CacheEntry{
						Key:   fmt.Sprintf("cpu-%d-item-%d", id, j),
						Value: "test-data",
						TTL:   time.Now(),
					})
				}

				// Force some GC work
				runtime.GC()
			}(i)
		}

		wg.Wait()
		duration := time.Since(start)

		throughput := float64(cpus*10000) / duration.Seconds()
		fmt.Printf("  • CPUs: %d, Time: %v, Throughput: %.0f objects/sec\n",
			cpus, duration, throughput)
	}
}

func showMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("\nMemory Statistics:")
	fmt.Printf("  • Total GC pause time: %.2f ms\n", float64(m.PauseTotalNs)/1_000_000)
	fmt.Printf("  • Objects allocated: %d\n", m.Mallocs)
	fmt.Printf("  • Objects freed: %d\n", m.Frees)
	fmt.Printf("  • Live objects: %d\n", m.Mallocs-m.Frees)
}

func isGreenTeaGC() bool {
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, setting := range info.Settings {
			if setting.Key == "GOEXPERIMENT" &&
				strings.Contains(setting.Value, "greenteagc") {
				return true
			}
		}
	}
	return false
}
