package go_lodash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dropTests = []struct {
	name  string
	array []int
	n     int
	want  []int
}{
	{"drop", []int{1, 2}, 1, []int{2}},
	{"drop 1", []int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5}},
	{"exceed drop", []int{1, 2, 3}, 5, []int{}},
}

func TestDrop(t *testing.T) {
	is := assert.New(t)
	for _, tt := range dropTests {
		t.Run(tt.name, func(t *testing.T) {
			res := Drop(tt.array, tt.n)
			is.Equal(tt.want, res)
		})
	}
}

var dropRightTests = []struct {
	name  string
	array []int
	n     int
	want  []int
}{
	{"drop", []int{1, 2}, 1, []int{1}},
	{"drop 1", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3}},
	{"exceed drop", []int{1, 2, 3}, 5, []int{}},
	{"none drop", []int{1, 2}, 0, []int{1, 2}},
}

func TestDropRight(t *testing.T) {
	is := assert.New(t)
	for _, tt := range dropRightTests {
		t.Run(tt.name, func(t *testing.T) {
			res := DropRight(tt.array, tt.n)
			is.Equal(tt.want, res)
		})
	}
}

func TestDropWhileFunc(t *testing.T) {
	is := assert.New(t)
	res := DropWhileFunc([]int{1, 3, 4}, func(x int, i int, array []int) bool {
		if x <= 3 {
			return true
		}
		return false
	})

	is.Equal([]int{4}, res)
}

func TestDropWhileObject(t *testing.T) {
	is := assert.New(t)
	input := []map[string]interface{}{
		{"user": "barney", "active": false},
		{"user": "fred", "active": false},
		{"user": "pebbles", "active": true},
	}
	res := DropWhileFunc(input, func(x map[string]interface{}, i int, array []map[string]interface{}) bool {
		val := x["active"].(bool)
		return !val
	})
	is.Equal("pebbles", res[0]["user"].(string))
}

func TestDropRightWhileFunc(t *testing.T) {
	is := assert.New(t)
	res := DropWhileRightFunc([]int{1, 3, 4}, func(x int, i int, array []int) bool {
		if x >= 3 {
			return true
		}
		return false
	})

	is.Equal([]int{1}, res)
}

func TestDropRightWhileObject(t *testing.T) {
	is := assert.New(t)
	input := []map[string]interface{}{
		{"user": "barney", "active": false},
		{"user": "fred", "active": false},
		{"user": "pebbles", "active": true},
	}
	res := DropWhileRightFunc(input, func(x map[string]interface{}, i int, array []map[string]interface{}) bool {
		fmt.Println(x["user"])
		val := x["active"].(bool)
		return val
	})
	is.Equal("fred", res[1]["user"].(string))
}
