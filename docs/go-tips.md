# Go Programming Tips

## Idiomatic Go

### 1. Error Handling

```go
// Good: Check errors immediately
result, err := doSomething()
if err != nil {
    return err
}

// Bad: Ignoring errors
result, _ := doSomething()
```

### 2. Variable Names

```go
// Good: Short, clear names
for i, v := range values {
    sum += v
}

// Bad: Overly verbose
for index, value := range values {
    sum += value
}
```

### 3. Struct Initialization

```go
// Good: Named fields
person := Person{
    Name: "Alice",
    Age:  30,
}

// Avoid: Positional (fragile)
person := Person{"Alice", 30}
```

### 4. Zero Values

```go
// Good: Leverage zero values
var count int     // 0
var name string   // ""
var ptr *int      // nil

// Unnecessary explicit initialization
var count int = 0
```

## Testing Best Practices

### Table-Driven Tests

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d, want %d", 
                    tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

### Test Helpers

```go
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper() // Marks this as helper
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

## Concurrency Patterns

### 1. Don't Use Goroutines Without Control

```go
// Bad: Unbounded goroutines
for i := 0; i < 1000000; i++ {
    go process(i)
}

// Good: Use worker pool or semaphore
sem := make(chan struct{}, 100)
for i := 0; i < 1000000; i++ {
    sem <- struct{}{}
    go func(id int) {
        defer func() { <-sem }()
        process(id)
    }(i)
}
```

### 2. Always Clean Up Goroutines

```go
// Good: Proper cleanup
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

go func() {
    for {
        select {
        case <-ctx.Done():
            return // Clean exit
        default:
            doWork()
        }
    }
}()
```

### 3. Use sync.WaitGroup for Multiple Goroutines

```go
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        process(id)
    }(i)
}
wg.Wait()
```

## Performance Tips

### 1. Preallocate Slices When Size is Known

```go
// Good: Preallocate
result := make([]int, 0, 1000)

// Less efficient: Grow dynamically
result := []int{}
```

### 2. Use strings.Builder for String Concatenation

```go
// Good: strings.Builder
var b strings.Builder
for _, s := range words {
    b.WriteString(s)
}
result := b.String()

// Bad: String concatenation
result := ""
for _, s := range words {
    result += s
}
```

### 3. Avoid Memory Allocations in Loops

```go
// Good: Reuse buffer
buf := make([]byte, 1024)
for {
    n, err := reader.Read(buf)
    // Use buf[:n]
}

// Bad: Allocate every iteration
for {
    buf := make([]byte, 1024)
    n, err := reader.Read(buf)
}
```

### 4. Use Pointers for Large Structs

```go
// Good: Avoid copying large structs
func Process(data *LargeStruct) {
    // ...
}

// Inefficient: Copy on every call
func Process(data LargeStruct) {
    // ...
}
```

## Common Pitfalls

### 1. Loop Variable Capture

```go
// Wrong: All goroutines see same i
for i := 0; i < 10; i++ {
    go func() {
        fmt.Println(i) // Likely prints 10 ten times
    }()
}

// Correct: Pass as parameter
for i := 0; i < 10; i++ {
    go func(id int) {
        fmt.Println(id)
    }(i)
}
```

### 2. Nil Slice vs Empty Slice

```go
var s1 []int        // nil slice
s2 := []int{}       // empty slice
s3 := make([]int, 0) // empty slice

// All have len 0, but:
s1 == nil  // true
s2 == nil  // false
s3 == nil  // false
```

### 3. Mutex With Defer

```go
// Good: Defer unlock
func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

// Risky: Forget to unlock
func (c *Counter) Inc() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock() // What if panic?
}
```

### 4. Range Over Map

```go
// Warning: Order not guaranteed
for k, v := range myMap {
    fmt.Println(k, v) // Order varies
}
```

## Interface Tips

### 1. Accept Interfaces, Return Structs

```go
// Good design
func Process(r io.Reader) *Result {
    // ...
}

// Less flexible
func Process(f *os.File) *Result {
    // ...
}
```

### 2. Small Interfaces

```go
// Good: Small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Avoid: Large interfaces
type EverythingDoer interface {
    Read() error
    Write() error
    Close() error
    // ... 20 more methods
}
```

## Code Organization

### 1. Package Names

```go
// Good: Short, clear
package http
package strings

// Bad: Redundant
package httputil  // Use http/util instead
package stringhelper
```

### 2. Internal Packages

```
myproject/
  cmd/
  internal/         // Only accessible within myproject
    common/
  pkg/             // Public packages
```

### 3. File Organization

- One package per directory
- Related functionality in same file
- Keep files reasonably sized (< 500 lines)

## Memory Management

### 1. Escape Analysis

```go
// Stays on stack (fast)
func small() int {
    x := 42
    return x
}

// Escapes to heap (slower)
func escape() *int {
    x := 42
    return &x
}
```

### 2. Pool for Temporary Objects

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

// Get from pool
buf := bufferPool.Get().([]byte)
defer bufferPool.Put(buf)
```

## Debugging and Profiling

### 1. Race Detector

```bash
go test -race ./...
go run -race main.go
```

### 2. CPU Profiling

```bash
go test -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### 3. Memory Profiling

```bash
go test -memprofile=mem.prof
go tool pprof mem.prof
```

### 4. Benchmarking

```bash
go test -bench=. -benchmem
```

## Best Practices Summary

1. ✅ Handle all errors
2. ✅ Use `defer` for cleanup
3. ✅ Keep functions small and focused
4. ✅ Use table-driven tests
5. ✅ Document exported functions
6. ✅ Use `go fmt` and `go vet`
7. ✅ Leverage zero values
8. ✅ Prefer composition over inheritance
9. ✅ Write clear, simple code
10. ✅ Profile before optimizing

## Tools Every Go Developer Should Know

- `go fmt` - Format code
- `go vet` - Static analysis
- `go test` - Run tests
- `go mod` - Dependency management
- `golangci-lint` - Comprehensive linting
- `pprof` - Performance profiling
- `delve` - Debugger
- `go doc` - Documentation

## Resources

- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Proverbs](https://go-proverbs.github.io/)
