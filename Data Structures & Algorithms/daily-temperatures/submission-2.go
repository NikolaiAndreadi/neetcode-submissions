// 30,38,30,36,35,40,28
	//          ^
	// res: [1, _, 1]
	// stack: [38@1,36@3
	// NEW IS BIGGER => update stack, calc a diff for old idx:
	// res[oldIdx]: currIdx-oldIdx
	// do until monotonically decreasing

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	type stackValue struct {
		val int
		idx int
	}

	mds := make([]stackValue, 0, len(temperatures))
	
	for i, currTemp := range temperatures {
		for len(mds) > 0 {
			top := mds[len(mds)-1]
			if top.val < currTemp {
				result[top.idx] = i-top.idx
				mds = mds[:len(mds)-1]
				continue
			}
			break
		}
		mds = append(mds, stackValue{val: currTemp, idx: i})
	}
	return result
}
