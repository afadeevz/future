package future

import (
	"testing"
)

func benchmarkWaiter(newWaiter func() Waiter, n int) {
	for i := 0; i < n; i++ {
		w := newWaiter()
		go func() {
			w.Done()
		}()
		w.Wait()
		w.Wait()
	}
}

func BenchmarkWaiter(b *testing.B) {
	benchmarkWaiter(NewWaiter, b.N)
}
