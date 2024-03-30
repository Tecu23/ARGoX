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
	ParseFEN(&board, TrickyPosition)

	// for i := 1; i <= 8; i++ {
	// 	Nodes = 0
	// 	start := GetTimeInMiliseconds()
	// 	perftDriver(&board, i)
	// 	fmt.Printf("Time taken to execute in ms: %d\n", GetTimeInMiliseconds()-start)
	// 	fmt.Printf("Nodes: %d\n\n", Nodes)
	// }

	perftTest(&board, 5)
}
