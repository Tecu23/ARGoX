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

	debug := true

	if debug {
		board := BoardStruct{}

		ParseFEN(&board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 ")

		board.PrintBoard()

		fmt.Printf("Score: %d\n", board.EvaluatePosition())
	} else {
		fmt.Println("Starting ARGoX")
		Uci(input())
		fmt.Println("Quiting ARGoX")
	}
}
