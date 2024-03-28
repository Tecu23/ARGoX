// Package main contains the startup function and logic of the engine
package main

func main() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()
	InitFen2Sq()
	FillOptimalMagicsB()
	FillOptimalMagicsR()

	GenerateSliderPiecesAttacks(Bishop) // bishop
	GenerateSliderPiecesAttacks(Rook)   // rook

	board := BoardStruct{}

	ParseFEN(&board, TrickyPosition)
	board.SideToMove = BLACK
	board.SetSq(BP, G2)
	board.SetSq(Empty, A2)
	board.SetSq(WP, A4)
	board.EnPassant = A3

	board.PrintBoard()

	board.generateMoves()
}
