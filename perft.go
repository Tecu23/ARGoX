package main

import "fmt"

// Nodes are the leaf nodes (number of positions reached during the test of the
//
//	move generator at a given depth)
var Nodes int64

// perft driver
func perftDriver(b *BoardStruct, depth int) {
	// base cases
	if depth == 0 {
		Nodes++
		return
	}

	// create move list instance
	var mvlst Movelist

	// generate moves
	b.GenerateMoves(&mvlst)

	for _, mv := range mvlst {
		copyB := b.CopyBoard()

		if !b.MakeMove(mv, AllMoves) {
			continue
		}

		perftDriver(b, depth-1)

		// take back move
		b.TakeBack(copyB)
	}
}

func PerftTest(b *BoardStruct, depth int) {
	var whiteMoves Movelist

	totalMoves := int64(0)

	b.GenerateMoves(&whiteMoves)

	fmt.Printf("\n  Performance test\n\n")
	start := GetTimeInMiliseconds()

	for _, m := range whiteMoves {
		Nodes = 0

		copyB := b.CopyBoard()

		if !b.MakeMove(m, AllMoves) {
			continue
		}
		perftDriver(b, depth-1)

		// take back move
		b.TakeBack(copyB)

		fmt.Printf("%s: %d\n", m, Nodes)

		totalMoves += Nodes
	}
	// print results
	fmt.Printf("\n Nodes: %d Time: %d\n\n ", totalMoves, GetTimeInMiliseconds()-start)
}
