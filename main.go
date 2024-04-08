// Package main contains the startup function and logic of the engine
package main

import (
	"flag"
	"fmt"
)

// TODO: Makefile
// TODO: Make a helper function to print to STDIN
// TODO: Fix optimal magic problem
func main() {
	debug := flag.Bool("d", false, "debugging option")

	flag.Parse()

	initHelpers()

	if *debug {
		board := BoardStruct{}

		ParseFEN(&board, TrickyPosition)
		board.PrintBoard()

	} else {
		fmt.Println("Starting ARGoX")
		Uci(input())
		fmt.Println("Quiting ARGoX")
	}
}

func initHelpers() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()
	InitFen2Sq()
	// FillOptimalMagicsB()
	// FillOptimalMagicsR()

	GenerateSliderPiecesAttacks(Bishop) // bishop
	GenerateSliderPiecesAttacks(Rook)   // rook

	InitMaterialScore()

	InitRandomHashKeys()
}
