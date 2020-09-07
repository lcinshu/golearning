package Linkedlist

import "fmt"

func MainOut1() {
	l1 := buildListNode([]int{1, 2, 3, 4, 5, 6})
	l2 := reverseList1(l1)
	fmt.Print(l2)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode = nil
	for ; head != nil; head.Next, pre, head = pre, head, head.Next {
	}
	return pre
}
