// Package main contains the startup function and logic of the engine
package main

import "fmt"

func main() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()
	InitFen2Sq()
	// FillOptimalMagicsB()
	// FillOptimalMagicsR()

	GenerateSliderPiecesAttacks(Bishop) // bishop
	GenerateSliderPiecesAttacks(Rook)   // rook

	board := BoardStruct{}
	ParseFEN(&board, TrickyPosition)

	board.PrintBoard()

	move := board.ParseMove("d5d6")

	if move != NoMove {
		board.MakeMove(move, AllMoves)
		board.PrintBoard()
	} else {
		fmt.Print("illegal move")
	}
}
