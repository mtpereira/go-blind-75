package main

import (
	"reflect"
	"slices"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "ex1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "ex2",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "ex3",
			args: args{strs: []string{"a"}},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			for _, s := range got {
				slices.Sort(s)
			}
			slices.SortFunc(got, func(a, b []string) int {
				return slices.Compare(a, b)
			})
			for _, s := range tt.want {
				slices.Sort(s)
			}
			slices.SortFunc(tt.want, func(a, b []string) int {
				return slices.Compare(a, b)
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
