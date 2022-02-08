package go_lodash

import (
	"sort"
)

type UniqType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Uniq[V UniqType](array []V) []V {
	if len(array) == 0 {
		return array
	}

	out := make([]V, len(array))
	copy(out, array)

	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	result := out[:0]
	for i := 0; i < len(array); i++ {
		if len(result) == 0 || result[len(result)-1] != array[i] {
			result = append(result, array[i])
		}
	}

	return result
}

type UniqByFunc[V UniqType, X UniqType] func(value V, index int, array []V) X

func UniqBy[V UniqType, X UniqType](array []V, uniqFunc UniqByFunc[V, X]) []V {
	lookup := map[X]struct{}{}
	result := make([]V, 0, len(array))

	for i := 0; i < len(array); i++ {
		val := uniqFunc(array[i], i, array)
		if _, ok := lookup[val]; !ok {
			lookup[val] = struct{}{}
			result = append(result, array[i])
		}
	}
	return result
}
