package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var concatTests = []struct {
	name   string
	arrays [][]int
	want   []int
}{
	{"test 1", [][]int{{1, 2, 3, 4}, {5}}, []int{1, 2, 3, 4, 5}},
	{"test empty", [][]int{}, []int{}},
	{"test nil", nil, []int{}},
}

func TestConcat(t *testing.T) {
	is := assert.New(t)
	for _, tt := range concatTests {
		t.Run(tt.name, func(t *testing.T) {
			res := Concat(tt.arrays...)
			is.Equal(tt.want, res)
		})
	}
}

func TestConcatMap(t *testing.T) {
	is := assert.New(t)
	input := [][]map[string]interface{}{
		{
			{"a": "b"},
			{"c": "d"},
			{"e": "f"},
		},
		{
			{"g": "h"},
		},
	}

	expected := []map[string]interface{}{
		{"a": "b"},
		{"c": "d"},
		{"e": "f"},
		{"g": "h"},
	}

	res := Concat(input...)
	is.Equal(expected, res)

}
