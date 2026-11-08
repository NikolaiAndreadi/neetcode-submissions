func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	l, r := 0, 1
	result := 0
	for r < len(prices) {
		lp, rp := prices[l], prices[r]
		if rp <= lp {
			l = r
			r++
			continue
		}
		result = max(result, rp-lp)
		r++
	}
	return result
}
