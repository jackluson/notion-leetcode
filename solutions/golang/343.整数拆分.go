/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 *
 * https://leetcode.cn/problems/integer-break/description/
 *
 * algorithms
 * Medium (62.11%)
 * Likes:    942
 * Dislikes: 0
 * Total Accepted:    206K
 * Total Submissions: 331.8K
 * Testcase Example:  '2'
 *
 * 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
 *
 * 返回 你可以获得的最大乘积 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = 2
 * 输出: 1
 * 解释: 2 = 1 + 1, 1 × 1 = 1。
 *
 * 示例 2:
 *
 *
 * 输入: n = 10
 * 输出: 36
 * 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 *
 *
 *
 * 提示:
 *
 *
 * 2 <= n <= 58
 *
 *
 */
// @lc code=start

package main

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func integerBreak(n int) int {
	dp := []int{0, 0, 1, 2, 4}
	for i := 5; i <= n; i++ {
		maxNum := max(2*max(i-2, dp[i-2]), 3*max(i-3, dp[i-3]))
		dp = append(dp, maxNum)
	}
	return dp[n]
}

// @lc code=end
