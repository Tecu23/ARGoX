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
	ParseFEN(&board, "r3k2r/p1ppqpb1/bn2pnp1/3PN3/Pp2P3/2N2Q1p/1PPBBPPP/R3K2R b KQkq a3 0 1 ")

	board.PrintBoard()

	var mvlst Movelist
	board.generateMoves(&mvlst)

	for _, mv := range mvlst {
		copyB := board.CopyBoard()

		if mv.GetEnpassant() != 0 {
			mv.PrintMove()
			board.PrintBoard()
			board.MakeMove(mv, AllMoves)
			board.PrintBoard()
			board.TakeBack(copyB)
			board.PrintBoard()
		}

	}
}
