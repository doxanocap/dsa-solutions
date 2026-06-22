package mocks

import (
	"container/heap"
	"math"
)

type Item struct {
	node int
	cost int
	disc int
}

type PriorityQueue []Item

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumCost(n int, highways [][]int, discounts int) int {
	adj := make([][][2]int, n)
	for _, h := range highways {
		u, v, w := h[0], h[1], h[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// dist[node][discounts_used]
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, discounts+1)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt64
		}
	}

	pq := &PriorityQueue{}
	heap.Push(pq, Item{0, 0, 0})
	dist[0][0] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Item)

		if curr.cost > dist[curr.node][curr.disc] {
			continue
		}
		if curr.node == n-1 {
			return curr.cost
		}

		for _, edge := range adj[curr.node] {
			next, weight := edge[0], edge[1]

			// Вариант 1: Едем без скидки
			if dist[next][curr.disc] > curr.cost+weight {
				dist[next][curr.disc] = curr.cost + weight
				heap.Push(pq, Item{dist[next][curr.disc], next, curr.disc})
			}

			// Вариант 2: Едем со скидкой (если они остались)
			if curr.disc < discounts {
				discountCost := curr.cost + (weight / 2)
				if dist[next][curr.disc+1] > discountCost {
					dist[next][curr.disc+1] = discountCost
					heap.Push(pq, Item{discountCost, next, curr.disc + 1})
				}
			}
		}
	}

	return -1
}

func generateDraw(n int) []int {

}

func solve(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])

	dirs := [][]int{
		{+1, 0},
		{0, +1},
		{-1, 0},
		{0, -1},
	}

	queue := [][]int{{0, 0, k, 0}}

	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		i, j, curr_k, steps := curr[0], curr[1], curr[2], curr[3]

		if i == n-1 && j == m-1 {
			return steps
		}

		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]

			if ni < 0 || ni >= n || nj < 0 || nj >= m {
				continue
			}

			nk := curr_k - grid[ni][nj]

			if nk >= 0 && visited[ni][nj] < nk {
				visited[ni][nj] = nk
				queue = append(queue, []int{ni, nj, nk, steps + 1})
			}

			if grid[ni][nj] == 1 {

			}
		}
	}

	return ans
}
