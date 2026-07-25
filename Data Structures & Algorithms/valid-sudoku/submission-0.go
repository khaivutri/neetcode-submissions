func isValidSudoku(board [][]byte) bool {
	checkRow := [9][9]bool{}
	checkCol := [9][9]bool{}
	checkSquare := [9][9]bool{}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			value := board[r][c]
			if value == '.' {
				continue
			}

			num := value - '1'

			squareIndex := (r/3)*3 + (c/3)
			if checkRow[r][num] || checkCol[c][num] || checkSquare[squareIndex][num] {
				return false
			}

			checkRow[r][num] = true
			checkCol[c][num] = true
			checkSquare[squareIndex][num] = true
		}
	}
	return true
}