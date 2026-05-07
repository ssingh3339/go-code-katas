# Problem Solving Patterns

## 1. Two Pointers

### When to Use

- Sorted arrays
- Finding pairs with specific sum
- Removing duplicates
- Palindrome problems
- Container with most water

### Pattern

```go
left, right := 0, len(arr)-1
for left < right {
    // Process current pair
    // Move pointers based on condition
    if condition {
        left++
    } else {
        right--
    }
}
```

### Example Problems

- Two Sum (sorted array)
- Three Sum
- Container With Most Water
- Valid Palindrome
- Remove Duplicates

---

## 2. Sliding Window

### When to Use

- Contiguous subarrays/substrings
- Maximum/minimum in subarrays
- Longest/shortest subarray with condition
- Fixed or variable window size

### Pattern (Fixed Size)

```go
windowSum := 0
for i := 0; i < k; i++ {
    windowSum += arr[i]
}
maxSum := windowSum

for i := k; i < len(arr); i++ {
    windowSum = windowSum - arr[i-k] + arr[i]
    maxSum = max(maxSum, windowSum)
}
```

### Pattern (Variable Size)

```go
left := 0
for right := 0; right < len(arr); right++ {
    // Add arr[right] to window
    
    // Shrink window if needed
    for windowNotValid() {
        // Remove arr[left] from window
        left++
    }
    
    // Update result
}
```

### Example Problems

- Maximum Sum Subarray of Size K
- Longest Substring Without Repeating Characters
- Minimum Window Substring
- Longest Substring with K Distinct Characters

---

## 3. Fast & Slow Pointers

### When to Use

- Linked list cycle detection
- Finding middle of linked list
- Palindrome linked list
- Happy number problem

### Pattern

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    
    if slow == fast {
        // Cycle detected or other condition
    }
}
```

### Example Problems

- Linked List Cycle
- Find Middle of Linked List
- Happy Number
- Palindrome Linked List

---

## 4. Binary Search

### When to Use

- Sorted array
- Search space can be divided
- Finding boundary conditions
- Optimization problems (minimize/maximize)

### Pattern (Classic)

```go
left, right := 0, len(arr)-1
for left <= right {
    mid := left + (right-left)/2
    
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        left = mid + 1
    } else {
        right = mid - 1
    }
}
```

### Pattern (Search on Answer)

```go
left, right := minPossible, maxPossible
for left < right {
    mid := left + (right-left)/2
    
    if isValid(mid) {
        right = mid  // Try smaller
    } else {
        left = mid + 1  // Need larger
    }
}
return left
```

### Example Problems

- Binary Search
- First Bad Version
- Search in Rotated Array
- Find Minimum in Rotated Array
- Capacity To Ship Packages

---

## 5. Top K Elements (Heap)

### When to Use

- Finding K largest/smallest elements
- K-way merge
- Priority-based problems
- Running median

### Pattern

```go
import "container/heap"

h := &MinHeap{}
heap.Init(h)

for _, num := range nums {
    heap.Push(h, num)
    if h.Len() > k {
        heap.Pop(h)
    }
}

// h now contains top k elements
```

### Example Problems

- Kth Largest Element
- Top K Frequent Elements
- K Closest Points to Origin
- Merge K Sorted Lists

---

## 6. Breadth-First Search (BFS)

### When to Use

- Level-order traversal
- Shortest path in unweighted graph
- Distance from source
- Connected components

### Pattern (Graph)

```go
queue := []int{start}
visited := make(map[int]bool)
visited[start] = true

for len(queue) > 0 {
    node := queue[0]
    queue = queue[1:]
    
    // Process node
    
    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            visited[neighbor] = true
            queue = append(queue, neighbor)
        }
    }
}
```

### Pattern (Tree Level Order)

```go
if root == nil {
    return nil
}

queue := []*TreeNode{root}
for len(queue) > 0 {
    levelSize := len(queue)
    level := []int{}
    
    for i := 0; i < levelSize; i++ {
        node := queue[0]
        queue = queue[1:]
        level = append(level, node.Val)
        
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}
```

### Example Problems

- Binary Tree Level Order Traversal
- Rotting Oranges
- Word Ladder
- Number of Islands

---

## 7. Depth-First Search (DFS)

### When to Use

- Tree/graph traversal
- Finding all paths
- Backtracking problems
- Cycle detection

### Pattern (Recursive)

```go
func dfs(node int, visited map[int]bool) {
    if visited[node] {
        return
    }
    
    visited[node] = true
    // Process node
    
    for _, neighbor := range graph[node] {
        dfs(neighbor, visited)
    }
}
```

### Pattern (Iterative)

```go
stack := []int{start}
visited := make(map[int]bool)

for len(stack) > 0 {
    node := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    
    if visited[node] {
        continue
    }
    
    visited[node] = true
    // Process node
    
    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            stack = append(stack, neighbor)
        }
    }
}
```

### Example Problems

- Tree Traversals (Inorder, Preorder, Postorder)
- Number of Islands
- Path Sum
- Clone Graph

---

## 8. Backtracking

### When to Use

- All possible combinations
- All possible permutations
- Constraint satisfaction
- Subset problems

### Pattern

```go
func backtrack(path []int, remaining []int, result *[][]int) {
    // Base case: found valid solution
    if isValid(path) {
        temp := make([]int, len(path))
        copy(temp, path)
        *result = append(*result, temp)
        return
    }
    
    // Try all choices
    for i, choice := range remaining {
        // Make choice
        path = append(path, choice)
        
        // Recurse
        backtrack(path, remaining[i+1:], result)
        
        // Undo choice (backtrack)
        path = path[:len(path)-1]
    }
}
```

### Example Problems

- Subsets
- Permutations
- Combinations
- N-Queens
- Sudoku Solver

---

## 9. Dynamic Programming

### When to Use

- Optimization problems (min/max)
- Counting problems
- Overlapping subproblems
- Optimal substructure

### Pattern (Top-Down - Memoization)

```go
memo := make(map[string]int)

func dp(state string) int {
    if val, exists := memo[state]; exists {
        return val
    }
    
    // Base case
    if baseCase(state) {
        return baseValue
    }
    
    // Recursive case
    result := compute(dp(nextState1), dp(nextState2))
    memo[state] = result
    return result
}
```

### Pattern (Bottom-Up - Tabulation)

```go
dp := make([]int, n+1)
dp[0] = baseValue

for i := 1; i <= n; i++ {
    for j := 0; j < i; j++ {
        dp[i] = compute(dp[i], dp[j])
    }
}

return dp[n]
```

### Example Problems

- Fibonacci
- Climbing Stairs
- Coin Change
- Longest Increasing Subsequence
- Edit Distance

---

## 10. Union Find (Disjoint Set)

### When to Use

- Connected components
- Cycle detection in undirected graph
- Network connectivity
- Minimum spanning tree

### Pattern

```go
type UnionFind struct {
    parent []int
    rank   []int
}

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
    rootX, rootY := uf.Find(x), uf.Find(y)
    if rootX == rootY {
        return false
    }
    
    if uf.rank[rootX] < uf.rank[rootY] {
        uf.parent[rootX] = rootY
    } else if uf.rank[rootX] > uf.rank[rootY] {
        uf.parent[rootY] = rootX
    } else {
        uf.parent[rootY] = rootX
        uf.rank[rootX]++
    }
    return true
}
```

### Example Problems

- Number of Connected Components
- Redundant Connection
- Accounts Merge
- Most Stones Removed

---

## 11. Greedy

### When to Use

- Optimization problems with greedy choice property
- Local optimal leads to global optimal
- Activity selection
- Interval scheduling

### Pattern

```go
// Sort by some criteria
sort.Slice(items, func(i, j int) bool {
    return items[i].someField < items[j].someField
})

result := 0
for _, item := range items {
    if satisfiesCondition(item) {
        // Make greedy choice
        result += item.value
    }
}
```

### Example Problems

- Jump Game
- Gas Station
- Interval Scheduling
- Minimum Coins

---

## 12. Monotonic Stack/Queue

### When to Use

- Next greater/smaller element
- Histogram problems
- Temperature problems
- Sliding window maximum

### Pattern (Next Greater Element)

```go
stack := []int{}
result := make([]int, len(arr))

for i := len(arr) - 1; i >= 0; i-- {
    // Pop smaller elements
    for len(stack) > 0 && arr[stack[len(stack)-1]] <= arr[i] {
        stack = stack[:len(stack)-1]
    }
    
    if len(stack) == 0 {
        result[i] = -1
    } else {
        result[i] = arr[stack[len(stack)-1]]
    }
    
    stack = append(stack, i)
}
```

### Example Problems

- Next Greater Element
- Daily Temperatures
- Largest Rectangle in Histogram
- Trapping Rain Water

---

## General Problem-Solving Strategy

1. **Understand the problem**
   - Read carefully
   - Identify inputs and outputs
   - List constraints

2. **Examples**
   - Walk through examples
   - Consider edge cases
   - Test understanding

3. **Identify pattern**
   - Does it fit a known pattern?
   - What data structure helps?
   - What's the brute force approach?

4. **Design algorithm**
   - Start with brute force
   - Optimize using patterns
   - Analyze time/space complexity

5. **Code**
   - Write clean, readable code
   - Handle edge cases
   - Use meaningful names

6. **Test**
   - Run through examples
   - Test edge cases
   - Verify complexity

7. **Optimize**
   - Can we do better?
   - Trade time for space or vice versa
   - Is there a mathematical insight?
