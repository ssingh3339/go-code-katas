package templates

import "container/heap"

// MinHeap template
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// HeapTemplate demonstrates using a heap to find K largest/smallest elements
func TopKElements(nums []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)

	// Maintain a heap of size k
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Extract elements from heap
	result := make([]int, 0, k)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(int))
	}

	return result
}

// PriorityQueueItem represents an item with priority
type PriorityQueueItem struct {
	value    int
	priority int
	index    int
}

// PriorityQueue implements a priority queue
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Lower priority value means higher priority
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// PriorityQueueTemplate demonstrates using a priority queue
func PriorityQueueTemplate() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Add items
	heap.Push(&pq, &PriorityQueueItem{value: 10, priority: 3})
	heap.Push(&pq, &PriorityQueueItem{value: 20, priority: 1})
	heap.Push(&pq, &PriorityQueueItem{value: 30, priority: 2})

	// Process items by priority
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityQueueItem)
		_ = item // Process item
	}
}
