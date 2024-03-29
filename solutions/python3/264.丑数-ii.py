#
# @lc app=leetcode.cn id=264 lang=python3
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

# @lc code=start
class Solution:
    def nthUglyNumber(self, n: int) -> int:
        p2 = 1
        p3 = 1
        p5 = 1
        dp = [1]
        for index in range(1, n):
            cur_min = min(2 * dp[p2-1], 3 *dp[p3-1], 5 * dp[p5-1])
            if cur_min == dp[p2-1] * 2:
                p2 = p2+1
            if cur_min == dp[p3-1] * 3:
                p3 = p3+1
            if cur_min == dp[p5-1] * 5:
                p5 = p5+1
            dp.append(cur_min)
        return dp[n-1]
# @lc code=end

