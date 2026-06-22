package main

func solution(S int) int {
	ans := 0
	count := make([]int, 19)
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			count[i+j]++
		}
	}

	for l := 0; l <= S; l++ {
		r := S - l
		if l > 18 || r > 18 {
			continue
		}
		ans += count[l] * count[r]
	}

	return ans
}

func solution1(S int) int {
	k := 4
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, S+1)
	}
	dp[0][0] = 1

	for i := 1; i <= k; i++ {
		for s := 0; s <= S; s++ {
			for d := 0; d <= 9; d++ {
				if s >= d {
					dp[i][s] += dp[i-1][s-d]
				}
			}
		}
	}

	return dp[k][S]
}
