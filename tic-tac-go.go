package main

import (
	"fmt"
)

// TicTacToe represents a Tic-Tac-Toe game
type TicTacToe struct {
	board [9]rune
	turn  rune
}

// NewGame initializes a new Tic-Tac-Toe game
func NewGame() *TicTacToe {
	return &TicTacToe{
		board: [9]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
		turn:  'X',
	}
}

// PrintBoard prints the current state of the board
func (t *TicTacToe) PrintBoard() {
	fmt.Println("-------------")
	for i := 0; i < 9; i += 3 {
		fmt.Printf("| %s | %s | %s |\n", colorize(t.board[i]), colorize(t.board[i+1]), colorize(t.board[i+2]))
		fmt.Println("-------------")
	}
}

// MakeMove makes a move on the board
func (t *TicTacToe) MakeMove(position int) bool {
	if position < 1 || position > 9 || t.board[position-1] == 'X' || t.board[position-1] == 'O' {
		return false
	}
	t.board[position-1] = t.turn
	return true
}

// SwitchTurn switches the turn between players
func (t *TicTacToe) SwitchTurn() {
	if t.turn == 'X' {
		t.turn = 'O'
	} else {
		t.turn = 'X'
	}
}

// CheckWin checks if the current player has won
func (t *TicTacToe) CheckWin() bool {
	winningCombinations := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // columns
		{0, 4, 8}, {2, 4, 6}, // diagonals
	}

	for _, combo := range winningCombinations {
		if t.board[combo[0]] == t.turn && t.board[combo[1]] == t.turn && t.board[combo[2]] == t.turn {
			return true
		}
	}
	return false
}

// CheckDraw checks if the game is a draw
func (t *TicTacToe) CheckDraw() bool {
	for _, cell := range t.board {
		if cell != 'X' && cell != 'O' {
			return false
		}
	}
	return true
}

// colorize adds color to X and O
func colorize(s rune) string {
	switch s {
	case 'X':
		return red + string(s) + reset
	case 'O':
		return blue + string(s) + reset
	default:
		return string(s)
	}
}

func main() {
	game := NewGame()
	var position int
	var err error

	for {
		game.PrintBoard()
		if game.turn == 'X' {
			fmt.Printf("Player %s, enter your move (1-9): ", game.turn)
			_, err = fmt.Scanf("%d\n", &position)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number between 1 and 9.")
				fmt.Scanf("%s\n", new(string)) // clear the invalid input
				continue
			}
		} else {
			fmt.Println("Computer is thinking...")
			position = game.BestMove()
			fmt.Printf("Computer chose position %d\n", position)
		}

		if !game.MakeMove(position) {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		if game.CheckWin() {
			game.PrintBoard()
			fmt.Printf("Player %s wins!\n", game.turn)
			break
		}

		if game.CheckDraw() {
			game.PrintBoard()
			fmt.Println("It's a draw!")
			break
		}

		game.SwitchTurn()
	}

	fmt.Println("Game over.")
}
