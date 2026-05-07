# Two Sum

## Problem Statement

Given an array of integers `nums` and an integer `target`, return indices of the two numbers in the array such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.

## Examples

**Example 1:**
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

## Constraints

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- Only one valid answer exists

## Approach

### Brute Force (O(n²))
Check every pair of numbers to see if they sum to target.

### Hash Map (O(n))
Use a hash map to store numbers we've seen along with their indices. For each number, check if `target - number` exists in the map.

## Complexity Analysis

### Brute Force
- **Time Complexity:** O(n²) - nested loops
- **Space Complexity:** O(1) - no extra space

### Hash Map
- **Time Complexity:** O(n) - single pass through array
- **Space Complexity:** O(n) - hash map storage

## Edge Cases

1. Minimum array size (2 elements)
2. Negative numbers
3. Zero values
4. Large numbers
5. Solution at beginning/end of array

## Learning Notes

- Hash map pattern is very common for array problems
- Trading space for time is a key optimization technique
- Always consider the brute force approach first, then optimize
- Remember to handle edge cases in your solution
