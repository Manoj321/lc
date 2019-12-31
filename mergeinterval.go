package main

import (
	"fmt"
	"sort"
)

type intervals [][]int

func (in intervals) Len() int { return len(in) }
func (in intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
func (in intervals) Less(i, j int) bool {
	return (in[i][0] <= in[j][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(iv intervals) [][]int {
	sort.Sort(iv)

	var newIntervals [][]int
	if len(iv) == 1 {
		return iv
	}

	newIntervals = append(newIntervals, iv[0])

	for i, j := 1, 0; i < len(iv); i++ {
		if iv[i][0] <= newIntervals[j][1] {
			newIntervals[j][1] = max(newIntervals[j][1], iv[i][1])
		} else {
			newIntervals = append(newIntervals, iv[i])
			j++
		}
	}
	return newIntervals
}

func main() {
	testCase := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {0, 0}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {2, 3}},
		{{1, 4}, {0, 2}, {3, 5}},
		{{5, 5}, {1, 3}, {3, 5}, {4, 6}, {1, 1}, {3, 3}, {5, 6}, {3, 3}, {2, 4}, {0, 0}},
		{{0, 0}, {1, 1}, {1, 6}},
	}
	for _, tc := range testCase {
		fmt.Printf("%v -- %v\n", tc, merge(tc))
	}

}
