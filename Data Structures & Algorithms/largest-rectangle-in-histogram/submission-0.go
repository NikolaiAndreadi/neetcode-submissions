func largestRectangleArea(heights []int) int {
	type IndexHeight struct {
		i int
		h int
	}
	mds := make([]IndexHeight, 0, len(heights))
	maxArea := 0
	
	for i, h := range heights {
		stratIdx := i
		for len(mds) != 0 && (mds[len(mds)-1].h > h) {
			popped := mds[len(mds)-1]
			mds = mds[:len(mds)-1]

			maxArea = max(maxArea, (i-popped.i)*popped.h)
			stratIdx = popped.i
		}
		mds = append(mds, IndexHeight{i: stratIdx, h: h})
	}

	for _, leftover := range mds {
		maxArea = max(maxArea, leftover.h * (len(heights) - leftover.i))
	}

	return maxArea
}
