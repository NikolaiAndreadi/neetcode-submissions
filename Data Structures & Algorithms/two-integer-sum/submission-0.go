func twoSum(nums []int, target int) []int {
    memo := make(map[int]int)
	for i, num := range nums {
		delta := target - num
		if anotherIdx, ok := memo[num]; ok {
			return []int{anotherIdx, i}
		}
		memo[delta] = i
	}
	return []int{-1, -1}
}

// 3 4 5 6 .... 7
// target - 3 = 4
