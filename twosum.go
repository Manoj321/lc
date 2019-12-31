package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) == 1 {
		return nums
	}

	m := map[int]int{}
	for i, n := range nums {
		m[n] = i
	}

	for i := 0; i < len(nums); i++ {
		b := target - nums[i]
		j, ok := m[b]
		if !ok {
			continue
		}
		if i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
	}

	for _, tc := range testCases {
		res := twoSum(tc.nums, tc.target)
		fmt.Printf(" %v, %v, Got: %v, Expected: %v",
			tc.nums, tc.target, res, tc.output)
	}
}
