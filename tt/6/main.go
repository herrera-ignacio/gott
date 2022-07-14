/**
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListNodeVal(n *ListNode) int {
	if n != nil {
		return n.Val
	} else {
		return 0
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	curr := head
	carry := 0
	var sum, l1Val, l2Val int

	for l1 != nil || l2 != nil || carry != 0 {
		// Get new val
		l1Val = getListNodeVal(l1)
		l2Val = getListNodeVal(l2)
		sum = l1Val + l2Val + carry

		// Append new node
		curr.Next = &ListNode{sum % 10, nil}
		// Update current pointer
		curr = curr.Next
		// Update carry
		carry = (sum / 10) % 10
		// Move lists pointers
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head.Next
}

