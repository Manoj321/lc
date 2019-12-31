package main

import "fmt"

func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := []rune{}
	matchingParen := map[rune]rune{'}': '{', ']': '[', ')': '('}

	for _, v := range s {
		n := len(stack)
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if n > 0 && stack[n-1] == matchingParen[v] {
			stack = stack[:n-1]
		} else {
			return false
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	tc := []string{"()", "{}[](", "{([]})", "{([])}"}
	for _, t := range tc {
		fmt.Printf("%v : %v\n", t, isValid(t))
	}
}
