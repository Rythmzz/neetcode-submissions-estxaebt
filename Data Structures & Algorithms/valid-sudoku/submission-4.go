func isValidSudoku(board [][]byte) bool {
	checker := map[byte]bool{}

	for i:= 0 ; i< len(board);i++{
	checker = map[byte]bool{}
		for j:= 0; j < len(board); j++{
			if checker[board[i][j]] && board[i][j] != '.' {
				return false
			}
			checker[board[i][j]] = true
		}
	}

	for i:= 0 ; i< len(board);i++{
	checker = map[byte]bool{}
		for j:= 0; j < len(board); j++{
			if checker[board[j][i]] && board[j][i] != '.' {
				return false
			}
			checker[board[j][i]] = true
		}
	}

	for startRow := 0 ; startRow <  len(board);  startRow +=3 {
		for startCol :=0 ; startCol < len(board); startCol +=3 {
			checker = map[byte]bool{}

			for i:= startRow; i < startRow+3;i++{
				for j:= startCol; j< startCol+3; j++{
					if checker[board[i][j]] && board[i][j] != '.' {
						return false
					}
					checker[board[i][j]] = true
				}
			}
		}
	}

	return true
}
