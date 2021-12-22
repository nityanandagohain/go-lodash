package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lastIndexOfTests = []struct {
	name    string
	array   []int
	value   int
	options *LastIndexOfOptions
	want    int
}{
	{"without options", []int{1, 2, 1, 2}, 2, nil, 3},
	{"with options", []int{1, 2, 1, 2}, 2, &LastIndexOfOptions{FromIndex: GetPointer(2)}, 1},
}

func TestLastIndexOf(t *testing.T) {
	is := assert.New(t)
	for _, tt := range lastIndexOfTests {
		t.Run(tt.name, func(t *testing.T) {
			is.Equal(tt.want, LastIndexOf(tt.array, tt.value, tt.options))
		})
	}
}
