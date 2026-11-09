package main

import (
	"testing"

	"github.com/FLUXparticle/tic-tac-go/game"
)

func BenchmarkBestMove(b *testing.B) {
	g := game.NewGame()
	g.MakeMove(5)
	g.SwitchTurn()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.BestMove()
	}
}
