package game_test

import "github.com/FLUXparticle/tic-tac-go/game"

func ExampleTicTacToe_PrintBoard() {
	g := game.NewGame()
	g.PrintBoard()
	// Output:
	// -------------
	// | 1 | 2 | 3 |
	// -------------
	// | 4 | 5 | 6 |
	// -------------
	// | 7 | 8 | 9 |
	// -------------
}
