package twosum

// TwoSum returns indices of two numbers that add up to target using hash map
// Time: O(n), Space: O(n)
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index

	for i, num := range nums {
		complement := target - num
		if idx, found := seen[complement]; found {
			return []int{idx, i}
		}
		seen[num] = i
	}

	return nil // no solution found (shouldn't happen per problem statement)
}

// TwoSumBruteForce returns indices using brute force approach
// Time: O(n²), Space: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
