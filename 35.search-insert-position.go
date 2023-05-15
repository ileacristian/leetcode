// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.
// https://leetcode.com/problems/search-insert-position/

package search_insert_position

func searchInsert(nums []int, target int) int {
	for i := range nums {
		if target == nums[i] {
			return i
		} else if i < len(nums)-1 && target > nums[i] && target < nums[i+1] {
			return i + 1
		} else if i == len(nums)-1 && target > nums[i] {
			return i + 1
		}
	}

	return 0
}
