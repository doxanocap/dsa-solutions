package lc_3885

import (
	"container/heap"
)

type Item struct {
	eventID  int
	priority int
	version  int
}

type MaxHeap []Item

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }

func (h MaxHeap) Less(i, j int) bool {
	if h[i].priority == h[j].priority {
		return h[i].eventID < h[j].eventID
	}
	return h[i].priority > h[j].priority
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type EventManager struct {
	events map[int]Item
	heap   MaxHeap
}

func Constructor(events [][]int) EventManager {
	emap := make(map[int]Item, len(events))

	items := make([]Item, 0, len(events))

	for _, e := range events {
		i := Item{
			eventID:  e[0],
			priority: e[1],
			version:  1,
		}
		items = append(items, i)
		emap[e[0]] = i
	}

	h := MaxHeap(items)
	heap.Init(&h)

	return EventManager{
		events: emap,
		heap:   h,
	}
}

func (this *EventManager) UpdatePriority(eventId int, newPriority int) {
	prev := this.events[eventId]

	item := Item{
		eventID:  eventId,
		priority: newPriority,
		version:  prev.version + 1,
	}

	this.events[eventId] = item
	heap.Push(&this.heap, item)
}

func (this *EventManager) PollHighest() int {

	for this.heap.Len() > 0 {
		item := heap.Pop(&this.heap).(Item)
		last := this.events[item.eventID]

		if item.version == last.version {
			return item.eventID
		}
	}

	return -1
}

/**
 * Your EventManager object will be instantiated and called as such:
 * obj := Constructor(events);
 * obj.UpdatePriority(eventId,newPriority);
 * param_2 := obj.PollHighest();
 */
