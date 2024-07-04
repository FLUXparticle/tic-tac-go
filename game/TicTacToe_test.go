package game

func Example_PrintBoard() {
	t := NewTicTacToe()
	t.PrintBoard()
	// Output:
	// -------------
	// | 1 | 2 | 3 |
	// -------------
	// | 4 | 5 | 6 |
	// -------------
	// | 7 | 8 | 9 |
	// -------------
}
