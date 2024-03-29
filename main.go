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

	start := GetTimeInMiliseconds()
	var mvlst Movelist
	board.generateMoves(&mvlst)

	for _, mv := range mvlst {
		copyB := board.CopyBoard()

		if !board.MakeMove(mv, AllMoves) {
			continue
		}
		board.TakeBack(copyB)
	}
	fmt.Printf("Time taken to execute in ms: %d\n", GetTimeInMiliseconds()-start)
}
