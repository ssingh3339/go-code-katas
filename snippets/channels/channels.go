package channels

import (
	"fmt"
	"time"
)

// BasicChannel demonstrates basic channel operations
func BasicChannel() {
	ch := make(chan int)

	// Send in a goroutine
	go func() {
		ch <- 42
	}()

	// Receive
	value := <-ch
	fmt.Println("Received:", value)
}

// BufferedChannel demonstrates buffered channels
func BufferedChannel() {
	ch := make(chan int, 3)

	// Can send up to buffer size without blocking
	ch <- 1
	ch <- 2
	ch <- 3

	// Receive values
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// ChannelDirection demonstrates directional channels
func ChannelDirection() {
	ch := make(chan int)

	// Send-only channel
	go sender(ch)

	// Receive-only channel
	receiver(ch)
}

func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
	}
}

// SelectStatement demonstrates select for multiple channels
func SelectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

// SelectWithTimeout demonstrates timeout pattern
func SelectWithTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case result := <-ch:
		fmt.Println("Got:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}

// SelectWithDefault demonstrates non-blocking operations
func SelectWithDefault() {
	ch := make(chan int, 1)

	// Non-blocking send
	select {
	case ch <- 42:
		fmt.Println("Sent")
	default:
		fmt.Println("Channel full, skipping")
	}

	// Non-blocking receive
	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	default:
		fmt.Println("No data available")
	}
}

// FanIn merges multiple channels into one
func FanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)

	for _, ch := range channels {
		go func(c <-chan int) {
			for val := range c {
				out <- val
			}
		}(ch)
	}

	return out
}

// OrDone pattern provides early cancellation
func OrDone(done <-chan struct{}, ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case val, ok := <-ch:
				if !ok {
					return
				}
				select {
				case out <- val:
				case <-done:
					return
				}
			}
		}
	}()
	return out
}

// Broadcast sends a value to multiple channels
func Broadcast(value int, channels ...chan<- int) {
	for _, ch := range channels {
		go func(c chan<- int) {
			c <- value
		}(ch)
	}
}

// RateLimiter demonstrates rate limiting with channels
func RateLimiter() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Rate limit: one request per 200ms
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()

	for req := range requests {
		<-limiter.C
		fmt.Println("Processing request", req, time.Now())
	}
}

// BurstLimiter allows bursts of requests
func BurstLimiter() {
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the bucket with 3 tokens
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Refill 1 token every 200ms
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for t := range ticker.C {
			burstyLimiter <- t
		}
	}()

	// Simulate 5 requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-burstyLimiter
		fmt.Println("Processing bursty request", req, time.Now())
	}
}
