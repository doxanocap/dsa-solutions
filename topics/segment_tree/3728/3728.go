package lc_3728

// LazySegmentTree — сегментное дерево с ленивыми обновлениями (lazy propagation).
type LazySegmentTree struct {
	n   int
	mx  []int // максимум на отрезке
	mn  []int // минимум на отрезке
	tag []int // отложенные (lazy) прибавления
	arr []int // исходный массив
}

// NewLazySegmentTree создаёт дерево.
func NewLazySegmentTree(n int, arr []int) *LazySegmentTree {
	tree := &LazySegmentTree{
		n:   n,
		mx:  make([]int, n*4),
		mn:  make([]int, n*4),
		tag: make([]int, n*4),
		arr: make([]int, n),
	}
	if arr != nil && len(arr) == n {
		copy(tree.arr, arr)
		tree.build(1, 0, n)
	}
	return tree
}

func (st *LazySegmentTree) build(pos, s, t int) {
	if s+1 == t {
		st.mx[pos] = st.arr[s]
		st.mn[pos] = st.arr[s]
		return
	}
	mid := (s + t) / 2
	st.build(pos*2, s, mid)
	st.build(pos*2+1, mid, t)
	st.merge(pos)
}

func (st *LazySegmentTree) merge(pos int) {
	left := pos * 2
	right := pos*2 + 1
	st.mx[pos] = max(st.mx[left], st.mx[right])
	st.mn[pos] = min(st.mn[left], st.mn[right])
}

func (st *LazySegmentTree) updatePos(pos, val int) {
	st.mx[pos] += val
	st.mn[pos] += val
	st.tag[pos] += val
}

func (st *LazySegmentTree) push(pos, s, t int) {
	if st.tag[pos] == 0 {
		return
	}

	left := pos * 2
	right := pos*2 + 1

	st.updatePos(left, st.tag[pos])
	st.updatePos(right, st.tag[pos])
	st.tag[pos] = 0
}

func (st *LazySegmentTree) query(pos, s, t, x int) int {
	if s+1 == t {
		return s
	}
	st.push(pos, s, t)
	mid := (s + t) / 2
	left := pos * 2
	right := pos*2 + 1

	if st.mn[left] <= x && x <= st.mx[left] {
		return st.query(left, s, mid, x)
	}
	return st.query(right, mid, t, x)
}

// modify выполняет диапазонное обновление [qs, qt).
func (st *LazySegmentTree) modify(pos, s, t, qs, qt, val int) {
	if qt <= s || t <= qs {
		return
	}
	if qs <= s && t <= qt {
		st.updatePos(pos, val)
		return
	}

	st.push(pos, s, t)
	mid := (s + t) / 2
	left := pos * 2
	right := pos*2 + 1

	st.modify(left, s, mid, qs, qt, val)
	st.modify(right, mid, t, qs, qt, val)
	st.merge(pos)
}

func longestBalanced(nums []int) int {
	n := len(nums) + 1
	last := make(map[int]int)
	seg := NewLazySegmentTree(n, nil)

	ans, cur := 0, 0

	for i := 1; i < n; i++ {
		v := nums[i-1]

		delta := 1
		if v%2 == 0 {
			delta = -1
		}

		if pi, ok := last[v]; ok {
			cur -= delta
			seg.modify(1, 0, n, pi, n, -delta)
		}

		cur += delta
		seg.modify(1, 0, n, i, n, delta)
		last[v] = i

		left := seg.query(1, 0, n, cur)
		ans = max(ans, i-left)
	}

	return ans
}
