package main

import (
	"fmt"
	"math/bits"
)

// half move counter
var Ply int

// best move
var bestMove Move

var nodes int

func (b *BoardStruct) quiescence(alpha, beta int) int {
	evaluation := b.EvaluatePosition()

	if evaluation >= beta {
		return beta
	}

	if evaluation > alpha {
		alpha = evaluation
	}

	// generate moves
	var moves Movelist
	b.generateMoves(&moves)

	for _, m := range moves {
		copyB := b.CopyBoard()

		Ply++
		if !b.MakeMove(m, OnlyCaptures) {
			Ply--
			continue
		}
		score := -b.quiescence(-beta, -alpha) // score current move

		Ply--

		b.TakeBack(copyB)

		// fail-hard beta cutoff
		if score >= beta {
			return beta // node (move) fails high
		}

		if score > alpha {
			alpha = score // PV node (move)
		}
	}
	return alpha // node (move) fails low
}

// negamax alpha beta search
func (b *BoardStruct) negamax(alpha, beta, depth int) int {
	if depth == 0 { // base case
		return b.quiescence(alpha, beta)
	}

	nodes++ // traversed nodes

	var bit int
	if b.SideToMove == WHITE {
		bit = bits.TrailingZeros64(uint64(b.Bitboards[WK]))
	} else {
		bit = bits.TrailingZeros64(uint64(b.Bitboards[BK]))
	}

	inCheck := b.isSquareAttacked(bit, b.SideToMove.Opp())

	legalMoves, bestSoFar := 0, NoMove

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
		legalMoves++

		score := -b.negamax(-beta, -alpha, depth-1) // score current move
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
				bestSoFar = m // Associate best move with the best score
			}
		}
	}

	if legalMoves == 0 {
		if inCheck {
			return -49000 + Ply
		}
		return 0 // if not check then stalemate
	}

	if oldAlpha != alpha {
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
