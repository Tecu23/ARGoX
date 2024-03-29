package main

import "fmt"

// generate all moves
func (b *BoardStruct) generateMoves(movelist *Movelist) {
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, WQ, 0, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, WR, 0, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, WB, 0, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, WN, 0, 0, 0, 0))
						} else {

							// one square ahead move
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))

							// two square ahead move
							if (sourceSq >= A2 && sourceSq <= H2) && !b.Occupancies[BOTH].Test(targetSq+N) {
								movelist.AddMove(EncodeMove(sourceSq, targetSq+N, piece, 0, 0, 1, 0, 0))
							}
						}
					}

					// init pawn attacks bb
					attacks = PawnAttacks[WHITE][sourceSq] & b.Occupancies[BLACK]

					// generate pawn captures
					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if sourceSq >= A7 && sourceSq <= H7 {

							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, WQ, 1, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, WR, 1, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, WB, 1, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, WN, 1, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
						}
					}

					// generate EnPassant captures
					if b.EnPassant != -1 {
						enpassantAttacks := PawnAttacks[WHITE][sourceSq] & (1 << b.EnPassant)

						// check enpassant capture
						if enpassantAttacks != 0 {
							// init enpassant capture target square
							targetEnpassant := enpassantAttacks.FirstOne()
							movelist.AddMove(
								EncodeMove(sourceSq, targetEnpassant, piece, 0, 1, 0, 1, 0),
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
							movelist.AddMove(EncodeMove(E1, G1, piece, 0, 0, 0, 0, 1))
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
							movelist.AddMove(EncodeMove(E1, C1, piece, 0, 0, 0, 0, 1))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, BQ, 0, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, BR, 0, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, BB, 0, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, BN, 0, 0, 0, 0))
						} else {
							// one square ahead move
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))

							// two square ahead move
							if (sourceSq >= A7 && sourceSq <= H7) && !b.Occupancies[BOTH].Test(targetSq+S) {
								movelist.AddMove(EncodeMove(sourceSq, targetSq+S, piece, 0, 0, 1, 0, 0))
							}
						}
					}

					// init pawn attacks bb
					attacks = PawnAttacks[BLACK][sourceSq] & b.Occupancies[WHITE]

					// generate pawn captures
					for attacks != 0 {
						targetSq = attacks.FirstOne()

						if sourceSq >= A2 && sourceSq <= H2 {

							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, BQ, 1, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, BR, 1, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, BB, 1, 0, 0, 0))
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, BN, 1, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
						}
					}

					// generate EnPassant captures
					if b.EnPassant != -1 {
						enpassantAttacks := PawnAttacks[BLACK][sourceSq] & (1 << b.EnPassant)

						// check enpassant capture
						if enpassantAttacks != 0 {
							// init enpassant capture target square
							targetEnpassant := enpassantAttacks.FirstOne()
							movelist.AddMove(EncodeMove(sourceSq, targetEnpassant, piece, 0, 1, 0, 1, 0))
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
							movelist.AddMove(EncodeMove(E8, G8, piece, 0, 0, 0, 0, 1))
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
							movelist.AddMove(EncodeMove(E8, C8, piece, 0, 0, 0, 0, 1))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
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
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 0, 0, 0, 0))
						} else {
							movelist.AddMove(EncodeMove(sourceSq, targetSq, piece, 0, 1, 0, 0, 0))
						}
					}
				}
			}
		}
	}
}
