package main

func groupAnagrams(strs []string) [][]string {
	index := make(map[[26]int][]string, len(strs))
	for _, s := range strs {
		count := [26]int{}
		for _, c := range s {
			count[c-'a'] += 1
		}
		index[count] = append(index[count], s)
	}

	var result [][]string
	for _, s := range index {
		result = append(result, s)
	}

	return result
}
