package main

import (
	"fmt"
	"math/bits"
)

// Ply is the half move counter
var Ply int

// best move
var bestMove Move

// TODO: Figure out why there are more nodes than it should (1241 vs 1067)
var nodes int64

func (b *BoardStruct) sortMoves(mvlist Movelist) {
	// TODO: Add a faster sorting algorithm

	var moveScores []int

	for _, m := range mvlist {
		moveScores = append(moveScores, b.scoreMove(m))
	}

	for i := 0; i < len(moveScores)-1; i++ {
		for j := i + 1; j < len(moveScores); j++ {
			if moveScores[i] < moveScores[j] {
				moveScores[i], moveScores[j] = moveScores[j], moveScores[i]
				mvlist[i], mvlist[j] = mvlist[j], mvlist[i]
			}
		}
	}
}

func (b *BoardStruct) scoreMove(mv Move) int {
	if mv.GetCapture() != 0 {
		tgtPc := 0
		tgtSq := mv.GetTarget()

		var startPc, endPc int

		if b.SideToMove == BLACK {
			startPc = WP
			endPc = WK
		} else {
			startPc = BP
			endPc = BK
		}

		for p := startPc; p <= endPc; p++ {
			if b.Bitboards[p].Test(tgtSq) {
				tgtPc = p
				break
			}
		}

		return MvvLva[mv.GetPiece()][tgtPc] // score move by MVV LVA lookup [source][target]

	} else { // quiet move
	}

	return 0
}

// ListScoreMoves should list all scores for all moves in a given position
func (b *BoardStruct) ListScoreMoves(mvlist Movelist) {
	fmt.Printf("   Move Scores: \n")
	for _, m := range mvlist {
		fmt.Printf("   %s  Score: %d\n", m, b.scoreMove(m))
	}
}

func (b *BoardStruct) quiescence(alpha, beta int) int {
	nodes++

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

	b.sortMoves(moves)

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

	// increase search depth if the king has been exposed to a check
	if inCheck {
		depth++
	}

	legalMoves, bestSoFar := 0, NoMove

	oldAlpha := alpha

	// generate moves
	var moves Movelist
	b.generateMoves(&moves)

	b.sortMoves(moves)

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
	score := b.negamax(-50000, 50000, depth)

	fmt.Printf("bestmove %s\n", bestMove)
	if bestMove != 0 {
		fmt.Printf("info score cp %d depth %d nodes %d\n", score, depth, nodes)

		// best move placeholder
		fmt.Printf("bestmove %s\n", bestMove)
	}
}
