// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
// https://leetcode.com/problems/two-sum/

package two_sum

import "fmt"

func twoSum(nums []int, target int) []int {
	n := len(nums)
	m := make(map[int]int, n)

	for index, value := range nums {
		complementIndex, exists := m[target-value]
		if exists {
			return []int{index, complementIndex}
		}

		m[value] = index
	}

	return []int{}
}

func main() {

	// 	Example 1:
	// Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

	// Example 2:
	// Input: nums = [3,2,4], target = 6
	// Output: [1,2]
	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))

	// Example 3:
	// Input: nums = [3,3], target = 6
	// Output: [0,1]
	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))
}
