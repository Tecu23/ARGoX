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
							fmt.Printf("Pawn push: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])

							// two square ahead move
							if (sourceSq >= A2 && sourceSq <= H2) && !b.Occupancies[BOTH].Test(targetSq+N) {
								fmt.Printf("Pawn double push: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq+N])
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
							fmt.Printf("Pawn capture: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}

					// generate EnPassant captures
					if b.EnPassant != -1 {
						enpassantAttacks := PawnAttacks[WHITE][sourceSq] & (1 << b.EnPassant)

						// check enpassant capture
						if enpassantAttacks != 0 {
							// init enpassant capture target square
							targetEnpassant := enpassantAttacks.FirstOne()
							fmt.Printf(
								"Pawn enpassant capture: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetEnpassant],
							)

						}
					}

				}
			}
			// Castlings moves
			if piece == WK {
				// King side castling is available
				if uint(b.Castlings)&ShortW != 0 {
					// make sure square between king and king's rook are empty
					if !b.Occupancies[BOTH].Test(F1) && !b.Occupancies[BOTH].Test(G1) {
						// make sure king and the f1 square are not under attack
						if !b.isSquareAttacked(E1, BLACK) && !b.isSquareAttacked(F1, BLACK) {
							fmt.Printf("Castling move: e1g1\n")
						}
					}
				}

				// Queen side castling is available
				if uint(b.Castlings)&LongW != 0 {
					// make sure square between king and queens's rook are empty
					if !b.Occupancies[BOTH].Test(D1) && !b.Occupancies[BOTH].Test(C1) &&
						!b.Occupancies[BOTH].Test(B1) {
						// make sure king and the f1 square are not under attack
						if !b.isSquareAttacked(E1, BLACK) && !b.isSquareAttacked(D1, BLACK) {
							fmt.Printf("Castling move: e1c1\n")
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

					// init pawn attacks bb
					attacks = PawnAttacks[BLACK][sourceSq] & b.Occupancies[WHITE]

					// generate pawn captures
					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if sourceSq >= A2 && sourceSq <= H2 {

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
							fmt.Printf("Pawn capture: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}

					// generate EnPassant captures
					if b.EnPassant != -1 {
						enpassantAttacks := PawnAttacks[BLACK][sourceSq] & (1 << b.EnPassant)

						// check enpassant capture
						if enpassantAttacks != 0 {
							// init enpassant capture target square
							targetEnpassant := enpassantAttacks.FirstOne()
							fmt.Printf(
								"Pawn enpassant capture: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetEnpassant],
							)

						}
					}

				}
			}

			// Castlings moves
			if piece == BK {
				// King side castling is available
				if uint(b.Castlings)&ShortB != 0 {
					// make sure square between king and king's rook are empty
					if !b.Occupancies[BOTH].Test(F8) && !b.Occupancies[BOTH].Test(G8) {
						// make sure king and the f1 square are not under attack
						if !b.isSquareAttacked(E8, WHITE) && !b.isSquareAttacked(F8, WHITE) {
							fmt.Printf("Castling move: e8g8\n")
						}
					}
				}

				// Queen side castling is available
				if uint(b.Castlings)&LongB != 0 {
					// make sure square between king and queens's rook are empty
					if !b.Occupancies[BOTH].Test(D8) && !b.Occupancies[BOTH].Test(C8) &&
						!b.Occupancies[BOTH].Test(B8) {
						// make sure king and the f1 square are not under attack
						if !b.isSquareAttacked(E8, WHITE) && !b.isSquareAttacked(D8, WHITE) {
							fmt.Printf("Castling move: e8c8\n")
						}
					}
				}
			}
		}

		// generate knight moves
		if b.SideToMove == WHITE {
			if piece == WN {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = KnightAttacks[sourceSq] & (^b.Occupancies[WHITE])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[BLACK].Test(targetSq) {
							fmt.Printf(
								"Knight quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("Knight capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		} else {
			if piece == BN {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = KnightAttacks[sourceSq] & (^b.Occupancies[BLACK])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[WHITE].Test(targetSq) {
							fmt.Printf(
								"Knight quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("Knight capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		}

		// generate bihop moves
		if b.SideToMove == WHITE {
			if piece == WB {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = getBishopAttacks(
						sourceSq,
						b.Occupancies[BOTH],
					) & (^b.Occupancies[WHITE])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[BLACK].Test(targetSq) {
							fmt.Printf(
								"Bishop quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("Bishop capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		} else {
			if piece == BB {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = getBishopAttacks(sourceSq, b.Occupancies[BOTH]) & (^b.Occupancies[BLACK])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[WHITE].Test(targetSq) {
							fmt.Printf(
								"Bishop quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("Bishop capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		}

		// generate rook moves
		if b.SideToMove == WHITE {
			if piece == WR {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = getRookAttacks(
						sourceSq,
						b.Occupancies[BOTH],
					) & (^b.Occupancies[WHITE])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[BLACK].Test(targetSq) {
							fmt.Printf(
								"Rook quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("Rook capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		} else {
			if piece == BR {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = getRookAttacks(sourceSq, b.Occupancies[BOTH]) & (^b.Occupancies[BLACK])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[WHITE].Test(targetSq) {
							fmt.Printf(
								"Rook quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("Rook capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		}

		// generate queen moves
		if b.SideToMove == WHITE {
			if piece == WQ {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = getQueenAttacks(
						sourceSq,
						b.Occupancies[BOTH],
					) & (^b.Occupancies[WHITE])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[BLACK].Test(targetSq) {
							fmt.Printf(
								"Queen quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("Queen capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		} else {
			if piece == BQ {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = getQueenAttacks(sourceSq, b.Occupancies[BOTH]) & (^b.Occupancies[BLACK])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[WHITE].Test(targetSq) {
							fmt.Printf(
								"Queen quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("Queen capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		}

		// generate king moves
		if b.SideToMove == WHITE {
			if piece == WK {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = KingAttacks[sourceSq] & (^b.Occupancies[WHITE])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[BLACK].Test(targetSq) {
							fmt.Printf(
								"King quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("King capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		} else {
			if piece == BK {
				for bitboard != 0 {

					sourceSq = bitboard.FirstOne()

					// init piece attacks
					attacks = KingAttacks[sourceSq] & (^b.Occupancies[BLACK])

					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if !b.Occupancies[WHITE].Test(targetSq) {
							fmt.Printf(
								"King quiet move: %s to %s\n",
								Sq2Fen[sourceSq],
								Sq2Fen[targetSq],
							)
						} else {
							fmt.Printf("King capture move: %s to %s\n", Sq2Fen[sourceSq], Sq2Fen[targetSq])
						}
					}
				}
			}
		}
	}
}
