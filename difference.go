package go_lodash

type DifferenceType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Difference[V IntersectionType](array []V, values []V) []V {

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
