package templates

// BFS template for graph traversal
// Replace types and logic as needed for your specific problem

type Queue []int

func (q *Queue) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, true
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// BFSTemplate is a generic BFS traversal template
func BFSTemplate(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	queue := &Queue{}

	// Initialize: add start node to queue
	queue.Enqueue(start)
	visited[start] = true

	// Process nodes level by level
	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()

		// Process current node
		result = append(result, node)

		// Explore neighbors
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.Enqueue(neighbor)
			}
		}
	}

	return result
}

// LevelOrderBFS performs BFS with level tracking
func LevelOrderBFS(graph map[int][]int, start int) [][]int {
	result := [][]int{}
	if len(graph) == 0 {
		return result
	}

	visited := make(map[int]bool)
	queue := &Queue{}

	queue.Enqueue(start)
	visited[start] = true

	for !queue.IsEmpty() {
		levelSize := len(*queue)
		currentLevel := []int{}

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node, _ := queue.Dequeue()
			currentLevel = append(currentLevel, node)

			// Add unvisited neighbors for next level
			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue.Enqueue(neighbor)
				}
			}
		}

		result = append(result, currentLevel)
	}

	return result
}
