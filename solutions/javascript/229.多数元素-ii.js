/*
 * @lc app=leetcode.cn id=229 lang=javascript
 *
 * [229] 多数元素 II
 *
 * https://leetcode.cn/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (53.77%)
 * Likes:    658
 * Dislikes: 0
 * Total Accepted:    91.4K
 * Total Submissions: 169.3K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,2,3]
 * 输出：[3]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1]
 * 输出：[1]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,2]
 * 输出：[1,2]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 *
 *
 * 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
 *
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var majorityElement = function (nums) {
  let candidate1;
  let candidate2;
  let count1 = 0;
  let count2 = 0;
  for (let i = 0; i < nums.length; i++) {
    if (candidate1 === nums[i]) {
      count1++;
    } else if (candidate2 === nums[i]) {
      count2++;
    } else if (count2 === 0) {
      candidate2 = nums[i];
      count2 = 1;
    } else if (count1 === 0) {
      candidate1 = nums[i];
      count1 = 1;
    } else {
      count1--;
      count2--;
    }
  }
  count1 = count2 = 0;
  for (let i = 0; i < nums.length; i++) {
    if (candidate1 === nums[i]) {
      count1++;
    } else if (candidate2 === nums[i]) {
      count2++;
    }
  }
  let candidates = [];
  if (count1 > nums.length / 3) {
    candidates.push(candidate1);
  }
  if (count2 > nums.length / 3) {
    candidates.push(candidate2);
  }
  return candidates;
};
// @lc code=end
