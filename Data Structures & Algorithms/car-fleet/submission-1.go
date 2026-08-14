func carFleet(target int, position []int, speed []int) int {
	type posSpeed struct {
		x int
		v int
	}
	posSpeeds := make([]posSpeed, len(position))
	for i := range position {
		posSpeeds[i].x = position[i]
		posSpeeds[i].v = speed[i]
	}
	sort.Slice(posSpeeds, func(i, j int) bool {
		return posSpeeds[i].x < posSpeeds[j].x
	})

	mds := make([]float64, 0, len(posSpeeds))
	for _, ps := range posSpeeds {
		val := float64(target-ps.x) / float64(ps.v)
		for len(mds) > 0 && mds[len(mds)-1] <= val {
			mds = mds[:len(mds)-1]
		}
		mds = append(mds, val)
	}
	return len(mds)
}

// 0  1 4 7
// 1  2 2 1
// 10 5 3 3

// position[4 1 0 7]
// speed.  [2 2 1 1]
// target 10
// (target-position) / speed [3 5 10 3] 