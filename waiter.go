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
	w := &waiter{}
	w.wg.Add(1)
	return w
}

func (w *waiter) Wait() {
	w.wg.Wait()
}

func (w *waiter) Done() {
	w.wg.Done()
}
