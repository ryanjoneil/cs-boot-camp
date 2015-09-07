package main

import "fmt"

func conflicts(row int, board []int) bool {
	// Determines if the current row positions conflicts with another.
	for i := row - 1; i >= 0; i-- {
		// Since everything beyond the current position hasn't been
		// determined yet, we only need to look at the past.

		// Row conflicts can't happen as we are putting one queen in each
		// row, so we don't have to detect those.

		// This detects column conflicts.
		if board[i] == board[row] {
			return true
		}

		// These detect diagonal conflicts.
		if board[i] == board[row]-(row-i) {
			return true
		}

		if board[i] == board[row]+(row-i) {
			return true
		}
	}

	return false
}

func nQueens(row int, board []int) bool {
	// Returns true if a non-conflicting board layout could be found.

	// If there are no more positions left to fill, we're done.
	if row >= len(board) {
		return true
	}

	// Recursively assign columns to each row. Stop as soon as we find
	// a layout in which no queen threatens any other queen.
	for i := 0; i < len(board); i++ {
		board[row] = i
		if !conflicts(row, board) {
			if nQueens(row+1, board) {
				return true
			}
		}
	}

	return false
}

func main() {
	size := 8
	board := make([]int, size)

	if nQueens(0, board) {
		// Print out a human readable version of the baord.
		for _, i := range board {
			for j := 0; j < size; j++ {
				if i == j {
					fmt.Print("Q ")
				} else {
					fmt.Print("- ")
				}
			}
			fmt.Println()
		}

	} else {
		fmt.Printf("Unable to solve %d-queens.\n", size)
	}
}
