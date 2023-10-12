package sort_colors

func SortColors(nums []int) {
	sortColors(nums)
}

func sortColors(nums []int) {
	red := 0
	white := 0
	blue := len(nums) - 1

	for white <= blue {
		switch nums[white] {
		case 0:
			nums[red], nums[white] = nums[white], nums[red]
			white++
			red++
		case 1:
			white++
		default:
			nums[blue], nums[white] = nums[white], nums[blue]
			blue--
		}
	}
}
