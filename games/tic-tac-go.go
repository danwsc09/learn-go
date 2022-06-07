package main

import (
	"fmt"
)

func main() {
	var player string
	OToMove := true
	var row, col int
	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	for {
		printBoard(board)
		if OToMove {
			player = "O"
		} else {
			player = "X"
		}
		fmt.Printf("Enter coordinates (vertical then horizontal).\nPlayer %s: ", player)
		fmt.Scanln(&row, &col)

		if !makeMove(board, row, col, player) {
			fmt.Println("Invalid move! Try again.")
		} else if getWinner(board) != "-" {
			printBoard(board)
			fmt.Printf("Winner: %s\n", getWinner(board))
			break
		} else if gameIsOver(board) {
			fmt.Println("Draw!")
			break
		} else {
			OToMove = !OToMove
		}
	}
}

/*
1. invalid move
2. valid move
	a. there is a winner
		- break and show winner
	b. there is no winner
		i. all squares are full - stop game
		ii. change player
*/

func getWinner(board [][]string) string {
	// row + col
	for i := 0; i < 3; i++ {
		if (board[i][0] != "-" && board[i][1] == board[i][0] && board[i][2] == board[i][0]) ||
			(board[0][i] != "-" && board[1][i] == board[0][i] && board[2][i] == board[0][i]) {
			return board[i][0]
		}
	}
	if (board[0][0] != "-" && board[1][1] == board[0][0] && board[2][2] == board[0][0]) ||
		(board[0][2] != "-" && board[0][2] == board[1][1] && board[2][0] == board[1][1]) {
		return board[1][1]
	}
	return "-"
}

func printBoard(board [][]string) {
	fmt.Println()
	fmt.Printf("  1 2 3\n")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func makeMove(board [][]string, row, col int, player string) bool {
	if board[row-1][col-1] != "-" {
		return false
	}
	board[row-1][col-1] = player

	return true
}

func gameIsOver(board [][]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}
