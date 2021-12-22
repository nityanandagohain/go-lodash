package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var joinTests = []struct {
	name      string
	array     []int
	seperator string
	want      string
}{
	{"join", []int{1, 2, 3, 9}, "-", "1-2-3-9"},
}

func TestJoin(t *testing.T) {
	is := assert.New(t)
	for _, tt := range joinTests {
		t.Run(tt.name, func(t *testing.T) {
			is.Equal(tt.want, Join(tt.array, tt.seperator))
		})
	}
}

var joinTest_string = []struct {
	name      string
	array     []string
	seperator string
	want      string
}{
	{"join", []string{"abc", "def", "ghi", "jkl"}, "-", "abc-def-ghi-jkl"},
}

func TestJoin_string(t *testing.T) {
	is := assert.New(t)
	for _, tt := range joinTest_string {
		t.Run(tt.name, func(t *testing.T) {
			is.Equal(tt.want, Join(tt.array, tt.seperator))
		})
	}
}
