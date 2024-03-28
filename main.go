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

	ParseFEN(&board, "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1 ")

	move := Move(0)
	move.EncodeMove(E2, E4, WP, WQ, 0, 0, 0, 0)

	var list Movelist

	list.AddMove(move)
	list.PrintMovelist()
}
