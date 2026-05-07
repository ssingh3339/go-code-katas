package templates

// UnionFind template for disjoint set operations
type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

// NewUnionFind creates a new union-find structure
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
		count:  n,
	}
}

// Find returns the root of element x with path compression
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges sets containing x and y
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	// Union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	uf.count--
	return true
}

// Connected checks if x and y are in the same set
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Count returns the number of disjoint sets
func (uf *UnionFind) Count() int {
	return uf.count
}

// UnionFindTemplate demonstrates a typical union-find problem
// Example: Find number of connected components in a graph
func CountComponents(n int, edges [][]int) int {
	uf := NewUnionFind(n)

	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}

	return uf.Count()
}

// DetectCycle uses union-find to detect cycle in undirected graph
func DetectCycle(n int, edges [][]int) bool {
	uf := NewUnionFind(n)

	for _, edge := range edges {
		// If both nodes already in same set, adding edge creates cycle
		if !uf.Union(edge[0], edge[1]) {
			return true
		}
	}

	return false
}
