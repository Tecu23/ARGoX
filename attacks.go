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

// InitKnightAttacks should initialize the KinghtAttacks table
func InitKnightAttacks() {
	for sq := A1; sq <= H8; sq++ {
		KnightAttacks[sq] = generateKnightAttacks(sq)
	}
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

// InitKingAttacks should initialize the KingAttacks table
func InitKingAttacks() {
	for sq := A1; sq <= H8; sq++ {
		KingAttacks[sq] = generateKingAttacks(sq)
	}
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

// generateBishopAttacks should generate all attacks for a bishop in a certain square
func generateBishopAttacks(square int) Bitboard {
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

// generateRookAttacks should generate all attacks for a rook in a certain square
func generateRookAttacks(square int) Bitboard {
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
