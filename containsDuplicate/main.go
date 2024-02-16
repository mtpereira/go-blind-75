package main

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	previous := map[int]bool{}
	for _, num := range nums {
		if previous[num] == true {
			return true
		}
		previous[num] = true
	}
	return false
}
