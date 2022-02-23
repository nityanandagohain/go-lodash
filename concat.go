package go_lodash

type ConcatType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | map[string]interface{}
}

// Concat Creates a new array concatenating array with any additional arrays and/or values.
func Concat[V ConcatType](arrays ...[]V) []V {
	array := make([]V, 0)
	for _, arr := range arrays {
		array = append(array, arr...)
	}
	return array
}
