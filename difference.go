package go_lodash

type DifferenceType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Difference[V DifferenceType](array []V, values []V) []V {

	result := []V{}

	lookup := map[V]bool{}

	// create a lookup
	for _, val := range values {
		lookup[val] = true
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

	result := []V{}

	lookup := map[X]bool{}

	// create a lookup
	for _, val := range values {
		lookup[differenceByFunc(val)] = true
	}

	for _, val := range array {
		if _, ok := lookup[differenceByFunc(val)]; !ok {
			result = append(result, val)
		}
	}

	return result
}
