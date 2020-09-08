package Linkedlist

import (
	"fmt"
)

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//
//
// 说明：
//
// 给定的 n 保证是有效的。
//
// 进阶：
//
// 你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针
// 👍 978 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func RemoveNthFromEnd() {
	head := buildListNode([]int{1, 2})
	result := removeNthFromEnd(head, 2)
	fmt.Print(result)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1. 定义双指针，让其从 head.pre 开始进行迭代
	pre := &ListNode{}
	pre.Next = head
	start, end := pre, pre
	// 2. before指针走到第N个节点
	for n != 0 {
		start = start.Next
		n--
	}
	// 3. end指针开始后移，直到 start 指针为空
	for start.Next != nil {
		end = end.Next
		start = start.Next
	}
	end.Next = end.Next.Next
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
