package main

// colors
const (
	WHITE = Color(0)
	BLACK = Color(1)
)

const (
	Row1  = Bitboard(0x00000000000000FF)
	Row2  = Bitboard(0x000000000000FF00)
	Row3  = Bitboard(0x0000000000FF0000)
	Row4  = Bitboard(0x00000000FF000000)
	Row5  = Bitboard(0x000000FF00000000)
	Row6  = Bitboard(0x0000FF0000000000)
	Row7  = Bitboard(0x00FF000000000000)
	Row8  = Bitboard(0xFF00000000000000)
	FileA = Bitboard(0x0101010101010101)
	FileB = Bitboard(0x0202020202020202)
	FileC = Bitboard(0x0404040404040404)
	FileD = Bitboard(0x0808080808080808)
	FileE = Bitboard(0x1010101010101010)
	FileF = Bitboard(0x2020202020202020)
	FileG = Bitboard(0x4040404040404040)
	FileH = Bitboard(0x8080808080808080)
)

// directions
const (
	E  = +1
	W  = -1
	N  = 8
	S  = -8
	NW = +7
	NE = +9
	SW = -NE
	SE = -NW
)

// square names
const (
	A1 = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1

	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2

	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3

	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4

	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5

	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6

	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7

	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)
