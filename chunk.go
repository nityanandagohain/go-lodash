package go_lodash

type ChunkType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | map[string]interface{}
}

func Chunk[V ChunkType](array []V, size int) [][]V {
	res := [][]V{}
	start := 0
	length := len(array)
	for end := 1; end < length; end++ {
		if ((end - start) % size) == 0 {
			res = append(res, array[start:end])
			start = end
		}
	}

	if start < length {
		res = append(res, array[start:])
	}
	return res
}
