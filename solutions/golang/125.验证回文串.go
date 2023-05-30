/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode.cn/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.85%)
 * Likes:    639
 * Dislikes: 0
 * Total Accepted:    471.3K
 * Total Submissions: 1M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
 *
 * 字母和数字都属于字母数字字符。
 *
 * 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: s = "A man, a plan, a canal: Panama"
 * 输出：true
 * 解释："amanaplanacanalpanama" 是回文串。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "race a car"
 * 输出：false
 * 解释："raceacar" 不是回文串。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = " "
 * 输出：true
 * 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
 * 由于空字符串正着反着读都一样，所以是回文串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s 仅由可打印的 ASCII 字符组成
 *
 *
 */
package main

// @lc code=start
import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	for start, end := 0, len(runes)-1; start < end; {
		// fmt.Printf("Rune %v is '%c'\n", i, runes[i])
		c := runes[start]
		h := runes[end]
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			start++
			continue
		} else if !unicode.IsLetter(h) && !unicode.IsNumber(h) {
			end--
			continue
		}
		if unicode.ToLower(c)-unicode.ToLower(h) != 0 {
			return false
		}
		start++
		end--
	}
	return true
}

// @lc code=end

// func main() {
// 	isPalindrome("jisdjfi你你好")
// }
