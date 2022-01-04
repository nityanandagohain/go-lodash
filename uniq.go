package go_lodash

type UniqType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Uniq[V UniqType](array []V) []V {
	lookup := map[V]bool{}
	for i := 0; i < len(array); i++ {
		if _, ok := lookup[array[i]]; !ok {
			lookup[array[i]] = true
		}
	}

	result := make([]V, 0, len(lookup))
	for key, _ := range lookup {
		result = append(result, key)
	}

	return result
}

type UniqByFunc[V UniqType, X UniqType] func(value V, index int, array []V) X

func UniqBy[V UniqType, X UniqType](array []V, uniqFunc UniqByFunc[V, X]) []V {
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
