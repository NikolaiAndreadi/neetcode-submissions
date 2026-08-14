func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers)-1
	for {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l+1,r+1}
		}
		if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{0,0}
}
