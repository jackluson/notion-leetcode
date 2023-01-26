/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 *
 * https://leetcode.cn/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (45.43%)
 * Likes:    1786
 * Dislikes: 0
 * Total Accepted:    289.8K
 * Total Submissions: 638.7K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1], k = 2
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3], k = 3
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -1000 <= nums[i] <= 1000
 * -10^7 <= k <= 10^7
 *
 *
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	memo := map[int]int{0: 1}
	sum := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		key := sum - k
		count, has := memo[key]
		if has {
			ans += count
		}
		pre, preHas := memo[sum]
		if preHas {
			memo[sum] = pre + 1
		} else {
			memo[sum] = 1
		}
	}
	return ans
}

// @lc code=end

//[1,2,1,2,1,2] // 3
