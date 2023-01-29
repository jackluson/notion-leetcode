/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode.cn/problems/next-permutation/description/
 *
 * algorithms
 * Medium (37.93%)
 * Likes:    2042
 * Dislikes: 0
 * Total Accepted:    394.9K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3]'
 *
 * 整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。
 *
 *
 * 例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
 *
 *
 * 整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列
 * 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
 *
 *
 * 例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
 * 类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
 * 而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
 *
 *
 * 给你一个整数数组 nums ，找出 nums 的下一个排列。
 *
 * 必须 原地 修改，只允许使用额外常数空间。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1]
 * 输出：[1,2,3]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,1,5]
 * 输出：[1,5,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 100
 *
 *
 */
package main

// @lc code=start
func nextPermutation(nums []int) {
	lgt := len(nums)
	if lgt <= 1 {
		return
	}
	i, j, k := lgt-2, lgt-1, lgt-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// for j := i + 1; j < len(nums); j++ {
	// 	for k := len(nums) - 1; k >= j+1; k-- {
	// 		if nums[k] < nums[k-1] {
	// 			nums[k], nums[k-1] = nums[k-1], nums[k]
	// 		}
	// 	}
	// }
	for i, j := j, lgt-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i = i + 1
		j = j - 1
	}
}

// func main() {
// 	arr := []int{1, 1}
// 	fmt.Println(arr)
// 	nextPermutation(arr)
// 	fmt.Println(arr)
// }

// @lc code=end
// [4,5,8,7,1] --> [4,7,8,5,1] --> [4,7,1,5,8]
// [1,3,2] --> [2,3,1] --> [2,1,3]
