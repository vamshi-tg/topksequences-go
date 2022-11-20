package topksequences

import (
	"container/heap"
)

type SequenceCount struct {
	Key   string
	Value int
}

// SequenceCountMinHeap is used to store the top K sequences
type SequenceCountMinHeap []*SequenceCount

func newSequenceCountMinHeap() *SequenceCountMinHeap {
	min := &SequenceCountMinHeap{}
	heap.Init(min)
	return min
}

func (h *SequenceCountMinHeap) Len() int { return len(*h) }

func (h *SequenceCountMinHeap) Empty() bool { return len(*h) == 0 }

func (h *SequenceCountMinHeap) Top() *SequenceCount { return (*h)[0] }

func (h *SequenceCountMinHeap) Less(i, j int) bool {
	// when values are equal, compare the keys
	if (*h)[i].Value == (*h)[j].Value {
		return (*h)[i].Key > (*h)[j].Key
	}
	return (*h)[i].Value < (*h)[j].Value
}

func (h *SequenceCountMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *SequenceCountMinHeap) Push(x interface{}) {
	*h = append(*h, x.(*SequenceCount))
}

func (h *SequenceCountMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}
