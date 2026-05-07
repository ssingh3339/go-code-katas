package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "with zero",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
		{
			name:   "large numbers",
			nums:   []int{1000000000, 2000000000, -1000000000},
			target: 1000000000,
			want:   []int{1, 2},
		},
		{
			name:   "solution at end",
			nums:   []int{1, 2, 3, 4, 5},
			target: 9,
			want:   []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumBruteForce(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSumBruteForce(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	target := 1500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}

func BenchmarkTwoSumBruteForce(b *testing.B) {
	nums := make([]int, 100) // Smaller size for brute force
	for i := range nums {
		nums[i] = i
	}
	target := 150

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSumBruteForce(nums, target)
	}
}

func BenchmarkTwoSumSmallInput(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}

func BenchmarkTwoSumLargeInput(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i
	}
	target := 15000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}
