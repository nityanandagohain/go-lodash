package go_lodash

type PointerType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func GetPointer[V JoinType](val V) *V {
	return &val
}
