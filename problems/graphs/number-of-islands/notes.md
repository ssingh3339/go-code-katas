# Number of Islands - Learning Notes

## Key Insights

### 1. Graph Representation

- 2D grid is an implicit graph where each cell is a node
- Edges exist between adjacent cells (up, down, left, right)
- This is different from explicit graph representations (adjacency list/matrix)

### 2. Connected Components

- Each island is a connected component in the graph
- Finding islands = counting connected components
- Classic application of graph traversal

### 3. In-Place Modification

- Can modify the grid to mark visited cells (change '1' to '0')
- Saves space by avoiding separate visited set
- Trade-off: destroys original data

## Algorithm Comparison

### DFS (Recursion)

**Pros:**

- Clean, intuitive code
- Natural fit for connected component problems
- Easy to implement

**Cons:**

- Stack overflow risk for very large grids
- Uses implicit call stack

### BFS (Queue)

**Pros:**

- No recursion, no stack overflow
- Iterative approach
- Better for finding shortest paths

**Cons:**

- Slightly more code
- Requires explicit queue

### Union-Find

**Pros:**

- Natural fit for connectivity problems
- Can answer "are these connected?" queries efficiently
- Good when you need to track parent relationships

**Cons:**

- More complex implementation
- Overkill for simple island counting
- Extra space for parent/rank arrays

## Common Mistakes

1. **Forgetting boundary checks:** Always validate i, j before accessing grid[i][j]
2. **Diagonal connections:** Problem only considers 4-directional connectivity, not 8
3. **Not marking visited:** Will count same island multiple times
4. **Modifying while iterating:** Make sure to mark visited in the right order

## Complexity Deep Dive

### Why O(m × n)?

- Each cell is visited at most twice:
  1. Once in the main iteration
  2. Once during DFS/BFS exploration
- Marking as visited ensures we don't revisit

### Space Complexity

- **DFS:** O(m × n) worst case if entire grid is one island (call stack depth)
- **BFS:** O(min(m, n)) worst case (queue size at any time)
- **Union-Find:** O(m × n) for parent/rank arrays

## Variations and Extensions

### Similar Problems

1. **Max Area of Island:** Find the largest island
2. **Number of Closed Islands:** Islands not touching the border
3. **Island Perimeter:** Calculate the perimeter of islands
4. **Number of Distinct Islands:** Count unique shapes

### Modifications

1. **8-directional:** Include diagonal connections
2. **Different characters:** Instead of '0'/'1', use other markers
3. **Multiple queries:** Efficiently handle adding/removing land
4. **Sinking islands:** Remove islands as you find them

## Optimization Tips

1. **Early termination:** If you only need to know if islands exist, stop at 1
2. **Parallel processing:** Different sections can be processed in parallel
3. **Visited set vs modification:** Choose based on whether you can modify input

## Interview Strategy

1. **Clarify problem:**
   - 4-directional or 8-directional?
   - Can I modify the input grid?
   - What about empty grid?

2. **Choose approach:**
   - Default to DFS (simplest)
   - Use BFS if explicitly asked
   - Mention Union-Find as alternative

3. **Walk through example:**
   - Draw the grid
   - Show how you mark visited cells
   - Count islands step by step

4. **Consider edge cases:**
   - Empty grid
   - All water
   - All land
   - Single row/column

## Pattern Recognition

This problem demonstrates the **"flood fill"** pattern:

- Start from a seed point
- Explore all connected points
- Mark as visited
- Count number of separate fills needed

Similar to:

- Flood fill in paint programs
- Region coloring
- Maze solving
- Network connectivity
