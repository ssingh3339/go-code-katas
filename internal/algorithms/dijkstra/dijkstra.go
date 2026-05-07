package dijkstra

import (
	"container/heap"
	"math"
)

// Edge represents a weighted edge
type Edge struct {
	To     int
	Weight int
}

// Item represents a node in the priority queue
type Item struct {
	node     int
	distance int
	index    int
}

// PriorityQueue implements heap.Interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
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

// ShortestPath finds the shortest path from start to all other nodes using Dijkstra's algorithm
// Returns a map of distances from start to each node
func ShortestPath(graph map[int][]Edge, start, numNodes int) map[int]int {
	dist := make(map[int]int)
	for i := 0; i < numNodes; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		node := item.node
		distance := item.distance

		if distance > dist[node] {
			continue
		}

		for _, edge := range graph[node] {
			newDist := dist[node] + edge.Weight
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				heap.Push(&pq, &Item{node: edge.To, distance: newDist})
			}
		}
	}

	return dist
}

// FindPath finds the actual path from start to end using Dijkstra's algorithm
// Returns the path and the total distance, or nil if no path exists
func FindPath(graph map[int][]Edge, start, end, numNodes int) ([]int, int) {
	dist := make(map[int]int)
	parent := make(map[int]int)
	for i := 0; i < numNodes; i++ {
		dist[i] = math.MaxInt32
		parent[i] = -1
	}
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		node := item.node
		distance := item.distance

		if node == end {
			break
		}

		if distance > dist[node] {
			continue
		}

		for _, edge := range graph[node] {
			newDist := dist[node] + edge.Weight
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				parent[edge.To] = node
				heap.Push(&pq, &Item{node: edge.To, distance: newDist})
			}
		}
	}

	if dist[end] == math.MaxInt32 {
		return nil, -1
	}

	// Reconstruct path
	path := []int{}
	for node := end; node != -1; node = parent[node] {
		path = append([]int{node}, path...)
	}

	return path, dist[end]
}
