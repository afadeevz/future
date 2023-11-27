package future

type Value[T any] interface {
	Future[T]
	Resolve(T, error)
}

type value[T any] struct {
	value  T
	err    error
	waiter Waiter
}

func NewValue[T any]() Value[T] {
	return &value[T]{
		waiter: NewWaiter(),
	}
}

func (f *value[T]) Wait() (T, error) {
	f.waiter.Wait()
	return f.value, f.err
}

func (f *value[T]) Resolve(value T, err error) {
	f.value = value
	f.err = err
	f.waiter.Done()
}
