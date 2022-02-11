package main

func hasDuplicateInRows(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		cache := map[byte]struct{}{}
		for j := 0; j < 9; j++ {
			cellVal := board[i][j]

			if _, ok := cache[cellVal]; cellVal != '.' && ok {
				return true
			}
			cache[cellVal] = struct{}{}
		}
	}
	return false
}

func hasDuplicateInColumns(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		cache := map[byte]struct{}{}
		for j := 0; j < 9; j++ {
			cellVal := board[j][i]

			if _, ok := cache[cellVal]; cellVal != '.' && ok {
				return true
			}
			cache[cellVal] = struct{}{}
		}
	}
	return false
}

func hasDuplicateInSubGrid(board [][]byte) bool {
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			cache := map[byte]struct{}{}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					cellVal := board[a*3+i][b*3+j]

					if _, ok := cache[cellVal]; cellVal != '.' && ok {
						return true
					}
					cache[cellVal] = struct{}{}
				}
			}
		}
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	return !hasDuplicateInRows(board) && !hasDuplicateInColumns(board) && !hasDuplicateInSubGrid(board)
}
