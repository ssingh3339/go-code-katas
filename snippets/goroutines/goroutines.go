package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// BasicGoroutine demonstrates launching a simple goroutine
func BasicGoroutine() {
	go func() {
		fmt.Println("Hello from goroutine")
	}()

	// Wait a bit for goroutine to complete (not recommended in production)
	time.Sleep(time.Millisecond * 100)
}

// WaitGroupExample shows how to wait for multiple goroutines
func WaitGroupExample() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is working\n", id)
			time.Sleep(time.Millisecond * 100)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")
}

// WorkerPool demonstrates a worker pool pattern
func WorkerPool(jobs []int, numWorkers int) []int {
	jobsChan := make(chan int, len(jobs))
	resultsChan := make(chan int, len(jobs))

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobsChan, resultsChan, &wg)
	}

	// Send jobs
	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)

	// Wait for workers to complete
	wg.Wait()
	close(resultsChan)

	// Collect results
	results := []int{}
	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Process job
		result := job * 2
		results <- result
	}
}

// FanOut demonstrates fan-out pattern (one input, multiple processors)
func FanOut(input []int, numWorkers int) []int {
	inputChan := make(chan int, len(input))
	outputChans := make([]chan int, numWorkers)

	// Create output channels
	for i := 0; i < numWorkers; i++ {
		outputChans[i] = make(chan int)
	}

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for val := range inputChan {
				// Process value
				result := val * val
				outputChans[workerID] <- result
			}
			close(outputChans[workerID])
		}(i)
	}

	// Send input
	for _, val := range input {
		inputChan <- val
	}
	close(inputChan)

	// Collect results from all workers
	results := []int{}
	for _, ch := range outputChans {
		for result := range ch {
			results = append(results, result)
		}
	}

	return results
}

// ParallelProcessing demonstrates processing items in parallel
func ParallelProcessing(items []int, processFunc func(int) int) []int {
	numWorkers := len(items)
	if numWorkers > 10 {
		numWorkers = 10 // Limit concurrency
	}

	results := make([]int, len(items))
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, numWorkers)

	for i, item := range items {
		wg.Add(1)
		go func(index, value int) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire
			defer func() { <-semaphore }() // Release

			results[index] = processFunc(value)
		}(i, item)
	}

	wg.Wait()
	return results
}

// Pipeline demonstrates a pipeline pattern
func Pipeline(input []int) []int {
	// Stage 1: Generate
	gen := func(nums []int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	// Stage 2: Square
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Stage 3: Add one
	addOne := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n + 1
			}
			close(out)
		}()
		return out
	}

	// Connect pipeline stages
	c1 := gen(input)
	c2 := square(c1)
	c3 := addOne(c2)

	// Collect results
	results := []int{}
	for result := range c3 {
		results = append(results, result)
	}

	return results
}
