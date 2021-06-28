package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) print() {
	fmt.Print(l.Val)
	if l.Next != nil {
		l.Next.print()
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var result *ListNode
	if l1.Val < l2.Val {
		result = l1
		result.Next = mergeTwoLists(l1.Next, l2)
	} else {
		result = l2
		result.Next = mergeTwoLists(l1, l2.Next)
	}

	return result
}

func main() {
	fmt.Println("Test 1")
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}

	result := mergeTwoLists(l1, l2)
	result.print()
	fmt.Println()
}
