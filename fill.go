package go_lodash

import "errors"

type FillType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type FillOptions struct {
	start *int
	end   *int
}

// Fills elements of array with value from start up to, but not including, end.
func Fill[V FillType](array []V, value V, options *FillOptions) ([]V, error) {
	start, end := 0, len(array)

	if options != nil {
		if options.start != nil {
			start = *options.start
			if start < 0 {
				return nil, errors.New("start is out of index")
			}
		}

		if options.end != nil {
			if *options.end > end {
				return nil, errors.New("end is out of index")
			}
			end = *options.end
		}
	}

	for ; start < end; start++ {
		array[start] = value
	}
	return array, nil
}
