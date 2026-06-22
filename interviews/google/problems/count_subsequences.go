package problems

import (
	"slices"
	"sort"
)

/*
2.Find the count of subsequences of the array containing positive integers in which each number is
divisible by its position in the subsequence. A subsequence is a sequence (not necessarily consecutive) derived from an original sequence by deleting
zero or more elements without changing the order of the remaining elements.

input:
a []int

output:
int

test case:

a = [4 9 3 6 5 9]

4 6 9
9 6 9
3 6 9

6 -> 1, 2, 3, 6
4 -> 1, 2, 4
9 -> 1, 3, 9

dividers < O(sqrt(N))

for i in a:
dp[div_i] += dp[div_i - 1]

[1 2 3 4 5]

seq number or
divider 	1	2
count		5	2 -> 5 * 2 combs

*/

func countSubsequences(nums []int) int64 {
	n := len(nums)
	copy
	ans := int64(0)

	dp := make([]int, n+1)
	dp[0] = 1

	for _, num := range nums {

		divs := make([]int, 0)
		for d := 1; d*d <= num; d++ {
			if num%d == 0 {
				divs = append(divs, d)

				if d*d != num && num/d <= n {
					divs = append(divs, num/d)
				}
			}
		}

		sort.Ints(divs)
		slices.Reverse(divs)

		for _, d := range divs {
			dp[d] += dp[d-1]
		}
	}

	for _, v := range dp {
		ans += int64(v)
	}

	return ans - 1
}
