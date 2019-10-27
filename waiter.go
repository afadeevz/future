package future

import (
	"sync"
)

type Waiter interface {
	Wait()
	Done()
}

type waiter struct {
	wg sync.WaitGroup
}

func NewWaiter() Waiter {
	var wg sync.WaitGroup
	wg.Add(1)

	return &waiter{
		wg: wg,
	}
}

func (w *waiter) Wait() {
	w.wg.Wait()
}

func (w *waiter) Done() {
	w.wg.Done()
}
