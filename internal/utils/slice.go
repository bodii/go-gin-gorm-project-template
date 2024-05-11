package utils

import (
	"math/rand"
	"time"
)

// RandSlice 随机打乱Slice，并从中抽取count个元素
func RandSlice[T any](origin []T, count int) []T {
	tmpOrigin := make([]T, len(origin))
	copy(tmpOrigin, origin)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(tmpOrigin), func(i int, j int) {
		tmpOrigin[i], tmpOrigin[j] = tmpOrigin[j], tmpOrigin[i]
	})

	result := make([]T, 0, count)
	for index, value := range tmpOrigin {
		if index == count {
			break
		}
		result = append(result, value)
	}

	return result
}
