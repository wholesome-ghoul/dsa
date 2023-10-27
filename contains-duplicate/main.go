package contains_duplicate

type void struct{}

var member void

func ContainsDuplicate(nums []int) bool {
	return containsDuplicate(nums)
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]void)

	for _, num := range nums {
		set[num] = member
	}

	if len(set) != len(nums) {
		return true
	}

	return false
}
