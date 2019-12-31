package main

import (
	"fmt"
)

func main() {
	var p []int
	q := []int{}
	fmt.Println(len(p), len(q))
	fmt.Println(cap(p), cap(q))
}
