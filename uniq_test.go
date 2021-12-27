package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var UniqTests = []struct {
	name  string
	array []int
	want  []int
}{
	{"test 1", []int{1, 2, 1, 2}, []int{1, 2}},
	{"test 2", []int{1, 3, 5, 9}, []int{1, 3, 5, 9}},
	{"test empty", []int{}, []int{}},
}

func TestUniq(t *testing.T) {
	is := assert.New(t)
	for _, tt := range UniqTests {
		t.Run(tt.name, func(t *testing.T) {
			is.ElementsMatch(tt.want, Uniq(tt.array))
		})
	}
}

func TestUniq_string(t *testing.T) {
	is := assert.New(t)
	input := []string{"hello", "world", "world"}
	res := Uniq(input)
	is.ElementsMatch([]string{"hello", "world"}, res)
}
