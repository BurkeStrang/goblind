package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func sort2DSlice(s [][]int) {
	for _, v := range s {
		sort.Ints(v)
	}
	sort.Slice(s, func(i, j int) bool {
		for x := range s[i] {
			if x >= len(s[j]) {
				return false
			}
			if s[i][x] != s[j][x] {
				return s[i][x] < s[j][x]
			}
		}
		return len(s[i]) < len(s[j])
	})
}

func TestCombinationSumII(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}
	for _, tt := range tests {
		got := CombinationSumII(tt.candidates, tt.target)
		sort2DSlice(got)
		sort2DSlice(tt.expected)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("CombinationSumII(%v, %d) = %v; want %v", tt.candidates, tt.target, got, tt.expected)
		}
	}
}
