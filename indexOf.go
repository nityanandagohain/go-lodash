package go_lodash

type IndexOfType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | bool
}

type IndexOfOptions struct {
	FromIndex *int
}

func IndexOf[V IndexOfType](array []V, value V, options *IndexOfOptions) int {
	fromIndex := 0
	if options != nil && options.FromIndex != nil {
		fromIndex = *options.FromIndex
	}

	if fromIndex < 0 {
		fromIndex = 0
	}

	for i := fromIndex; i < len(array); i++ {
		if array[i] == value {
			return i
		}
	}

	return -1
}
