package main

import (
	"fmt"
)

func main() {
	a := []int{1, 4, 7, 2, 3, 9, 0}
	fmt.Println(MergeSort(a))
}

func MergeSort(a []int) []int {

	if len(a) == 1 {
		return a
	}
	mid := len(a) / 2
	return Merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

func Merge(left, right []int) []int {
	size := len(left) + len(right)
	slice := make([]int, size)
	l := 0
	r := 0
	i := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			slice[i] = left[l]
			l += 1
		} else {
			slice[i] = right[r]
			r += 1
		}
		i++
	}
	for ; l < len(left); l++ {
		slice[i] = left[l]
		i++
	}
	for ; r < len(right); r++ {
		slice[i] = right[r]
		i++
	}
	return slice
}
