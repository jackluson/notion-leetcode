/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 *
 * https://leetcode.cn/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (58.76%)
 * Likes:    985
 * Dislikes: 0
 * Total Accepted:    142.1K
 * Total Submissions: 241.3K
 * Testcase Example:  '10'
 *
 * 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
 *
 * 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 10
 * 输出：12
 * 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 * 解释：1 通常被视为丑数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */
package main

// @lc code=start
func nthUglyNumber(n int) int {
	dp := []int{1}
	cnt2 := 0
	cnt3 := 0
	cnt5 := 0
	for i := 1; i <= n; i++ {
		min := Min(Min(dp[cnt2]*2, dp[cnt3]*3), dp[cnt5]*5)
		if min == dp[cnt2]*2 {
			cnt2++
		}
		if min == dp[cnt3]*3 {
			cnt3++
		}
		if min == dp[cnt5]*5 {
			cnt5++
		}
		dp = append(dp, min)
	}
	return dp[n-1]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
