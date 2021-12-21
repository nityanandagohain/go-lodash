package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var differenceTests = []struct {
	name   string
	array  []int
	values []int
	want   []int
}{
	{"check difference", []int{1, 2}, []int{2, 3}, []int{1}},
	{"same result", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}},
	{"emoty result", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
}

func TestDifference(t *testing.T) {
	is := assert.New(t)
	for _, tt := range differenceTests {
		t.Run(tt.name, func(t *testing.T) {
			res := Difference(tt.array, tt.values)
			is.Equal(tt.want, res)
		})
	}
}

var differenceTestsString = []struct {
	name   string
	array  []string
	values []string
	want   []string
}{
	{"check difference", []string{"a", "b"}, []string{"c", "d"}, []string{"a"}},
	{"check difference multiple chars", []string{"hello", "world"}, []string{"hello", "john"}, []string{"hello"}},
	{"same result", []string{"a", "b", "c"}, []string{"d", "e", "f"}, []string{"a", "b", "c"}},
	{"emoty result", []string{"a", "b", "c"}, []string{"a", "b", "c"}, []string{}},
}

func TestDifference_String(t *testing.T) {
	is := assert.New(t)
	for _, tt := range differenceTests {
		t.Run(tt.name, func(t *testing.T) {
			res := Difference(tt.array, tt.values)
			is.Equal(tt.want, res)
		})
	}
}
