package main

// twoSum implements a solution for
// https://leetcode.com/problems/two-sum/
// For each number, calculate the difference between it and the target,
// then look for the difference in nums, and return index of both.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		d := target - n
		if j, ok := seen[d]; ok && j != i {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
