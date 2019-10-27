package future

type GoFunc func() (interface{}, error)
type ThenFunc func(interface{}) (interface{}, error)

type Future interface {
	Wait() (interface{}, error)
	Then(ThenFunc) Future
}

func Go(fn GoFunc) Future {
	result := NewValue()

	go func() {
		val, err := fn()
		result.Resolve(val, err)
	}()

	return result
}
