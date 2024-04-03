package main

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

	return score
}
