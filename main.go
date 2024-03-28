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

	ParseFEN(&board, "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1 ")

	move := Move(0)
	move.EncodeMove(E2, E4, WP, WQ, 1, 1, 1, 1)

	fmt.Printf(
		"%s %s %c %c %d %d %d %d",
		Sq2Fen[move.GetSource()],
		Sq2Fen[move.GetTarget()],
		AciiPieces[move.GetPiece()],
		AciiPieces[move.GetPromoted()],
		move.GetCapture(),
		move.GetDoublePush(),
		move.GetEnpassant(),
		move.GetCastling(),
	)
}
