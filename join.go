package go_lodash

import (
	"fmt"
)

type JoinType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Join[V JoinType](array []V, separator string) string {
	result := ""
	for i, v := range array {
		if i > 0 {
			result += separator
		}
		result += fmt.Sprintf("%v", v)
	}

	return result
}
