package main

// Board - A Chess board
type Board [][]byte

// QUEEN - A Queen on the Board
const QUEEN = byte('Q')

// TAKEN - A place on the Board that cannot have a Queen on
const TAKEN = byte('X')

// FREE - A free place on the Board
const FREE = 0

// MakeBoard - Create a n by n Board
func MakeBoard(n int) Board {
	result := make(Board, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, n)
	}
	return result
}

// Copy - Get a copy of a board
func (board Board) Copy() Board {
	result := make([][]byte, len(board))
	for i, row := range board {
		result[i] = make([]byte, len(row))
		copy(result[i], row)
	}
	return result
}

// ToString - Get a good-looking string representation of the board
func (board Board) ToString() string {
	result := "[ "
	for i, row := range board {
		result += "[ "
		for _, place := range row {
			switch place {
			case QUEEN:
				result += string(place)
			default:
				result += "."
			}
			result += " "
		}
		if i+1 < len(board) {
			result += "]\n  "
		}
	}
	result += "] ]"
	return result
}
