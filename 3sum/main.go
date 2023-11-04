package _3sum

import (
	"fmt"
	"slices"
)

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

// optimized solution
func threeSum(nums []int) [][]int {
	triplets := [][]int{}

	tripletsSet := make(map[[3]int]bool)
	numsSet := make(map[int]bool)

	slices.Sort(nums)
	for i := range nums {
		if numsSet[nums[i]] {
			continue
		}

		numsSet[nums[i]] = true
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				tripletsSet[triplet] = true

				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	for k := range tripletsSet {
		triplets = append(triplets, []int{k[0], k[1], k[2]})
	}

	return triplets
}

func _threeSum(nums []int) [][]int {
	triplets := [][]int{}
	target := int(0)

	twoSums := make(map[int][][2]int)
	setOfNums := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		// for optimization purposes, skip duplicated nums
		if setOfNums[nums[i]] {
			continue
		}
		setOfNums[nums[i]] = true

		for j := i + 1; j < len(nums); j++ {
			sum := int(nums[i] + nums[j])
			twoSums[sum] = append(twoSums[sum], [2]int{i, j})
		}
	}

	rawTriplets := make(map[[3]int]bool)
	for i, num := range nums {
		missingMembers := twoSums[target-int(num)]

		if len(missingMembers) == 0 {
			continue
		}

		for _, twoSum := range missingMembers {
			index1, index2 := twoSum[0], twoSum[1]
			if index1 == i || index2 == i {
				continue
			}

			sorted := [3]int{nums[index1], nums[index2], num}
			slices.Sort(sorted[:])
			rawTriplets[sorted] = true
		}
	}

	for k := range rawTriplets {
		triplets = append(triplets, []int{k[0], k[1], k[2]})
	}

	return triplets
}

func main() {
	fmt.Println("OI")
}
