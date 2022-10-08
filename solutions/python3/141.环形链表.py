'''
Desc:
File: /141.环形链表.py
File Created: Sunday, 9th October 2022 12:14:33 am
Author: luxuemin2108@gmail.com
-----
Copyright (c) 2022 Camel Lu
'''
#
# @lc app=leetcode.cn id=141 lang=python3
#
# [141] 环形链表
#
# https://leetcode.cn/problems/linked-list-cycle/description/
#
# algorithms
# Easy (51.52%)
# Likes:    1604
# Dislikes: 0
# Total Accepted:    831.2K
# Total Submissions: 1.6M
# Testcase Example:  '[3,2,0,-4]\n1'
#
# 给你一个链表的头节点 head ，判断链表中是否有环。
#
# 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos
# 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
#
# 如果链表中存在环 ，则返回 true 。 否则，返回 false 。
#
#
#
# 示例 1：
#
#
#
#
# 输入：head = [3,2,0,-4], pos = 1
# 输出：true
# 解释：链表中有一个环，其尾部连接到第二个节点。
#
#
# 示例 2：
#
#
#
#
# 输入：head = [1,2], pos = 0
# 输出：true
# 解释：链表中有一个环，其尾部连接到第一个节点。
#
#
# 示例 3：
#
#
#
#
# 输入：head = [1], pos = -1
# 输出：false
# 解释：链表中没有环。
#
#
#
#
# 提示：
#
#
# 链表中节点的数目范围是 [0, 10^4]
# -10^5 <= Node.val <= 10^5
# pos 为 -1 或者链表中的一个 有效索引 。
#
#
#
#
# 进阶：你能用 O(1)（即，常量）内存解决此问题吗？
#
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
from typing import Optional


class ListNode:
    def __init__(self, x, next=None):
        self.val = x
        self.next = next


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if head == None:
            return False
        fast = slow = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if fast == slow:
                return True
        return False
        # @lc code=end


node_4 = ListNode(4)
node_0 = ListNode(0, node_4)
node_2 = ListNode(2, node_0)
node_3 = ListNode(3, node_2)

node_4.next = node_2

res = Solution().hasCycle(node_3)
print('res:', res)
