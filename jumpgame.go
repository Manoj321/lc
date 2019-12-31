package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	max := nums[0]
	for i := 0; i <= max; i++ {
		if max >= len(nums)-1 {
			return true
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return false
}
func main() {
	A := []int{0, 0, 8, 2, 0, 0, 1}
	fmt.Println(canJump(A))
}
