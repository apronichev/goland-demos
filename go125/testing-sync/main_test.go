package main

import (
	"testing"
	"testing/synctest"
	"time"
)

// A simple rate limiter that allows 1 request per second
type RateLimiter struct {
	lastRequest time.Time
}

func (r *RateLimiter) Allow() bool {
	now := time.Now()
	if now.Sub(r.lastRequest) < time.Second {
		return false
	}
	r.lastRequest = now
	return true
}

// Worker that processes items with rate limiting
func processItems(limiter *RateLimiter, items []string, results chan<- string) {
	for _, item := range items {
		// Wait until rate limiter allows
		for !limiter.Allow() {
			time.Sleep(100 * time.Millisecond)
		}
		results <- "processed: " + item
	}
	close(results)
}

func TestRateLimiterWithSynctest(t *testing.T) {
	start := time.Now()

	synctest.Test(t, func(t *testing.T) {
		limiter := &RateLimiter{}
		items := []string{"A", "B", "C", "D", "E"} // 5 items = 4 seconds of waiting
		results := make(chan string)

		// Start processing in a goroutine
		go processItems(limiter, items, results)

		// Process all items
		for i := 0; i < 5; i++ {
			got := <-results
			t.Logf("Time %v: %s", time.Now(), got)

			if i < 4 { // Don't wait after last item
				// Wait for the goroutine to block on rate limiting
				synctest.Wait()

				// Advance fake clock by 1 second
				time.Sleep(1 * time.Second)
			}
		}
	})

	elapsed := time.Since(start)
	t.Logf("🚀 Synctest version completed in: %v", elapsed)
}

// Without synctest, this test takes real time!
func TestRateLimiterNormal(t *testing.T) {
	start := time.Now()

	limiter := &RateLimiter{}
	items := []string{"A", "B", "C", "D", "E"} // 5 items = 4 seconds of waiting
	results := make(chan string)

	// Start processing in a goroutine
	go processItems(limiter, items, results)

	// Process all items
	for i := 0; i < 5; i++ {
		got := <-results
		t.Logf("Time %v: %s", time.Now(), got)
	}

	elapsed := time.Since(start)
	t.Logf("🐌 Normal version completed in: %v", elapsed)
}

// Test with even more dramatic timing - 10 second delays
type SlowService struct {
	lastCall time.Time
}

func (s *SlowService) DoWork() string {
	now := time.Now()
	if !s.lastCall.IsZero() && now.Sub(s.lastCall) < 10*time.Second {
		time.Sleep(10*time.Second - now.Sub(s.lastCall))
	}
	s.lastCall = time.Now()
	return "work done"
}

func TestSlowServiceWithSynctest(t *testing.T) {
	start := time.Now()

	synctest.Test(t, func(t *testing.T) {
		service := &SlowService{}

		// First call is immediate
		result1 := service.DoWork()
		t.Logf("First result: %s", result1)

		// Start second call in goroutine
		done := make(chan string)
		go func() {
			done <- service.DoWork()
		}()

		// Wait for goroutine to block
		synctest.Wait()

		// Advance time by 10 seconds instantly!
		time.Sleep(10 * time.Second)

		result2 := <-done
		t.Logf("Second result after 10 second delay: %s", result2)
	})

	elapsed := time.Since(start)
	t.Logf("🚀 Synctest with 10-second delay completed in: %v", elapsed)
}

func TestSlowServiceNormal(t *testing.T) {
	start := time.Now()

	service := &SlowService{}

	// First call is immediate
	result1 := service.DoWork()
	t.Logf("First result: %s", result1)

	// Second call waits 10 real seconds!
	result2 := service.DoWork()
	t.Logf("Second result after 10 second delay: %s", result2)

	elapsed := time.Since(start)
	t.Logf("🐌 Normal test with 10-second delay completed in: %v", elapsed)
}
