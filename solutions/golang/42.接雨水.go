/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode.cn/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (61.94%)
 * Likes:    4097
 * Dislikes: 0
 * Total Accepted:    632.3K
 * Total Submissions: 1M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */
package main

// @lc code=start
func trap(height []int) int {
	lenght := len(height)
	if lenght <= 2 {
		return 0
	}
	lMax := height[0]
	rMax := height[lenght-1]
	left := 0
	right := lenght - 1
	sum := 0
	for left <= right {
		if lMax > rMax {
			if height[right] > rMax {
				rMax = height[right]
			} else {
				sum = sum + rMax - height[right]
			}
			right--
		} else {
			if height[left] > lMax {
				lMax = height[left]
			} else {
				sum = sum + lMax - height[left]
			}
			left++
		}
	}
	return sum
}

// func main() {
// 	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
// 	trap(height)
// }

// @lc code=end
