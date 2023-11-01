package search_in_rotated_sorted_array

func Search(nums []int, target int) int {
	return search(nums, target)
}

func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	rotationIndex := left
	left, right = 0, n-1
	for left <= right {
		mid := (left + right) / 2
		actualMid := (mid + rotationIndex) % n
		if nums[actualMid] == target {
			return actualMid
		}

		if nums[actualMid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
