package game

import (
	"fmt"
	"math"
	"os"
	"runtime/pprof"
)

const (
	red   = "\033[31m"
	blue  = "\033[34m"
	reset = "\033[0m"
)

// TicTacToe represents a Tic-Tac-Toe game
type TicTacToe struct {
	board [9]rune
	Turn  rune
}

// NewTicTacToe initializes a new Tic-Tac-Toe game
func NewTicTacToe() *TicTacToe {
	return &TicTacToe{
		board: [9]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
		Turn:  'X',
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
	t.board[position-1] = t.Turn
	return true
}

// SwitchTurn switches the turn between players
func (t *TicTacToe) SwitchTurn() {
	if t.Turn == 'X' {
		t.Turn = 'O'
	} else {
		t.Turn = 'X'
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
		if t.board[combo[0]] == t.Turn && t.board[combo[1]] == t.Turn && t.board[combo[2]] == t.Turn {
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

// Minimax implements the Minimax algorithm
func (t *TicTacToe) Minimax() int {
	if t.CheckWin() {
		return 1
	}
	if t.CheckDraw() {
		return 0
	}

	t.SwitchTurn()

	bestScore := math.Inf(-1)
	for i := 0; i < 9; i++ {
		if t.board[i] != 'X' && t.board[i] != 'O' {
			// Create a copy of the current game state
			c := new(TicTacToe)
			*c = *t
			c.board[i] = c.Turn
			score := -c.Minimax()
			if float64(score) > bestScore {
				bestScore = float64(score)
			}
		}
	}
	return int(bestScore)
}

// BestMove calculates the best move using the Minimax algorithm
func (t *TicTacToe) BestMove() int {
	f, err := os.Create("minmax.prof")
	if err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		fmt.Println(err)
	} else {
		defer pprof.StopCPUProfile()
	}

	bestScore := math.Inf(-1)
	move := -1
	for x := 0; x < 1_000; x++ {
		for i := 0; i < 9; i++ {
			if t.board[i] != 'X' && t.board[i] != 'O' {
				// Create a copy of the current game state
				c := new(TicTacToe)
				*c = *t
				c.board[i] = c.Turn
				score := c.Minimax()
				if float64(score) > bestScore {
					bestScore = float64(score)
					move = i
				}
			}
		}
	}
	return move + 1
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
