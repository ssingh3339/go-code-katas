# Two Sum - Learning Notes

## Key Insights

### 1. Hash Map Pattern
The hash map approach is a classic example of trading space for time. By storing previously seen numbers, we can check for the complement in O(1) time.

### 2. Single Pass
We don't need to populate the entire hash map first. We can check and insert in a single pass through the array.

### 3. Why This Works
For each number `x`, we need to find `target - x`. By storing numbers as we go, when we reach the complement, we've already seen the first number.

## Common Mistakes

1. **Using the same element twice:** Make sure to check `i != j` in brute force, or insert after checking in hash map approach
2. **Forgetting to return indices:** The problem asks for indices, not the values themselves
3. **Off-by-one errors:** Be careful with loop boundaries in brute force

## Time Complexity Breakdown

### Hash Map Approach
- **Best case:** O(1) - solution is the first two elements
- **Average case:** O(n) - need to check roughly half the array
- **Worst case:** O(n) - solution is at the end

### Space Complexity
- Hash map can store up to n-1 elements in worst case
- Each entry is a key-value pair (int -> int)

## Extensions

1. **Three Sum:** Find three numbers that sum to target (O(n²) with sorted array)
2. **K Sum:** Generalize to k numbers
3. **All Pairs:** Return all pairs instead of just one
4. **Without Hash Map:** Use two pointers on sorted array (O(n log n))

## Related Problems

- Three Sum
- Four Sum
- Two Sum II (sorted array)
- Two Sum III (data structure design)
- Two Sum IV (binary search tree)

## Optimization Notes

- For very small arrays, brute force might be faster due to no hash map overhead
- Consider using arrays instead of maps if the number range is small
- For multiple queries on same array, consider preprocessing

## Interview Tips

1. Always start with brute force to show you understand the problem
2. Explain the optimization: "We're doing repeated lookups, can we make them faster?"
3. Draw out an example to visualize the hash map approach
4. Discuss trade-offs between time and space
5. Test with edge cases: negatives, zeros, duplicates
