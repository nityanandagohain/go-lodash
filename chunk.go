package go_lodash

type ChunkType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | map[string]interface{}
}

func Chunk[V ChunkType](array []V, size int) [][]V {
	res := make([][]V, 0)
	start := 0
	length := len(array)

	for start < length {
		end := start + size
		if end > length {
			end = length
		}
		res = append(res, array[start:end])
		start = end
	}

	return res
}
