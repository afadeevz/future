package future

import (
	"testing"
	"strings"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	result, err := NewPromise(func() (string, error) {
		return "3 14 15 92 65", nil
	}).Then(func(s string) ([]string, error) {
		return strings.Split(s, " "), nil
	}).Then(func(s []string) ([]int64, error) {
		var ints []int64

		for _, si := range s {
			val, err := strconv.ParseInt(si, 10, 64)
			if err != nil {
				return nil, err
			}

			ints = append(ints, val)
		}

		return ints, nil
	}).Result()

	assert.Nil(t, err)
	assert.Equal(t, []int64{3, 14, 15, 92, 65}, result)
}
