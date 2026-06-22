package lc_3806

import (
	"math/bits"
	"sort"
)

func maximumAND(nums []int, k int, m int) int {
	ans := 0
	n := len(nums)

	for i := 29; i >= 0; i-- {
		candidate := ans | (1 << i)

		costs := make([]int, n)

		for j, x := range nums {
			if (x & candidate) == candidate {
				costs[j] = 0
				continue
			}

			missing := candidate & ^x

			msb := bits.Len(uint(missing)) - 1
			mid := 1 << msb

			prefix := candidate & ^((1 << (msb + 1)) - 1)
			suffix := candidate & ((1 << msb) - 1)

			target := prefix | mid | suffix

			costs[j] = target - x
		}

		sort.Ints(costs)

		var cost int
		for j := 0; j < m; j++ {
			cost += costs[j]
		}

		if cost <= k {
			ans = candidate
		}

	}

	return ans
}
