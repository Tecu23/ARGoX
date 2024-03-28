package main

import "fmt"

// BoardStruct represent a board representation
type BoardStruct struct {
	Bitboards [12]Bitboard // define piece bitboards

	Occupancies [3]Bitboard // define occupancy bitboards (black, white and all occupancies)

	SideToMove Color
	EnPassant  int
	Castlings
}

// Clear should clear the board, flags, bitboards etc
func (b *BoardStruct) Clear() {
	b.SideToMove = WHITE
	// b.Rule50 = 0
	b.EnPassant = -1
	b.Castlings = 0

	for i := 0; i < 12; i++ {
		b.Bitboards[i] = 0
	}

	for i := 0; i < 3; i++ {
		b.Occupancies[i] = 0
	}
}

// SetSq should set a square sq to a particular piece pc
func (b *BoardStruct) SetSq(piece, sq int) {
	pieceColor := PcColor(piece)

	if piece == Empty {
		return
	}

	if b.Occupancies[BOTH].Test(sq) {
		// need to remove the piece
		for p := WP; p <= BK; p++ {
			if b.Bitboards[p].Test(sq) {
				b.Bitboards[p].Clear(sq)
			}
		}

		b.Occupancies[BOTH].Clear(sq)
		b.Occupancies[WHITE].Clear(sq)
		b.Occupancies[BLACK].Clear(sq)
	}

	b.Bitboards[piece].Set(sq)

	if pieceColor == WHITE {
		b.Occupancies[WHITE].Set(sq)
	} else {
		b.Occupancies[BLACK].Set(sq)
	}

	b.Occupancies[BOTH] |= b.Occupancies[WHITE]
	b.Occupancies[BOTH] |= b.Occupancies[BLACK]
}

// isSquareAttacked should return whether the given square is attacked by the curren given side
func (b *BoardStruct) isSquareAttacked(sq int, side Color) bool {
	if side == WHITE {
		if PawnAttacks[BLACK][sq]&b.Bitboards[WP] != 0 {
			return true
		}

		if KnightAttacks[sq]&b.Bitboards[WN] != 0 {
			return true
		}

		if KingAttacks[sq]&b.Bitboards[WK] != 0 {
			return true
		}

		bishopAttacks := getBishopAttacks(sq, b.Occupancies[BOTH])

		if bishopAttacks&b.Bitboards[WB] != 0 {
			return true
		}
		rookAttacks := getRookAttacks(sq, b.Occupancies[BOTH])

		if rookAttacks&b.Bitboards[WR] != 0 {
			return true
		}
		queenAttacks := getQueenAttacks(sq, b.Occupancies[BOTH])

		if queenAttacks&b.Bitboards[WQ] != 0 {
			return true
		}
	} else {

		if PawnAttacks[WHITE][sq]&b.Bitboards[BP] != 0 {
			return true
		}

		if KnightAttacks[sq]&b.Bitboards[BN] != 0 {
			return true
		}

		if KingAttacks[sq]&b.Bitboards[BK] != 0 {
			return true
		}

		bishopAttacks := getBishopAttacks(sq, b.Occupancies[BOTH])

		if bishopAttacks&b.Bitboards[BB] != 0 {
			return true
		}
		rookAttacks := getRookAttacks(sq, b.Occupancies[BOTH])

		if rookAttacks&b.Bitboards[BR] != 0 {
			return true
		}
		queenAttacks := getQueenAttacks(sq, b.Occupancies[BOTH])

		if queenAttacks&b.Bitboards[BQ] != 0 {
			return true
		}
	}
	return false
}

// PrintAttackedSquares should print all squares that are currently being attacked
func (b *BoardStruct) PrintAttackedSquares(side Color) {
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			sq := rank*8 + file

			if file == 0 {
				fmt.Printf(" %d ", rank+1)
			}

			if b.isSquareAttacked(sq, side) {
				fmt.Printf(" %d", 1)
			} else {
				fmt.Printf(" %d", 0)
			}
		}

		fmt.Printf("\n")
	}
	fmt.Printf("\n     a b c d e f g h\n\n")
}

// PrintBoard should print the current position of the board
func (b BoardStruct) PrintBoard() {
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			if file == 0 {
				fmt.Printf("%d  ", rank+1)
			}
			piece := -1

			// loop over all piece bitboards
			for bb := WP; bb <= BK; bb++ {
				if b.Bitboards[bb].Test(rank*8 + file) {
					piece = bb
				}
			}

			if piece == -1 {
				fmt.Printf(" %c", '.')
			} else {
				fmt.Printf(" %c", AciiPieces[piece])
			}

		}
		fmt.Println()
	}

	fmt.Printf("\n    a b c d e f g h\n\n")

	// print side to move
	if b.SideToMove == WHITE {
		fmt.Printf(" Side:     %s\n", "white")
	} else {
		fmt.Printf(" Side:     %s\n", "black")
	}

	// print enpassant square
	fmt.Printf(" Enpassant:   %s\n", Sq2Fen[b.EnPassant])

	// print castling rights
	fmt.Printf(" Castling:  %s\n\n", b.Castlings.String())
}
