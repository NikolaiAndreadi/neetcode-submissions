func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num] += 1
	}

	invFreqArr := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		invFreqArr[freq] = append(invFreqArr[freq], num)
	}
	
	res := make([]int, 0, k)
	for i := len(nums); i >= 0; i-- {
		if len(invFreqArr[i]) > 0 {
			res = append(res, invFreqArr[i]...)
		}
		if len(res) >= k {
			break
		}
	}
	return res
}
