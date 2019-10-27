package future

import (
	"sync"
	"testing"
)

func TestWaitAfterDone(t *testing.T) {
	t.Parallel()
	w := NewWaiter()
	w.Done()
	w.Wait()
}

func TestWaitBeforeDone(t *testing.T) {
	t.Parallel()
	w := NewWaiter()

	go func() {
		w.Done()
	}()

	w.Wait()
}

func TestWaitMultiple(t *testing.T) {
	t.Parallel()
	w := NewWaiter()

	const count = 42

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			w.Wait()
			wg.Done()
		}()
	}

	w.Done()
	wg.Wait()
}
