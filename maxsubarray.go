package main

import "fmt"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func find_max_crossing_sub_array(A []int, low, mid, high int) (int, int, int) {
	left_sum := MinInt
	sum := 0
	var max_left int
	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > left_sum {
			left_sum = sum
			max_left = i
		}
	}
	sum = 0
	right_sum := MinInt
	var max_right int
	for j := mid + 1; j <= high; j++ {
		sum += A[j]
		if sum > right_sum {
			right_sum = sum
			max_right = j
		}
	}
	return max_left, max_right, left_sum + right_sum
}

func find_max_sub_array(A []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, A[low]
	}
	mid := int((low + high) / 2)
	left_low, left_high, left_sum := find_max_sub_array(A, low, mid)
	right_low, right_high, right_sum := find_max_sub_array(A, mid+1, high)
	cross_low, cross_high, cross_sum := find_max_crossing_sub_array(A, low, mid, high)
	if left_sum > right_sum && left_sum > cross_sum {
		return left_low, left_high, left_sum
	}
	if right_sum > left_sum && right_sum > cross_sum {
		return right_low, right_high, right_sum
	}
	return cross_low, cross_high, cross_sum
}
func main() {
	// A := []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97}
	// B := []int{}
	// for i := 1; i < len(A); i++ {
	// 	B = append(B, A[i]-A[i-1])
	// }
	// fmt.Println(B)
	B := []int{-1, -1, -2, -2}
	fmt.Println(find_max_sub_array(B, 0, len(B)-1))
}
