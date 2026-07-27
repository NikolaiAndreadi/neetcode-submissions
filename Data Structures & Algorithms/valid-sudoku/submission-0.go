func isValidSudoku(board [][]byte) bool {
	type unique [9]map[byte]struct{}
	var rows, cols, boxes unique

	checkAndSave := func(u *unique, pos int, el byte) bool {
		if el == '.' {
			return true
		}
		if _, ok := u[pos][el]; ok {
			return false
		}
		if u[pos] == nil {
			u[pos] = make(map[byte]struct{}, 9)
		}
		u[pos][el] = struct{}{}
		return true
	}

	for r := range board {
		for c, el := range board[r] {
			if ok := checkAndSave(&rows, r, el); !ok {
				return false
			}
			if ok := checkAndSave(&cols, c, el); !ok {
				return false
			}
			boxNum := 3*(r/3) + (c/3)
			fmt.Println(r,c,boxNum)
			if ok := checkAndSave(&boxes, boxNum, el); !ok {
				return false
			}
		}
	}
	return true
}

