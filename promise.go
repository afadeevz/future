package future

import (
	"reflect"
	"fmt"
)

type Promise interface {
	Then(fn interface{}) Promise // fn should be a func(value T1) (T2, error)
	Wait()
	Result() (interface{}, error)
}

type promise struct {
	fn     interface{}
	result interface{}
	err    error
	Waiter
}

// fn should be a func() (T, error)
func NewPromise(fn interface{}) Promise {
	p := &promise{
		fn:     fn,
		Waiter: NewWaiter(),
	}

	p.checkTypes()
	go p.resolve()

	return p
}

func (p *promise) checkTypes() {
	checkIsValidHandlerFunction(p.fn, 0)
}

func (p *promise) resolve() {
	rv := reflect.ValueOf(p.fn)
	retValues := rv.Call(nil)

	p.result = retValues[0].Interface()

	iErr := retValues[1].Interface()
	if iErr != nil {
		p.err = iErr.(error)
	}

	p.Done()
}

type thenHandler struct {
	impl interface{}
}

func (h *thenHandler) handle(v interface{}) (interface{}, error) {
	h.validate(v)

	rv := reflect.ValueOf(h.impl)
	retValues := rv.Call([]reflect.Value{reflect.ValueOf(v)})

	iRes := retValues[0].Interface()
	iErr := retValues[1].Interface()
	if iErr != nil {
		return iRes, iErr.(error)
	}
	return iRes, nil
}

func (h *thenHandler) validate(v interface{}) {
	checkIsValidHandlerFunction(h.impl, 1)

	rto := reflect.TypeOf(v)
	rti := reflect.TypeOf(h.impl).In(0)

	if !rto.ConvertibleTo(rti) {
		panic(fmt.Sprintf("Mismatching types %s and %s", rto.Name(), rti.Name()))
	}
}

func (p *promise) Then(fn interface{}) Promise {
	return NewPromise(func() (interface{}, error) {
		p.Wait()
		if p.err != nil {
			return nil, p.err
		}

		h := thenHandler{fn}
		return h.handle(p.result)
	})
}

func (p *promise) Result() (interface{}, error) {
	p.Wait()
	return p.result, p.err
}

func checkIsValidHandlerFunction(fn interface{}, numOfInArgs int) {
	rt := reflect.TypeOf(fn)

	if rk := rt.Kind(); rk != reflect.Func {
		panic(fmt.Sprintf("Expected function but got %s %s", rk, rt.Name()))
	}

	if n := rt.NumIn(); n != numOfInArgs {
		panic(fmt.Sprintf("Expected function with %d arguments, got %d", numOfInArgs, n))
	}

	if n := rt.NumOut(); n != 2 {
		panic(fmt.Sprintf("Expected function with exactly 2 return values, got %d", n))
	}

	rtErr := reflect.TypeOf((*error)(nil)).Elem()
	if t := rt.Out(1); !t.Implements(rtErr) {
		panic(fmt.Sprintf("Expected function error second return value, got %s", t.Name()))
	}
}
