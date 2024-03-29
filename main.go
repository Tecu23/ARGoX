// Package main contains the startup function and logic of the engine
package main

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
	ParseFEN(&board, "r3k2r/p1ppqpb1/1n2pnp1/3PN3/1p2P3/2N2Q1p/PPPBrPPP/R3K2R w KQkq - 0 1 ")
	board.PrintBoard()

	var mvlst Movelist
	board.generateMoves(&mvlst)

	for _, mv := range mvlst {
		copyB := board.CopyBoard()

		mv.PrintMove()
		if !board.MakeMove(mv, AllMoves) {
			continue
		}
		board.PrintBoard()
		board.TakeBack(copyB)
		board.PrintBoard()
	}
}
