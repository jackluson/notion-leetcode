/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 *
 * https://leetcode.cn/problems/happy-number/description/
 *
 * algorithms
 * Easy (63.16%)
 * Likes:    1191
 * Dislikes: 0
 * Total Accepted:    340.4K
 * Total Submissions: 538.1K
 * Testcase Example:  '19'
 *
 * 编写一个算法来判断一个数 n 是不是快乐数。
 *
 * 「快乐数」 定义为：
 *
 *
 * 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
 * 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
 * 如果这个过程 结果为 1，那么这个数就是快乐数。
 *
 *
 * 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 19
 * 输出：true
 * 解释：
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 2
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 2^31 - 1
 *
 *
 */
package main

// @lc code=start
func getNext(k int) int {
	sum := 0
	for k > 0 {
		mod := k % 10
		k = k / 10
		sum += mod * mod
	}
	return sum
}
func isHappy(n int) bool {
	// 2 > 4 > 16 > 1 + 36 = 37 > 9 + 49 = 58
	// 25 + 64 = 81 > 64 + 1 = 65 > 36 + 35 = 71
	// 49 + 1 = 50 > 25 > 29 > 8
	memoMap := make(map[int]bool)
	memoMap[1] = true
	next := n
	_, isExist := memoMap[next]
	for isExist == false {
		next = getNext(next)
		_, isExist = memoMap[next]
		if isExist == false {
			memoMap[next] = true
		}
	}
	if next == 1 {
		return true
	}
	return false
}

// @lc code=end
