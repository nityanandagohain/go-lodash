package go_lodash

type DifferenceType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Difference[V DifferenceType](array []V, values []V) []V {
	result := make([]V, 0)
	lookup := map[V]struct{}{}

	// create a lookup
	for _, val := range values {
		lookup[val] = struct{}{}
	}

	for _, val := range array {
		if _, ok := lookup[val]; !ok {
			result = append(result, val)
		}
	}
	return result
}

type DifferenceFunc[V DifferenceType, X DifferenceType] func(value V) X

func DifferenceBy[V DifferenceType, X DifferenceType](array []V, values []V, differenceByFunc DifferenceFunc[V, X]) []V {
	result := make([]V, 0)
	lookup := map[X]struct{}{}

	// create a lookup
	for _, val := range values {
		lookup[differenceByFunc(val)] = struct{}{}
	}

	for _, val := range array {
		if _, ok := lookup[differenceByFunc(val)]; !ok {
			result = append(result, val)
		}
	}

	return result
}
