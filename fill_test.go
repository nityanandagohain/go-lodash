package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fillTests = []struct {
	name      string
	array     []int
	value     int
	options   *FillOptions
	want      []int
	isError   *bool
	wantError *string
}{
	{"t1", []int{1, 2, 3}, 10, &FillOptions{}, []int{10, 10, 10}, nil, nil},
	{"t2", []int{1, 2, 3}, 10, &FillOptions{start: GetPointer(1)}, []int{1, 10, 10}, nil, nil},
}

func TestFill(t *testing.T) {
	is := assert.New(t)
	for _, tt := range fillTests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Fill(tt.array, tt.value, tt.options)
			if tt.isError != nil && *tt.isError {
				is.Equal(tt.wantError, err.Error())
			} else {
				is.Equal(tt.want, res)
			}
		})
	}
}
