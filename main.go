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

	board.Bitboards[WP].Set(A2)
	board.Bitboards[WP].Set(B2)
	board.Bitboards[WP].Set(C2)
	board.Bitboards[WP].Set(D2)
	board.Bitboards[WP].Set(E2)
	board.Bitboards[WP].Set(F2)
	board.Bitboards[WP].Set(G2)
	board.Bitboards[WP].Set(H2)

	board.Bitboards[WN].Set(B1)
	board.Bitboards[WN].Set(G1)

	board.Bitboards[WB].Set(C1)
	board.Bitboards[WB].Set(F1)

	board.Bitboards[WR].Set(A1)
	board.Bitboards[WR].Set(H1)

	board.Bitboards[WQ].Set(D1)

	board.Bitboards[WK].Set(E1)

	board.Bitboards[BP].Set(A7)
	board.Bitboards[BP].Set(B7)
	board.Bitboards[BP].Set(C7)
	board.Bitboards[BP].Set(D7)
	board.Bitboards[BP].Set(E7)
	board.Bitboards[BP].Set(F7)
	board.Bitboards[BP].Set(G7)
	board.Bitboards[BP].Set(H7)

	board.Bitboards[BN].Set(B8)
	board.Bitboards[BN].Set(G8)

	board.Bitboards[BB].Set(C8)
	board.Bitboards[BB].Set(F8)

	board.Bitboards[BR].Set(A8)
	board.Bitboards[BR].Set(H8)

	board.Bitboards[BQ].Set(D8)

	board.Bitboards[BK].Set(E8)

	board.SideToMove = BLACK

	board.EnPassant = E3

	board.Castlings |= Castlings(ShortW)
	board.Castlings |= Castlings(LongW)
	board.Castlings |= Castlings(ShortB)
	board.Castlings |= Castlings(LongB)

	board.PrintBoard()

	// for _, bb := range board.Bitboards {
	// 	bb.PrintBitboard()
	// }
}
