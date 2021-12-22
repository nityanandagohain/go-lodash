package go_lodash

import (
	"fmt"
	"strings"
)

type JoinType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Join[V JoinType](array []V, seperator string) string {
	result := ""
	for _, val := range array {
		result += seperator + fmt.Sprint(val)
	}
	return strings.Trim(result, seperator)
}
