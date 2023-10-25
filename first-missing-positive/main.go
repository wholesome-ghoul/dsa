package first_missing_positive

import "math"

func FirstMissingPositive(nums []int) int {
	return firstMissingPositive(nums)
}

func firstMissingPositive(nums []int) int {
	missing := 1

	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = 0
		}
	}

	for _, num := range nums {
		if num-1 > len(nums)-1 || num == 0 {
			continue
		}

		abs := int(math.Abs(float64(num)))
		index := abs - 1

		if nums[index] == 0 {
			nums[index] = -abs
			continue
		}

		if nums[index] < 0 {
			continue
		}

		nums[index] = -nums[index]
	}

	for _, num := range nums {
		if num >= 0 {
			break
		}

		missing++
	}

	return missing
}
