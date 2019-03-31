package future

type Future interface {
	Wait()
	Value() interface{}
}

type Resolver interface {
	Future

	Resolve(value interface{})
}

type future struct {
	Waiter

	value interface{}
}

func NewFuture() Resolver {
	return &future{
		Waiter: NewWaiter(),
	}
}

func (f *future) Value() interface{} {
	f.Wait()
	return f.value
}

func (f *future) Resolve(value interface{}) {
	f.value = value
	f.Done()
}
