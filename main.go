// Package main contains the startup function and logic of the engine
package main

import "fmt"

func main() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()
	InitFen2Sq()

	for sq := A1; sq <= H8; sq++ {
		b := generateBishopAttacks(sq)
		fmt.Printf(" %d", b.Count())
	}
}
