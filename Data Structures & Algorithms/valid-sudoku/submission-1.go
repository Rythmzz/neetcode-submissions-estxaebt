func isValidSudoku(board [][]byte) bool {
	for i:= 0 ;i < len(board); i++{
		checkRow := map[byte]bool{}
		for j:= 0; j < len(board); j++ {
			if checkRow[board[i][j]] && board[i][j] != '.' {
				return false
			} else {
				checkRow[board[i][j]] = true
			}
		}
	}

	for i:= 0 ;i < len(board); i++{
		checkColumn := map[byte]bool{}
		for j:= 0; j < len(board); j++ {
			if checkColumn[board[j][i]] && board[j][i] != '.' {
				return false
			} else {
				checkColumn[board[j][i]] = true
			}
		}
	}

	for startRow := 0; startRow < len(board); startRow += 3 {
		for startColumn := 0; startColumn < len(board); startColumn +=3 {
			checkSquare := map[byte]bool{}

			for i:= startRow; i < startRow + 3;i++ {
				for j:= startColumn; j < startColumn + 3; j++ {
					if checkSquare[board[i][j]] && board[i][j] != '.' {
						return false
					} else {
						checkSquare[board[i][j]] = true
					}
				}
			}
		}
	}

	return true

}
