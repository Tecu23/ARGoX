package main

import "strings"

// Castlings represents the castling possibilities for a given position
type Castlings uint

/*
                                castling  move    in       in
                                   right update   binary   decimal

   king & rooks didn't move:       1111 & 1111  =  1111    15

           white king moved:       1111 & 1100  =  1100    12
    white king's rook moved:       1111 & 1110  =  1110    14
   white queen's rook moved:       1111 & 1101  =  1101    13

           black king moved:       1111 & 0011  =  0011    3
    black king's rook moved:       1111 & 1011  =  1011    11
   black queen's rook moved:       1111 & 0111  =  0111    7
*/

// CastlingRights update constants
var CastlingRights = []uint{
	13, 15, 15, 15, 12, 15, 15, 14,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	7, 15, 15, 15, 3, 15, 15, 11,
}

// ParseCastlings should parse the castlings part of a FEN string
func ParseCastlings(fenCastl string) Castlings {
	c := uint(0)

	if fenCastl == "-" {
		return Castlings(0)
	}

	if strings.Contains(fenCastl, "K") {
		c |= ShortW
	}
	if strings.Contains(fenCastl, "Q") {
		c |= LongW
	}
	if strings.Contains(fenCastl, "k") {
		c |= ShortB
	}
	if strings.Contains(fenCastl, "q") {
		c |= LongB
	}

	return Castlings(c)
}

// String returns the string represantion of the castling
func (c Castlings) String() string {
	flags := ""
	if uint(c)&ShortW != 0 {
		flags = "K"
	}
	if uint(c)&LongW != 0 {
		flags += "Q"
	}
	if uint(c)&ShortB != 0 {
		flags += "k"
	}
	if uint(c)&LongB != 0 {
		flags += "q"
	}
	if flags == "" {
		flags = "-"
	}
	return flags
}
