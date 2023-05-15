// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent
// houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

// Example 1:
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.

// Example 2:
// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
// Total amount you can rob = 2 + 9 + 1 = 12.

// https://leetcode.com/problems/house-robber/

import Foundation

class Solution {
  func rob(_ nums: [Int]) -> Int {
    hashmap = [:]
    return max(rob2(nums, pos: 0), rob2(nums, pos: 1))
  }

  var hashmap: [Int: Int] = [:]

  // top down
  func rob2(_ nums: [Int], pos: Int) -> Int {
    if pos >= nums.count {
      return 0
    }

    if let cached_sum = hashmap[pos] {
      return cached_sum
    }

    let sum = nums[pos] + max(rob2(nums, pos: pos + 2), rob2(nums, pos: pos + 3))
    hashmap[pos] = sum
    return sum

  }

  // bottom up
  func rob1(_ nums: [Int]) -> Int {
    var dp: [Int] = Array(repeating: 0, count: nums.count + 3)
    for i in stride(from: nums.count - 1, through: 0, by: -1) {
      dp[i] = nums[i] + max(dp[i + 2], dp[i + 3])
    }
    return max(dp[0], dp[1])
  }
}

let solution = Solution()
print(solution.rob1([1, 2, 3, 1]))  // 4
print(solution.rob1([2, 7, 9, 3, 1]))  // 12
