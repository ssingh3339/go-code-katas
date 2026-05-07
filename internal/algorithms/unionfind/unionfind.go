package unionfind

// UnionFind represents a union-find (disjoint set) data structure
type UnionFind struct {
	parent []int
	rank   []int
	count  int // number of connected components
}

// NewUnionFind creates a new UnionFind structure with n elements
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

// Find returns the root of the element x with path compression
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y
// Returns true if x and y were in different sets, false otherwise
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

// Count returns the number of connected components
func (uf *UnionFind) Count() int {
	return uf.count
}
