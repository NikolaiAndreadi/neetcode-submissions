func carFleet(target int, position []int, speed []int) int {
	type car struct {
		x int
		time float64
	}
	cars := make([]car, len(position))
	for i := range position {
		cars[i].x = position[i]
		cars[i].time = float64(target-position[i]) / float64(speed[i])
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].x < cars[j].x
	})

	mds := make([]float64, 0, len(cars))
	for _, c := range cars {
		for len(mds) > 0 && mds[len(mds)-1] <= c.time {
			mds = mds[:len(mds)-1]
		}
		mds = append(mds, c.time)
	}
	return len(mds)
}
