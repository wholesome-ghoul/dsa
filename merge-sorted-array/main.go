package merge_sorted_array

func Merge(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p0 := len(nums1) - 1
	p1 := m - 1 // can be less than 0
	p2 := n - 1 // can be less than 0

	for p1 >= 0 && p2 >= 0 {
		if nums2[p2] >= nums1[p1] {
			nums1[p0] = nums2[p2]
			p2--
		} else {
			nums1[p0] = nums1[p1]
			p1--
		}
		p0--
	}

	for p1 >= 0 {
		nums1[p0] = nums1[p1]
		p0--
		p1--
	}

	for p2 >= 0 {
		nums1[p0] = nums2[p2]
		p0--
		p2--
	}
}
