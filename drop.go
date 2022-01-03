package go_lodash

type DropType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | map[string]interface{}
}

func Drop[V DropType](array []V, n int) []V {
	if n > len(array) {
		return []V{}
	}
	return array[n:]
}

func DropRight[V DropType](array []V, n int) []V {
	if n > len(array) {
		return []V{}
	}
	return array[:len(array)-n]
}

type whileFunc[V DropType] func(value V, index int, array []V) bool

func DropWhileFunc[V DropType](array []V, whilefunc whileFunc[V]) []V {
	till := 0
	for index, val := range array {
		if !whilefunc(val, index, array) {
			till = index
			break
		}
	}
	return array[till:]
}

func DropWhileRightFunc[V DropType](array []V, whilefunc whileFunc[V]) []V {
	from := 0
	index := len(array) - 1
	for ; index >= 0; index-- {
		val := array[index]
		if !whilefunc(val, index, array) {
			from = index
			break
		}
	}
	return array[:from+1]
}
