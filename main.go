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

	// board := BoardStruct{}

	occ := Bitboard(0)
	occ.Set(B6)
	occ.Set(D6)
	occ.Set(F6)

	b := getQueenAttacks(D4, occ)
	b.PrintBitboard()
}
