package function

import (
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	lastRun = make(map[string]time.Time)
)

// Debounce ensures the given function is executed at most once within the specified duration.
func Debounce(fn func(), d time.Duration) func() {
	fnName := Name(fn)
	return func() {
		mu.Lock()
		lastExecution, exists := lastRun[fnName]
		mu.Unlock()
		if !exists || time.Since(lastExecution) >= d {
			lastRun[fnName] = time.Now()
			fn()
		}
	}
}
