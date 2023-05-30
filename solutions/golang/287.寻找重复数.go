/*
 * @lc app=leetcode.cn id=287 lang=golang
 * @lcpr version=21902
 *
 * [287] 寻找重复数
 */
package main

// @lc code=start
func findDuplicate(nums []int) int {

	// for index, num := range nums {
	// 	for i := index + 1; i < len(nums); i++ {
	// 		if num == nums[i] {
	// 			return num
	// 		}
	// 	}
	// }
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			cur := 0
			for nums[cur] != nums[slow] {
				cur = nums[cur]
				slow = nums[slow]
			}
			return nums[cur]
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,4,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,1,3,4,2]\n
// @lcpr case=end

*/
