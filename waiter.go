package future

import "sync"

type Waiter interface {
	Done()
	Wait()
}

type waiter struct {
	wg sync.WaitGroup
}

func NewWaiter() Waiter {
	wg := sync.WaitGroup{}
	wg.Add(1)
	return &waiter{
		wg: wg,
	}
}

func (w *waiter) Done() {
	w.wg.Done()
}

func (w *waiter) Wait() {
	w.wg.Wait()
}
