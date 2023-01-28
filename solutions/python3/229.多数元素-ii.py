#
# @lc app=leetcode.cn id=229 lang=python3
#
# [229] 多数元素 II
#
# https://leetcode.cn/problems/majority-element-ii/description/
#
# algorithms
# Medium (53.77%)
# Likes:    658
# Dislikes: 0
# Total Accepted:    91.3K
# Total Submissions: 169.3K
# Testcase Example:  '[3,2,3]'
#
# 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
# 
# 
# 
# 示例 1：
# 
# 
# 输入：nums = [3,2,3]
# 输出：[3]
# 
# 示例 2：
# 
# 
# 输入：nums = [1]
# 输出：[1]
# 
# 
# 示例 3：
# 
# 
# 输入：nums = [1,2]
# 输出：[1,2]
# 
# 
# 
# 提示：
# 
# 
# 1 <= nums.length <= 5 * 10^4
# -10^9 <= nums[i] <= 10^9
# 
# 
# 
# 
# 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
# 
#

# @lc code=start
class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        candidate1 = candidate2 = None
        vote1 = vote2 = 0
        for num in nums:
            if num == candidate2:
                vote2 +=1
            elif num == candidate1:
                vote1 += 1
            elif vote1 == 0:
                candidate1 = num
                vote1 = 1
            elif vote2 == 0:
                candidate2 = num
                vote2 = 1
            else:
                vote1 -=1
                vote2 -=1
        cnt1 = cnt2 = 0
        for num in nums:
            if num == candidate1 and vote1 > 0:
                cnt1 +=1
            elif num == candidate2 and vote2 > 0:
                cnt2 +=1
        ans = []
        if(cnt1 > len(nums)/3):
            ans.append(candidate1)
        if(cnt2 > len(nums)/3):
            ans.append(candidate2)
        return ans
# @lc code=end

