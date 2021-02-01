package linkedlist

import "fmt"

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
// 示例 1:
//
// 输入: 1->1->2
//输出: 1->2
//
//
// 示例 2:
//
// 输入: 1->1->2->3->3
//输出: 1->2->3
// Related Topics 链表
// 👍 389 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteRepetNode() {
	l1 := buildListNode([]int{1, 1, 1})
	l2 := deleteDuplicatesCurAndNext(l1)
	fmt.Print(l2)
}

/**
 关于链表的技巧：
1. 如果有删除节点的操作时，视情况选择指正组合。这里需要删除重复节点，
*/

/**
如果节点值相等，不需要移动head节点，只需要修改当前节点后继节点的指向
可以想象成，当前节点不动，后续节点往前缩减

head = head.next 这里如果不放else里面，会提前后移head节点，导致head.next = nil
导致最后一个元素无法进行判断迭代，比如{1,1,1}，最后结果为：{1,1}

*/
func deleteDuplicatesOrigin(head *ListNode) *ListNode {
	headTmp := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return headTmp
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 1. 如果节点为空，直接返回头结点
	if head == nil {
		return head
		// return nil
	}
	// 2. 通过map中的key校验元素是否重复
	maps := make(map[int]int)

	// 3. 新增两个节点，headPre节点用于删除重复节点时
	headPre := &ListNode{-1, nil}
	headPreTmp := headPre
	headPre.Next = head
	// 4. 通过map中key判断当前节点是否已存在，如果不存在则存入map中，并赋值1，同时修改pre和cur指向下一节点
	//    如果存在，则修改pre指针指向cur.next，仅修改cur节点指向下一节点
	for head != nil {
		if 0 == maps[head.Val] {
			maps[head.Val] = 1
			headPre = headPre.Next
		} else {
			headPre.Next = headPre.Next.Next
		}
		head = head.Next
	}
	return headPreTmp.Next
}

func deleteDuplicatesCurAndNext(head *ListNode) *ListNode {
	// 1. 如果节点为空，直接返回头结点
	if head == nil {
		return head
		// return nil
	}
	// 2. 通过map中的key校验元素是否重复
	maps := make(map[int]int)
	maps[head.Val] = 1
	// 3. 新增两个节点，headPre节点用于删除重复节点时
	cur := head

	// 4. 通过map中key判断当前节点是否已存在，如果不存在则存入map中，并赋值1，同时修改pre和cur指向下一节点
	//    如果存在，则修改pre指针指向cur.next，仅修改cur节点指向下一节点
	for cur != nil && cur.Next != nil {
		if 0 != maps[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			maps[head.Next.Val] = 1
			cur = cur.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
