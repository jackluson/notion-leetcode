#
# @lc app=leetcode.cn id=144 lang=python3
#
# [144] 二叉树的前序遍历
#
# https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
#
# algorithms
# Easy (71.27%)
# Likes:    923
# Dislikes: 0
# Total Accepted:    723.8K
# Total Submissions: 1M
# Testcase Example:  '[1,null,2,3]'
#
# 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
#
#
#
# 示例 1：
#
#
# 输入：root = [1,null,2,3]
# 输出：[1,2,3]
#
#
# 示例 2：
#
#
# 输入：root = []
# 输出：[]
#
#
# 示例 3：
#
#
# 输入：root = [1]
# 输出：[1]
#
#
# 示例 4：
#
#
# 输入：root = [1,2]
# 输出：[1,2]
#
#
# 示例 5：
#
#
# 输入：root = [1,null,2]
# 输出：[1,2]
#
#
#
#
# 提示：
#
#
# 树中节点数目在范围 [0, 100] 内
# -100
#
#
#
#
# 进阶：递归算法很简单，你可以通过迭代算法完成吗？
#
#

# @lc code=start
# Definition for a binary tree node.
from typing import Optional, List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        if root == None:
            return []
        res = []
        stack = []
        # stack.append(root)
        cur = root
        while len(stack) or cur:
            if cur:
                res.append(cur.val)
                if cur.right:
                    stack.append(cur.right)
                if cur.left:
                    cur = cur.left
                else:
                    cur = None
            else:
                cur = stack.pop()

        return res
# @lc code=end
