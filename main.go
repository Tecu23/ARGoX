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

	InitMaterialScore()

	debug := false

	if debug {
		board := BoardStruct{}

		ParseFEN(&board, StartPosition)
		board.PrintBoard()
		board.SearchPosition(7)

		fmt.Printf("Score: %d\n", board.EvaluatePosition())
	} else {
		fmt.Println("Starting ARGoX")
		Uci(input())
		fmt.Println("Quiting ARGoX")
	}
}
