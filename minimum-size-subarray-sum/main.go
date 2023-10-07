package minimum_size_subarray_sum

import "math"

func MinSubArrayLen(target int, nums []int) int {
	return minSubArrayLen(target, nums)
}

// https://leetcode.com/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	prefix := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	minLen := math.MaxInt

	i := 0
	j := 1

	for j < len(prefix) {
		diff := prefix[j] - prefix[i]
		if diff < target {
			j++
		} else {
			if j-i < minLen {
				minLen = j - i
			}
			i++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}
