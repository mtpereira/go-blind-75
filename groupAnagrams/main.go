package main

import (
	"sort"
)

type word struct {
	count int
	index []int
}

func (w *word) seen(index int) {
	w.count++
	w.index = append(w.index, index)
}

func (w *word) fetch(strs []string) []string {
	var result []string
	for _, i := range w.index {
		result = append(result, strs[i])
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	index := make(map[string]*word, len(strs))
	for i, s := range strs {
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		s = string(r)
		if index[s] == nil {
			index[s] = &word{count: 0, index: []int{}}
		}
		index[s].seen(i)
	}

	var result [][]string
	for _, w := range index {
		result = append(result, w.fetch(strs))
	}

	return result
}
