package main

import (
	"fmt"
	"tic-tac-go/game"
)

func main() {
	g := game.NewTicTacToe()
	var position int
	var err error

	for {
		g.PrintBoard()
		if g.Turn == 'X' {
			fmt.Printf("Player %s, enter your move (1-9): ", g.Turn)
			_, err = fmt.Scanf("%d\n", &position)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number between 1 and 9.")
				fmt.Scanf("%s\n", new(string)) // clear the invalid input
				continue
			}
		} else {
			fmt.Println("Computer is thinking...")
			position = g.BestMove()
			fmt.Printf("Computer chose position %d\n", position)
		}

		if !g.MakeMove(position) {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		if g.CheckWin() {
			g.PrintBoard()
			fmt.Printf("Player %s wins!\n", g.Turn)
			break
		}

		if g.CheckDraw() {
			g.PrintBoard()
			fmt.Println("It's a draw!")
			break
		}

		g.SwitchTurn()
	}

	fmt.Println("Game over.")
}
