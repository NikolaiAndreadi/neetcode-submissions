type cars struct {
	position []int
	speed    []int
}

func (c cars) Len() int {
	return len(c.position)
}

func (c cars) Less(i, j int) bool {
	return c.position[i] > c.position[j]
}

func (c cars) Swap(i, j int) {
	c.position[i], c.position[j] = c.position[j], c.position[i]
	c.speed[i], c.speed[j] = c.speed[j], c.speed[i]
}

func carFleet(target int, position []int, speed []int) int {
	sort.Sort(cars{
		position: position,
		speed:    speed,
	})

	fleets := 0

	var maxNumerator, maxDenominator int64

	for i := range position {
		distance := int64(target - position[i])
		velocity := int64(speed[i])

		if fleets == 0 ||
			distance*maxDenominator > maxNumerator*velocity {

			fleets++
			maxNumerator = distance
			maxDenominator = velocity
		}
	}

	return fleets
}