#
# @lc app=leetcode.cn id=264 lang=python
#
# [264] 丑数 II
#
# https://leetcode.cn/problems/ugly-number-ii/description/
#
# algorithms
# Medium (58.76%)
# Likes:    985
# Dislikes: 0
# Total Accepted:    142.1K
# Total Submissions: 241.3K
# Testcase Example:  '10'
#
# 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
# 
# 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
# 
# 
# 
# 示例 1：
# 
# 
# 输入：n = 10
# 输出：12
# 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
# 
# 
# 示例 2：
# 
# 
# 输入：n = 1
# 输出：1
# 解释：1 通常被视为丑数。
# 
# 
# 
# 
# 提示：
# 
# 
# 1 
# 
# 
#
import heapq
# @lc code=start
class Solution(object):
    def nthUglyNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        cur = 1
        seen = []
        minHeap = [cur]
        heapq.heappush(minHeap, 1)
        factors = [2, 3, 5]
        for i in range(n+1):
            cur = heapq.heappop(minHeap)
            for factor in factors:
                temp = factor * cur
                if temp not in seen:
                    heapq.heappush(minHeap, temp)
                    seen.append(temp)
        return cur
# @lc code=end
