package context

import (
	"context"
	"fmt"
	"time"
)

// BasicContext demonstrates basic context usage
func BasicContext() {
	// Background context - used as a starting point
	ctx1 := context.Background()
	_ = ctx1

	// TODO context - placeholder when context is unclear
	ctx2 := context.TODO()
	_ = ctx2
}

// ContextWithTimeout demonstrates timeout context
func ContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}
}

// ContextWithCancel demonstrates manual cancellation
func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		cancel() // Cancel after 1 second
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Cancelled:", ctx.Err())
	}
}

// ContextWithDeadline demonstrates deadline context
func ContextWithDeadline() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Deadline exceeded:", ctx.Err())
	}
}

// ContextWithValue demonstrates passing values through context
func ContextWithValue() {
	type key string
	const userKey key = "user"

	ctx := context.WithValue(context.Background(), userKey, "john_doe")

	// Retrieve value
	if user, ok := ctx.Value(userKey).(string); ok {
		fmt.Println("User:", user)
	}
}

// WorkerWithContext demonstrates a worker that respects context
func WorkerWithContext(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// Simulate work
			time.Sleep(time.Millisecond * 100)
			results <- job * 2
		}
	}
}

// CancellableWorkersDemo demonstrates cancelling workers with context
func CancellableWorkersDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	for i := 0; i < 3; i++ {
		go WorkerWithContext(ctx, i, jobs, results)
	}

	// Send jobs
	go func() {
		for i := 0; i < 10; i++ {
			jobs <- i
			time.Sleep(time.Millisecond * 300)
		}
		close(jobs)
	}()

	// Collect results until context is done
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done, stopping")
			return
		case result, ok := <-results:
			if !ok {
				return
			}
			fmt.Println("Result:", result)
		}
	}
}

// HTTPRequestWithContext simulates an HTTP request with context
func HTTPRequestWithContext(ctx context.Context, url string) error {
	resultChan := make(chan error, 1)

	go func() {
		// Simulate HTTP request
		time.Sleep(3 * time.Second)
		resultChan <- nil // success
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-resultChan:
		return err
	}
}

// RequestWithTimeoutDemo demonstrates HTTP request with timeout
func RequestWithTimeoutDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := HTTPRequestWithContext(ctx, "https://example.com")
	if err != nil {
		fmt.Println("Request error:", err)
	} else {
		fmt.Println("Request successful")
	}
}

// ContextChaining demonstrates context value inheritance
func ContextChaining() {
	type key string
	const requestIDKey key = "requestID"
	const userKey key = "user"

	// Create base context with request ID
	ctx := context.WithValue(context.Background(), requestIDKey, "req-123")

	// Create derived context with user info and timeout
	ctx = context.WithValue(ctx, userKey, "alice")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Both values are accessible
	fmt.Println("Request ID:", ctx.Value(requestIDKey))
	fmt.Println("User:", ctx.Value(userKey))
}

// CascadingCancellation demonstrates parent-child context cancellation
func CascadingCancellation() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	// Child context
	childCtx, childCancel := context.WithCancel(parentCtx)
	defer childCancel()

	go func() {
		<-childCtx.Done()
		fmt.Println("Child context cancelled")
	}()

	// Cancelling parent also cancels child
	parentCancel()
	time.Sleep(time.Millisecond * 100)
}
