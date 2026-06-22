package main

import "sort"

func update(tree, set []int, l, r int, idx, num int) {
	if l < 0 || r >= len(set) || l > r {
		return
	}
	if set[l] <= num && num <= set[r] {
		tree[idx] += 1
	} else {
		return
	}
	mid := (l + r) / 2
	update(tree, set, l, mid-1, 2*idx+1, num)
	update(tree, set, mid, r, 2*idx+2, num)
}

func query(tree, set []int, l, r int, idx, num int) int {
	if l < 0 || r >= len(set) || l > r {
		return 0
	}
	ans := 0
	if num < set[l] {
		return ans
	}
	if set[r] <= num {
		ans += tree[idx]
		return ans
	}
	mid := (l + r) / 2

	return query(tree, set, l, mid-1, 2*idx+1, num) +
		query(tree, set, mid, r, 2*idx+2, num)
}

func countSmaller(nums []int) []int {
	n := len(nums)
	set := make([]int, 0, n)
	m := map[int]int{}
	for i, v := range nums {
		if _, ok := m[v]; !ok {
			set = append(set, v)
		}
		m[v] = i
	}

	tree := make([]int, n*4)
	sort.Ints(set)

	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[n-i-1] = query(tree, set, 0, len(set), 0, nums[i]-1)
		update(tree, set, 0, len(set), 0, nums[i])
	}

	return res
}
