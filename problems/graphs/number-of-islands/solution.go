package numberofislands

// NumIslands counts islands using DFS
// Time: O(m*n), Space: O(m*n) for recursion stack
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i, j int) {
	// Boundary check
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}

	// Already visited or water
	if grid[i][j] != '1' {
		return
	}

	// Mark as visited
	grid[i][j] = '0'

	// Explore all 4 directions
	dfs(grid, i-1, j) // up
	dfs(grid, i+1, j) // down
	dfs(grid, i, j-1) // left
	dfs(grid, i, j+1) // right
}

// NumIslandsBFS counts islands using BFS
// Time: O(m*n), Space: O(min(m,n)) for queue
func NumIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				bfs(grid, i, j)
			}
		}
	}

	return count
}

type point struct {
	row, col int
}

func bfs(grid [][]byte, startRow, startCol int) {
	rows := len(grid)
	cols := len(grid[0])
	queue := []point{{startRow, startCol}}
	grid[startRow][startCol] = '0'

	directions := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newRow := curr.row + dir.row
			newCol := curr.col + dir.col

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == '1' {
				grid[newRow][newCol] = '0'
				queue = append(queue, point{newRow, newCol})
			}
		}
	}
}

// NumIslandsUnionFind counts islands using Union-Find
// Time: O(m*n), Space: O(m*n)
func NumIslandsUnionFind(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	uf := newUnionFind(rows * cols)

	// Count water cells
	waterCount := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				waterCount++
				continue
			}

			// Union with right neighbor
			if j+1 < cols && grid[i][j+1] == '1' {
				uf.union(i*cols+j, i*cols+j+1)
			}

			// Union with down neighbor
			if i+1 < rows && grid[i+1][j] == '1' {
				uf.union(i*cols+j, (i+1)*cols+j)
			}
		}
	}

	return uf.count - waterCount
}

type unionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{
		parent: parent,
		rank:   rank,
		count:  n,
	}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootX == rootY {
		return
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	uf.count--
}
