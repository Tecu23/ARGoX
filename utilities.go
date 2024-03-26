package main

import (
	"fmt"
	"strconv"
	"strings"
)

// piece char definitions
const (
	PcFen = "PpNnBbRrQqKk     "
)

// Ascii representation of chess pieces
const (
	AciiPieces = "PNBRQKpnbrqk"
)

// ParseFEN should parse a FEN string and retrieve the board
// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -
func ParseFEN(Board BoardStruct, FEN string) {
	// board.Clear()

	fenIdx := 0
	sq := 0

	// parsing the FEN from the start and setting the from top to bottom
	for row := 7; row >= 0; row-- {
		for sq = row * 8; sq < row*8+8; {

			char := string(FEN[fenIdx])
			fenIdx++

			if char == "/" {
				continue
			}

			// if we find a number we should skip that many squares from our current board
			if i, err := strconv.Atoi(char); err == nil {
				for j := 0; j < i; j++ {
					Board.SetSq(Empty, sq)
					sq++
				}
				continue
			}

			// if we find an invalid piece we skip
			if strings.IndexAny(PcFen, char) == -1 {
				fmt.Printf("Invalid piece %s try next one", char)
				// log.Errorf("error string invalid piece %s try next one", char)
				continue
			}

			Board.SetSq(Fen2pc(char), sq)

			sq++
		}
	}

	remaining := strings.Split(strings.TrimSpace(FEN[fenIdx:]), " ")

	// Setting the Side to Move
	if len(remaining) > 0 {
		if remaining[0] == "w" {
			Board.SideToMove = WHITE
		} else if remaining[0] == "b" {
			Board.SideToMove = BLACK
		} else {
			fmt.Printf("Remaining=%v; sq=%v;  fenIx=%v;", strings.Join(remaining, " "), sq, fenIdx)
			fmt.Printf("%s invalid side to move color", remaining[0])
			// log.Errorf("info string remaining=%v; sq=%v;  fenIx=%v;", strings.Join(remaining, " "), sq, fenIdx)
			// log.Errorf("info string %s invalid stm color", remaining[0])
			Board.SideToMove = WHITE
		}
	}

	// Checking for castling
	Board.Castlings = 0
	if len(remaining) > 1 {
		Board.Castlings = ParseCastlings(remaining[1])
	}

	// En Passant
	Board.EnPassant = 0
	if len(remaining) > 2 {
		if remaining[2] != "-" {
			Board.EnPassant = Fen2Sq[remaining[2]]
		}
	}

	// Cheking for 50 move rule
	// Board.Rule50 = 0
	// if len(remaining) > 3 {
	// 	Board.Rule50 = parse50(remaining[3])
	// }
}

// Fen2pc convert pieceString to pc int
func Fen2pc(c string) int {
	for p, x := range PcFen {
		if string(x) == c {
			return p
		}
	}
	return Empty
}

// PcColor returns the color of a piece
func PcColor(pc int) Color {
	return Color(pc & 0x1)
}
