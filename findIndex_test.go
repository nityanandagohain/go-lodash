package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndex(t *testing.T) {
	is := assert.New(t)
	res := FindIndex([]int{1, 2, 3, 4}, func(value int, index int, array []int) bool {
		if value == 3 {
			return true
		}
		return false
	},
		nil,
	)
	is.Equal(2, res)
}
