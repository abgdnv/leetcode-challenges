package main

import (
	"container/heap"
	"strings"
)

type Item struct {
	c     string
	count int
}

func longestDiverseString(a int, b int, c int) string {

	var result strings.Builder
	pq := &PriorityQueue{}
	if a > 0 {
		heap.Push(pq, &Item{"a", a})
	}
	if b > 0 {
		heap.Push(pq, &Item{"b", b})
	}
	if c > 0 {
		heap.Push(pq, &Item{"c", c})
	}

	prev0 := ""
	prev1 := ""
	for len(*pq) > 0 {
		prevLen := result.Len()

		v := heap.Pop(pq).(*Item)
		if result.Len() >= 2 && prev0 == v.c && prev1 == v.c {
			if pq.Len() == 0 {
				break
			}
			v2 := heap.Pop(pq).(*Item)
			result.WriteString(v2.c)
			prev0, prev1 = prev1, v2.c
			v2.count--
			if v2.count > 0 {
				heap.Push(pq, v2)
			}
		} else {
			result.WriteString(v.c)
			prev0, prev1 = prev1, v.c
			v.count--
		}
		if v.count > 0 {
			heap.Push(pq, v)
		}

		if prevLen == result.Len() {
			break
		}
	}

	if result.Len() > 0 {
		return result.String()
	}
	return ""
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
