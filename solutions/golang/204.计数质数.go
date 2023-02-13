/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode.cn/problems/count-primes/description/
 *
 * algorithms
 * Medium (37.52%)
 * Likes:    960
 * Dislikes: 0
 * Total Accepted:    227.7K
 * Total Submissions: 608K
 * Testcase Example:  '10'
 *
 * 给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 10
 * 输出：4
 * 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 5 * 10^6
 *
 *
 */

package main

// @lc code=start
//
//	func isPrime(n int) bool {
//		for i := 2; i*i <= n; i++ {
//			mod := n % i
//			if mod == 0 {
//				return false
//			}
//		}
//		return true
//	}
func countPrimes(n int) int {
	// var s = make([]int, n)
	if n <= 1 {
		return 0
	}
	res := make([]bool, n)
	cnt := 0
	res[0] = true
	res[1] = true
	for i := 2; i < n; i++ {
		for j := 2; i*j < n; j++ {
			res[i*j] = true
		}
	}
	for _, val := range res {
		if val == false {
			cnt++
		}
	}
	return cnt
}

// @lc code=end
