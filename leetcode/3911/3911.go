package _911

import (
	"fmt"
	"slices"
	"sort"
)

func kthRemainingInteger(nums []int, queries [][]int) []int {
	const maxr = int(2e9)

	n := len(nums)
	ans := make([]int, n)

	count := make([]int, n+1)
	for i := 0; i < n; i++ {
		count[i+1] = count[i]
		if nums[i]%2 == 0 {
			count[i+1]++
		}
	}

	query := func(l, r, x int) int {
		evens := x / 2

		idx := sort.Search(r-l+1, func(i int) bool {
			return nums[l+i] > x
		})

		removed := count[idx] - count[l]

		return evens - removed
	}

	for i, q := range queries {
		l, r, k := q[0], q[1], q[2]

		low, high := 1, maxr
		res := high

		for low <= high {
			mid := low + (high-low)/2
			if query(l, r, mid) >= k {
				res = mid
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		if res%2 != 0 {
			res--
		}
		ans[i] = res
	}

	return ans
}

func sortVowels(s string) string {
	n := len(s)
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}

	order := make([][]int, 0, 5) // byte,cnt

	for i := 0; i < n; i++ {
		if !slices.Contains(vowels, s[i]) {
			continue
		}

		intv := int(s[i] - 'a')

		idx := len(order)
		for i, v := range order {
			if v[0] == intv {
				idx = i
				break
			}
		}

		if idx == len(order) {
			order = append(order, []int{intv, 0})
		}

		order[idx][1] += 1
	}

	nextMax := func() (int, int) {
		maxb, maxcnt := 0, 0

		for _, v := range order {
			if v[1] > maxcnt {
				maxcnt = v[1]
				maxb = v[0]
			}
		}
		return maxb, maxcnt
	}

	sb := []byte(s)
	b, cnt := nextMax()

	for i := 0; i < n; i++ {
		if !slices.Contains(vowels, s[i]) {
			continue
		}
		if cnt == 0 {
			b, cnt = nextMax()
		}

		sb[i] = byte(b + 'a')
	}

	fmt.Println(order)

	return string(sb)
}
