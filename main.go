// Package main contains the startup function and logic of the engine
package main

func main() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()

	for i := A1; i <= H8; i++ {
		b := generateRookAttacks(i)
		b.PrintBitboard()
	}
}
