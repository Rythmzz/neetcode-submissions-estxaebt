func isValidSudoku(board [][]byte) bool {
	for i:=0;i<len(board);i++{
		rowDup := map[byte]bool{}
		for j:=0 ;j< len(board);j++{
			if rowDup[board[i][j]] && board[i][j] != '.' {
				return false
			} else {
				rowDup[board[i][j]] = true
			}
		}
	}

	for i:=0;i<len(board);i++{
		rowDup := map[byte]bool{}
		for j:=0 ;j< len(board);j++{
			if rowDup[board[j][i]] && board[j][i] != '.' {
				return false
			} else {
				rowDup[board[j][i]] = true
			}
		}
	}

	for startRow:=0;startRow<len(board);startRow +=3{
		for startCol:=0;startCol<len(board);startCol+=3{
			squareDup:= map[byte]bool{}
			for i:= startRow;i < startRow+3;i++ {
				for j:= startCol;j < startCol+3;j++ {
					if squareDup[board[i][j]] && board[i][j] != '.' {
						return false
					} else {
						squareDup[board[i][j]] = true
					}
				}
			}
		}
	}
return true

}

