package future

type Waiter interface {
	Wait()
	Done()
}

type waiter struct {
	c    chan struct{}
	done bool
}

func NewWaiter() Waiter {
	return &waiter{
		c: make(chan struct{}, 1),
	}
}

func (w *waiter) Wait() {
	if !w.done {
		w.c <- <-w.c
	}
}

func (w *waiter) Done() {
	w.c <- struct{}{}
	w.done = true
}
