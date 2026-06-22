package leetcode_1509

import "sort"

func minDifference(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	if n <= 4 {
		return 0
	}

	diff := nums[n-1] - nums[0]
	for k := 0; k < 4; k++ {
		diff = min(diff, nums[n-4+k]-nums[k])
	}

	return diff
}
