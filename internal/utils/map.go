package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// SortMapPair sort map pair
type SortMapPairT struct {
	Key string
	Val any
}

// SortOrderT sort order mode(asc/desc) type
type SortOrderT int

const (
	// SortOrderAsc sort order asc
	SortOrderAsc SortOrderT = 1
	// SortOrderDesc sort order desc
	SortOrderDesc SortOrderT = -1
)

// SortMapByKey sort map by key
func SortMapByKey[T constraints.Ordered](m map[string]T, sortord SortOrderT) []SortMapPairT {
	pairs := []SortMapPairT{}
	for k, v := range m {
		pairs = append(pairs, SortMapPairT{k, v})
	}

	if sortord == SortOrderDesc {
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].Key > pairs[j].Key
		})
	} else if sortord == SortOrderAsc {
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].Key < pairs[j].Key
		})
	}

	return pairs
}

// SortMapByKeyToKeys sort map by key to keys
//
//	For example:
//	m := map[string]any{
//		"a": 1,
//		"c": 3,
//		"b": 2,
//	}
//	keys := SortMapByKeyToKeys(m, SortOrderAsc)
//	fmt.Println(keys)
//	Output: [a b c]
func SortMapByKeyToKeys[T constraints.Ordered](m map[string]T, sortord SortOrderT) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	if sortord == SortOrderDesc {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else if sortord == SortOrderAsc {
		sort.Strings(keys)
	}

	return keys
}

// SortMapByValueToKeys sort map by value to keys
//
//		For example:
//		m := map[string]int{
//	 	"a": 1,
//	 	"v": 3,
//	 	"p": 2,
//		}
//	 keys := SortMapByValueToKeys(m, SortOrderAsc)
//	 fmt.Println(keys)
//	 Output: [a p v]
func SortMapByValueToKeys[T constraints.Ordered](m map[string]T, sortord SortOrderT) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	if sortord == SortOrderDesc {
		sort.SliceStable(keys, func(i, j int) bool {
			return m[keys[i]] > m[keys[j]]
		})
	} else if sortord == SortOrderAsc {
		sort.SliceStable(keys, func(i, j int) bool {
			return m[keys[i]] < m[keys[j]]
		})
	}

	return keys
}
