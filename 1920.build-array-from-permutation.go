// Given a zero-based permutation nums (0-indexed), build an array ans of the same length where ans[i] = nums[nums[i]] for each 0 <= i < nums.length and return it.
// A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).
// https://leetcode.com/problems/build-array-from-permutation/

package build_array_from_permutation

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = nums[nums[i]]
	}

	return ans
}
