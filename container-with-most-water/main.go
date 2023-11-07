package container_with_most_water

import "math"

func MaxArea(height []int) int {
	return maxArea(height)
}
func maxArea(height []int) int {
	product := 0

	left := 0
	right := len(height) - 1
	for left < right {
		indexDiff := right - left
		leftHeight := height[left]
		rightHeight := height[right]
		minEdge := int(math.Min(float64(leftHeight), float64(rightHeight)))

		if leftHeight > rightHeight {
			right--
		} else {
			left++
		}

		product = int(math.Max(float64(product), float64(indexDiff*minEdge)))
	}

	return product
}
