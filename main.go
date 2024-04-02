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

	err := board.ParsePosition(
		"position startpos moves e2e4 e7e5 g1f3",
	)
	if err != nil {
		fmt.Println(err)
	}

	board.PrintBoard()

	err = board.ParseGo("go depth 7")
	if err != nil {
		fmt.Println(err)
	}
}
