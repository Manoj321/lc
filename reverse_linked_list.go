package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) append(v int) {
	if l.Head == nil {
		l.Head = &ListNode{v, nil}
		return
	}
	n := l.Head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &ListNode{v, nil}
}

func (l *LinkedList) print() {
	n := l.Head
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
func (l *LinkedList) reverse() *LinkedList {
	head := l.Head

	newHead := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return &LinkedList{newHead}
}

func main() {
	ll := LinkedList{}
	ll.append(1)
	ll.append(10)
	ll.append(12)
	ll.print()
	newll := ll.reverse()
	newll.print()elf

}
