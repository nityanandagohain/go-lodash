package go_lodash

type LastIndexOfType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type LastIndexOfOptions struct {
	FromIndex *int
}

func LastIndexOf[V LastIndexOfType](array []V, value V, options *LastIndexOfOptions) int {
	fromIndex := 0
	if options != nil && options.FromIndex != nil {
		fromIndex = *options.FromIndex
	}
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == value {
			return i - fromIndex
		}
	}
	return -1
}
