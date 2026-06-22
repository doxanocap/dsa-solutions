package lc_3801

import (
	"math"
	"math/bits"
)

func minMergeCost(lists [][]int) int64 {
	n := len(lists)
	limit := 1 << n

	medians := make([]int, limit)
	lengths := make([]int, limit)
	allSorted := make([][]int, limit)

	for mask := 1; mask < limit; mask++ {
		lowbit := mask & -mask
		rest := mask ^ lowbit
		idx := bits.TrailingZeros(uint(lowbit))

		if rest == 0 {
			allSorted[mask] = make([]int, len(lists[idx]))
			copy(allSorted[mask], lists[idx])
		} else {
			allSorted[mask] = merge(allSorted[rest], lists[idx])
		}

		lengths[mask] = len(allSorted[mask])
		if lengths[mask] > 0 {
			medians[mask] = allSorted[mask][(lengths[mask]-1)/2]
		}
	}

	// DP Phase (O(3^N))
	dp := make([]int64, limit)
	for i := range dp {
		dp[i] = math.MaxInt64
	}

	for mask := 1; mask < limit; mask++ {
		if (mask & (mask - 1)) == 0 {
			dp[mask] = 0
			continue
		}

		for sub := (mask - 1) & mask; sub > 0; sub = (sub - 1) & mask {
			complement := mask ^ sub

			if sub > complement {
				continue
			}

			if dp[sub] == math.MaxInt64 || dp[complement] == math.MaxInt64 {
				continue
			}

			cost := int64(lengths[sub] + lengths[complement] + abs(medians[sub]-medians[complement]))
			total := dp[sub] + dp[complement] + cost

			if total < dp[mask] {
				dp[mask] = total
			}
		}
	}

	return dp[limit-1]
}

func merge(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
