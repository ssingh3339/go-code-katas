package templates

// DFS template for graph traversal
// Replace types and logic as needed for your specific problem

// DFSTemplate is a recursive DFS traversal template
func DFSTemplate(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	dfsHelper(graph, start, visited, &result)
	return result
}

func dfsHelper(graph map[int][]int, node int, visited map[int]bool, result *[]int) {
	// Mark current node as visited
	visited[node] = true

	// Process current node
	*result = append(*result, node)

	// Explore neighbors recursively
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfsHelper(graph, neighbor, visited, result)
		}
	}
}

// DFSIterativeTemplate is an iterative DFS using a stack
func DFSIterativeTemplate(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	stack := []int{start}

	for len(stack) > 0 {
		// Pop from stack
		n := len(stack)
		node := stack[n-1]
		stack = stack[:n-1]

		// Skip if already visited
		if visited[node] {
			continue
		}

		// Mark as visited and process
		visited[node] = true
		result = append(result, node)

		// Push neighbors to stack
		// Push in reverse order to maintain left-to-right traversal
		neighbors := graph[node]
		for i := len(neighbors) - 1; i >= 0; i-- {
			if !visited[neighbors[i]] {
				stack = append(stack, neighbors[i])
			}
		}
	}

	return result
}

// BacktrackingTemplate demonstrates a backtracking pattern
func BacktrackingTemplate(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}
	backtrack(candidates, target, 0, current, &result)
	return result
}

func backtrack(candidates []int, target int, start int, current []int, result *[][]int) {
	// Base case: found a valid solution
	if target == 0 {
		// Make a copy of current path
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	// Explore candidates
	for i := start; i < len(candidates); i++ {
		// Prune: skip if candidate is too large
		if candidates[i] > target {
			continue
		}

		// Make choice
		current = append(current, candidates[i])

		// Explore
		backtrack(candidates, target-candidates[i], i+1, current, result)

		// Undo choice (backtrack)
		current = current[:len(current)-1]
	}
}
