package main

import (
	"container/heap"
	"slices"
)

func smallestChair(times [][]int, targetFriend int) int {
	events := make([]Event, 2*len(times))
	availableChairs := &IntHeap{}
	for i := range times {
		events[i] = Event{times[i][0], 1, i}
		events[len(times)+i] = Event{times[i][1], -1, i}
		availableChairs.Push(i)
	}
	heap.Init(availableChairs)
	slices.SortFunc(events, func(i, j Event) int {
		return i.time - j.time
	})
	occupiedChairs := &ChairsHeap{}
	for _, event := range events {
		if occupiedChairs.Len() != 0 && occupiedChairs.Peek().timeLeave <= event.time {
			newAvail := heap.Pop(occupiedChairs).(Item)
			heap.Push(availableChairs, newAvail.chair)
		}
		if event.eType > 0 {
			chair := heap.Pop(availableChairs).(int)
			if event.friendIdx == targetFriend {
				return chair
			}
			heap.Push(occupiedChairs, Item{timeLeave: times[event.friendIdx][1], chair: chair})
		}

	}
	return -1
}

type Event struct {
	time      int
	eType     int
	friendIdx int
}

// ************* Chairs heap ******************
type Item struct {
	timeLeave int
	chair     int
}

type ChairsHeap []Item

func (h ChairsHeap) Len() int { return len(h) }
func (h ChairsHeap) Less(i, j int) bool {
	if h[i].timeLeave == h[j].timeLeave {
		return h[i].chair < h[j].chair //When a friend arrives at the party, they sit on the unoccupied chair with the smallest number.
	}
	return h[i].timeLeave < h[j].timeLeave
}
func (h ChairsHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (pq *ChairsHeap) Push(x any) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *ChairsHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (h *ChairsHeap) Peek() Item {
	return (*h)[0]
}

// ********************************************
// ***************** int heap *****************
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() any {
	return (*h)[len(*h)-1]
}

//********************************************

//simulating
// func smallestChair(times [][]int, targetFriend int) int {
// 	idx := make([]int, len(times))
// 	for i := range idx {
// 		idx[i] = i
// 	}
// 	slices.SortFunc(idx, func(i, j int) int {
// 		return times[i][0] - times[j][0]
// 	})
// 	slices.SortFunc(times, func(i, j []int) int {
// 		return i[0] - j[0]
// 	})
// 	chairs := make([]int, len(times))
// 	for i := range times {
// 		for j := range chairs {
// 			if chairs[j] <= times[i][0] {
// 				if targetFriend == idx[i] {
// 					return j
// 				}
// 				chairs[j] = times[i][1]
// 				break
// 			}
// 		}

// 	}
// 	return 0
// }
