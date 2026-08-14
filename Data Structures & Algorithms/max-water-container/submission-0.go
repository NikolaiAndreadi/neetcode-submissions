func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	result := 0
	for l < r {
		onL := heights[l]
		onR := heights[r]

		minHeight := min(onL, onR)
		width := r-l
		volume := minHeight * width
		result = max(result, volume)

		if onL > onR {
			r--
		} else {
			l++
		}
	}
	return result
}
