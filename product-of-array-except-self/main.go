package product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	return productExceptSelf(nums)
}

// without division
func productExceptSelf(nums []int) []int {
	fromLeftProds := make([]int, len(nums)+2)
	fromRightProds := make([]int, len(nums)+2)

	fromLeftProds[0] = 1
	fromLeftProds[len(fromLeftProds)-1] = 1
	fromRightProds[0] = 1
	fromRightProds[len(fromRightProds)-1] = 1

	prods := make([]int, len(nums))
	prod := 1

	occurrencesOfZero := []int{}
	for i, num := range nums {
		if num != 0 {
			prod *= num
			fromLeftProds[i+1] = prod
			continue
		}
		occurrencesOfZero = append(occurrencesOfZero, i)
	}

	if len(occurrencesOfZero) == 1 {
		prods[occurrencesOfZero[0]] = prod
		return prods
	}

	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			prod *= nums[i]
			fromRightProds[i+1] = prod
			continue
		}
	}

	if len(occurrencesOfZero) < 1 {
		for i := range nums {
			prods[i] = fromLeftProds[i] * fromRightProds[i+2]
		}
	}

	return prods
}

func __productExceptSelf(nums []int) []int {
	prods := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		prods[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		prods[i] *= postfix
		postfix *= nums[i]
	}

	return prods
}

func _productExceptSelf(nums []int) []int {
	prods := make([]int, len(nums))
	prod := 1

	occurrencesOfZero := []int{}
	for i, num := range nums {
		if num != 0 {
			prod *= num
			continue
		}
		occurrencesOfZero = append(occurrencesOfZero, i)
	}

	if len(occurrencesOfZero) == 1 {
		prods[occurrencesOfZero[0]] = prod
	} else if len(occurrencesOfZero) < 1 {
		for i := 0; i < len(nums); i++ {
			prods[i] = prod / nums[i]
		}
	}

	return prods
}
