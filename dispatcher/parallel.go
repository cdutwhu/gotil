package dispatcher

import "sync"

// SyncParallel : Async dispatch n threads of func. Once all done, Sync then return.
// DO NOT apply to compute with a shared variable as slow than normal.
// [f] start with `defer func() { done <- tid }()`
func SyncParallel(n int, f func(dim, tid int, done chan int, params ...interface{}), params ...interface{}) {
	if n < 1 {
		n = 1
	}
	wg, done := sync.WaitGroup{}, make(chan int, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go f(n, i, done, params...)
	}
	i := 0
	for range done {
		if i == n-1 {
			close(done)
		}
		i++
	}
	wg.Done()
}
