func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num] += 1
	}

	invFreqArr := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		invFreqArr[freq] = append(invFreqArr[freq], num)
	}
	fmt.Println(invFreqArr)
	
	res := make([]int, 0, k)
	for i := len(nums); i >= 0; i-- {
		if len(invFreqArr[i]) == 0 {
			continue
		}
		for _, candidate := range invFreqArr[i] {
			res = append(res, candidate)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
