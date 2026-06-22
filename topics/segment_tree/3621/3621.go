package _621

import (
	"math/bits"
)

type SegmentTree struct {
	n    int
	tree []int
}

func (t *SegmentTree) Update(idx, delta int) {
	t.update(0, t.n-1, 1, idx, delta)
}

func (t *SegmentTree) Query(l, r int) int {
	return t.query(0, t.n-1, 1, l, r)
}

func (t *SegmentTree) update(tl, tr, tIdx, idx, delta int) {
	if tl == tr {
		t.tree[tIdx] += delta
		return
	}

	tm := (tl + tr) / 2
	if idx <= tm {
		t.update(tl, tm, 2*tIdx, idx, delta)
	} else {
		t.update(tm+1, tr, 2*tIdx+1, idx, delta)
	}
	t.tree[tIdx] = t.tree[2*tIdx] + t.tree[2*tIdx+1]
}

func (t *SegmentTree) query(tl, tr, tIdx, ql, qr int) int {
	if ql > qr || tr < ql || qr < tl {
		return 0
	}
	if ql <= tl && tr <= qr {
		return t.tree[tIdx]
	}

	tm := (tl + tr) / 2
	left := t.query(tl, tm, 2*tIdx, ql, qr)
	right := t.query(tm+1, tr, 2*tIdx+1, ql, qr)
	return left + right
}

func popcountDepth(nums []int64, queries [][]int64) []int {
	n := len(nums)

	computed := map[int64]int{}
	sts := make([]SegmentTree, 6)

	for i := 0; i < 6; i++ {
		sts[i] = SegmentTree{
			n:    n,
			tree: make([]int, 4*n),
		}
	}

	for idx, val := range nums {
		pd, ok := computed[val]
		if !ok {
			pd = popDepth(val)
			computed[val] = pd
		}
		sts[pd].Update(idx, 1)
	}

	ans := make([]int, 0, n)

	for _, q := range queries {
		if q[0] == 2 {
			idx, val := int(q[1]), q[2]

			prevVal := nums[idx]
			pd := computed[prevVal]
			sts[pd].Update(idx, -1)

			pd = popDepth(val)
			computed[val] = pd
			nums[idx] = val

			sts[pd].Update(idx, 1)
			continue
		}

		l, r, pd := int(q[1]), int(q[2]), q[3]
		c := sts[pd].Query(l, r)

		ans = append(ans, c)
	}

	return ans
}

func popDepth(x int64) int {
	depth := 0
	for x > 1 {
		x = int64(bits.OnesCount64(uint64(x)))
		depth++
	}
	return depth
}
