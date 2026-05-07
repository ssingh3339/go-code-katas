package webcrawler

import (
	"fmt"
	"sync"
)

// Fetcher interface for fetching URLs
type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

// CrawlMutex implements concurrent web crawler using mutex
func CrawlMutex(url string, depth int, fetcher Fetcher) map[string]bool {
	visited := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	var crawl func(url string, depth int)
	crawl = func(url string, depth int) {
		defer wg.Done()

		if depth <= 0 {
			return
		}

		// Check if already visited
		mu.Lock()
		if visited[url] {
			mu.Unlock()
			return
		}
		visited[url] = true
		mu.Unlock()

		// Fetch the URL
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", url, err)
			return
		}

		fmt.Printf("Found: %s %q\n", url, body)

		// Crawl child URLs
		for _, u := range urls {
			wg.Add(1)
			go crawl(u, depth-1)
		}
	}

	wg.Add(1)
	go crawl(url, depth)
	wg.Wait()

	return visited
}

// CrawlChannel implements concurrent web crawler using channels
func CrawlChannel(url string, maxDepth int, fetcher Fetcher) map[string]bool {
	visited := make(map[string]bool)
	type task struct {
		url   string
		depth int
	}

	tasks := make(chan task, 100)
	done := make(chan bool)

	// Coordinator goroutine
	go func() {
		tasks <- task{url, maxDepth}
		activeWorkers := 1

		for activeWorkers > 0 {
			t := <-tasks
			activeWorkers--

			// Check if already visited
			if visited[t.url] || t.depth <= 0 {
				if activeWorkers == 0 && len(tasks) == 0 {
					break
				}
				continue
			}

			visited[t.url] = true

			// Fetch URLs
			go func(t task) {
				body, urls, err := fetcher.Fetch(t.url)
				if err != nil {
					fmt.Printf("Error fetching %s: %v\n", t.url, err)
					done <- true
					return
				}

				fmt.Printf("Found: %s %q\n", t.url, body)

				// Send child tasks
				for _, u := range urls {
					tasks <- task{u, t.depth - 1}
				}
				done <- true
			}(t)

			activeWorkers++

			// Check for completion signals
			select {
			case <-done:
				activeWorkers--
			default:
			}
		}

		close(tasks)
		close(done)
	}()

	// Wait for completion
	for range done {
	}

	return visited
}

// SafeCounter is a thread-safe counter
type SafeCounter struct {
	mu    sync.Mutex
	count map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count[key]++
}

// Value returns the current count for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count[key]
}

// CrawlWorkerPool implements crawler with fixed worker pool
func CrawlWorkerPool(url string, maxDepth int, fetcher Fetcher, numWorkers int) map[string]bool {
	type task struct {
		url   string
		depth int
	}

	visited := make(map[string]bool)
	var mu sync.Mutex

	tasks := make(chan task, 100)
	var wg sync.WaitGroup

	// Start worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for t := range tasks {
				// Check if already visited
				mu.Lock()
				if visited[t.url] || t.depth <= 0 {
					mu.Unlock()
					continue
				}
				visited[t.url] = true
				mu.Unlock()

				// Fetch URL
				body, urls, err := fetcher.Fetch(t.url)
				if err != nil {
					fmt.Printf("Worker %d: Error fetching %s: %v\n", workerID, t.url, err)
					continue
				}

				fmt.Printf("Worker %d: Found %s %q\n", workerID, t.url, body)

				// Add child tasks
				for _, u := range urls {
					tasks <- task{u, t.depth - 1}
				}
			}
		}(i)
	}

	// Send initial task
	tasks <- task{url, maxDepth}

	// Wait a bit for tasks to complete
	// In production, you'd use a more sophisticated completion detection
	go func() {
		wg.Wait()
		close(tasks)
	}()

	// Give some time for processing
	wg.Wait()

	return visited
}

// CrawlSequential implements a sequential (non-concurrent) crawler for comparison
func CrawlSequential(url string, depth int, fetcher Fetcher, visited map[string]bool) {
	if depth <= 0 || visited[url] {
		return
	}

	visited[url] = true

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}

	fmt.Printf("Found: %s %q\n", url, body)

	for _, u := range urls {
		CrawlSequential(u, depth-1, fetcher, visited)
	}
}
