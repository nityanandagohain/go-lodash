package go_lodash

import (
	"fmt"
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

func TestUniqBy(t *testing.T) {
	is := assert.New(t)
	input := []float32{2.1, 1.2, 2.3}
	res := UniqBy(input, func(value float32, index int, array []float32) int {
		fmt.Println(int(value))
		return int(value)
	})
	is.ElementsMatch([]float32{2.1, 1.2}, res)
}

func BenchmarkUniq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range UniqTests {
			Uniq(tt.array)
		}
	}
}

func BenchmarkUniq1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range UniqTests {
			Uniq1(tt.array)
		}
	}
}

func BenchmarkUniqBy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range UniqTests {
			UniqBy(tt.array, func(value int, index int, array []int) int {
				return value
			})
		}
	}
}

func BenchmarkUniqBy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range UniqTests {
			UniqBy1(tt.array, func(value int, index int, array []int) int {
				return value
			})
		}
	}
}

func Uniq1[V UniqType](array []V) []V {
	lookup := map[V]bool{}
	for i := 0; i < len(array); i++ {
		if _, ok := lookup[array[i]]; !ok {
			lookup[array[i]] = true
		}
	}

	result := make([]V, 0, len(lookup))
	for key := range lookup {
		result = append(result, key)
	}

	return result
}

func UniqBy1[V UniqType, X UniqType](array []V, uniqFunc UniqByFunc[V, X]) []V {
	lookup := map[X]bool{}
	result := make([]V, 0, len(lookup))
	for i := 0; i < len(array); i++ {
		val := uniqFunc(array[i], i, array)
		if _, ok := lookup[val]; !ok {
			lookup[val] = true
			result = append(result, array[i])
		}
	}
	return result
}
