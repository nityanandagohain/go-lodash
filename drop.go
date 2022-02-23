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

type WhileFunc[V DropType] func(value V, index int, array []V) bool

func DropWhileFunc[V DropType](array []V, whilefunc WhileFunc[V]) []V {
	for i, v := range array {
		if !whilefunc(v, i, array) {
			return array[i:]
		}
	}
	return []V{}
}

func DropWhileRightFunc[V DropType](array []V, whilefunc WhileFunc[V]) []V {
	for i := len(array) - 1; i >= 0; i-- {
		if !whilefunc(array[i], i, array) {
			return array[:i+1]
		}
	}
	return []V{}
}
