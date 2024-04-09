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

// Evaluation Masks
var (
	FileMasks        [64]Bitboard
	RankMasks        [64]Bitboard
	IsolatedMasks    [64]Bitboard
	WhitePassedMasks [64]Bitboard
	BlackPassedMasks [64]Bitboard
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
	score := 0

	bb := Bitboard(0)

	pc, sq := 0, 0

	for bbPc := WP; bbPc <= BK; bbPc++ {
		bb = b.Bitboards[bbPc]

		for bb != 0 {
			pc = bbPc

			sq = bb.FirstOne()

			score += MaterialScore[pc]

			switch pc {
			// evaluare white pieces
			case WP:
				score += PawnScore[sq]
			case WN:
				score += KnightScore[sq]
			case WB:
				score += BishopScore[sq]
			case WR:
				score += RookScore[sq]
			case WK:
				score += KingScore[sq]

			// evaluate black pieces
			case BP:
				score -= PawnScore[MirrorScore[sq]]
			case BN:
				score -= KnightScore[MirrorScore[sq]]
			case BB:
				score -= BishopScore[MirrorScore[sq]]
			case BR:
				score -= RookScore[MirrorScore[sq]]
			case BK:
				score -= KingScore[MirrorScore[sq]]
			}
		}
	}

	if b.SideToMove == WHITE {
		return score
	}

	return -score
}
