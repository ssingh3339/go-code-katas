package graph

// Graph represents an adjacency list graph
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// NewGraph creates a new graph with n vertices
func NewGraph(n int) *Graph {
	return &Graph{
		Vertices: n,
		AdjList:  make(map[int][]int),
	}
}

// AddEdge adds an edge to the graph (directed)
func (g *Graph) AddEdge(from, to int) {
	g.AdjList[from] = append(g.AdjList[from], to)
}

// AddUndirectedEdge adds an undirected edge to the graph
func (g *Graph) AddUndirectedEdge(u, v int) {
	g.AddEdge(u, v)
	g.AddEdge(v, u)
}

// GetNeighbors returns the neighbors of a vertex
func (g *Graph) GetNeighbors(vertex int) []int {
	return g.AdjList[vertex]
}

// WeightedGraph represents a weighted graph
type WeightedGraph struct {
	Vertices int
	Edges    map[int][]Edge
}

// Edge represents a weighted edge
type Edge struct {
	To     int
	Weight int
}

// NewWeightedGraph creates a new weighted graph with n vertices
func NewWeightedGraph(n int) *WeightedGraph {
	return &WeightedGraph{
		Vertices: n,
		Edges:    make(map[int][]Edge),
	}
}

// AddEdge adds a weighted edge to the graph (directed)
func (g *WeightedGraph) AddEdge(from, to, weight int) {
	g.Edges[from] = append(g.Edges[from], Edge{To: to, Weight: weight})
}

// AddUndirectedEdge adds an undirected weighted edge to the graph
func (g *WeightedGraph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// GetEdges returns the edges from a vertex
func (g *WeightedGraph) GetEdges(vertex int) []Edge {
	return g.Edges[vertex]
}
