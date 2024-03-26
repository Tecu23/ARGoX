// Package main contains the startup function and logic of the engine
package main

import "fmt"

func main() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()
	InitFen2Sq()
	FillOptimalMagicsB()
	FillOptimalMagicsR()

	generateSliderPieces(Bishop) // bishop
	generateSliderPieces(Rook)   // rook

	bishopOccupancy := Bitboard(0)
	bishopOccupancy.Set(G7)
	bishopOccupancy.Set(F6)
	bishopOccupancy.Set(C5)
	bishopOccupancy.Set(B2)
	bishopOccupancy.Set(G1)

	fmt.Printf("\n    Bishop occupancy\n")
	bishopOccupancy.PrintBitboard()

	fmt.Printf("\n    Bishop Attacks\n")
	b := getBishopAttacks(D4, bishopOccupancy)
	b.PrintBitboard()

	rookOccupancy := Bitboard(0)
	rookOccupancy.Set(D7)
	rookOccupancy.Set(D6)
	rookOccupancy.Set(D3)
	rookOccupancy.Set(A4)
	rookOccupancy.Set(F4)

	fmt.Printf("\n    Rook occupancy\n")
	rookOccupancy.PrintBitboard()

	fmt.Printf("\n    Rook Attacks\n")
	b = getRookAttacks(D4, rookOccupancy)
	b.PrintBitboard()
}
