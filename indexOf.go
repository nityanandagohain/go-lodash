package go_lodash

type IndexOfType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | bool
}

func IndexOf[V IndexOfType](array []V, value V) int {
	for index, val := range array {
		if value == val {
			return index
		}
	}
	return -1
}

func IndexOfFrom[V IndexOfType](array []V, value V, from int) int {
	for i := from; i < len(array); i++ {
		if array[i] == value {
			return i
		}
	}
	return -1
}
