package go_lodash

type ReverseType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Reverse[V ReverseType](array []V) []V {
	for i := len(array)/2 - 1; i >= 0; i-- {
		opp := len(array) - 1 - i
		array[i], array[opp] = array[opp], array[i]
	}

	return array
}
