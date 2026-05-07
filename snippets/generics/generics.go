package generics

import (
	"fmt"
)

// GenericMin returns the minimum of two comparable values
func GenericMin[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// GenericMax returns the maximum of two comparable values
func GenericMax[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Stack is a generic stack data structure
type Stack[T any] struct {
	items []T
}

// NewStack creates a new generic stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: []T{},
	}
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// IsEmpty returns true if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// StackDemo demonstrates generic stack usage
func StackDemo() {
	// Integer stack
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	val, _ := intStack.Pop()
	fmt.Println("Popped:", val)

	// String stack
	strStack := NewStack[string]()
	strStack.Push("hello")
	strStack.Push("world")
	str, _ := strStack.Pop()
	fmt.Println("Popped:", str)
}

// Map applies a function to each element of a slice
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter filters a slice based on a predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces a slice to a single value
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// MapFilterReduceDemo demonstrates higher-order functions with generics
func MapFilterReduceDemo() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Square all numbers
	squared := Map(nums, func(n int) int { return n * n })
	fmt.Println("Squared:", squared)

	// Filter even numbers
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens)

	// Sum all numbers
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Sum:", sum)
}

// Pair is a generic pair type
type Pair[T, U any] struct {
	First  T
	Second U
}

// NewPair creates a new pair
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

// PairDemo demonstrates generic pairs
func PairDemo() {
	p1 := NewPair(1, "one")
	fmt.Printf("Pair: %d -> %s\n", p1.First, p1.Second)

	p2 := NewPair("key", 42)
	fmt.Printf("Pair: %s -> %d\n", p2.First, p2.Second)
}

// Ordered is a constraint for ordered types
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// FindMax finds the maximum value in a slice
func FindMax[T Ordered](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}

	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max, true
}

// FindMin finds the minimum value in a slice
func FindMin[T Ordered](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}

	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min, true
}

// OrderedDemo demonstrates ordered constraint
func OrderedDemo() {
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6}
	maxInt, _ := FindMax(ints)
	minInt, _ := FindMin(ints)
	fmt.Printf("Max: %d, Min: %d\n", maxInt, minInt)

	floats := []float64{3.14, 2.71, 1.41, 1.73}
	maxFloat, _ := FindMax(floats)
	minFloat, _ := FindMin(floats)
	fmt.Printf("Max: %.2f, Min: %.2f\n", maxFloat, minFloat)
}

// Contains checks if a slice contains a value
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Index returns the index of value in slice, or -1 if not found
func Index[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
