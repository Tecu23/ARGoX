package main

/*
         Rank mask            File mask           Isolated mask        Passed pawn mask
       for square a6        for square f2         for square g2          for square c4

   8  0 0 0 0 0 0 0 0    8  0 0 0 0 0 1 0 0    8  0 0 0 0 0 1 0 1     8  0 1 1 1 0 0 0 0
   7  0 0 0 0 0 0 0 0    7  0 0 0 0 0 1 0 0    7  0 0 0 0 0 1 0 1     7  0 1 1 1 0 0 0 0
   6  1 1 1 1 1 1 1 1    6  0 0 0 0 0 1 0 0    6  0 0 0 0 0 1 0 1     6  0 1 1 1 0 0 0 0
   5  0 0 0 0 0 0 0 0    5  0 0 0 0 0 1 0 0    5  0 0 0 0 0 1 0 1     5  0 1 1 1 0 0 0 0
   4  0 0 0 0 0 0 0 0    4  0 0 0 0 0 1 0 0    4  0 0 0 0 0 1 0 1     4  0 0 0 0 0 0 0 0
   3  0 0 0 0 0 0 0 0    3  0 0 0 0 0 1 0 0    3  0 0 0 0 0 1 0 1     3  0 0 0 0 0 0 0 0
   2  0 0 0 0 0 0 0 0    2  0 0 0 0 0 1 0 0    2  0 0 0 0 0 1 0 1     2  0 0 0 0 0 0 0 0
   1  0 0 0 0 0 0 0 0    1  0 0 0 0 0 1 0 0    1  0 0 0 0 0 1 0 1     1  0 0 0 0 0 0 0 0

      a b c d e f g h       a b c d e f g h       a b c d e f g h        a b c d e f g h
*/
// PawnScore keeps the pawn positional score
var pawnScore = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, -10, -10, 0, 0, 0,
	0, 0, 0, 5, 5, 0, 0, 0,
	5, 5, 10, 20, 20, 5, 5, 5,
	10, 10, 10, 20, 20, 10, 10, 10,
	20, 20, 20, 30, 30, 30, 20, 20,
	30, 30, 30, 40, 40, 30, 30, 30,
	90, 90, 90, 90, 90, 90, 90, 90,
}

// KnightScore keeps the knight positional score
var knightScore = [64]int{
	-5, -10, 0, 0, 0, 0, -10, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 5, 20, 10, 10, 20, 5, -5,
	-5, 10, 20, 30, 30, 20, 10, -5,
	-5, 10, 20, 30, 30, 20, 10, -5,
	-5, 5, 20, 20, 20, 20, 5, -5,
	-5, 0, 0, 10, 10, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
}

// BishopScore keeps the bishop positional score
var bishopScore = [64]int{
	0, 0, -10, 0, 0, -10, 0, 0,
	0, 30, 0, 0, 0, 0, 30, 0,
	0, 10, 0, 0, 0, 0, 10, 0,
	0, 0, 10, 20, 20, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

// RookScore keeps the rook positional score
var rookScore = [64]int{
	0, 0, 0, 20, 20, 0, 0, 0,
	0, 0, 10, 20, 20, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 0, 0,
	50, 50, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50,
}

// KingScore keeps the king positional score
var kingScore = [64]int{
	0, 0, 5, 0, -15, 0, 10, 0,
	0, 5, 5, -5, -5, 0, 5, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 5, 10, 20, 20, 10, 5, 0,
	0, 5, 10, 20, 20, 10, 5, 0,
	0, 5, 5, 10, 10, 5, 5, 0,
	0, 0, 5, 5, 5, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

// MirrorScore keeps the mirror positional score tables for opposite side
var mirrorScore = [128]int{
	A8, B8, C8, D8, E8, F8, G8, H8,
	A7, B7, C7, D7, E7, F7, G7, H7,
	A6, B6, C6, D6, E6, F6, G6, H6,
	A5, B5, C5, D5, E5, F5, G5, H5,
	A4, B4, C4, D4, E4, F4, G4, H4,
	A3, B3, C3, D3, E3, F3, G3, H3,
	A2, B2, C2, D2, E2, F2, G2, H2,
	A1, B1, C1, D1, E1, F1, G1, H1,
}

// Evaluation Masks
var (
	FileMasks        [64]Bitboard
	RankMasks        [64]Bitboard
	IsolatedMasks    [64]Bitboard
	WhitePassedMasks [64]Bitboard
	BlackPassedMasks [64]Bitboard
)

// GetRank should extract rank from a square [square]
var getRank = [64]int{
	7, 7, 7, 7, 7, 7, 7, 7,
	6, 6, 6, 6, 6, 6, 6, 6,
	5, 5, 5, 5, 5, 5, 5, 5,
	4, 4, 4, 4, 4, 4, 4, 4,
	3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 0, 0,
}

// PassedPawnBonus is the bonus foe each rank passed
var passedPawnBonus = [8]int{0, 10, 30, 50, 75, 100, 150, 200}

const (
	// Penalties for pawn positions
	doublePawnPenalty   = -10
	isolatedPawnPenalty = -10
	// Open and semi open file scores
	openFileScore     = 15
	semiOpenFileScore = 10
	kingShieldBonus   = 5
)

func setFileRankMask(fileNum, rankNum int) Bitboard {
	mask := Bitboard(0)
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			sq := r*8 + f
			if fileNum != -1 {
				if f == fileNum {
					mask.Set(sq)
					mask |= mask
				}
			} else if rankNum != -1 {
				if r == rankNum {
					mask.Set(sq)
					mask |= mask
				}
			}
		}
	}
	return mask
}

// InitEvaluationMasks should initialize the evaluation masks
func InitEvaluationMasks() {
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			sq := r*8 + f

			FileMasks[sq] |= setFileRankMask(f, -1)
			RankMasks[sq] |= setFileRankMask(-1, r)

			IsolatedMasks[sq] |= setFileRankMask(f-1, -1)
			IsolatedMasks[sq] |= setFileRankMask(f+1, -1)
		}
	}
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			sq := r*8 + f
			WhitePassedMasks[sq] |= setFileRankMask(f-1, -1)
			WhitePassedMasks[sq] |= setFileRankMask(f, -1)
			WhitePassedMasks[sq] |= setFileRankMask(f+1, -1)

			for i := 0; i < r+1; i++ {
				WhitePassedMasks[sq] &= ^RankMasks[i*8+f]
			}

			BlackPassedMasks[sq] |= setFileRankMask(f-1, -1)
			BlackPassedMasks[sq] |= setFileRankMask(f, -1)
			BlackPassedMasks[sq] |= setFileRankMask(f+1, -1)

			for i := r; i < 8; i++ {
				BlackPassedMasks[sq] &= ^RankMasks[i*8+f]
			}
		}
	}
}

// EvaluatePosition should evaluate a certain position
func (b *BoardStruct) EvaluatePosition() int {
	bb := Bitboard(0)
	score := 0
	pc, sq := 0, 0
	doublePawns := 0

	for bbPc := WP; bbPc <= BK; bbPc++ {
		bb = b.Bitboards[bbPc]

		for bb != 0 {
			pc = bbPc
			sq = bb.FirstOne()

			score += MaterialScore[pc]

			switch pc {
			// evaluare white pieces
			case WP:
				score += pawnScore[sq]

				doublePawns = (b.Bitboards[WP] & FileMasks[sq]).Count()
				if doublePawns > 1 {
					score += doublePawns * doublePawnPenalty
				}

				if (b.Bitboards[WP] & IsolatedMasks[sq]) == 0 {
					score += isolatedPawnPenalty
				}

				if (WhitePassedMasks[sq] & b.Bitboards[BP]) == 0 {
					score += passedPawnBonus[getRank[sq]]
				}

			case WN:
				score += knightScore[sq]
			case WB:
				score += bishopScore[sq]

				score += getBishopAttacks(sq, b.Occupancies[BOTH]).Count()
			case WR:
				score += rookScore[sq]

				if (b.Bitboards[WP] & FileMasks[sq]) == 0 {
					score += semiOpenFileScore
				}

				if ((b.Bitboards[WP] | b.Bitboards[BP]) & FileMasks[sq]) == 0 {
					score += openFileScore
				}
			case WQ:
				score += getQueenAttacks(sq, b.Occupancies[BOTH]).Count()
			case WK:
				score += kingScore[sq]

				if (b.Bitboards[WP] & FileMasks[sq]) == 0 {
					score -= semiOpenFileScore
				}

				if ((b.Bitboards[WP] | b.Bitboards[BP]) & FileMasks[sq]) == 0 {
					score -= openFileScore
				}

				score += (KingAttacks[sq] & b.Occupancies[WHITE]).Count() * kingShieldBonus

			// evaluate black pieces
			case BP:
				score -= pawnScore[mirrorScore[sq]]

				doublePawns = (b.Bitboards[BP] & FileMasks[sq]).Count()
				if doublePawns > 1 {
					score -= doublePawns * doublePawnPenalty
				}

				if (b.Bitboards[BP] & IsolatedMasks[sq]) == 0 {
					score -= isolatedPawnPenalty
				}

				if (BlackPassedMasks[sq] & b.Bitboards[WP]) == 0 {
					score -= passedPawnBonus[getRank[mirrorScore[sq]]]
				}
			case BN:
				score -= knightScore[mirrorScore[sq]]
			case BB:
				score -= bishopScore[mirrorScore[sq]]

				score -= getBishopAttacks(sq, b.Occupancies[BOTH]).Count()
			case BR:
				score -= rookScore[mirrorScore[sq]]

				if (b.Bitboards[BP] & FileMasks[sq]) == 0 {
					score -= semiOpenFileScore
				}

				if ((b.Bitboards[WP] | b.Bitboards[BP]) & FileMasks[sq]) == 0 {
					score -= openFileScore
				}
			case BQ:
				score += getQueenAttacks(sq, b.Occupancies[BOTH]).Count()

			case BK:
				score -= kingScore[mirrorScore[sq]]

				if (b.Bitboards[BP] & FileMasks[sq]) == 0 {
					score += semiOpenFileScore
				}

				if ((b.Bitboards[WP] | b.Bitboards[BP]) & FileMasks[sq]) == 0 {
					score += openFileScore
				}

				score -= (KingAttacks[sq] & b.Occupancies[BLACK]).Count() * kingShieldBonus
			}
		}
	}
	if b.SideToMove == WHITE {
		return score
	}
	return -score
}
