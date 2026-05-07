package numberofislands

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "example 1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "example 2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "all water",
			grid: [][]byte{
				{'0', '0'},
				{'0', '0'},
			},
			want: 0,
		},
		{
			name: "all land",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 1,
		},
		{
			name: "single cell land",
			grid: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: "single cell water",
			grid: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			name: "checkerboard pattern",
			grid: [][]byte{
				{'1', '0', '1', '0'},
				{'0', '1', '0', '1'},
				{'1', '0', '1', '0'},
				{'0', '1', '0', '1'},
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" DFS", func(t *testing.T) {
			grid := copyGrid(tt.grid)
			got := NumIslands(grid)
			if got != tt.want {
				t.Errorf("NumIslands() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" BFS", func(t *testing.T) {
			grid := copyGrid(tt.grid)
			got := NumIslandsBFS(grid)
			if got != tt.want {
				t.Errorf("NumIslandsBFS() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" UnionFind", func(t *testing.T) {
			grid := copyGrid(tt.grid)
			got := NumIslandsUnionFind(grid)
			if got != tt.want {
				t.Errorf("NumIslandsUnionFind() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to copy grid since we modify it in place
func copyGrid(original [][]byte) [][]byte {
	if original == nil {
		return nil
	}
	copy := make([][]byte, len(original))
	for i := range original {
		copy[i] = make([]byte, len(original[i]))
		for j := range original[i] {
			copy[i][j] = original[i][j]
		}
	}
	return copy
}

// Benchmark tests
func BenchmarkNumIslandsDFS(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testGrid := copyGrid(grid)
		NumIslands(testGrid)
	}
}

func BenchmarkNumIslandsBFS(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testGrid := copyGrid(grid)
		NumIslandsBFS(testGrid)
	}
}

func BenchmarkNumIslandsUnionFind(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testGrid := copyGrid(grid)
		NumIslandsUnionFind(testGrid)
	}
}

func BenchmarkNumIslandsLargeGrid(b *testing.B) {
	// Create a 100x100 grid with some islands
	grid := make([][]byte, 100)
	for i := range grid {
		grid[i] = make([]byte, 100)
		for j := range grid[i] {
			if (i+j)%3 == 0 {
				grid[i][j] = '1'
			} else {
				grid[i][j] = '0'
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testGrid := copyGrid(grid)
		NumIslands(testGrid)
	}
}
