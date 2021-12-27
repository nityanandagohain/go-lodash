package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var indexOfTests = []struct {
	name    string
	array   []int
	options *IndexOfOptions
	value   int
	want    int
}{
	{"value present", []int{1, 2, 3, 9}, nil, 2, 1},
	{"using from", []int{1, 2, 3, 9}, &IndexOfOptions{FromIndex: GetPointer(1)}, 3, 2},
	{"out of index 2", []int{1, 2, 3, 9}, nil, 0, -1},
}

func TestIndexOf(t *testing.T) {
	is := assert.New(t)
	for _, tt := range indexOfTests {
		t.Run(tt.name, func(t *testing.T) {
			is.Equal(IndexOf(tt.array, tt.value, tt.options), tt.want)
		})
	}
}

func TestIndexOf_string(t *testing.T) {
	is := assert.New(t)
	input := []string{"hello", "world", "how", "are", "you"}
	res := IndexOf(input, "are", nil)
	is.Equal(res, 3)
}

func TestIndexOf_float(t *testing.T) {
	is := assert.New(t)
	input := []float64{1.1, 2.3, 5.5}
	res := IndexOf(input, 2.3, nil)
	is.Equal(res, 1)
}
