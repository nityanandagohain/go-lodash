package go_lodash

type ReverseType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Reverse[V LastIndexOfType](array []V) []V {
	var temp V
	i, j := 0, len(array)-1
	for ; i < (len(array)-1)/2; i++ {
		temp = array[j]
		array[j] = array[i]
		array[i] = temp
		j--
	}
	return array
}
