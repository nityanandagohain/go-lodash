package go_lodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var chunkTests = []struct {
	name  string
	array []int
	size  int
	want  [][]int
}{
	{"even chunk", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
	{"odd chunk", []int{1, 2, 3, 4}, 3, [][]int{{1, 2, 3}, {4}}},
	{"odd chunk", []int{1}, 3, [][]int{{1}}},
	{"odd chunk", []int{}, 3, [][]int{}},
}

func TestChunk(t *testing.T) {
	is := assert.New(t)
	for _, tt := range chunkTests {
		t.Run(tt.name, func(t *testing.T) {
			res := Chunk(tt.array, tt.size)
			is.Equal(tt.want, res)
		})
	}
}

func TestChunkMap(t *testing.T) {
	is := assert.New(t)
	input := []map[string]interface{}{
		{"a": "b"},
		{"c": "d"},
		{"e": "f"},
	}

	expected := [][]map[string]interface{}{
		{
			{"a": "b"},
			{"c": "d"},
		},
		{
			{"e": "f"},
		},
	}

	res := Chunk(input, 2)
	is.Equal(expected, res)

}
