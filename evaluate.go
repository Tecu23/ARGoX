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

// MaterialScore [game phase][piece]
var MaterialScore = [2][12]int{
	// opening material score
	{82, 337, 365, 477, 1025, 12000, -82, -337, -365, -477, -1025, -12000},

	// endgame material score
	{94, 281, 297, 512, 936, 12000, -94, -281, -297, -512, -936, -12000},
}

// OpeningPhaseScore represents the opening game phase scores
const OpeningPhaseScore = 6192

// EndgamePhaseScore represents the endgame game phase scores
const EndgamePhaseScore = 518

// game phases
const (
	Opening int = iota
	Endgame
	Middlegame
)

// PositionalScore [game phase][piece][square]
var PositionalScore = [2][6][64]int{

	// opening positional piece scores //
	//pawn
	{
		{
			0, 0, 0, 0, 0, 0, 0, 0,
			-35, -1, -20, -23, -15, 24, 38, -22,
			-26, -4, -4, -10, 3, 3, 33, -12,
			-27, -2, -5, 12, 17, 6, 10, -25,
			-14, 13, 6, 21, 23, 12, 17, -23,
			-6, 7, 26, 31, 65, 56, 25, -20,
			98, 134, 61, 95, 68, 126, 34, -11,
			0, 0, 0, 0, 0, 0, 0, 0,
		},

		// knight
		{
			-105, -21, -58, -33, -17, -28, -19, -23,
			-29, -53, -12, -3, -1, 18, -14, -19,
			-23, -9, 12, 10, 19, 17, 25, -16,
			-13, 4, 16, 13, 28, 19, 21, -8,
			-9, 17, 19, 53, 37, 69, 18, 22,
			-47, 60, 37, 65, 84, 129, 73, 44,
			-73, -41, 72, 36, 23, 62, 7, -17,
			-167, -89, -34, -49, 61, -97, -15, -107,
		},

		// bishop
		{
			-33, -3, -14, -21, -13, -12, -39, -21,
			4, 15, 16, 0, 7, 21, 33, 1,
			0, 15, 15, 15, 14, 27, 18, 10,
			-6, 13, 13, 26, 34, 12, 10, 4,
			-4, 5, 19, 50, 37, 37, 7, -2,
			-16, 37, 43, 40, 35, 50, 37, -2,
			-26, 16, -18, -13, 30, 59, 18, -47,
			-29, 4, -82, -37, -25, -42, 7, -8,
		},

		// rook
		{
			-19, -13, 1, 17, 16, 7, -37, -26,
			-44, -16, -20, -9, -1, 11, -6, -71,
			-45, -25, -16, -17, 3, 0, -5, -33,
			-36, -26, -12, -1, 9, -7, 6, -23,
			-24, -11, 7, 26, 24, 35, -8, -20,
			-5, 19, 26, 36, 17, 45, 61, 16,
			27, 32, 58, 62, 80, 67, 26, 44,
			32, 42, 32, 51, 63, 9, 31, 43,
		},

		// queen
		{
			-1, -18, -9, 10, -15, -25, -31, -50,
			-35, -8, 11, 2, 8, 15, -3, 1,
			-14, 2, -11, -2, -5, 2, 14, 5,
			-9, -26, -9, -10, -2, -4, 3, -3,
			-27, -27, -16, -16, -1, 17, -2, 1,
			-13, -17, 7, 8, 29, 56, 47, 57,
			-24, -39, -5, 1, -16, 57, 28, 54,
			-28, 0, 29, 12, 59, 44, 43, 45,
		},

		// king
		{
			-15, 36, 12, -54, 8, -28, 24, 14,
			1, 7, -8, -64, -43, -16, 9, 8,
			-14, -14, -22, -46, -44, -30, -15, -27,
			-49, -1, -27, -39, -46, -44, -33, -51,
			-17, -20, -12, -27, -30, -25, -14, -36,
			-9, 24, 2, -16, -20, 6, 22, -22,
			29, -1, -20, -7, -8, -4, -38, -29,
			-65, 23, 16, -15, -56, -34, 2, 13,
		},
	},

	// Endgame positional piece scores //

	//pawn
	{
		{
			0, 0, 0, 0, 0, 0, 0, 0,
			13, 8, 8, 10, 13, 0, 2, -7,
			4, 7, -6, 1, 0, -5, -1, -8,
			13, 9, -3, -7, -7, -8, 3, -1,
			32, 24, 13, 5, -2, 4, 17, 17,
			94, 100, 85, 67, 56, 53, 82, 84,
			178, 173, 158, 134, 147, 132, 165, 187,
			0, 0, 0, 0, 0, 0, 0, 0,
		},

		// knight
		{
			-29, -51, -23, -15, -22, -18, -50, -64,
			-42, -20, -10, -5, -2, -20, -23, -44,
			-23, -3, -1, 15, 10, -3, -20, -22,
			-18, -6, 16, 25, 16, 17, 4, -18,
			-17, 3, 22, 22, 22, 11, 8, -18,
			-24, -20, 10, 9, -1, -9, -19, -41,
			-25, -8, -25, -2, -9, -25, -24, -52,
			-58, -38, -13, -28, -31, -27, -63, -99,
		},

		// bishop
		{
			-23, -9, -23, -5, -9, -16, -5, -17,
			-14, -18, -7, -1, 4, -9, -15, -27,
			-12, -3, 8, 10, 13, 3, -7, -15,
			-6, 3, 13, 19, 7, 10, -3, -9,
			-3, 9, 12, 9, 14, 10, 3, 2,
			2, -8, 0, -1, -2, 6, 0, 4,
			-8, -4, 7, -12, -3, -13, -4, -14,
			-14, -21, -11, -8, -7, -9, -17, -24,
		},

		// rook
		{
			-9, 2, 3, -1, -5, -13, 4, -20,
			-6, -6, 0, 2, -9, -9, -11, -3,
			-4, 0, -5, -1, -7, -12, -8, -16,
			3, 5, 8, 4, -5, -6, -8, -11,
			4, 3, 13, 1, 2, 1, -1, 2,
			7, 7, 7, 5, 4, -3, -5, -3,
			11, 13, 13, 11, -3, 3, 8, 3,
			13, 10, 18, 15, 12, 12, 8, 5,
		},

		// queen
		{
			-33, -28, -22, -43, -5, -32, -20, -41,
			-22, -23, -30, -16, -16, -23, -36, -32,
			-16, -27, 15, 6, 9, 17, 10, 5,
			-18, 28, 19, 47, 31, 34, 39, 23,
			3, 22, 24, 45, 57, 40, 57, 36,
			-20, 6, 9, 49, 47, 35, 19, 9,
			-17, 20, 32, 41, 58, 25, 30, 0,
			-9, 22, 22, 27, 27, 19, 10, 20,
		},

		// king
		{
			-53, -34, -21, -11, -28, -14, -24, -43,
			-27, -11, 4, 13, 14, 4, -5, -17,
			-19, -3, 11, 21, 23, 16, 7, -9,
			-18, -4, 21, 24, 27, 23, 9, -11,
			-8, 22, 24, 27, 26, 33, 26, 3,
			10, 17, 23, 15, 20, 45, 44, 13,
			-12, 17, 14, 17, 17, 38, 23, 11,
			-74, -35, -18, -18, -11, 15, 4, -17,
		},
	},
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

// getGamePhaseScore  should return the game score based on the game phase
func (b *BoardStruct) getGamePhaseScore() int {
	/*
	   The game phase score of the game is derived from the pieces
	   (not counting pawns and kings) that are still on the board
	   The full material starting position game phase score is:

	   4 * knight material score in the opening +
	   4 * bishop material score in the opening +
	   4 * rook material score in the opening +
	   2 * queen material score in the opening
	*/

	wPieces, bPieces := 0, 0

	for p := WN; p <= WQ; p++ {
		wPieces += b.Bitboards[p].Count() * MaterialScore[Opening][p]
	}

	for p := BN; p <= BQ; p++ {
		bPieces += b.Bitboards[p].Count() * -MaterialScore[Opening][p]
	}

	return wPieces + bPieces
}

// EvaluatePosition should evaluate a certain position
func (b *BoardStruct) EvaluatePosition() int {
	bb := Bitboard(0)

	gamePhaseScore := b.getGamePhaseScore()

	gamePhase := -1

	if gamePhaseScore > OpeningPhaseScore {
		gamePhase = Opening
	} else if gamePhaseScore > EndgamePhaseScore {
		gamePhase = Middlegame
	} else {
		gamePhase = Endgame
	}

	score, scoreOpening, scoreEndgame := 0, 0, 0
	pc, sq := 0, 0
	// doublePawns := 0

	for bbPc := WP; bbPc <= BK; bbPc++ {
		bb = b.Bitboards[bbPc]

		for bb != 0 {
			pc = bbPc
			sq = bb.FirstOne()

			scoreOpening += MaterialScore[Opening][pc]
			scoreEndgame += MaterialScore[Endgame][pc]

			switch pc {
			// evaluare white pieces
			case WP:
				scoreOpening += PositionalScore[Opening][Pawn][sq]
				scoreEndgame += PositionalScore[Endgame][Pawn][sq]

				// doublePawns = (b.Bitboards[WP] & FileMasks[sq]).Count()
				// if doublePawns > 1 {
				// 	score += doublePawns * doublePawnPenalty
				// }
				//
				// if (b.Bitboards[WP] & IsolatedMasks[sq]) == 0 {
				// 	score += isolatedPawnPenalty
				// }
				//
				// if (WhitePassedMasks[sq] & b.Bitboards[BP]) == 0 {
				// 	score += passedPawnBonus[getRank[sq]]
				// }
				//
			case WN:
				scoreOpening += PositionalScore[Opening][Knight][sq]
				scoreEndgame += PositionalScore[Endgame][Knight][sq]

			case WB:
				scoreOpening += PositionalScore[Opening][Bishop][sq]
				scoreEndgame += PositionalScore[Endgame][Bishop][sq]

				// score += GetBishopAttacks(sq, b.Occupancies[BOTH]).Count()
			case WR:
				scoreOpening += PositionalScore[Opening][Rook][sq]
				scoreEndgame += PositionalScore[Endgame][Rook][sq]

				// if (b.Bitboards[WP] & FileMasks[sq]) == 0 {
				// 	score += semiOpenFileScore
				// }
				//
				// if ((b.Bitboards[WP] | b.Bitboards[BP]) & FileMasks[sq]) == 0 {
				// 	score += openFileScore
				// }
			case WQ:
				scoreOpening += PositionalScore[Opening][Queen][sq]
				scoreEndgame += PositionalScore[Endgame][Queen][sq]

				// score += GetQueenAttacks(sq, b.Occupancies[BOTH]).Count()
			case WK:
				scoreOpening += PositionalScore[Opening][King][sq]
				scoreEndgame += PositionalScore[Endgame][King][sq]

				// if (b.Bitboards[WP] & FileMasks[sq]) == 0 {
				// 	score -= semiOpenFileScore
				// }
				//
				// if ((b.Bitboards[WP] | b.Bitboards[BP]) & FileMasks[sq]) == 0 {
				// 	score -= openFileScore
				// }
				//
				// score += (KingAttacks[sq] & b.Occupancies[WHITE]).Count() * kingShieldBonus
			// evaluate black pieces
			case BP:
				scoreOpening -= PositionalScore[Opening][Pawn][mirrorScore[sq]]
				scoreEndgame -= PositionalScore[Endgame][Pawn][mirrorScore[sq]]

				// doublePawns = (b.Bitboards[BP] & FileMasks[sq]).Count()
				// if doublePawns > 1 {
				// 	score -= doublePawns * doublePawnPenalty
				// }
				//
				// if (b.Bitboards[BP] & IsolatedMasks[sq]) == 0 {
				// 	score -= isolatedPawnPenalty
				// }
				//
				// if (BlackPassedMasks[sq] & b.Bitboards[WP]) == 0 {
				// 	score -= passedPawnBonus[getRank[mirrorScore[sq]]]
				// }
			case BN:
				scoreOpening -= PositionalScore[Opening][Knight][mirrorScore[sq]]
				scoreEndgame -= PositionalScore[Endgame][Knight][mirrorScore[sq]]

			case BB:
				scoreOpening -= PositionalScore[Opening][Bishop][mirrorScore[sq]]
				scoreEndgame -= PositionalScore[Endgame][Bishop][mirrorScore[sq]]

				// score -= GetBishopAttacks(sq, b.Occupancies[BOTH]).Count()
			case BR:
				scoreOpening -= PositionalScore[Opening][Rook][mirrorScore[sq]]
				scoreEndgame -= PositionalScore[Endgame][Rook][mirrorScore[sq]]

				// if (b.Bitboards[BP] & FileMasks[sq]) == 0 {
				// 	score -= semiOpenFileScore
				// }
				//
				// if ((b.Bitboards[WP] | b.Bitboards[BP]) & FileMasks[sq]) == 0 {
				// 	score -= openFileScore
				// }
			case BQ:
				scoreOpening -= PositionalScore[Opening][Queen][mirrorScore[sq]]
				scoreEndgame -= PositionalScore[Endgame][Queen][mirrorScore[sq]]

				// score += GetQueenAttacks(sq, b.Occupancies[BOTH]).Count()
			case BK:
				scoreOpening -= PositionalScore[Opening][King][mirrorScore[sq]]
				scoreEndgame -= PositionalScore[Endgame][King][mirrorScore[sq]]

				// if (b.Bitboards[BP] & FileMasks[sq]) == 0 {
				// 	score += semiOpenFileScore
				// }
				//
				// if ((b.Bitboards[WP] | b.Bitboards[BP]) & FileMasks[sq]) == 0 {
				// 	score += openFileScore
				// }
				//
				// score -= (KingAttacks[sq] & b.Occupancies[BLACK]).Count() * kingShieldBonus
			}
		}
	}
	/*
	   Now in order to calculate interpolated score
	   for a given game phase we use this formula
	   (same for material and positional scores):

	   (
	     score_opening * game_phase_score +
	     score_endgame * (opening_phase_score - game_phase_score)
	   ) / opening_phase_score

	   E.g. the score for pawn on d4 at phase say 5000 would be
	   interpolated_score = (12 * 5000 + (-7) * (6192 - 5000)) / 6192 = 8,342377261
	*/

	// interpolate score in the middlegame
	if gamePhase == Middlegame {
		score = (scoreOpening*gamePhaseScore + scoreEndgame*(OpeningPhaseScore-gamePhaseScore)) / OpeningPhaseScore
	} else if gamePhase == Opening {
		score = scoreOpening
	} else {
		score = scoreEndgame
	}

	if b.SideToMove == WHITE {
		return score
	}
	return -score
}

// PRIVATE Methods

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
