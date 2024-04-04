// Package main contains the startup function and logic of the engine
package main

import "fmt"

// TODO: Make debug as a flag
// TODO: Fix optimal magic problem
func main() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()
	InitFen2Sq()
	// FillOptimalMagicsB()
	// FillOptimalMagicsR()

	GenerateSliderPiecesAttacks(Bishop) // bishop
	GenerateSliderPiecesAttacks(Rook)   // rook

	InitMaterialScore()

	debug := true

	if debug {
		board := BoardStruct{}

		ParseFEN(&board, TrickyPosition)
		board.SearchPosition(1)

	} else {
		fmt.Println("Starting ARGoX")
		Uci(input())
		fmt.Println("Quiting ARGoX")
	}
}
