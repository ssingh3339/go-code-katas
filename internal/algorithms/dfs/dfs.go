package dfs

// DFS performs depth-first search on a graph
// Returns the order of visited nodes
func DFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	dfsHelper(graph, start, visited, &result)
	return result
}

func dfsHelper(graph map[int][]int, node int, visited map[int]bool, result *[]int) {
	visited[node] = true
	*result = append(*result, node)

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfsHelper(graph, neighbor, visited, result)
		}
	}
}

// DFSIterative performs depth-first search iteratively using a stack
func DFSIterative(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	stack := []int{start}

	for len(stack) > 0 {
		// Pop from stack
		n := len(stack)
		node := stack[n-1]
		stack = stack[:n-1]

		if visited[node] {
			continue
		}

		visited[node] = true
		result = append(result, node)

		// Push neighbors in reverse order to maintain left-to-right traversal
		neighbors := graph[node]
		for i := len(neighbors) - 1; i >= 0; i-- {
			if !visited[neighbors[i]] {
				stack = append(stack, neighbors[i])
			}
		}
	}

	return result
}

// HasCycle detects if there's a cycle in a directed graph using DFS
func HasCycle(graph map[int][]int, numNodes int) bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for node := 0; node < numNodes; node++ {
		if !visited[node] {
			if hasCycleHelper(graph, node, visited, recStack) {
				return true
			}
		}
	}

	return false
}

func hasCycleHelper(graph map[int][]int, node int, visited, recStack map[int]bool) bool {
	visited[node] = true
	recStack[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			if hasCycleHelper(graph, neighbor, visited, recStack) {
				return true
			}
		} else if recStack[neighbor] {
			return true
		}
	}

	recStack[node] = false
	return false
}
