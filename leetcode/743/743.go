package lc743

import (
	"container/heap"
	"math"
	"sort"
)

type Node struct {
	id   int
	dist int
}

type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := map[int][]Node{}
	h := &MinHeap{}
	heap.Init(h)

	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})
	heap.Push(h, Node{id: k, dist: 0})

	distances := make([]int, n+1)
	for i := range distances {
		distances[i] = math.MaxInt32
	}

	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], Node{
			id:   t[1],
			dist: t[2],
		})

	}

	for h.Len() > 0 {
		node := heap.Pop(h).(Node)

		if distances[node.id] < node.dist {
			continue
		}

		for _, leaf := range graph[node.id] {
			leaf.dist = node.dist + leaf.dist

			if leaf.dist < distances[leaf.id] {
				distances[leaf.id] = leaf.dist
				heap.Push(h, leaf)
			}
		}
	}

	res := 0
	for i := 1; i < len(distances); i++ {
		if distances[i] == math.MaxInt32 {
			return -1
		}
		res = max(res, distances[i])
	}
	return res
}
