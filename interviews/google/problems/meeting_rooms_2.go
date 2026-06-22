package problems

import "sort"

func minChairs(s, e []int) int {
	sort.Ints(s)
	sort.Ints(e)

	res, count := 0, 0
	i, j := 0, 0

	for i < len(s) && j < len(e) {
		if s[i] < e[j] {
			count++
			res = max(res, count)
			i++
		} else {
			count--
			j++
		}
	}
	return res
}
