package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var indexOfTests = []struct {
	name  string
	array []int
	value int
	want  int
}{
	{"value present", []int{1, 2, 3, 9}, 2, 1},
	{"out of index 1", []int{1, 2, 3, 9}, 10, -1},
	{"out of index 2", []int{1, 2, 3, 9}, 0, -1},
}

func TestIndexOf(t *testing.T) {
	is := assert.New(t)
	for _, tt := range indexOfTests {
		t.Run(tt.name, func(t *testing.T) {
			is.Equal(IndexOf(tt.array, tt.value), tt.want)
		})
	}
}

var indexOfFromTests = []struct {
	name  string
	array []int
	value int
	from  int
	want  int
}{
	{"value present", []int{1, 2, 3, 9}, 3, 2, 2},
	{"out of index 1", []int{1, 2, 3, 9}, 10, 2, -1},
	{"out of index 2", []int{1, 2, 3, 9}, 1, 3, -1},
}

func TestIndexOfFrom(t *testing.T) {
	is := assert.New(t)
	for _, tt := range indexOfTests {
		t.Run(tt.name, func(t *testing.T) {
			is.Equal(IndexOf(tt.array, tt.value), tt.want)
		})
	}
}

func TestIndexOf_string(t *testing.T) {
	is := assert.New(t)
	input := []string{"hello", "world", "how", "are", "you"}
	res := IndexOf(input, "are")
	is.Equal(res, 3)
}

func TestIndexOf_float(t *testing.T) {
	is := assert.New(t)
	input := []float64{1.1, 2.3, 5.5}
	res := IndexOf(input, 2.3)
	is.Equal(res, 1)
}
