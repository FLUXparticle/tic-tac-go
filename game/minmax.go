package game

import "math"

const (
	red   = "\033[31m"
	blue  = "\033[34m"
	reset = "\033[0m"
)

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
		if t.board[i] != "X" && t.board[i] != "O" {
			// Create a copy of the current game state
			c := new(TicTacToe)
			*c = *t
			c.board[i] = c.turn
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
	bestScore := math.Inf(-1)
	move := -1
	for i := 0; i < 9; i++ {
		if t.board[i] != "X" && t.board[i] != "O" {
			// Create a copy of the current game state
			c := new(TicTacToe)
			*c = *t
			c.board[i] = c.turn
			score := c.Minimax()
			if float64(score) > bestScore {
				bestScore = float64(score)
				move = i
			}
		}
	}
	return move + 1
}
