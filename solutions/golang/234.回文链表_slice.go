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
package main

func isPalindrome(head *ListNode) bool {
	// var str1 string
	node := head
	// for node != nil {
	// 	str1 = str1 + strconv.Itoa(node.Val)
	// 	node = node.Next
	// }
	// start, end := 0, len(str1)-1
	// for start < end {
	// 	if str1[start] != str1[end] {
	// 		return false
	// 	}
	// 	start++
	// 	end--
	// }

	var nums = make([]int, 0)
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	start, end := 0, len(nums)-1
	for start < end {
		if nums[start] != nums[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// @lc code=end
