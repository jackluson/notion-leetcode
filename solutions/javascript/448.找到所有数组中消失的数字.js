/*
 * @lc app=leetcode.cn id=448 lang=javascript
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (65.84%)
 * Likes:    1106
 * Dislikes: 0
 * Total Accepted:    239.8K
 * Total Submissions: 364.3K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums
 * 中的数字，并以数组的形式返回结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,3,2,7,8,2,3,1]
 * 输出：[5,6]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,1]
 * 输出：[2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1
 * 1
 *
 *
 * 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
 *
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findDisappearedNumbers = function (nums) {
  for (const num of nums) {
    // Now we choose an index in the nums array based on the value we're currently looking at.
    // We do have to use Math.abs on it first, in case we've already visited this index and made it negative.
    // Then, since arrays are 0 indexed, subtract 1 from it
    const indexBasedOnThisValue = Math.abs(num) - 1;
    if (nums[indexBasedOnThisValue] > 0) {
      nums[indexBasedOnThisValue] = nums[indexBasedOnThisValue] * -1;
    }
  }
  // Now that we've marked the array with negative numbers, loop through it again.
  // This time, we'll be building our result
  const result = [];
  console.log('nums', nums);
  for (let i = 0; i < nums.length; i++) {
    // Check if the number at this position is positive or negative.
    // It doesn't matter what the number is necessarily, we just want to use the index of this value to check what we visited.
    // And again, since arrays are 0-indexed, we'll be off by one, so add 1 when we push to results.
    if (nums[i] >= 0) {
      result.push(i + 1);
    }
  }

  return result;
};
// @lc code=end
