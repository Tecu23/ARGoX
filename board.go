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
	b.EnPassant = 0
	b.Castlings = 0

	for i := 0; i < 12; i++ {
		b.Bitboards[i] = 0
	}

	for i := 0; i < 3; i++ {
		b.Occupancies[i] = 0
	}

	// b.WbBB[WHITE], b.WbBB[BLACK] = 0, 0
	// for i := 0; i < NoPiecesT; i++ {
	// 	b.PieceBB[i] = 0
	// }
}

// SetSq should set a square sq to a particular piece pc
func (b *BoardStruct) SetSq(pc, sq int) {
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
