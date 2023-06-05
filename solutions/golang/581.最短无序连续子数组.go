/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 *
 * https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/description/
 *
 * algorithms
 * Medium (41.39%)
 * Likes:    1064
 * Dislikes: 0
 * Total Accepted:    179.2K
 * Total Submissions: 429.1K
 * Testcase Example:  '[2,6,4,8,10,9,15]'
 *
 * 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
 *
 * 请你找出符合题意的 最短 子数组，并输出它的长度。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,6,4,8,10,9,15]
 * 输出：5
 * 解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,4]
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^5
 *
 *
 *
 *
 * 进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？
 *
 *
 *
 */
package main

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	start := 0
	min := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			start = i
		}
	}
	end := len(nums) - 1
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
	}
	if min == nums[start] && max == nums[end] {
		return 0
	}
	return end - start + 1
}

// @lc code=end
