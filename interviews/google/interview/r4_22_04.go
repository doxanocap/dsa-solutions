package interview

import "container/heap"

/*
 * Implement an ad server with two APIs, InsertAd(content, score) that inserts
 * an ad to the set of available ads, and GetAd() that returns the best ad among
 * them. These are the rules:
 *
 * + 1. An ad is better if it has a higher score.
 * + 2. Do not return the same ad consecutively.
 * + 3. Once served, the ad's score is reduced by 1.
 */

insert
- priority queue
- updatable

getad
- pop
- store
- put back

type MaxHeap []Add

func (h *MaxHeap) Less(a,b Add) bool {
	return a.content > b.content
}

type Add struct {
	content string
	score int
}

type server struct {
	stack []Add
	heap  MaxHeap
}

func NewServer() *server {
	// ...
}

InsertAd("A",10)
InsertAd("B",9)
InsertAd("C",0)

// heap = ["A","B","C"]
// stack = []
GetAd() > "A"
// heap = ["B","C"]
// stack = ["A"]

GetAd() > "B"
// stack = ["A","B"]
// heap = ["C"]

// stack = ["B"]
// heap = ["A", "C"]

// O(log(n))
func (s *server) InsertAd(a Add) {
	heap.Push(s.heap, a)
}

// O(log(n))
func (s *server) GetAd() *Add {
	length := heap.Len(s.heap)

	defer func() {
		if len(s.stack) > 0 {
			heap.Push(s.heap, s.stack[0])
			s.stack = s.stack[1:]
		}
	}

	if length == 0 {
		return nil
	}

	item := heap.Pop(s.heap).(Add) // O(log(n))

	// put it into stack
	updItem := item
	updItem.score = updItem.score - 1
	s.stack = append(s.stack, updItem) // O(1)

	return &item
}














