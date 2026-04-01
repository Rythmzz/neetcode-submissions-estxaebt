func isValidSudoku(board [][]byte) bool {
	for i:= 0; i < len(board);i++{
		dup := map[byte]bool{}
		for j:= 0 ; j< len(board);j++ {
			if dup[board[i][j]] && board[i][j] != '.' {
				return false
			}
			dup[board[i][j]] = true
		}
	}

	for i:= 0; i < len(board);i++{
		dup := map[byte]bool{}
		for j:= 0 ; j< len(board);j++ {
			if dup[board[j][i]] && board[j][i] != '.' {
				return false
			}
			dup[board[j][i]] = true
		}
	}

	for startRow := 0 ;startRow < len(board); startRow += 3 {
		for startCol := 0 ;startCol < len(board); startCol +=3 {
		dup := map[byte]bool{}
		for i:= startRow; i < startRow+3;i++{
			for j:= startCol ; j< startCol+3;j++ {
			if dup[board[i][j]] && board[i][j] != '.' {
				return false
			}
			dup[board[i][j]] = true
		}
	}
		}
	}

	return true
}
