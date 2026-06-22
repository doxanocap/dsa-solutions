package lc_3691

import (
	"container/heap"
	"math"
)

type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SegmentTree struct {
	isMax  bool
	defVal int
	n      int
	tree   []int
}

func NewSegmentTree(a []int, isMax bool) *SegmentTree {
	st := &SegmentTree{
		n:      len(a),
		isMax:  isMax,
		defVal: 0,
		tree:   make([]int, 4*len(a)),
	}

	if !isMax {
		st.defVal = math.MaxInt32
	}

	st.build(a, 1, 0, st.n-1)
	return st
}

func (st *SegmentTree) build(arr []int, idx int, l, r int) {
	if l == r {
		st.tree[idx] = arr[l]
		return
	}

	m := (l + r) / 2
	st.build(arr, 2*idx, l, m)
	st.build(arr, 2*idx+1, m+1, r)

	lv, rv := st.tree[2*idx], st.tree[2*idx+1]
	if st.isMax {
		st.tree[idx] = max(lv, rv)
		return
	}
	st.tree[idx] = min(lv, rv)
}

func (st *SegmentTree) Query(l, r int) int {
	return st.query(0, st.n-1, 1, l, r)
}

func (st *SegmentTree) query(tL, tR, idx, l, r int) int {
	if tR < l || r < tL {
		return st.defVal
	}
	if l <= tL && tR <= r {
		return st.tree[idx]
	}

	m := (tL + tR) / 2
	lv := st.query(tL, m, 2*idx, l, r)
	rv := st.query(m+1, tR, 2*idx+1, l, r)
	if st.isMax {
		return max(lv, rv)
	}
	return min(lv, rv)
}

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)

	maxst := NewSegmentTree(nums, true)
	minst := NewSegmentTree(nums, false)

	l, r := 0, n-1
	seen := map[[2]int]bool{[2]int{l, r}: true}

	h := &MaxHeap{}
	heap.Init(h)

	v := maxst.Query(l, r) - minst.Query(l, r)
	heap.Push(h, []int{v, l, r})

	ans := int64(0)
	for ; k > 0 && h.Len() > 0; k-- {
		s := heap.Pop(h).([]int)
		v, l, r = s[0], s[1], s[2]
		ans += int64(v)

		for _, p := range [][2]int{{l + 1, r}, {l, r - 1}} {
			if p[0] >= p[1] || p[0] < 0 || p[1] > n-1 || seen[p] {
				continue
			}

			v = maxst.Query(p[0], p[1]) - minst.Query(p[0], p[1])
			heap.Push(h, []int{v, p[0], p[1]})
			seen[p] = true
		}
	}

	return ans
}
