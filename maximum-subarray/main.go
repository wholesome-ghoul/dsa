package maximum_subarray

import "math"

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}

func maxSubArray(nums []int) int {
	sum := nums[0]

	right := 1

	tmpSum := sum
	for right < len(nums) {
		tmpSum += nums[right]

		if tmpSum < nums[right] {
			tmpSum = nums[right]
		}

		sum = int(math.Max(float64(sum), float64(tmpSum)))
		right++
	}

	return sum
}
