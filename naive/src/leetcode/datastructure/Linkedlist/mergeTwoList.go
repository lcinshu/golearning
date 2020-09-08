package Linkedlist

import "fmt"

func MainOut2() {
	l1 := buildListNode([]int{1, 3, 5, 7})
	l2 := buildListNode([]int{2, 4, 6, 8})
	l3 := mergeTwoLists(l1, l2)
	fmt.Print(l3)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// 1. 建立头结点用于最后返回，建立哨兵节点用于迭代
	// 头结点独立于当前节点链表之外，为单独内存。如果考虑直接使用l1 或者l2的头结点，则需要先判断l1,l2是否为空
	// 代码不灵活
	head := &ListNode{}
	cur := head

	// 2. 节点均不为空时
	// 2.1 比较两节点值大小
	// 2.2 将较小节点加到cur所在节点之后
	// 2.3 较小节点后移
	// 2.4 当前节点后移
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	// 3. 链表先跑完，则把未跑完节点添加到当前节点之后即可，后续均有序
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}

	// 4. 返回头结点下一节点
	return head.Next
}

func mergeTwoListsDiGui(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoListsDiGui(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsDiGui(l1, l2.Next)
		return l2
	}
}
