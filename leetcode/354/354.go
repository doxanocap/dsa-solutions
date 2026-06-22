package _54

import (
	"fmt"
	"slices"
)

func maxEnvelopes(e [][]int) int {
	slices.SortFunc(e, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	for _, v := range e {
		fmt.Println(v)
	}

	n := len(e)
	dp := make([]int, n)

	for i := range e {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if e[j][0] < e[i][0] && e[j][1] < e[i][1] {
				dp[i] = max(
					dp[i],
					dp[j]+1,
				)
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	return ans
}
