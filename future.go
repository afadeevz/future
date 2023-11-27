package future

type Future[T any] interface {
	Wait() (T, error)
}

type GoFunc[T any] func() (T, error)

func Go[T any](fn GoFunc[T]) Future[T] {
	result := NewValue[T]()

	go func() {
		val, err := fn()
		result.Resolve(val, err)
	}()

	return result
}

type ThenFunc[T1 any, T2 any] func(T1) (T2, error)

func Then[T1 any, T2 any](f Future[T1], thenFunc ThenFunc[T1, T2]) Future[T2] {
	return Go(func() (T2, error) {
		val, err := f.Wait()
		if err != nil {
			var zero T2
			return zero, err
		}

		return thenFunc(val)
	})
}
