package topksequences

import (
	"container/heap"
)

type sequenceCount struct {
	Key   string
	Value int
}

// sequenceCountMinHeap is used to store the top K sequences
type sequenceCountMinHeap []*sequenceCount

func newSequenceCountMinHeap() *sequenceCountMinHeap {
	min := &sequenceCountMinHeap{}
	heap.Init(min)
	return min
}

func (h *sequenceCountMinHeap) Len() int { return len(*h) }

func (h *sequenceCountMinHeap) Empty() bool { return len(*h) == 0 }

func (h *sequenceCountMinHeap) Top() *sequenceCount { return (*h)[0] }

func (h *sequenceCountMinHeap) Less(i, j int) bool {
	// when values are equal, compare the keys
	if (*h)[i].Value == (*h)[j].Value {
		return (*h)[i].Key > (*h)[j].Key
	}
	return (*h)[i].Value < (*h)[j].Value
}

func (h *sequenceCountMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *sequenceCountMinHeap) Push(x interface{}) {
	*h = append(*h, x.(*sequenceCount))
}

func (h *sequenceCountMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}
