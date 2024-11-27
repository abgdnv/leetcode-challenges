package main

import (
	"container/heap"
	"math"
)

func maxKelements(nums []int, k int) int64 {
	h := &IntHeap{}
	heap.Init(h)
	for i := range nums {
		heap.Push(h, nums[i])
	}
	var sum int64
	for range k {
		v := heap.Pop(h).(int)
		sum += int64(v)
		v = int(math.Ceil(float64(v) / 3.0))
		heap.Push(h, v)
	}
	return sum
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
