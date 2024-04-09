package main

import (
	"fmt"
	"math/bits"
)

// Ply is the half move counter
var Ply int

var nodes int64

// LMR constants
const (
	FullDepthMoves = 4
	ReductionLimit = 3
)

// Score bounds for the range of the mating scores
const (
	Infinity  = 50000
	MateValue = 49000
	MateScore = 48000
)

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

/*  =======================
         Move ordering
    =======================

    1. PV move
    2. Captures in MVV/LVA
    3. 1st killer move
    4. 2nd killer move
    5. History moves
    6. Unsorted moves

*/

func (b *BoardStruct) scoreMove(mv Move) int {
	if ScorePv {
		if PvTable[0][Ply] == mv {
			ScorePv = false
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
	if nodes&2047 == 0 {
		// stop the search if the time passed
		if limits.Timeset && GetTimeInMiliseconds() > limits.StopTime {
			limits.setStop(true)
		}
	}
	nodes++
	if Ply > MaxPly-1 {
		return b.EvaluatePosition()
	}

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

		if limits.Stop {
			return 0
		}

		if score > alpha {
			alpha = score // PV node (move)

			// fail-hard beta cutoff
			if score >= beta {
				return beta // node (move) fails high
			}
		}
	}
	return alpha // node (move) fails low
}

// negamax alpha beta search
func (b *BoardStruct) negamax(alpha, beta, depth int) int {
	score := 0

	hashFlag := HashfAlpha

	// hack to figure out wether the curr node is PV or not
	pvNode := beta-alpha > 1

	// read hash entry and if the move has already been searched => return the score
	score = TransTable.ReadEntry(alpha, beta, depth, b.Key)
	if Ply > 0 && score != NoHashEntry && !pvNode {
		return score
	}

	if nodes&2047 == 0 {
		// stop the search if the time passed
		if limits.Timeset && GetTimeInMiliseconds() > limits.StopTime {
			limits.setStop(true)
		}
	}

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

	// Null Move Pruning
	if depth >= 3 && !inCheck && Ply != 0 {
		copyBoard := b.CopyBoard()
		Ply++
		if b.EnPassant != -1 {
			b.Key ^= EnpassantKeys[b.EnPassant]
		}

		b.EnPassant = -1
		b.SideToMove = b.SideToMove.Opp()
		b.Key ^= SideKey
		// find beta cuttoffs
		sc := -b.negamax(-beta, -beta+1, depth-1-2) // depth - 1 -R where R is reduction limit

		Ply--
		b.TakeBack(copyBoard)

		if limits.Stop {
			return 0
		}

		if sc >= beta {
			return beta
		}
	}

	// generate moves
	var moves Movelist
	b.generateMoves(&moves)

	if FollowPv {
		b.enablePvScoring(moves) // enable PV move scoring
	}

	b.sortMoves(moves)
	movesSearched := 0

	for _, m := range moves {
		copyB := b.CopyBoard()
		Ply++

		if !b.MakeMove(m, AllMoves) {
			Ply--
			continue
		}
		legalMoves++

		if movesSearched == 0 {
			score = -b.negamax(-beta, -alpha, depth-1) // do normal search
		} else {
			// Late Move Reduction
			if movesSearched >= FullDepthMoves && depth >= ReductionLimit && !inCheck && m.GetCapture() == 0 && m.GetPromoted() == 0 {
				score = -b.negamax(-alpha-1, -alpha, depth-2)
			} else {
				score = alpha + 1 // Hack to ensure the full-depth search
			}

			if score > alpha { // aspiration window
				// if move > alpha & move < beta => prove rest of moves are bad
				score = -b.negamax(-alpha-1, -alpha, depth-1)

				// if the alg was wrong then it needs to search again in the normal matter
				if score > alpha && score < beta {
					score = -b.negamax(-beta, -alpha, depth-1)
				}
			}
		}

		Ply--
		b.TakeBack(copyB)
		if limits.Stop {
			return 0
		}
		movesSearched++

		if score > alpha {
			hashFlag = HashfExact // switch from fail low to exact

			if m.GetCapture() == 0 {
				HistoryMove[m.GetPiece()][m.GetTarget()] += depth
			}
			alpha = score         // PV node (move)
			PvTable[Ply][Ply] = m // write PV move

			for nextPly := Ply + 1; nextPly < PvLength[Ply+1]; nextPly++ {
				PvTable[Ply][nextPly] = PvTable[Ply+1][nextPly] // copy from deeper ply to current ply
			}
			PvLength[Ply] = PvLength[Ply+1] // adjust PV length

			// fail-hard beta cutoff
			if score >= beta {
				// store hash entry with the score equal to beta
				TransTable.WriteEntry(beta, depth, HashfBeta, b.Key)
				if m.GetCapture() == 0 {
					KillerMove[1][Ply] = KillerMove[0][Ply]
					KillerMove[0][Ply] = m
				}
				return beta // node (move) fails high
			}
		}
	}

	if legalMoves == 0 {
		if inCheck {
			return -MateValue + Ply
		}
		return 0 // if not check then stalemate
	}

	TransTable.WriteEntry(alpha, depth, hashFlag, b.Key)

	return alpha // node (move) fails low
}

// SearchPosition should search the current board position for the best move
func (b *BoardStruct) SearchPosition(depth int) Move {
	nodes = 0

	limits.setStop(false)

	// clear helper data structures for search
	KillerMove = [2][64]Move{}
	HistoryMove = [12][64]int{}
	PvLength = [64]int{}
	PvTable = [64][64]Move{}

	FollowPv = false
	ScorePv = false

	score := 0

	alpha := -Infinity
	beta := Infinity

	for currDepth := 1; currDepth <= depth; currDepth++ {
		FollowPv = true // enable FollowPv

		if limits.Stop {
			break
		}

		// find best move within given position
		score = b.negamax(alpha, beta, currDepth)

		if score <= alpha || score >= beta {
			alpha = -Infinity
			beta = Infinity
			continue
		}

		alpha = score - 50
		beta = score + 50

		if score > -MateScore && score < -MateScore {
			fmt.Printf(
				"info score mate %d depth %d nodes %d pv ",
				-(score+MateValue)/2-1,
				currDepth,
				nodes,
			)
		} else if score > MateScore && score < MateValue {
			fmt.Printf("info score mate %d depth %d nodes %d pv ", (MateValue-score)/2+1, currDepth, nodes)
		} else {
			fmt.Printf("info score cp %d depth %d nodes %d pv ", score, currDepth, nodes)
		}

		for i := 0; i < PvLength[0]; i++ {
			fmt.Printf("%s ", PvTable[0][i])
		}
		fmt.Printf("\n")

	}

	// best move placeholder
	return PvTable[0][0]
}
