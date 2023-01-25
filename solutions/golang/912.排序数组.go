/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 *
 * https://leetcode.cn/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (55.68%)
 * Likes:    762
 * Dislikes: 0
 * Total Accepted:    489.7K
 * Total Submissions: 906.2K
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -5 * 10^4 <= nums[i] <= 5 * 10^4
 *
 *
 */
package main

// @lc code=start
func sortArray(nums []int) []int {
	low, high := 0, len(nums)-1
	sortQuick(nums, low, high)
	return nums
}

func sortQuick(nums []int, low int, high int) {
	if high <= low {
		return
	}
	pivot := partition(nums, low, high)
	sortQuick(nums, low, pivot-1)
	sortQuick(nums, pivot+1, high)
}
func partition(nums []int, low int, high int) int {
	pivot := nums[high]
	left := low
	for i := low; i < high; i++ {
		if nums[i] < pivot {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[high], nums[left] = nums[left], nums[high]
	return left
}

// [6,2,3,1,5], [2,3,1], [6,5]

// @lc code=end
