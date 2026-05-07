package mutex

import (
	"fmt"
	"sync"
)

// Counter demonstrates basic mutex usage
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value safely returns the counter value
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// ConcurrentIncrement demonstrates concurrent counter updates
func ConcurrentIncrement() {
	counter := &Counter{}
	var wg sync.WaitGroup

	// Launch 1000 goroutines to increment
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Value())
}

// RWMutexExample demonstrates read-write mutex
type Cache struct {
	mu    sync.RWMutex
	items map[string]string
}

// NewCache creates a new cache
func NewCache() *Cache {
	return &Cache{
		items: make(map[string]string),
	}
}

// Set writes to the cache
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

// Get reads from the cache
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.items[key]
	return val, ok
}

// RWMutexDemo demonstrates the performance benefit of RWMutex
func RWMutexDemo() {
	cache := NewCache()
	var wg sync.WaitGroup

	// One writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		}
	}()

	// Multiple readers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				cache.Get(fmt.Sprintf("key%d", j))
			}
		}()
	}

	wg.Wait()
}

// OnceExample demonstrates sync.Once
var instance *Singleton
var once sync.Once

type Singleton struct {
	data string
}

// GetInstance returns the singleton instance
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating singleton instance")
		instance = &Singleton{data: "singleton"}
	})
	return instance
}

// OnceDemo demonstrates sync.Once usage
func OnceDemo() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = GetInstance()
		}()
	}

	wg.Wait()
	// "Creating singleton instance" is printed only once
}

// SafeMap is a thread-safe map
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// NewSafeMap creates a new SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Store adds or updates a key-value pair
func (sm *SafeMap) Store(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Load retrieves a value by key
func (sm *SafeMap) Load(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}

// Delete removes a key
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

// Len returns the number of items
func (sm *SafeMap) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

// SyncMapExample demonstrates sync.Map
func SyncMapExample() {
	var sm sync.Map

	// Store
	sm.Store("key1", "value1")
	sm.Store("key2", "value2")

	// Load
	if value, ok := sm.Load("key1"); ok {
		fmt.Println("Found:", value)
	}

	// LoadOrStore
	actual, loaded := sm.LoadOrStore("key3", "value3")
	fmt.Println("Actual:", actual, "Loaded:", loaded)

	// Delete
	sm.Delete("key2")

	// Range
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %s\n", key, value)
		return true // continue iteration
	})
}
