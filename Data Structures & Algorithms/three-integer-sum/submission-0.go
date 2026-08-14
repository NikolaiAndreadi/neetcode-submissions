import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		pairs := twoSum(nums[i+1:], -nums[i])
		for _, pair := range pairs {
			result = append(result, []int{
				nums[i],
				pair[0],
				pair[1],
			})
		}
	}

	return result
}

func twoSum(nums []int, target int) [][]int {
	var result [][]int
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]

		switch {
		case sum < target:
			left++

		case sum > target:
			right--

		default:
			result = append(result, []int{
				nums[left],
				nums[right],
			})

			left++
			right--

			for left < right && nums[left] == nums[left-1] {
				left++
			}

			for left < right && nums[right] == nums[right+1] {
				right--
			}
		}
	}

	return result
}