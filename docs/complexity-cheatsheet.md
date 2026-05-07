# Time Complexity Cheatsheet

## Array Operations

| Operation | Average | Worst Case | Notes |
|-----------|---------|------------|-------|
| Access | O(1) | O(1) | Direct index access |
| Search | O(n) | O(n) | Linear search unsorted |
| Binary Search | O(log n) | O(log n) | Requires sorted array |
| Insert (end) | O(1) | O(1) | Amortized with append |
| Insert (middle) | O(n) | O(n) | Need to shift elements |
| Delete | O(n) | O(n) | Need to shift elements |

## Hash Map Operations

| Operation | Average | Worst Case | Notes |
|-----------|---------|------------|-------|
| Search | O(1) | O(n) | With good hash function |
| Insert | O(1) | O(n) | May trigger resize |
| Delete | O(1) | O(n) | With good hash function |

## Linked List Operations

| Operation | Average | Worst Case | Notes |
|-----------|---------|------------|-------|
| Access | O(n) | O(n) | Must traverse from head |
| Search | O(n) | O(n) | Linear traversal |
| Insert (head) | O(1) | O(1) | Direct pointer manipulation |
| Insert (tail) | O(n) | O(n) | Without tail pointer |
| Delete | O(n) | O(n) | Need to find node first |

## Stack/Queue Operations

| Operation | Average | Worst Case | Notes |
|-----------|---------|------------|-------|
| Push | O(1) | O(1) | Add to end |
| Pop | O(1) | O(1) | Remove from end/front |
| Peek | O(1) | O(1) | View without removing |

## Binary Search Tree Operations

| Operation | Average | Worst Case | Notes |
|-----------|---------|------------|-------|
| Search | O(log n) | O(n) | Balanced vs. skewed |
| Insert | O(log n) | O(n) | Balanced vs. skewed |
| Delete | O(log n) | O(n) | Balanced vs. skewed |

## Heap Operations

| Operation | Average | Worst Case | Notes |
|-----------|---------|------------|-------|
| Find Min/Max | O(1) | O(1) | Root element |
| Insert | O(log n) | O(log n) | Bubble up |
| Delete Min/Max | O(log n) | O(log n) | Bubble down |
| Build Heap | O(n) | O(n) | From array |

## Sorting Algorithms

| Algorithm | Best | Average | Worst | Space | Stable |
|-----------|------|---------|-------|-------|--------|
| Bubble Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |
| Insertion Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |
| Selection Sort | O(n²) | O(n²) | O(n²) | O(1) | No |
| Merge Sort | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes |
| Quick Sort | O(n log n) | O(n log n) | O(n²) | O(log n) | No |
| Heap Sort | O(n log n) | O(n log n) | O(n log n) | O(1) | No |
| Counting Sort | O(n+k) | O(n+k) | O(n+k) | O(k) | Yes |
| Radix Sort | O(nk) | O(nk) | O(nk) | O(n+k) | Yes |

## Graph Algorithms

| Algorithm | Time Complexity | Space | Notes |
|-----------|----------------|-------|-------|
| BFS | O(V + E) | O(V) | Queue + visited set |
| DFS | O(V + E) | O(V) | Stack + visited set |
| Dijkstra | O((V + E) log V) | O(V) | With min-heap |
| Bellman-Ford | O(VE) | O(V) | Handles negative weights |
| Floyd-Warshall | O(V³) | O(V²) | All pairs shortest path |
| Prim's MST | O(E log V) | O(V) | With min-heap |
| Kruskal's MST | O(E log E) | O(V) | With union-find |
| Topological Sort | O(V + E) | O(V) | DAG only |

## Common Algorithm Patterns

### Two Pointers
- **Time:** O(n)
- **Space:** O(1)
- **Use:** Sorted arrays, palindromes, pairs

### Sliding Window
- **Time:** O(n)
- **Space:** O(1) or O(k)
- **Use:** Subarrays, substrings

### Binary Search
- **Time:** O(log n)
- **Space:** O(1)
- **Use:** Sorted data, search space problems

### Dynamic Programming
- **Time:** O(n²) to O(2ⁿ) → O(n²) or better
- **Space:** O(n) to O(n²)
- **Use:** Optimization, counting problems

### Backtracking
- **Time:** O(2ⁿ) or O(n!)
- **Space:** O(n) for recursion
- **Use:** Combinations, permutations

### Greedy
- **Time:** Varies (often O(n log n))
- **Space:** O(1) to O(n)
- **Use:** Optimization with optimal substructure

## String Algorithms

| Algorithm | Time | Space | Notes |
|-----------|------|-------|-------|
| Naive Search | O(nm) | O(1) | Pattern length m, text length n |
| KMP | O(n + m) | O(m) | Knuth-Morris-Pratt |
| Rabin-Karp | O(n + m) | O(1) | Average case |
| Trie Insert | O(m) | O(alphabet × m × n) | m = word length |
| Trie Search | O(m) | O(1) | Just traversal |

## Math Operations

| Operation | Time | Notes |
|-----------|------|-------|
| Is Prime | O(√n) | Trial division |
| GCD | O(log(min(a,b))) | Euclidean algorithm |
| Power | O(log n) | Exponentiation by squaring |
| Factorial | O(n) | Iterative |

## Big-O Notation Hierarchy

From fastest to slowest:
1. **O(1)** - Constant
2. **O(log n)** - Logarithmic
3. **O(n)** - Linear
4. **O(n log n)** - Linearithmic
5. **O(n²)** - Quadratic
6. **O(n³)** - Cubic
7. **O(2ⁿ)** - Exponential
8. **O(n!)** - Factorial

## Rules of Thumb

1. **Drop constants:** O(2n) → O(n)
2. **Drop non-dominant terms:** O(n² + n) → O(n²)
3. **Different inputs use different variables:** O(a + b), not O(n)
4. **Amortized time:** Average time over sequence of operations
5. **Space complexity includes:** Input space, auxiliary space, and recursion stack

## Quick Reference

- Accessing array element: **O(1)**
- Binary search: **O(log n)**
- Linear search: **O(n)**
- Sorting (comparison): **O(n log n)** best average
- Matrix multiplication: **O(n³)** naive
- Fibonacci (recursive): **O(2ⁿ)**
- Fibonacci (DP): **O(n)**

## Tips for Analysis

1. **Count nested loops:** Usually multiply complexities
2. **Recursive calls:** Draw recursion tree
3. **Divide and conquer:** Often O(n log n)
4. **Multiple passes:** Add complexities
5. **Early termination:** Best case analysis
