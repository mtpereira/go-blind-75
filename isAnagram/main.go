package main

import (
	"fmt"
	"strings"
)

func isAnagram(s string, t string) bool {
	if s == t {
		return true
	}
	if len(s) != len(t) {
		return false
	}
	for _, c := range t {
		if !strings.ContainsRune(s, c) {
			return false
		}
		cString := fmt.Sprintf("%c", c)
		if strings.Count(s, cString) != strings.Count(t, cString) {
			return false
		}
	}
	return true
}
