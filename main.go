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
		ParseFEN(&board, StartPosition)
		board.SearchPosition(10)
		// TransTable.PrintAll()

		board.MakeMove(PvTable[0][0], AllMoves)

		board.SearchPosition(10)

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

	TransTable.Clear()
}
