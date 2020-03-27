package main

// Crawl over all the options of a board, starting from the given row.
// Outputs the solutions found into the results channel
func Crawl(board Board, results chan Board) {
	_crawl(board, 0, results)
	close(results)
}

// The underlying reqursive crawler function
func _crawl(board Board, row int, results chan Board) {
	if row >= len(board) {
		results <- board
		return
	}

	for i, place := range board[row] {
		if place != FREE {
			continue
		}

		boardCopy := board.Copy()
		boardCopy[row][i] = QUEEN
		updateTaken(boardCopy, row, i)
		_crawl(boardCopy, row+1, results)
	}
	// Check if there is a place for another queen on the board
	// If there is, Add a queen to the board and recurse until all rows are filled
	// Else, remove the queen and try a different position
}

// Update the relevant taken places on the board when a Queen is places on the specified position (row, column)
func updateTaken(board Board, row int, col int) {
	// There is no point in updating smaller rows because they all should be taken.
	// The same is true for the current row because it will never be checked

	// Update places in the column below Queen
	for r := row + 1; r < len(board); r++ {
		board[r][col] = TAKEN
	}

	// Update diagonal to the right
	for offset := 1; row+offset < len(board) && col+offset < len(board[0]); offset++ {
		board[row+offset][col+offset] = TAKEN
	}

	// Update diagonal to the left
	for offset := 1; row+offset < len(board) && col-offset >= 0; offset++ {
		board[row+offset][col-offset] = TAKEN
	}
}
