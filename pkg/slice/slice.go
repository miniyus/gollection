package slice

import "math"

func Copy[T interface{}](s []T) []T {
	copyS := make([]T, len(s))
	copy(copyS, s)
	return copyS
}

func Map[T interface{}](s []T, fn func(v T, i int) T) []T {
	var mapped []T

	for i, v := range s {
		mapped = append(mapped, fn(v, i))
	}

	return mapped
}

func Filter[T interface{}](s []T, fn func(v T, i int) bool) []T {
	var filtered []T

	for i, v := range s {
		f := fn(v, i)
		if f {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func Except[T interface{}](s []T, fn func(v T, i int) bool) []T {
	var excepted []T

	for i, v := range s {
		f := fn(v, i)
		if !f {
			excepted = append(excepted, v)
		}
	}

	return excepted
}

func Chunk[T interface{}](s []T, chunkSize int, fn ...func(v []T, i int)) [][]T {
	chunkSlice := make([]T, 0)
	chunkedSlice := make([][]T, 0)
	chunkedSize := int(math.Ceil(float64(len(s) / chunkSize)))

	var callback func(v []T, i int)
	if len(fn) != 0 {
		callback = fn[0]
	}

	for i := 0; i < chunkedSize; i++ {
		if (i*chunkSize)+chunkSize <= (len(s) - 1) {
			chunkSlice = s[(i * chunkSize) : (i*chunkSize)+chunkSize]
		} else {
			chunkSlice = s[(i * chunkSize):]
		}

		callback(chunkSlice, i)

		chunkedSlice = append(chunkedSlice, chunkSlice)
	}

	return chunkedSlice
}

func For[T interface{}](s []T, fn func(v T, i int)) []T {
	for i, v := range s {
		fn(v, i)
	}

	return s
}

func Add[T interface{}](s []T, v T) []T {
	return append(s, v)
}

func Remove[T interface{}](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func Concat[T interface{}](s []T, s2 []T) []T {
	if len(s2) == 0 {
		return s
	}

	for _, v := range s2 {
		s = append(s, v)
	}

	return s
}

func Push[T interface{}](s []T, i T) []T {
	return Add(s, i)
}

func Pop[T interface{}](s []T) (out []T, pop T) {
	pop = s[len(s)-1]
	out = s[:len(s)-1]

	return out, pop
}

func First[T interface{}](s []T) T {
	return s[0]
}

func Last[T interface{}](s []T) T {
	return s[len(s)-1]
}

func Merge[T interface{}](s1 []T, s2 []T) []T {
	return append(s1, s2...)
}

func Clear[T interface{}](s []T) []T {
	s = make([]T, 0)

	return s
}

func Reverse[T interface{}](s []T) []T {
	reverse := make([]T, 0)
	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, s[i])
	}

	return reverse
}

func Slice[T interface{}](s []T, start int, end int) []T {
	return s[start:end]
}
