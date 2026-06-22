package leetcode_1534

import (
	"container/heap"
	"math"
)

type Side struct {
	idx1 int
	idx2 int
	dist int
}

type MinHeap []Side

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Side))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return item
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := int(math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1])))
			heap.Push(h, Side{i, j, dist})
		}
	}

	visited := make(map[int]int, n)
	group := 1

	sum := 0
	edgesUsed := 0

	for edgesUsed < n-1 && h.Len() > 0 {
		side := heap.Pop(h).(Side)

		g1 := visited[side.idx1]
		g2 := visited[side.idx2]

		if g1 == 0 && g2 == 0 {
			group++
			visited[side.idx1] = group
			visited[side.idx2] = group
		} else if g1 == 0 {
			visited[side.idx1] = g2
		} else if g2 == 0 {
			visited[side.idx2] = g1
		} else if g1 == g2 {
			continue
		} else if g1 != g2 {
			for k, v := range visited {
				if v == g1 {
					visited[k] = g2
				}
			}
		}

		sum += side.dist
		edgesUsed++
	}

	return sum
}
