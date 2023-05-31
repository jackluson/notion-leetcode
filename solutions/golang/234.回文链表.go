/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode.cn/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (52.40%)
 * Likes:    1697
 * Dislikes: 0
 * Total Accepted:    594.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		tempNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = tempNext
	}
	return pre
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	slow = reverseList(slow)
	headNode := head
	tailNode := slow
	for tailNode != nil {
		if headNode.Val != tailNode.Val {
			return false
		}
		tailNode = tailNode.Next
		headNode = headNode.Next
	}
	return true
}

// @lc code=end

