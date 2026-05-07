# Number of Islands

## Problem Statement

Given an `m x n` 2D binary grid which represents a map of `'1'`s (land) and `'0'`s (water), return the number of islands.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## Examples

**Example 1:**

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

**Example 2:**

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
```

## Constraints

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`

## Approach

### DFS (Depth-First Search)

1. Iterate through each cell in the grid
2. When we find a '1', increment island count
3. Use DFS to mark all connected land cells as visited (change to '0' or use visited set)
4. Continue until all cells are processed

### BFS (Breadth-First Search)

1. Similar to DFS, but use a queue instead of recursion
2. When we find a '1', start BFS to explore all connected land
3. Mark cells as visited during exploration

### Union-Find

1. Treat each land cell as a node
2. Union adjacent land cells
3. Count the number of distinct root nodes

## Complexity Analysis

### DFS/BFS

- **Time Complexity:** O(m × n) - visit each cell once
- **Space Complexity:** O(m × n) worst case for recursion stack or queue

### Union-Find

- **Time Complexity:** O(m × n × α(m×n)) where α is inverse Ackermann function
- **Space Complexity:** O(m × n) for parent/rank arrays

## Edge Cases

1. Empty grid
2. All water ('0')
3. All land ('1') - single island
4. Single row or column
5. Diagonal lands (not connected)

## Learning Notes

- Classic graph traversal problem
- Both DFS and BFS work equally well
- Can modify grid in-place to avoid extra space for visited set
- Good practice for 2D grid traversal patterns
- Understanding connected components
