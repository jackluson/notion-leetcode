/*
 * @lc app=leetcode.cn id=204 lang=javascript
 *
 * [204] 计数质数
 *
 * https://leetcode.cn/problems/count-primes/description/
 *
 * algorithms
 * Medium (37.52%)
 * Likes:    960
 * Dislikes: 0
 * Total Accepted:    227.6K
 * Total Submissions: 607.6K
 * Testcase Example:  '10'
 *
 * 给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 10
 * 输出：4
 * 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 5 * 10^6
 *
 *
 */

// @lc code=start

const isPrime = (n) => {
  if (n <= 3) {
    return n > 1;
  }
  if (n % 6 !== 1 && n % 6 !== 5) {
    return false;
  }
  const sqrt_n = Math.sqrt(n);
  // 初始值是5 -- 6x+1 和 6x+5 (即等同于6x-1)
  for (let i = 5; i <= sqrt_n; i = i + 6) {
    if (n % i === 0 || n % (i + 2) === 0) {
      return false;
    }
  }
  return true;
};
/**
 * @param {number} n
 * @return {number}
 */
var countPrimes = function (n) {
  // const primes = new Array(n).fill(0);
  let count = 0;
  for (let i = 2; i < n; i++) {
    if (isPrime(i)) {
      count++;
    }
  }
  return count;
};
// @lc code=end
