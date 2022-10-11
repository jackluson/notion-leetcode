/*
 * @lc app=leetcode.cn id=264 lang=javascript
 *
 * [264] 丑数 II
 *
 * https://leetcode.cn/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (58.76%)
 * Likes:    985
 * Dislikes: 0
 * Total Accepted:    142.1K
 * Total Submissions: 241.3K
 * Testcase Example:  '10'
 *
 * 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
 *
 * 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 10
 * 输出：12
 * 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 * 解释：1 通常被视为丑数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

// @lc code=start
/**
 * @param {number} n
 * @return {number}
 */
var nthUglyNumber = function (n) {
  var dp = [1];
  let p2 = 0,
    p3 = 0,
    p5 = 0;
  for (let i = 1; i < n; i++) {
    const cur = Math.min(2 * dp[p2], 3 * dp[p3], 5 * dp[p5]);
    if (cur === 2 * dp[p2]) {
      p2++;
    }
    if (cur === 3 * dp[p3]) {
      p3++;
    }
    if (cur === 5 * dp[p5]) {
      p5++;
    }
    dp[i] = cur;
  }
  return dp[n - 1];
};
// @lc code=end
