package main

import (
	"slices"
)

// twoSum implements a solution for
// https://leetcode.com/problems/two-sum/
// For each number, calculate the difference between it and the target,
// then look for the difference in nums, and return index of both.
func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		d := target - n
		if j := slices.Index[[]int, int](nums, d); j != -1 && j != i {
			return []int{i, j}
		}
	}
	return nil
}
