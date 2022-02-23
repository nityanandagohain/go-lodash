package go_lodash

type FillType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type FillOptions struct {
	start *int
	end   *int
}

// Fill elements of array with value from start up to, but not including, end.
func Fill[V FillType](array []V, value V, options *FillOptions) ([]V, error) {
	start, end := 0, len(array)
	if options != nil {
		if options.start != nil {
			start = *options.start
		}

		if options.end != nil {
			end = *options.end
		}
	}

	if start < 0 {
		start = 0
	}
	if end > len(array) {
		end = len(array)
	}

	for i := start; i < end; i++ {
		array[i] = value
	}
	return array, nil
}
