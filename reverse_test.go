package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ReverseTests = []struct {
	name  string
	array []int
	want  []int
}{
	{"test 1", []int{1, 2, 1, 2}, []int{2, 1, 2, 1}},
	{"test 2", []int{1, 3, 5, 9}, []int{9, 5, 3, 1}},
	{"test odd", []int{1, 3, 5}, []int{5, 3, 1}},
	{"test empty", []int{}, []int{}},
}

func TestReverse(t *testing.T) {
	is := assert.New(t)
	for _, tt := range ReverseTests {
		t.Run(tt.name, func(t *testing.T) {
			is.ElementsMatch(tt.want, Reverse(tt.array))
		})
	}
}

func TestReverse_string(t *testing.T) {
	is := assert.New(t)
	input := []string{"hello", "world"}
	res := Reverse(input)
	is.ElementsMatch([]string{"world", "hello"}, res)
}
