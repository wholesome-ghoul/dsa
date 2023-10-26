package two_sum

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	indices := make([]int, 2)
	indexMaps := make(map[int]int)

	for i, num := range nums {
		indexMaps[num] = i
	}

	for i, num := range nums {
		opposite := target - num
		if _, ok := indexMaps[opposite]; ok {
			indices[0] = i

			if opposite == num {
				if i != indexMaps[opposite] {
					indices[1] = indexMaps[opposite]
					break
				}
				continue
			}

			indices[1] = indexMaps[opposite]
			break
		}
	}

	return indices
}
