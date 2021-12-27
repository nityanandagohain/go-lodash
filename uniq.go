package go_lodash

type UniqType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Uniq[V LastIndexOfType](array []V) []V {
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
