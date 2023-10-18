package number_of_visible_people_in_a_queue

func CanSeePersonsCount(heights []int) []int {
	return canSeePersonsCount(heights)
}

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	rightPersons := make([]int, n)

	stack := []int{heights[n-1]}

	for i := n - 2; i >= 0; i-- {
		for heights[i] > stack[len(stack)-1] {
			rightPersons[i]++
			stack = stack[0 : len(stack)-1]

			if len(stack) == 0 {
				break
			}
		}

		if len(stack) > 0 && heights[i] < stack[len(stack)-1] {
			rightPersons[i]++
		}

		stack = append(stack, heights[i])
	}

	return rightPersons
}
