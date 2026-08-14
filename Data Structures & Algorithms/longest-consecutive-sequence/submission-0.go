func longestConsecutive(nums []int) int {
	mappedNums := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		mappedNums[num] = struct{}{}
	}

	res := 0
	for num := range mappedNums {
		if _, ok := mappedNums[num-1]; ok {
			continue
		}
		subLen := 1 
		for next := num + 1; ; next++ {
			if _, exists := mappedNums[next]; !exists {
				break
			}
			subLen++
		}
		res = max(res, subLen)
	}

	return res
}