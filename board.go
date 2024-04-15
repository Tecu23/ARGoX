package main

import "fmt"

// PieceKeys keeps random piece keys [piece][square]
var PieceKeys [12][64]uint64

// EnpassantKeys keeps random enpassant keys
var EnpassantKeys [64]uint64

// CastlingKeys keeps radom castling keys
var CastlingKeys [16]uint64

// SideKey keeps random side key
var SideKey uint64

// InitRandomHashKeys should initialize random hash keys
func InitRandomHashKeys() {
	randomState := uint32(1804289383)

	for p := WP; p <= BK; p++ {
		for sq := A1; sq <= H8; sq++ {
			PieceKeys[p][sq] = generateRandomUint64Number(&randomState)
		}
	}
	for sq := A1; sq <= H8; sq++ {
		EnpassantKeys[sq] = generateRandomUint64Number(&randomState)
	}
	for i := 0; i < 16; i++ {
		CastlingKeys[i] = generateRandomUint64Number(&randomState)
	}
	SideKey = generateRandomUint64Number(&randomState)
}

func (b *BoardStruct) generateHashKey() uint64 {
	finalKey := uint64(0)

	bb := Bitboard(0)

	for p := WP; p <= BK; p++ {
		bb = b.Bitboards[p]

		for bb != 0 {
			sq := bb.FirstOne()
			finalKey ^= PieceKeys[p][sq]
		}
	}

	if b.EnPassant != -1 {
		finalKey ^= EnpassantKeys[b.EnPassant]
	}

	finalKey ^= CastlingKeys[b.Castlings]

	if b.SideToMove == BLACK {
		finalKey ^= SideKey
	}

	return finalKey
}

// BoardStruct represent a board representation
type BoardStruct struct {
	Bitboards [12]Bitboard // define piece bitboards

	Occupancies [3]Bitboard // define occupancy bitboards (black, white and all occupancies)

	SideToMove Color
	EnPassant  int
	Castlings

	Key uint64

	// positions repetition table & index
	RepetitionTable [1000]uint64 // 1000 is a number of plies (500 moves) in the entire game
	RepetitionIdx   int
}

// Clear should clear the board, flags, bitboards etc
func (b *BoardStruct) Clear() {
	b.SideToMove = WHITE
	// b.Rule50 = 0
	b.RepetitionTable = [1000]uint64{}
	b.RepetitionIdx = 0
	b.Key = 0
	b.EnPassant = -1
	b.Castlings = 0

	for i := 0; i < 12; i++ {
		b.Bitboards[i] = 0
	}

	for i := 0; i < 3; i++ {
		b.Occupancies[i] = 0
	}
}

// CopyBoard should take a copy of the current board position
func (b BoardStruct) CopyBoard() BoardStruct {
	boardCopy := b

	return boardCopy
}

// TakeBack should restore the board state to a copy
func (b *BoardStruct) TakeBack(copy BoardStruct) {
	*b = copy
}

// SetSq should set a square sq to a particular piece pc
func (b *BoardStruct) SetSq(piece, sq int) {
	pieceColor := PcColor(piece)

	if b.Occupancies[BOTH].Test(sq) {
		// need to remove the piece
		for p := WP; p <= BK; p++ {
			if b.Bitboards[p].Test(sq) {
				b.Bitboards[p].Clear(sq)
				b.Key ^= PieceKeys[p][sq]
			}
		}

		b.Occupancies[BOTH].Clear(sq)
		b.Occupancies[WHITE].Clear(sq)
		b.Occupancies[BLACK].Clear(sq)
	}

	if piece == Empty {
		return
	}
	b.Key ^= PieceKeys[piece][sq]
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

// ShowAttackedSquares should print all squares that are currently being attacked
func (b *BoardStruct) ShowAttackedSquares(side Color) {
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

// AllMoves and OnlyCaptures flags
const (
	AllMoves     = 0
	OnlyCaptures = 1
)

// MakeMove should make a move on the board
// TODO: Refactor this method for performance
func (b *BoardStruct) MakeMove(m Move, moveFlag int) bool {
	// quiet moves
	if moveFlag == AllMoves {
		// preserve board state
		copyB := b.CopyBoard()

		// parse the move
		src := m.GetSource()
		tgt := m.GetTarget()
		pc := m.GetPiece()
		color := PcColor(pc)
		prom := m.GetPromoted()
		// capt := m.GetCapture()
		dblPwn := m.GetDoublePush()
		ep := m.GetEnpassant()
		cast := m.GetCastling()

		if b.EnPassant != -1 {
			b.Key ^= EnpassantKeys[b.EnPassant]
		}
		b.EnPassant = -1

		if ep != 0 {
			b.SetSq(Empty, src)
			if color == WHITE {
				b.SetSq(Empty, tgt+S)
			} else {
				b.SetSq(Empty, tgt+N)
			}

			b.SetSq(pc, tgt)

			b.SideToMove = b.SideToMove.Opp()
			b.Key ^= SideKey
			// make sure the king was not exposed into a check
			var kingPos int

			if b.SideToMove == WHITE {
				if b.Bitboards[BK] == 0 {
					b.TakeBack(copyB)
					return false
				}
				kingPos = b.Bitboards[BK].FirstOne()
			} else {
				if b.Bitboards[WK] == 0 {
					b.TakeBack(copyB)
					return false
				}
				kingPos = b.Bitboards[WK].FirstOne()
			}
			if b.isSquareAttacked(kingPos, b.SideToMove) {
				// take back
				b.TakeBack(copyB)
				return false
			}
			if b.SideToMove == WHITE {
				b.Bitboards[BK].Set(kingPos)
			} else {
				b.Bitboards[WK].Set(kingPos)
			}

			return true
		}

		if cast != 0 {
			switch tgt {
			// WHITE Short Castle
			case G1:
				b.SetSq(Empty, H1)
				b.SetSq(WR, F1)
			// WHITE Long Castle
			case C1:
				b.SetSq(Empty, A1)
				b.SetSq(WR, D1)
			// BLACK Short Castle
			case G8:
				b.SetSq(Empty, H8)
				b.SetSq(BR, F8)
			// BLACK Long Castle
			case C8:
				b.SetSq(Empty, A8)
				b.SetSq(BR, D8)
			}
		}

		if dblPwn != 0 {
			if color == WHITE {
				b.EnPassant = src + N
				b.Key ^= EnpassantKeys[src+N]
			} else {
				b.EnPassant = src + S
				b.Key ^= EnpassantKeys[src+S]
			}
		}

		b.SetSq(Empty, src)

		if prom != 0 {
			b.SetSq(prom, tgt)
		} else {
			b.SetSq(pc, tgt)
		}

		b.Key ^= CastlingKeys[b.Castlings]

		// update castling rights
		b.Castlings &= Castlings(CastlingRights[src])
		b.Castlings &= Castlings(CastlingRights[tgt])

		b.Key ^= CastlingKeys[b.Castlings]

		// change side
		b.SideToMove = b.SideToMove.Opp()
		b.Key ^= SideKey

		// make sure the king was not exposed into a check
		var kingPos int

		if b.SideToMove == WHITE {
			if b.Bitboards[BK] == 0 {
				b.TakeBack(copyB)
				return false
			}
			kingPos = b.Bitboards[BK].FirstOne()
		} else {
			if b.Bitboards[WK] == 0 {
				b.TakeBack(copyB)
				return false
			}
			kingPos = b.Bitboards[WK].FirstOne()
		}
		if b.isSquareAttacked(kingPos, b.SideToMove) {
			// take back
			b.TakeBack(copyB)
			return false
		}
		if b.SideToMove == WHITE {
			b.Bitboards[BK].Set(kingPos)
		} else {
			b.Bitboards[WK].Set(kingPos)
		}

	} else { // capture moves
		if m.GetCapture() != 0 {
			return b.MakeMove(m, AllMoves)
		}

		return false // 0 means don't make it
	}

	return true
}

// ParseMove should parse user/GUI move string input (e.g. e7e8q)
func (b *BoardStruct) ParseMove(moveString string) Move {
	var moves Movelist

	b.generateMoves(&moves)

	src := Fen2Sq[moveString[:2]]
	tgt := Fen2Sq[moveString[2:4]]

	for cnt := 0; cnt < len(moves); cnt++ {
		mv := moves[cnt]

		if mv.GetSource() == src && mv.GetTarget() == tgt {
			prom := mv.GetPromoted()

			if prom != 0 {
				if (prom == WQ || prom == BQ) && moveString[4] == 'q' {
					return mv
				}

				if (prom == WR || prom == BR) && moveString[4] == 'r' {
					return mv
				}

				if (prom == WB || prom == BB) && moveString[4] == 'b' {
					return mv
				}

				if (prom == WN || prom == BN) && moveString[4] == 'n' {
					return mv
				}

				// continue the loop on wrong promotions
				continue
			}

			return mv
		}
	}

	return NoMove
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

	fmt.Printf(" HashKey: 0x%X\n\n", b.Key)
}
