package main

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
				c.board[i] = c.turn
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
