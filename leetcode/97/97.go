package _7

// isInterleave https://leetcode.com/problems/interleaving-string
func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	k := len(s3)

	if n+m != k {
		return false
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true

	for l := 1; l <= k; l++ {
		for i := 0; i <= l-1; i++ {
			j := l - 1 - i
			if i <= n && j <= m && dp[i][j] {
				if i < n && s1[i] == s3[l-1] {
					dp[i+1][j] = true
				}
				if j < m && s2[j] == s3[l-1] {
					dp[i][j+1] = true
				}
			}
		}
	}

	return dp[n][m]
}
