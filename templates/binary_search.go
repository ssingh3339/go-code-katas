package templates

// BinarySearchTemplate demonstrates binary search pattern
func BinarySearchTemplate(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // not found
}

// BinarySearchLeftmost finds leftmost (first) occurrence of target
func BinarySearchLeftmost(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// BinarySearchRightmost finds rightmost (last) occurrence of target
func BinarySearchRightmost(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left > 0 && nums[left-1] == target {
		return left - 1
	}
	return -1
}

// BinarySearchOnAnswer template for "search on answer" problems
// Example: Find minimum capacity to ship packages within D days
func MinCapacity(weights []int, days int) int {
	// Define search space
	left := max(weights)  // minimum possible capacity
	right := sum(weights) // maximum possible capacity

	// Binary search on answer
	for left < right {
		mid := left + (right-left)/2

		if canShip(weights, days, mid) {
			// Try smaller capacity
			right = mid
		} else {
			// Need larger capacity
			left = mid + 1
		}
	}

	return left
}

// Helper function to check if we can ship with given capacity
func canShip(weights []int, days int, capacity int) bool {
	daysNeeded := 1
	currentWeight := 0

	for _, w := range weights {
		if currentWeight+w > capacity {
			daysNeeded++
			currentWeight = 0
		}
		currentWeight += w
	}

	return daysNeeded <= days
}

// Helper functions
func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// BinarySearchRotated searches in a rotated sorted array
func BinarySearchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Determine which half is sorted
		if nums[left] <= nums[mid] {
			// Left half is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
