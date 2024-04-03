package main

import "fmt"

// half move counter
var Ply int

// best move
var bestMove Move

var nodes int

// negamax alpha beta search
func (b *BoardStruct) negamax(alpha, beta, depth int) int {
	// recursion base case
	if depth == 0 {
		return b.EvaluatePosition()
	}

	// increment nodes traversed count
	nodes++

	var bestSoFar Move

	// old value of alpha -> temporary
	oldAlpha := alpha

	// generate moves
	var moves Movelist
	b.generateMoves(&moves)

	for _, m := range moves {

		copyB := b.CopyBoard()

		Ply++

		if !b.MakeMove(m, AllMoves) {
			Ply--

			continue
		}

		// score current move
		score := -b.negamax(-beta, -alpha, depth-1)

		// fmt.Println(m, score)

		Ply--

		b.TakeBack(copyB)

		// fail-hard beta cutoff
		if score >= beta {
			return beta // node (move) fails high
		}

		if score > alpha {
			alpha = score // PV node (move)

			// if root move
			if Ply == 0 {
				// fmt.Println("Updating Best Move So Far", m)
				bestSoFar = m // Associate best move with the best score
			}
		}
	}

	if oldAlpha != alpha {
		// fmt.Println(bestSoFar)
		bestMove = bestSoFar
	}

	return alpha // node (move) fails low
}

// SearchPosition should search the current board position for the best move
func (b *BoardStruct) SearchPosition(depth int) {
	// find best move within given position
	b.negamax(-50000, 50000, depth)

	fmt.Printf("bestmove %s\n", bestMove)
}
