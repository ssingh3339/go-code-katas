package bfs

import "github.com/satyendrasingh/go-code-katas/internal/ds/queue"

// BFS performs breadth-first search on a graph
// Returns the order of visited nodes
func BFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	q := queue.NewQueue()

	q.Enqueue(start)
	visited[start] = true

	for !q.IsEmpty() {
		node, _ := q.Dequeue()
		result = append(result, node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				q.Enqueue(neighbor)
			}
		}
	}

	return result
}

// ShortestPath finds the shortest path from start to end using BFS
// Returns the path as a slice of nodes, or nil if no path exists
func ShortestPath(graph map[int][]int, start, end int) []int {
	if start == end {
		return []int{start}
	}

	visited := make(map[int]bool)
	parent := make(map[int]int)
	q := queue.NewQueue()

	q.Enqueue(start)
	visited[start] = true

	found := false
	for !q.IsEmpty() && !found {
		node, _ := q.Dequeue()

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = node
				q.Enqueue(neighbor)

				if neighbor == end {
					found = true
					break
				}
			}
		}
	}

	if !found {
		return nil
	}

	// Reconstruct path
	path := []int{}
	for node := end; node != start; node = parent[node] {
		path = append([]int{node}, path...)
	}
	path = append([]int{start}, path...)

	return path
}
