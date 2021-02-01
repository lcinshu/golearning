package linkedlist

import "fmt"

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例:
//
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
//
// Related Topics 链表
// 👍 613 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs() {
	list := buildListNode([]int{1, 2, 3, 4})
	result := swapPairs(list)
	fmt.Print(result)
}
func swapPairs(head *ListNode) *ListNode {
	node := &ListNode{}
	node.Next = head
	pre := node

	for pre.Next != nil && pre.Next.Next != nil {
		l1, l2 := pre.Next, pre.Next.Next
		next := l2.Next
		l2.Next = l1
		l1.Next = next
		pre.Next = l2

		pre = l1
	}
	return node.Next

}

/**
递归由于是两个节点一组，可以将3，4，……一(多)组节点看做是已完成交换完成的整体
把 1，2两个节点
*/

func swapPairsDg(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := head.Next

	firstNode.Next = swapPairsDg(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

//leetcode submit region end(Prohibit modification and deletion)
