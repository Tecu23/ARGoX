// Package main contains the startup function and logic of the engine
package main

import (
	"flag"
	"fmt"
)

// TODO: Make a helper function to print to STDIN
// TODO: Fix optimal magic problem
func main() {
	debug := flag.Bool("d", false, "debugging option")

	flag.Parse()

	initHelpers()

	if *debug {
		board := BoardStruct{}
		ParseFEN(&board, CmkPosition)
		board.PrintBoard()
		score := board.EvaluatePosition()
		fmt.Printf("Score:%d\n", score)

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

	InitSliderPiecesAttacks(Bishop) // bishop
	InitSliderPiecesAttacks(Rook)   // rook

	InitRandomHashKeys()
	TransTable.Clear()
	InitEvaluationMasks()
}
