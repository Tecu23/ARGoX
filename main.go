// Package main contains the startup function and logic of the engine
package main

import "fmt"

func main() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()
	InitFen2Sq()

	// for i := A1; i <= H8; i++ {
	// 	b := generateRookAttacks(i)
	// 	b.PrintBitboard()
	// }

	// init occupancy bitboard
	block := Bitboard(0)
	block.Set(D5)
	block.Set(D2)
	block.Set(C4)
	block.Set(E4)
	block.PrintBitboard()

	i := block.LastOne()
	fmt.Println(i)
	block.PrintBitboard()

	i = block.FirstOne()
	fmt.Println(i)
	block.PrintBitboard()

	b := generateRookAttacksOnTheFly(D4, block)
	b.PrintBitboard()
}
