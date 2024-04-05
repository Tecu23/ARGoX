package main

import (
	"fmt"
	"math/bits"
)

// Ply is the half move counter
var Ply int

// TODO: Figure out why there are more nodes than it should (1241 vs 1067)
var nodes int64

func (b *BoardStruct) enablePvScoring(mvlist Movelist) {
	FollowPv = false

	for _, m := range mvlist {
		if PvTable[0][Ply] == m {
			ScorePv = true
			FollowPv = true
		}
	}
}

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
	if ScorePv {
		if PvTable[0][Ply] == mv {
			ScorePv = false

			fmt.Printf("current Pv move: %s ply: %d\n", mv, Ply)

			return 20000 // give PV the highest score to search it first
		}
	}

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

		return MvvLva[mv.GetPiece()][tgtPc] + 10000 // score move by MVV LVA lookup [source][target]

	}

	// quiet moves
	if KillerMove[0][Ply] == mv { // score first killer move
		return 9000
	} else if KillerMove[1][Ply] == mv {
		return 8000 // second killer move
	}

	return HistoryMove[mv.GetPiece()][mv.GetTarget()] // score history move
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
	PvLength[Ply] = Ply // init PV length

	if depth == 0 { // base case
		return b.quiescence(alpha, beta)
	}

	if Ply > MaxPly-1 {
		return b.EvaluatePosition()
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

	legalMoves := 0

	// generate moves
	var moves Movelist
	b.generateMoves(&moves)

	if FollowPv {
		b.enablePvScoring(moves) // enable PV move scoring
	}

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

			if m.GetCapture() == 0 {
				KillerMove[1][Ply] = KillerMove[0][Ply]
				KillerMove[0][Ply] = m
			}

			return beta // node (move) fails high
		}

		if score > alpha {
			if m.GetCapture() == 0 {
				HistoryMove[m.GetPiece()][m.GetTarget()] += depth
			}
			alpha = score // PV node (move)

			PvTable[Ply][Ply] = m // write PV move

			for nextPly := Ply + 1; nextPly < PvLength[Ply+1]; nextPly++ {
				PvTable[Ply][nextPly] = PvTable[Ply+1][nextPly] // copy from deeper ply to current ply
			}

			PvLength[Ply] = PvLength[Ply+1] // adjust PV length
		}
	}

	if legalMoves == 0 {
		if inCheck {
			return -49000 + Ply
		}
		return 0 // if not check then stalemate
	}

	return alpha // node (move) fails low
}

// SearchPosition should search the current board position for the best move
func (b *BoardStruct) SearchPosition(depth int) {
	nodes = 0
	// clear helper data structures for search
	KillerMove = [2][64]Move{}
	HistoryMove = [12][64]int{}
	PvLength = [64]int{}
	PvTable = [64][64]Move{}

	FollowPv = false
	ScorePv = false

	score := 0

	for currDepth := 1; currDepth <= depth; currDepth++ {
		nodes = 0

		FollowPv = true // enable FollowPv

		// find best move within given position
		score = b.negamax(-50000, 50000, currDepth)

		fmt.Printf("info score cp %d depth %d nodes %d pv ", score, currDepth, nodes)

		for i := 0; i < PvLength[0]; i++ {
			fmt.Printf("%s ", PvTable[0][i])
		}
		fmt.Printf("\n")

	}

	// best move placeholder
	fmt.Printf("bestmove %s\n", PvTable[0][0])
	nodes = 0
	// clear helper data structures for search
	KillerMove = [2][64]Move{}
	HistoryMove = [12][64]int{}
	PvLength = [64]int{}
	PvTable = [64][64]Move{}

	FollowPv = false
	ScorePv = false

	// find best move within given position
	score = b.negamax(-50000, 50000, depth)

	fmt.Printf("info score cp %d depth %d nodes %d pv ", score, depth, nodes)

	for i := 0; i < PvLength[0]; i++ {
		fmt.Printf("%s ", PvTable[0][i])
	}
	fmt.Printf("\n")

	// best move placeholder
	fmt.Printf("bestmove %s\n", PvTable[0][0])
}
