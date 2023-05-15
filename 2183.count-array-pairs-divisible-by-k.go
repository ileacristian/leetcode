// Given a 0-indexed integer array nums of length n and an integer k, return the number of pairs (i, j) such that:
// 0 <= i < j <= n - 1 and
// nums[i] * nums[j] is divisible by k.
// https://leetcode.com/problems/count-array-pairs-divisible-by-k/

package count_array_pairs_divisible_by_k

import "math"

func countPairs(nums []int, k int) int64 {
	gcds := make(map[int]int64, int(math.Sqrt(float64(k))))

	for _, a := range nums {
		gcds[gcd(a, k)]++
	}

	var result int64 = 0
	for a, n1 := range gcds {
		for b, n2 := range gcds {
			if a < b || a*b%k != 0 {
				continue
			}

			if a != b {
				result += n1 * n2
			}

			if a == b {
				result += n1 * (n1 - 1) / 2
			}
		}
	}

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
