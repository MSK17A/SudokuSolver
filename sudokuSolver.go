package piscine

import "fmt"

type Board [9][9]int

func constructBoard(board *Board) {
	*board = Board{
		{0, 9, 6, 0, 4, 0, 0, 0, 1},
		{1, 0, 0, 0, 6, 0, 0, 0, 4},
		{5, 0, 4, 8, 1, 0, 3, 9, 0},
		{0, 0, 7, 9, 5, 0, 0, 4, 3},
		{0, 3, 0, 0, 8, 0, 0, 0, 0},
		{4, 0, 5, 0, 2, 3, 0, 1, 8},
		{0, 1, 0, 6, 3, 0, 0, 5, 9},
		{0, 5, 9, 0, 7, 0, 8, 3, 0},
		{0, 0, 3, 5, 9, 0, 0, 0, 7},
	}
	/*board[0] = [9]int{0, 9, 6, 0, 4, 0, 0, 0, 1}
	board[1] = [9]int{1, 0, 0, 0, 6, 0, 0, 0, 4}
	board[2] = [9]int{5, 0, 4, 8, 1, 0, 3, 9, 0}
	board[3] = [9]int{0, 0, 7, 9, 5, 0, 0, 4, 3}
	board[4] = [9]int{0, 3, 0, 0, 8, 0, 0, 0, 0}
	board[5] = [9]int{4, 0, 5, 0, 2, 3, 0, 1, 8}
	board[6] = [9]int{0, 1, 0, 6, 3, 0, 0, 5, 9}
	board[7] = [9]int{0, 5, 9, 0, 7, 0, 8, 3, 0}
	board[8] = [9]int{0, 0, 3, 5, 9, 0, 0, 0, 7}*/
}

func isValid(board *Board, num, pos_row, pos_col int) bool {
	for index_col := 0; index_col < 9; index_col++ {
		if board[pos_row][index_col] == num {
			return false
		}
	}
	for index_row := 0; index_row < 9; index_row++ {
		if board[index_row][pos_col] == num {
			return false
		}
	}

	// Find box
	box_row := int(pos_row / 3)
	box_col := int(pos_col / 3)

	/* The multiplying by 3 is because we want to get to first position in that box in of the board (possible box_row and col values are from 0 to 2)*/
	// Example, box_row = 1, the index_row should be 3, so index_row = box_row * 3
	for index_col := box_col * 3; index_col < box_col*3+3; index_col++ {
		for index_row := box_row * 3; index_row < box_row*3+3; index_row++ {
			if board[index_row][index_col] == num && index_row != pos_row && index_col != pos_col {
				return false
			}
		}
	}
	return true
}

func print_board(board *Board) {
	for i := 0; i < len(board); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("- - - - - - - - - - - - - - - -")
		}
		for j := 0; j < len(board[0]); j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print("|  ")
			}
			if j == 8 {
				fmt.Println(board[i][j])
			} else {
				fmt.Print(board[i][j])
				fmt.Print("  ")
			}
		}
	}
}

func empty_place(board *Board) (int, int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == 0 {
				return row, col
			}

		}
	}
	return -1, -1
}
func SudokuSolver(board *Board) bool {
	pos_row, pos_col := empty_place(board)
	if pos_row == -1 {
		/* if no empty places are found then we reached a solution*/
		// Print the board and khalaas
		print_board(board)
		return true
	}

	// If there are empty place try number from 1 to 9
	for input_num := 1; input_num < 10; input_num++ {
		// If the number is valid insert it in the empty position
		if isValid(board, input_num, pos_row, pos_col) {
			board[pos_row][pos_col] = input_num
			// One we put the number we can call the function again to choose the next number
			// And if the function returns true, that means there are no empty places and we have solution, so just return true
			// Note: SudokoSolver will return true only if there is a solutiuon
			if SudokuSolver(board) {
				return true
			}

			// Or backtrack (reset the current value and try another number)
			board[pos_row][pos_col] = 0
		}
	}
	return false
}
