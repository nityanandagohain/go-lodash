package go_lodash

type FindFunc[V any] func(value V, index int, array []V) bool

// This method is like _.find except that it returns the index of the first element predicate returns truthy for instead of the element itself.
// fromIndex default is 0 if nil
func FindIndex[V any](array []V, findFunc FindFunc[V], fromIndex *int) int {
	if fromIndex == nil {
		fromIndex = GetPointer(0)
	}
	for index, val := range array {
		if findFunc(val, index, array) {
			return index
		}
	}
	return -1
}
