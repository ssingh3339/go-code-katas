# Web Crawler

## Problem Statement

Implement a concurrent web crawler that:

1. Crawls URLs starting from a given URL
2. Extracts links from each page
3. Avoids crawling the same URL twice
4. Uses goroutines to crawl multiple URLs concurrently
5. Limits the maximum depth of crawling

The crawler should be thread-safe and efficiently utilize Go's concurrency primitives.

## Interface

```go
type Fetcher interface {
    // Fetch returns the body of URL and a slice of URLs found on that page
    Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) map[string]bool
```

## Requirements

1. **Concurrent execution:** Use goroutines to fetch multiple URLs in parallel
2. **Avoid duplicates:** Don't fetch the same URL twice
3. **Depth limiting:** Stop crawling when reaching maximum depth
4. **Synchronization:** Safely share state between goroutines
5. **Graceful completion:** Wait for all goroutines to finish

## Examples

**Example 1:**

```
Input: url = "https://golang.org/", depth = 2
Output: map of all crawled URLs
```

## Approaches

### 1. Mutex + WaitGroup

- Use sync.Mutex to protect visited map
- Use sync.WaitGroup to track active goroutines
- Simple and straightforward

### 2. Channel-Based

- Use channels to communicate between goroutines
- Centralized visited tracking
- More idiomatic Go

### 3. Worker Pool

- Fixed number of worker goroutines
- Process URLs from a queue
- Better resource control

## Complexity Analysis

### Mutex Approach

- **Time Complexity:** O(N) where N is number of unique URLs
- **Space Complexity:** O(N) for visited map
- **Concurrency:** Unbounded goroutines

### Channel Approach

- **Time Complexity:** O(N)
- **Space Complexity:** O(N)
- **Concurrency:** Bounded by channel buffer

## Edge Cases

1. Empty start URL
2. Depth of 0
3. No links found
4. Circular links (A -> B -> A)
5. Network errors
6. Very deep link chains

## Learning Notes

- Demonstrates practical use of goroutines and channels
- Shows different synchronization patterns
- Teaches about race conditions and thread safety
- Real-world concurrent programming example
