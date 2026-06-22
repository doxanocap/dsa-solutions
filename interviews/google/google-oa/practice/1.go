package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return item
}

func solution(A []int) int {
	n := len(A)
	ans := 0

	h := &MaxHeap{}
	heap.Init(h)

	used := make([]bool, n)

	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}

		h := &MaxHeap{}
		heap.Init(h)

		for j := i + 1; j < n; j++ {
			if used[j] {
				continue
			}
			heap.Push(h, []int{A[j], j})
		}

		lastIdx := -1

		for h.Len() > 0 {
			node := heap.Pop(h).([]int)
			if node[0] >= A[i] {
				continue
			}
			if node[1] < lastIdx {
				continue
			}

			used[node[1]] = true
			lastIdx = node[1]
		}

		used[i] = true
		ans++
	}

	return ans
}

func main() {
	// Read from stdin, solve the problem, and write the answer to stdout.
	var AS string
	fmt.Scan(&AS)
	A := toIntArray(AS)
	fmt.Print(solution(A))
}

func toIntArray(str string) []int {
	a := strings.Split(str, ",")
	b := make([]int, len(a))
	for i, v := range a {
		var err error
		b[i], err = strconv.Atoi(v)
		if err != nil {
		}
	}
	return b
}
