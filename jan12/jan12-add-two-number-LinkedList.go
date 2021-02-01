package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
type singlelinkedlist struct {
	length int
	Head   *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	for l1.Next != nil || l2.Next != nil {
		fmt.Println("value", l1.Val, "valu2", l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}
	return result
}

func main() {
	var l1 *ListNode
	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	var l2 *ListNode
	l1 = &ListNode{
		Val: 7,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	result := addTwoNumbers(l1, l2)
	fmt.Println("Result:", result)
}
