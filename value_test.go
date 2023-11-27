package future

import (
	"fmt"
	"testing"

	"github.com/AlexanderFadeev/myerrors"
	"github.com/stretchr/testify/assert"
)

func TestValueWaitAfterResolved(t *testing.T) {
	t.Parallel()

	v := NewValue[int]()
	v.Resolve(42, nil)
	val, err := v.Wait()

	assert.Equal(t, 42, val)
	assert.Nil(t, err)
}

func TestValueWaitBeforeResolved(t *testing.T) {
	t.Parallel()

	v := NewValue[int]()

	go func() {
		v.Resolve(42, nil)
	}()

	val, err := v.Wait()
	assert.Equal(t, 42, val)
	assert.Nil(t, err)
}

func TestValueThen(t *testing.T) {
	t.Parallel()

	v := NewValue[int]()

	go func() {
		v.Resolve(42, nil)
	}()

	val, err := Then(v, func(v int) (string, error) {
		str := fmt.Sprint(v)
		return str, nil
	}).Wait()

	assert.Equal(t, val, "42")
	assert.Nil(t, err)
}

func TestValueThenError(t *testing.T) {
	t.Parallel()

	v := NewValue[int]()

	go func() {
		v.Resolve(0, myerrors.New("err"))
	}()

	val, err := Then(v, func(v int) (string, error) {
		str := fmt.Sprint(v)
		return str, nil
	}).Wait()

	assert.Equal(t, "", val)
	assert.NotNil(t, err)
}
