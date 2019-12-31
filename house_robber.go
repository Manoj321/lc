package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	a, b := 0, 0
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, v+b)
		}
	}
	return max(a, b)
}
func main() {
	tc := [][]int{
		{1, 2, 9, 1, 2},
		{2, 7, 9, 3, 1},
		{1, 2, 3, 1},
		{2, 1, 1, 2},
	}
	for _, t := range tc {
		fmt.Printf("%v,%v", t, rob(t))
	}
}
