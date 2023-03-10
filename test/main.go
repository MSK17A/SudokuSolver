package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	var board piscine.Board
	args := os.Args[1:]
	// Only accept args that are from 1 to 9 otherwise return
	if len(args) < 9 {
		fmt.Println("Error")
		return
	}
	for row, value := range args {
		if len(value) < 9 {
			fmt.Println("Error")
			return
		}
		for col, value_char := range value {
			if '1' <= value_char && value_char <= '9' || value_char == '.' {
				if value_char == '.' {
					board[row][col] = 0
				} else {
					board[row][col] = int(value_char - '0')
				}
			} else {

				return
			}
		}
	}
	/*board := piscine.Board{
		{0, 9, 6, 0, 4, 0, 0, 0, 1},
		{1, 0, 0, 0, 6, 0, 0, 0, 4},
		{5, 0, 4, 8, 1, 0, 3, 9, 0},
		{0, 0, 7, 9, 5, 0, 0, 4, 3},
		{0, 3, 0, 0, 8, 0, 0, 0, 0},
		{4, 0, 5, 0, 2, 3, 0, 1, 8},
		{0, 1, 0, 6, 3, 0, 0, 5, 9},
		{0, 5, 9, 0, 7, 0, 8, 3, 0},
		{0, 0, 3, 5, 9, 0, 0, 0, 7},
	}*/
	if !piscine.SudokuSolver(&board) {
		fmt.Print("Error")
	}
}
