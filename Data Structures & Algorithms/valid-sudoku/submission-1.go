func isValidSudoku(board [][]byte) bool {
	checkRow := [9][9]bool{}
	checkCol := [9][9]bool{}
	checkSquare := [9][9]bool{}
    
	for r :=0 ; r < 9; r++{
        for c:= 0; c < 9; c++{
            value := board[r][c]

            if value == '.'{
                continue
            }

            index := value - '1' // 1001 - 0001 = 1000

            squareIndex := (r/3)*3 +(c/3)

            if (checkRow[r][index] || checkCol[c][index] || checkSquare[squareIndex][index]) {
                return false
            }

            
			checkRow[r][index] = true
			checkCol[c][index] = true
			checkSquare[squareIndex][index] = true


        }

    }
    return true
}


