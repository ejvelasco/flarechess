package main

import (
	"github.com/notnil/chess"
	"log"
)

func main() {
	game := chess.NewGame()
	moves := game.ValidMoves()
	log.Println(moves)
}
