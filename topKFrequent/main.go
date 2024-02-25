package main

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int, len(nums))
	for _, n := range nums {
		counter[n]++
	}

	maxes := make(map[int][]int, len(nums))
	for n, c := range counter {
		maxes[c] = append(maxes[c], n)
	}

	var result []int
	for i := len(nums); len(result) < k; i-- {
		if maxes[i] == nil {
			continue
		}
		result = append(result, maxes[i]...)
		if len(result) >= k {
			return result[:k]
		}
	}

	return result
}
