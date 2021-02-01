package linkedlist

import (
	"fmt"
	_ "fmt"
)

//编写一个程序，找到两个单链表相交的起始节点。
//
// 如下面的两个链表：
//
//
//
// 在节点 c1 开始相交。
//
//
//
// 示例 1：
//
//
//
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, s
//kipB = 3
//输出：Reference of the node with value = 8
//输入解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1
//,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
//
//
//
//
// 示例 2：
//
//
//
// 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB =
// 1
//输出：Reference of the node with value = 2
//输入解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4
//]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
//
//
//
//
// 示例 3：
//
//
//
// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
//输出：null
//输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而
// skipA 和 skipB 可以是任意值。
//解释：这两个链表不相交，因此返回 null。
//
//
//
//
// 注意：
//
//
// 如果两个链表没有交点，返回 null.
// 在返回结果后，两个链表仍须保持原有的结构。
// 可假定整个链表结构中没有循环。
// 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
//
// Related Topics 链表
// 👍 798 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
需要特殊说明的是：
相同节点的判断完全依赖于节点在内存中的分配方式
如果是相同值，但是通过new方式新建的，还是不同的节点

比如如下示例中，4, 1, 8, 4, 5  和 5, 6, 1, 8, 4, 5 ，但从节点值上来看，1节点是相交节点，但实际上8节点才是相交节点
*/

func MainOut() {
	l1 := buildListNode([]int{4, 1, 8, 4, 5})
	l2 := buildListNode([]int{5, 6, 1, 8, 4, 5})

	l3 := getIntersectionNode(l1, l2)
	fmt.Print(l3)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	l1 := headA
	l2 := headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1
}

func buildListNode(x []int) *ListNode {
	l1 := &ListNode{x[0], nil}
	head := l1
	for i := 1; i < len(x); i++ {
		head.Next = &ListNode{x[i], nil}
		head = head.Next
	}
	return l1
}

//leetcode submit region end(Prohibit modification and deletion)
