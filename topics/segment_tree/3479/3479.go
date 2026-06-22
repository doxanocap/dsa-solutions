package lc3479

type SegmentTree struct {
	n    int
	tree []int
}

func NewSegmentTree(basket []int) *SegmentTree {
	n := len(basket)
	st := &SegmentTree{
		n:    n,
		tree: make([]int, 4*n),
	}
	st.build(basket, 1, 0, n-1)
	return st
}

func (t *SegmentTree) build(arr []int, node, l, r int) {
	if l == r {
		t.tree[node] = arr[l]
		return
	}
	m := (l + r) / 2
	t.build(arr, 2*node, l, m)
	t.build(arr, 2*node+1, m+1, r)
	t.tree[node] = max(t.tree[2*node], t.tree[2*node+1])
}

func (st *SegmentTree) query(node, l, r, val int) int {
	if st.tree[node] < val {
		return -1
	}
	if l == r {
		st.tree[node] = 0
		return l
	}
	mid := (l + r) / 2
	res := -1
	if st.tree[2*node] >= val {
		res = st.query(2*node, l, mid, val)
	} else {
		res = st.query(2*node+1, mid+1, r, val)
	}
	st.tree[node] = max(st.tree[2*node], st.tree[2*node+1])
	return res
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	st := NewSegmentTree(baskets)
	unplaced := 0
	for _, f := range fruits {
		idx := st.query(1, 0, len(baskets)-1, f)
		if idx == -1 {
			unplaced++
		}
	}
	return unplaced
}
