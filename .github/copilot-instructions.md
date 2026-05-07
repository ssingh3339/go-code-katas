# GitHub Copilot Instructions for go-code-katas

## Repository Overview

This is a production-quality Go repository for practicing programming problems, algorithms, data structures, and concurrency exercises. The repository follows strict conventions for organization, testing, and documentation.

## Go Version and Standards

- **Go Version:** 1.26.2
- **Module Path:** `github.com/satyendrasingh/go-code-katas`
- Follow idiomatic Go conventions from [Effective Go](https://golang.org/doc/effective_go)
- Use standard library only - no external dependencies

## Repository Structure

```
go-code-katas/
├── problems/           # Problem implementations by category
├── internal/          # Reusable data structures and algorithms
├── templates/         # Algorithm templates
├── snippets/          # Concurrency and Go-specific patterns
├── benchmarks/        # Performance benchmarks
└── docs/             # Documentation
```

## Problem Structure Convention

Every problem MUST follow this structure:

```
problems/category/problem-name/
├── README.md          # Problem description, approach, complexity
├── solution.go        # Implementation
├── solution_test.go   # Tests and benchmarks
└── notes.md          # Learning notes and insights
```

### README.md Template

```markdown
# Problem Name

## Problem Statement
[Clear description of the problem]

## Examples
[At least 2 examples with input/output]

## Constraints
[List all constraints]

## Approach
[Explain the solution approach]

## Complexity Analysis
- Time Complexity: O(?)
- Space Complexity: O(?)

## Edge Cases
[List important edge cases]

## Learning Notes
[Key insights from solving this problem]
```

### solution.go Requirements

1. **Package Naming:**
   - Use lowercase, no underscores: `package twosum`
   - Match directory name

2. **Function Documentation:**
   ```go
   // FunctionName describes what it does
   // Time: O(?), Space: O(?)
   func FunctionName(params) returnType {
       // implementation
   }
   ```

3. **Multiple Solutions:**
   - Implement both brute force and optimized
   - Name variants clearly: `FunctionName`, `FunctionNameBruteForce`, `FunctionNameOptimized`

### solution_test.go Requirements

1. **Table-Driven Tests:**
   ```go
   func TestFunctionName(t *testing.T) {
       tests := []struct {
           name string
           input type
           want type
       }{
           {"descriptive name", input, expected},
       }
       
       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               got := FunctionName(tt.input)
               if got != tt.want {
                   t.Errorf("got %v, want %v", got, tt.want)
               }
           })
       }
   }
   ```

2. **Benchmark Tests:**
   ```go
   func BenchmarkFunctionName(b *testing.B) {
       // Setup
       data := setupData()
       
       b.ResetTimer()
       for i := 0; i < b.N; i++ {
           FunctionName(data)
       }
   }
   ```

3. **Test Coverage:**
   - Include edge cases
   - Test empty inputs
   - Test boundary conditions
   - Test error cases

### notes.md Template

```markdown
# Problem Name - Learning Notes

## Key Insights
[Main takeaways]

## Common Mistakes
[Pitfalls to avoid]

## Time Complexity Breakdown
[Detailed analysis]

## Extensions
[Related problems or variations]

## Interview Tips
[How to approach in interviews]
```

## Coding Conventions

### General Go Guidelines

1. **Error Handling:**
   ```go
   // Good
   result, err := doSomething()
   if err != nil {
       return err
   }
   
   // Bad - never ignore errors
   result, _ := doSomething()
   ```

2. **Variable Names:**
   - Short names in small scopes: `i`, `n`, `ch`
   - Descriptive names in larger scopes: `userCount`, `maxDepth`
   - No Hungarian notation

3. **Zero Values:**
   ```go
   // Good - leverage zero values
   var count int     // 0
   var items []int   // nil
   
   // Avoid unnecessary initialization
   var count int = 0
   ```

4. **Defer for Cleanup:**
   ```go
   mu.Lock()
   defer mu.Unlock()
   
   file, err := os.Open(path)
   if err != nil {
       return err
   }
   defer file.Close()
   ```

### Concurrency Patterns

1. **Always Use WaitGroup for Multiple Goroutines:**
   ```go
   var wg sync.WaitGroup
   for i := 0; i < n; i++ {
       wg.Add(1)
       go func(id int) {
           defer wg.Done()
           process(id)
       }(i)
   }
   wg.Wait()
   ```

2. **Avoid Goroutine Variable Capture:**
   ```go
   // Wrong
   for i := 0; i < 10; i++ {
       go func() {
           fmt.Println(i) // captures i
       }()
   }
   
   // Correct
   for i := 0; i < 10; i++ {
       go func(id int) {
           fmt.Println(id)
       }(i)
   }
   ```

3. **Use Context for Cancellation:**
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   
   select {
   case <-ctx.Done():
       return ctx.Err()
   case result := <-resultChan:
       return result
   }
   ```

### Data Structure Conventions

1. **Use `internal/ds/` for Reusable Structures:**
   - Heap: `internal/ds/heap/`
   - Trie: `internal/ds/trie/`
   - LinkedList: `internal/ds/linkedlist/`
   - Graph: `internal/ds/graph/`
   - Queue: `internal/ds/queue/`

2. **Use `internal/algorithms/` for Algorithms:**
   - BFS: `internal/algorithms/bfs/`
   - DFS: `internal/algorithms/dfs/`
   - Dijkstra: `internal/algorithms/dijkstra/`
   - UnionFind: `internal/algorithms/unionfind/`
   - Sorting: `internal/algorithms/sorting/`

3. **Import Reusable Components:**
   ```go
   import "github.com/satyendrasingh/go-code-katas/internal/ds/heap"
   import "github.com/satyendrasingh/go-code-katas/internal/algorithms/bfs"
   ```

## Testing Guidelines

### Test Organization

1. **One test file per implementation file:**
   - `solution.go` → `solution_test.go`

2. **Test function naming:**
   - `TestFunctionName` for unit tests
   - `BenchmarkFunctionName` for benchmarks

3. **Use `t.Helper()` for test utilities:**
   ```go
   func assertEqual(t *testing.T, got, want interface{}) {
       t.Helper()
       if got != want {
           t.Errorf("got %v, want %v", got, want)
       }
   }
   ```

### Running Tests

- `make test` - Run all tests with coverage
- `make benchmark` - Run all benchmarks
- `go test -race ./...` - Detect race conditions
- `go test -v ./...` - Verbose output

## Documentation Standards

### Code Comments

1. **Exported Functions:**
   ```go
   // TwoSum returns indices of two numbers that add up to target.
   // It uses a hash map for O(n) time complexity.
   func TwoSum(nums []int, target int) []int {
   ```

2. **Complexity Annotations:**
   ```go
   // Time: O(n), Space: O(n)
   ```

3. **Inline Comments:**
   - Explain "why", not "what"
   - Use for complex logic only
   - Keep comments up-to-date

### README Standards

- Use proper markdown formatting
- Include code examples with syntax highlighting
- Add complexity analysis for all solutions
- List edge cases explicitly

## Makefile Commands

- `make test` - Run tests with coverage
- `make benchmark` - Run benchmarks
- `make lint` - Run golangci-lint
- `make format` - Format code with gofmt
- `make vet` - Run go vet
- `make tidy` - Tidy go.mod
- `make check` - Run vet, format check, and tests
- `make clean` - Clean artifacts

## Common Patterns

### Array/Slice Operations

```go
// Preallocate when size is known
result := make([]int, 0, expectedSize)

// Copy slice
copy := make([]int, len(original))
copy(copy, original)

// Reverse slice in place
for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
    arr[i], arr[j] = arr[j], arr[i]
}
```

### Hash Map Patterns

```go
// Check and insert
if _, exists := seen[key]; !exists {
    seen[key] = value
}

// Get with default
value, ok := m[key]
if !ok {
    value = defaultValue
}
```

### Graph Traversal

```go
// BFS with queue
queue := []int{start}
visited := make(map[int]bool)
visited[start] = true

for len(queue) > 0 {
    node := queue[0]
    queue = queue[1:]
    
    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            visited[neighbor] = true
            queue = append(queue, neighbor)
        }
    }
}
```

## Problem Categories

When suggesting or implementing problems, use these categories:

1. **arrays** - Array and hash map problems
2. **strings** - String manipulation
3. **linkedlist** - Linked list operations
4. **trees** - Binary trees and BST
5. **graphs** - Graph algorithms
6. **heap** - Heap-based problems
7. **stack** - Stack problems
8. **queue** - Queue problems
9. **dynamic-programming** - DP problems
10. **greedy** - Greedy algorithms
11. **backtracking** - Backtracking problems
12. **sliding-window** - Sliding window pattern
13. **binary-search** - Binary search variations
14. **concurrency** - Goroutines and channels
15. **system-design** - System design exercises
16. **miscellaneous** - Other problems

## Quality Checklist

Before considering any problem complete, ensure:

- [ ] README.md with problem description and complexity analysis
- [ ] solution.go with documented functions
- [ ] solution_test.go with table-driven tests
- [ ] notes.md with learning insights
- [ ] At least 5 test cases including edge cases
- [ ] At least 1 benchmark test
- [ ] All tests passing: `make test`
- [ ] No lint errors: `make lint`
- [ ] Code formatted: `make format`
- [ ] Complexity analysis documented
- [ ] Edge cases identified and tested

## Performance Considerations

1. **Avoid Allocations in Loops:**
   ```go
   // Good
   buf := make([]byte, 1024)
   for {
       n, _ := reader.Read(buf)
       process(buf[:n])
   }
   ```

2. **Use strings.Builder for Concatenation:**
   ```go
   var b strings.Builder
   for _, s := range words {
       b.WriteString(s)
   }
   result := b.String()
   ```

3. **Benchmark Before Optimizing:**
   - Profile with `go test -bench=. -cpuprofile=cpu.prof`
   - Analyze with `go tool pprof cpu.prof`

## References

- Internal docs: `docs/`
  - `complexity-cheatsheet.md` - Time/space complexity reference
  - `go-tips.md` - Go best practices
  - `problem-solving-patterns.md` - Algorithmic patterns
- Templates: `templates/` - Reusable algorithm patterns
- Snippets: `snippets/` - Concurrency examples

## Philosophy

1. **Simplicity over cleverness** - Write clear, maintainable code
2. **Test first** - Write tests before implementation when possible
3. **Document thoroughly** - Help future developers (including yourself)
4. **Idiomatic Go** - Follow Go conventions and best practices
5. **Performance matters** - But measure before optimizing
6. **Learn and teach** - Notes should help others learn from your solutions
