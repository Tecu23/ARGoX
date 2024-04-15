package main

// PawnAttacks is a lookup table for app pawn attack for fast move generation
var (
	PawnAttacks   [2][64]Bitboard
	KnightAttacks [64]Bitboard
	KingAttacks   [64]Bitboard
)

// InitPawnAttacks should initialize the PawnAttacks table
func InitPawnAttacks() {
	for sq := A1; sq <= H8; sq++ {
		PawnAttacks[WHITE][sq] = generatePawnAttacks(sq, WHITE)
		PawnAttacks[BLACK][sq] = generatePawnAttacks(sq, BLACK)
	}
}

// InitKnightAttacks should initialize the KinghtAttacks table
func InitKnightAttacks() {
	for sq := A1; sq <= H8; sq++ {
		KnightAttacks[sq] = generateKnightAttacks(sq)
	}
}

// InitKingAttacks should initialize the KingAttacks table
func InitKingAttacks() {
	for sq := A1; sq <= H8; sq++ {
		KingAttacks[sq] = generateKingAttacks(sq)
	}
}

// InitSliderPiecesAttacks should generate all attacks for slider pieces
func InitSliderPiecesAttacks(piece int) {
	// loop over 64 board squares
	for sq := A1; sq <= H8; sq++ {
		// init bishop & rook masks
		BishopMasks[sq] = GenerateBishopAttacks(sq)
		RookMasks[sq] = GenerateRookAttacks(sq)

		// init current mask
		var attackMask Bitboard

		if piece == Bishop {
			attackMask = GenerateBishopAttacks(sq)
		} else {
			attackMask = GenerateRookAttacks(sq)
		}

		// count attack mask bits
		bitCount := attackMask.Count()

		// occupancy variations count
		occupancyVariations := 1 << bitCount

		// loop over occupancy variations
		for count := 0; count < occupancyVariations; count++ {
			if piece == Bishop {
				occupancy := SetOccupancy(count, bitCount, attackMask)

				magicIndex := occupancy * BishopMagicNumbers[sq] >> (64 - BishopRelevantBits[sq])
				BishopAttacks[sq][magicIndex] = GenerateBishopAttacksOnTheFly(sq, occupancy)
			} else {

				occupancy := SetOccupancy(count, bitCount, attackMask)

				magicIndex := occupancy * RookMagicNumbers[sq] >> (64 - RookRelevantBits[sq])
				RookAttacks[sq][magicIndex] = GenerateRookAttacksOnTheFly(sq, occupancy)
			}
		}
	}
}

// GenerateBishopAttacks should generate all attacks for a bishop in a certain square
func GenerateBishopAttacks(square int) Bitboard {
	attacks := Bitboard(0)

	// init rank & files
	r, f := 0, 0

	// init target rank & files
	tr := square / 8
	tf := square % 8

	// mask relevant bishop occupancy bits
	for r, f = tr+1, tf+1; r <= 6 && f <= 6; r, f = r+1, f+1 {
		attacks |= 1 << (r*8 + f)
	}

	for r, f = tr-1, tf+1; r >= 1 && f <= 6; r, f = r-1, f+1 {
		attacks |= 1 << (r*8 + f)
	}

	for r, f = tr+1, tf-1; r <= 6 && f >= 1; r, f = r+1, f-1 {
		attacks |= 1 << (r*8 + f)
	}

	for r, f = tr-1, tf-1; r >= 1 && f >= 1; r, f = r-1, f-1 {
		attacks |= 1 << (r*8 + f)
	}

	return attacks
}

// GenerateBishopAttacksOnTheFly should generate all attacks for a given blocker board
func GenerateBishopAttacksOnTheFly(square int, block Bitboard) Bitboard {
	attacks := Bitboard(0)

	// init rank & files
	r, f := 0, 0

	// init target rank & files
	tr := square / 8
	tf := square % 8

	// mask relevant bishop occupancy bits
	for r, f = tr+1, tf+1; r <= 7 && f <= 7; r, f = r+1, f+1 {
		attacks |= 1 << (r*8 + f)
		if (1<<(r*8+f))&block != 0 {
			break
		}
	}

	for r, f = tr-1, tf+1; r >= 0 && f <= 7; r, f = r-1, f+1 {
		attacks |= 1 << (r*8 + f)
		if (1<<(r*8+f))&block != 0 {
			break
		}
	}

	for r, f = tr+1, tf-1; r <= 7 && f >= 0; r, f = r+1, f-1 {
		attacks |= 1 << (r*8 + f)
		if (1<<(r*8+f))&block != 0 {
			break
		}
	}

	for r, f = tr-1, tf-1; r >= 0 && f >= 0; r, f = r-1, f-1 {
		attacks |= 1 << (r*8 + f)
		if (1<<(r*8+f))&block != 0 {
			break
		}
	}

	return attacks
}

// GenerateRookAttacks should generate all attacks for a rook in a certain square
func GenerateRookAttacks(square int) Bitboard {
	attacks := Bitboard(0)

	// init rank & files
	r, f := 0, 0

	// init target rank & files
	tr := square / 8
	tf := square % 8

	// mask relevant bishop occupancy bits
	for r = tr + 1; r <= 6; r++ {
		attacks |= 1 << (r*8 + tf)
	}

	for f = tf + 1; f <= 6; f++ {
		attacks |= 1 << (tr*8 + f)
	}

	for r = tr - 1; r >= 1; r-- {
		attacks |= 1 << (r*8 + tf)
	}

	for f = tf - 1; f >= 1; f-- {
		attacks |= 1 << (tr*8 + f)
	}

	return attacks
}

// GenerateRookAttacksOnTheFly should generate all attacks for a given blocker board
func GenerateRookAttacksOnTheFly(square int, block Bitboard) Bitboard {
	attacks := Bitboard(0)

	// init rank & files
	r, f := 0, 0

	// init target rank & files
	tr := square / 8
	tf := square % 8

	// mask relevant bishop occupancy bits
	for r = tr + 1; r <= 7; r++ {
		attacks |= 1 << (r*8 + tf)
		if (1<<(r*8+tf))&block != 0 {
			break
		}
	}

	for f = tf + 1; f <= 7; f++ {
		attacks |= 1 << (tr*8 + f)
		if (1<<(tr*8+f))&block != 0 {
			break
		}
	}

	for r = tr - 1; r >= 0; r-- {
		attacks |= 1 << (r*8 + tf)
		if (1<<(r*8+tf))&block != 0 {
			break
		}
	}

	for f = tf - 1; f >= 0; f-- {
		attacks |= 1 << (tr*8 + f)
		if (1<<(tr*8+f))&block != 0 {
			break
		}
	}

	return attacks
}

// GetBishopAttacks should return all the rook attacks for a certain square with a certain occup
func GetBishopAttacks(sq int, occupancy Bitboard) Bitboard {
	// calculate magic index
	occupancy &= BishopMasks[sq]
	occupancy *= BishopMagicNumbers[sq]
	occupancy >>= 64 - BishopRelevantBits[sq]

	return BishopAttacks[sq][occupancy]
}

// GetRookAttacks should return all the rook attacks for a certain square with a certain occup
func GetRookAttacks(sq int, occupancy Bitboard) Bitboard {
	// calculate magic index
	occupancy &= RookMasks[sq]
	occupancy *= RookMagicNumbers[sq]
	occupancy >>= 64 - RookRelevantBits[sq]

	return RookAttacks[sq][occupancy]
}

// GetQueenAttacks should return all the rook attacks for a certain square with a certain occup
func GetQueenAttacks(sq int, occupancy Bitboard) Bitboard {
	queenAttacks := Bitboard(0)

	bishopOccupancies := occupancy
	rookOccupancies := occupancy

	bishopOccupancies &= BishopMasks[sq]
	bishopOccupancies *= BishopMagicNumbers[sq]
	bishopOccupancies >>= 64 - BishopRelevantBits[sq]

	rookOccupancies &= RookMasks[sq]
	rookOccupancies *= RookMagicNumbers[sq]
	rookOccupancies >>= 64 - RookRelevantBits[sq]

	queenAttacks = BishopAttacks[sq][bishopOccupancies] | RookAttacks[sq][rookOccupancies]

	return queenAttacks
}

// PRIVATE Methods

// generatePawnAttacks should generate all pawn attacks provided the square and color
func generatePawnAttacks(square int, side Color) Bitboard {
	attacks := Bitboard(0)

	b := Bitboard(0)
	b.Set(square)

	if side == WHITE {
		attacks |= (b & ^FileA) << NW
		attacks |= (b & ^FileH) << NE
	} else {
		attacks |= (b & ^FileA) >> NE
		attacks |= (b & ^FileH) >> NW
	}

	return attacks
}

// generateKnightAttacks should generate all attacks for a knight in a certain position
func generateKnightAttacks(square int) Bitboard {
	attacks := Bitboard(0)

	b := Bitboard(0)
	b.Set(square)

	attacks |= (b & ^FileA) << (NW + N)
	attacks |= (b & ^FileA & ^FileB) << (NW + W)

	attacks |= (b & ^FileA & ^FileB) >> (NE + E)
	attacks |= (b & ^FileA) >> (NE + N)

	attacks |= (b & ^FileH) >> (NW + N)
	attacks |= (b & ^FileH & ^FileG) >> (NW + W)

	attacks |= (b & ^FileH & ^FileG) << (NE + E)
	attacks |= (b & ^FileH) << (NE + N)

	return attacks
}

// generateKingAttacks should generate all attacks for a king in a certain square
func generateKingAttacks(square int) Bitboard {
	attacks := Bitboard(0)

	b := Bitboard(0)
	b.Set(square)

	attacks |= (b & ^FileA) << NW
	attacks |= b << N
	attacks |= (b & ^FileH) << NE

	attacks |= (b & ^FileH) << E

	attacks |= (b & ^FileA) >> E

	attacks |= (b & ^FileH) >> NW
	attacks |= b >> N
	attacks |= (b & ^FileA) >> NE

	return attacks
}
