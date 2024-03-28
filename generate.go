package main

import "fmt"

// generate all moves
func (b *BoardStruct) generateMoves() {
	sourceSq, targetSq := 0, 0

	bitboard, attacks := Bitboard(0), Bitboard(0)
	fmt.Printf("\n %d \n", attacks)

	for piece := WP; piece <= BK; piece++ {

		bitboard = b.Bitboards[piece]

		// generate pawns and castling moves depending of size
		if b.SideToMove == WHITE {
			if piece == WP {
				for bitboard != 0 {
					sourceSq = bitboard.FirstOne()
					targetSq = sourceSq + N

					// quiet pawn moves
					if !(targetSq < A1) && !b.Occupancies[BOTH].Test(targetSq) {
						// pawn promotion
						if sourceSq >= A7 && sourceSq <= H7 {
							fmt.Printf(
								"Pawn Promotion: %s%sq\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Promotion: %s%sr\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Promotion: %s%sb\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Promotion: %s%sn\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							// one square ahead move
							fmt.Printf("pawn push: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])

							// two square ahead move
							if (sourceSq >= A2 && sourceSq <= H2) && !b.Occupancies[BOTH].Test(targetSq+N) {
								fmt.Printf("pawn double push: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq+N])
							}
						}
					}

					// init pawn attacks bb
					attacks = PawnAttacks[WHITE][sourceSq] & b.Occupancies[BLACK]

					// generate pawn captures
					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if sourceSq >= A7 && sourceSq <= H7 {

							fmt.Printf(
								"Pawn Capture Promotion: %s%sq\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Capture Promotion: %s%sr\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Capture Promotion: %s%sb\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Capture Promotion: %s%sn\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("pawn capture: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		} else {
			if piece == BP {
				for bitboard != 0 {
					sourceSq = bitboard.FirstOne()
					targetSq = sourceSq + S

					// quiet pawn moves
					if !(targetSq > H8) && !b.Occupancies[BOTH].Test(targetSq) {
						// pawn promotion
						if sourceSq >= A2 && sourceSq <= H2 {
							fmt.Printf(
								"Pawn Promotion: %s%sq\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Promotion: %s%sr\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Promotion: %s%sb\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
							fmt.Printf(
								"Pawn Promotion: %s%sn\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							// one square ahead move
							fmt.Printf("pawn push: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])

							// two square ahead move
							if (sourceSq >= A7 && sourceSq <= H7) && !b.Occupancies[BOTH].Test(targetSq+S) {
								fmt.Printf("pawn double push: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq+S])
							}
						}
					}
				}
			}
		}

		// generate knight moves

		// generate bihop moves

		// generate rook moves

		// generate queen moves

		// generate king moves
	}
}
