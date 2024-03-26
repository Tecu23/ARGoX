// Package main contains the startup function and logic of the engine
package main

func main() {
	InitPawnAttacks()
	InitKnightAttacks()
	InitKingAttacks()
	InitFen2Sq()

	InitMagic()

	// for sq := A1; sq <= H8; sq++ {
	// 	fmt.Printf(" %d", RookMagic)
	// }
}
