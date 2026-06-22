package leetcode

import (
	"slices"
	"sort"
)

func maxFixedPoints(nums []int) int {
	n := len(nums)
	ee := make([][2]int, 0, n)

	for i, v := range nums {
		ee = append(ee, [2]int{
			v,
			i - v,
		})
	}

	slices.SortFunc(ee, func(a, b [2]int) int {
		if a[0] == b[0] {
			return b[1] - a[1]
		}
		return a[0] - b[0]
	})

	pairs := make([]int, 0, n)

	for _, e := range ee {

		idx := sort.Search(len(pairs), func(i int) bool {
			return pairs[i] > e[1]
		})

		if idx == len(pairs) {
			pairs = append(pairs, e[1])
		} else {
			pairs[idx] = e[1]
		}
	}

	return len(pairs)
}
