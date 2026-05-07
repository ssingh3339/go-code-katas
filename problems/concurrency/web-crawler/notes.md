# Web Crawler - Learning Notes

## Key Concurrency Concepts

### 1. Race Conditions
Without proper synchronization, multiple goroutines accessing the `visited` map can cause:
- Lost updates
- Duplicate fetches
- Map corruption (panic)

### 2. Synchronization Primitives

**sync.Mutex:**
- Protects shared state (visited map)
- Lock before read/write, unlock after
- Simple and effective for protecting data structures

**sync.WaitGroup:**
- Tracks number of active goroutines
- Wait for all workers to complete
- Essential for knowing when crawl is done

**Channels:**
- Communicate between goroutines
- Can replace mutex + shared memory
- More idiomatic Go in some cases

### 3. Deadlock Prevention
Common deadlock scenarios:
- Forgetting to call `wg.Done()`
- Circular wait on channels
- Holding lock while waiting

## Implementation Strategies

### Mutex + WaitGroup Approach
**Pros:**
- Simple to understand
- Direct access to shared state
- Easy to debug

**Cons:**
- Lock contention under high concurrency
- Need careful lock/unlock pairing
- Potential for deadlocks

### Channel-Based Approach
**Pros:**
- No explicit locks
- Follows "share memory by communicating"
- Natural backpressure control

**Cons:**
- More complex control flow
- Need to handle channel closure
- Harder to reason about completion

### Worker Pool Pattern
**Pros:**
- Bounded concurrency
- Better resource control
- Predictable behavior

**Cons:**
- Fixed number of workers
- May underutilize with few URLs
- More setup code

## Common Mistakes

1. **Forgetting defer wg.Done():**
   ```go
   // Wrong
   go func() {
       wg.Done()  // If panic occurs, never called
       doWork()
   }()
   
   // Right
   go func() {
       defer wg.Done()
       doWork()
   }()
   ```

2. **Race on visited map:**
   ```go
   // Wrong - race condition
   if !visited[url] {
       visited[url] = true
   }
   
   // Right - atomic check-and-set
   mu.Lock()
   if !visited[url] {
       visited[url] = true
   }
   mu.Unlock()
   ```

3. **Unbounded goroutine creation:**
   - Can exhaust system resources
   - Use worker pools or semaphores

4. **Not handling errors:**
   - Network failures are common
   - Need proper error handling and retries

## Performance Considerations

### Concurrency vs. Parallelism
- **Concurrency:** Structure of program (can do multiple things)
- **Parallelism:** Execution (actually doing multiple things simultaneously)

### Optimal Worker Count
- Too few: Underutilize resources
- Too many: Context switching overhead
- Sweet spot: Usually 2-4x number of CPU cores for I/O bound tasks

### Bottlenecks
1. **Network I/O:** Usually the bottleneck in web crawling
2. **Lock contention:** If spending too much time waiting for locks
3. **Memory:** Storing visited URLs can grow large

## Testing Concurrent Code

### Challenges
- Non-deterministic execution
- Race conditions may not always manifest
- Timing-dependent bugs

### Tools
- `go test -race`: Detects race conditions
- `go test -count=100`: Run tests multiple times
- Mock interfaces for predictable testing

### Best Practices
- Test both sequential and concurrent versions
- Use fake/mock implementations
- Test edge cases (empty, depth 0, errors)
- Benchmark to compare approaches

## Real-World Extensions

### Production-Ready Crawler Would Need:
1. **Rate limiting:** Don't overwhelm servers
2. **Robots.txt:** Respect site policies
3. **Retries:** Handle transient failures
4. **Timeout:** Don't wait forever
5. **URL normalization:** Treat equivalent URLs the same
6. **Distributed crawling:** Scale across machines
7. **Storage:** Persist crawled data
8. **Duplicate detection:** Bloom filters for efficiency

## Go Concurrency Patterns Demonstrated

1. **Fan-out:** Create multiple goroutines from one
2. **Fan-in:** Collect results from multiple goroutines
3. **Worker pool:** Fixed set of workers processing tasks
4. **Pipeline:** Chain of processing stages
5. **Cancellation:** Stop work via context (not shown, but useful)

## Memory Management

### Visited Map Growth
- Can grow to millions of entries
- Consider:
  - Bloom filters for space efficiency
  - LRU cache with size limit
  - Database for persistence

### Goroutine Stack
- Each goroutine has a stack (starts small, grows)
- Too many goroutines = high memory usage
- Monitor with `runtime.NumGoroutine()`

## Debugging Tips

1. **Add logging:** Print goroutine IDs, timestamps
2. **Use channels for visibility:** Log state changes
3. **Visualize:** Draw goroutine interactions
4. **Race detector:** Always run with `-race` during development
5. **Profiling:** Use pprof for goroutine leaks

## Interview Tips

1. **Start simple:** Sequential solution first
2. **Identify parallelism:** Where can we run concurrently?
3. **Choose primitive:** Mutex or channels?
4. **Handle edge cases:** Empty, errors, circular links
5. **Discuss trade-offs:** Simplicity vs. performance
6. **Mention production concerns:** Rate limiting, politeness

## Related Problems

- Parallel file processing
- Concurrent download manager
- Map-reduce framework
- Task scheduler
- Load balancer

## Key Takeaways

1. **Shared state needs protection:** Use mutex or channels
2. **Track goroutine completion:** WaitGroup or done channel
3. **Avoid goroutine leaks:** Always ensure cleanup
4. **Test for races:** Use `-race` flag
5. **Consider worker pools:** For bounded concurrency
6. **Channels vs. mutexes:** Neither is always better, choose based on use case
