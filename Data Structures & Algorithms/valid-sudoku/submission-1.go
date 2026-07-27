func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9]uint16
	for r := range 9 {
		for c := range 9 {
			el := board[r][c]
			if el == '.' {
				continue
			}

			bit := uint16(1) << (el - '1')
			box := 3*(r/3) + c/3
			if rows[r]&bit != 0 || cols[c]&bit != 0 || boxes[box]&bit != 0 {
				return false
			}

			rows[r] |= bit
			cols[c] |= bit
			boxes[box] |= bit
		}
	}
	return true
}

