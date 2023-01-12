#
# @lc app=leetcode.cn id=263 lang=python3
#
# [263] 丑数
#
# https://leetcode.cn/problems/ugly-number/description/
#
# algorithms
# Easy (51.05%)
# Likes:    352
# Dislikes: 0
# Total Accepted:    138.7K
# Total Submissions: 272.2K
# Testcase Example:  '6'
#
# 丑数 就是只包含质因数 2、3 和 5 的正整数。
#
# 给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
#
#
#
# 示例 1：
#
#
# 输入：n = 6
# 输出：true
# 解释：6 = 2 × 3
#
# 示例 2：
#
#
# 输入：n = 1
# 输出：true
# 解释：1 没有质因数，因此它的全部质因数是 {2, 3, 5} 的空集。习惯上将其视作第一个丑数。
#
# 示例 3：
#
#
# 输入：n = 14
# 输出：false
# 解释：14 不是丑数，因为它包含了另外一个质因数 7 。
#
#
#
#
# 提示：
#
#
# -2^31 <= n <= 2^31 - 1
#
#
#

# @lc code=start


class Solution:
    def isUgly(self, n: int) -> bool:
        if n == 0:
            return False
        while n != 1:
            if n % 2 == 0:
                n = n / 2
            elif n % 3 == 0:
                n = n / 3
            elif n % 5 == 0:
                n = n / 5
            else:
                return False
        return True

# @lc code=end
