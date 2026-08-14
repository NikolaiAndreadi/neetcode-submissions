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
		return cars[i].x > cars[j].x
	})

	var maxArrivalTime float64
	var fleets int
	for _, c := range cars {
		if c.time > maxArrivalTime {
			fleets++
			maxArrivalTime = c.time
		}
	}
	return fleets
}
