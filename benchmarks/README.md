# Benchmarks

This directory is reserved for cross-problem benchmarking and performance comparisons.

## Purpose

Use this directory to:

1. **Compare Algorithms:** Benchmark different approaches to the same problem
2. **Data Structure Performance:** Compare performance of different data structures
3. **Concurrency Patterns:** Measure scalability of concurrent solutions
4. **System Benchmarks:** Test overall repository patterns

## Running Benchmarks

```bash
# Run all benchmarks
make benchmark

# Run specific benchmark
go test -bench=BenchmarkName ./benchmarks/...

# With memory stats
go test -bench=. -benchmem ./benchmarks/...

# Save results for comparison
go test -bench=. ./benchmarks/... > old.txt
# Make changes...
go test -bench=. ./benchmarks/... > new.txt
benchcmp old.txt new.txt
```

## Benchmark Guidelines

1. **Realistic Data:** Use representative input sizes
2. **Multiple Iterations:** Let `b.N` handle iteration count
3. **Reset Timer:** Use `b.ResetTimer()` after setup
4. **Consistent Environment:** Run on same machine, disable background tasks
5. **Statistical Significance:** Run multiple times, look for patterns

## Example Benchmark Structure

```go
package benchmarks

import "testing"

func BenchmarkApproach1(b *testing.B) {
    // Setup
    data := setupData()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        approach1(data)
    }
}

func BenchmarkApproach2(b *testing.B) {
    data := setupData()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        approach2(data)
    }
}
```

## Adding New Benchmarks

When adding benchmarks:

1. Create descriptive benchmark names
2. Document what is being measured
3. Include input size in benchmark name if relevant
4. Compare against baseline implementation
5. Document results in comments or separate file

## Profiling

For detailed profiling:

```bash
# CPU profile
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profile
go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof

# Trace
go test -bench=. -trace=trace.out
go tool trace trace.out
```

## Common Metrics

- **ns/op:** Nanoseconds per operation (lower is better)
- **B/op:** Bytes allocated per operation (lower is better)
- **allocs/op:** Number of allocations per operation (lower is better)

## Notes

- Individual problems have their own benchmarks in `*_test.go` files
- This directory is for comparative and system-wide benchmarking
- Consider using `benchstat` for statistical analysis of results
