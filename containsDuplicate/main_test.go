package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		desc string
		input []int
		output bool
	}{
		{
			desc: "ex1",
			input: []int{1, 2, 3, 1},
			output: true,
		},
		{
			desc: "ex2",
			input: []int{1, 2, 3, 4},
			output: false,
		},
		{
			desc: "ex3",
			input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			output: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := containsDuplicate(tC.input); result != tC.output {
				t.Fatalf("containsDuplicate(%v) returned %v", tC.input, result)
			}
		})
	}
}
