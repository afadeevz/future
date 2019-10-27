package future

type Value interface {
	Future
	Resolve(interface{}, error)
}

type value struct {
	value  interface{}
	err    error
	waiter Waiter
}

func NewValue() Value {
	return &value{
		waiter: NewWaiter(),
	}
}

func (f *value) Wait() (interface{}, error) {
	f.waiter.Wait()
	return f.value, f.err
}

func (f *value) Resolve(value interface{}, err error) {
	f.value = value
	f.err = err
	f.waiter.Done()
}

func (f *value) Then(thenFunc ThenFunc) Future {
	return Go(func() (interface{}, error) {
		val, err := f.Wait()
		if err != nil {
			return nil, err
		}

		return thenFunc(val)
	})
}
