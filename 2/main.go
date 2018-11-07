package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	head := &result
	overflow := 0
	for {
		temp := l1.Val + l2.Val + overflow

		overflow = temp / 10
		if overflow == 1 {
			temp = temp - 10
		}
		if l1.Next == nil && l2.Next == nil && overflow == 0 {
			head.Val = temp
			return &result
		}
		head.Val = temp
		head.Next = &ListNode{}
		head = head.Next
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}
		l1 = l1.Next
		if l2.Next == nil {
			l2.Next = &ListNode{}
		}
		l2 = l2.Next
	}
	return &ListNode{}
}

func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	addTwoNumbers(&l1, &l2)
}
