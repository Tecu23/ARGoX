package main

// colors
const (
	WHITE = Color(0)
	BLACK = Color(1)
	BOTH  = Color(2)
)

// 6 piece types - no color (P)
const (
	Pawn int = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

// Rows and Columns
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

// 12 pieces with color plus empty
const (
	WP = iota
	WN
	WB
	WR
	WQ
	WK
	BP
	BN
	BB
	BR
	BQ
	BK
	Empty = 15
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

// FEN dedug positions
const (
	EmptyBoard          = "8/8/8/8/8/8/8/8 w - - "
	StartPosition       = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 "
	TrickyPosition      = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1 "
	KillerPosition      = "rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1"
	CmkPosition         = "r2q1rk1/ppp2ppp/2n1bn2/2b1p3/3pP3/3P1NPP/PPP1NPB1/R1BQ1RK1 b - - 0 9 "
	RepetitionsPosition = "2r3k1/R7/8/1R6/8/8/P4KPP/8 w - - 0 40"
)

// Fen2Sq maps fen-sq to int
var Fen2Sq = make(map[string]int)

// Sq2Fen maps int-sq to fen
var Sq2Fen = make(map[int]string)

// InitFen2Sq should initialize the square map from string to int and int to string
func InitFen2Sq() {
	Fen2Sq["a1"] = A1
	Fen2Sq["a2"] = A2
	Fen2Sq["a3"] = A3
	Fen2Sq["a4"] = A4
	Fen2Sq["a5"] = A5
	Fen2Sq["a6"] = A6
	Fen2Sq["a7"] = A7
	Fen2Sq["a8"] = A8

	Fen2Sq["b1"] = B1
	Fen2Sq["b2"] = B2
	Fen2Sq["b3"] = B3
	Fen2Sq["b4"] = B4
	Fen2Sq["b5"] = B5
	Fen2Sq["b6"] = B6
	Fen2Sq["b7"] = B7
	Fen2Sq["b8"] = B8

	Fen2Sq["c1"] = C1
	Fen2Sq["c2"] = C2
	Fen2Sq["c3"] = C3
	Fen2Sq["c4"] = C4
	Fen2Sq["c5"] = C5
	Fen2Sq["c6"] = C6
	Fen2Sq["c7"] = C7
	Fen2Sq["c8"] = C8

	Fen2Sq["d1"] = D1
	Fen2Sq["d2"] = D2
	Fen2Sq["d3"] = D3
	Fen2Sq["d4"] = D4
	Fen2Sq["d5"] = D5
	Fen2Sq["d6"] = D6
	Fen2Sq["d7"] = D7
	Fen2Sq["d8"] = D8

	Fen2Sq["e1"] = E1
	Fen2Sq["e2"] = E2
	Fen2Sq["e3"] = E3
	Fen2Sq["e4"] = E4
	Fen2Sq["e5"] = E5
	Fen2Sq["e6"] = E6
	Fen2Sq["e7"] = E7
	Fen2Sq["e8"] = E8

	Fen2Sq["f1"] = F1
	Fen2Sq["f2"] = F2
	Fen2Sq["f3"] = F3
	Fen2Sq["f4"] = F4
	Fen2Sq["f5"] = F5
	Fen2Sq["f6"] = F6
	Fen2Sq["f7"] = F7
	Fen2Sq["f8"] = F8

	Fen2Sq["g1"] = G1
	Fen2Sq["g2"] = G2
	Fen2Sq["g3"] = G3
	Fen2Sq["g4"] = G4
	Fen2Sq["g5"] = G5
	Fen2Sq["g6"] = G6
	Fen2Sq["g7"] = G7
	Fen2Sq["g8"] = G8

	Fen2Sq["h1"] = H1
	Fen2Sq["h2"] = H2
	Fen2Sq["h3"] = H3
	Fen2Sq["h4"] = H4
	Fen2Sq["h5"] = H5
	Fen2Sq["h6"] = H6
	Fen2Sq["h7"] = H7
	Fen2Sq["h8"] = H8

	// -------------- Sq2Fen
	Sq2Fen[A1] = "a1"
	Sq2Fen[A2] = "a2"
	Sq2Fen[A3] = "a3"
	Sq2Fen[A4] = "a4"
	Sq2Fen[A5] = "a5"
	Sq2Fen[A6] = "a6"
	Sq2Fen[A7] = "a7"
	Sq2Fen[A8] = "a8"

	Sq2Fen[B1] = "b1"
	Sq2Fen[B2] = "b2"
	Sq2Fen[B3] = "b3"
	Sq2Fen[B4] = "b4"
	Sq2Fen[B5] = "b5"
	Sq2Fen[B6] = "b6"
	Sq2Fen[B7] = "b7"
	Sq2Fen[B8] = "b8"

	Sq2Fen[C1] = "c1"
	Sq2Fen[C2] = "c2"
	Sq2Fen[C3] = "c3"
	Sq2Fen[C4] = "c4"
	Sq2Fen[C5] = "c5"
	Sq2Fen[C6] = "c6"
	Sq2Fen[C7] = "c7"
	Sq2Fen[C8] = "c8"

	Sq2Fen[D1] = "d1"
	Sq2Fen[D2] = "d2"
	Sq2Fen[D3] = "d3"
	Sq2Fen[D4] = "d4"
	Sq2Fen[D5] = "d5"
	Sq2Fen[D6] = "d6"
	Sq2Fen[D7] = "d7"
	Sq2Fen[D8] = "d8"

	Sq2Fen[E1] = "e1"
	Sq2Fen[E2] = "e2"
	Sq2Fen[E3] = "e3"
	Sq2Fen[E4] = "e4"
	Sq2Fen[E5] = "e5"
	Sq2Fen[E6] = "e6"
	Sq2Fen[E7] = "e7"
	Sq2Fen[E8] = "e8"

	Sq2Fen[F1] = "f1"
	Sq2Fen[F2] = "f2"
	Sq2Fen[F3] = "f3"
	Sq2Fen[F4] = "f4"
	Sq2Fen[F5] = "f5"
	Sq2Fen[F6] = "f6"
	Sq2Fen[F7] = "f7"
	Sq2Fen[F8] = "f8"

	Sq2Fen[G1] = "g1"
	Sq2Fen[G2] = "g2"
	Sq2Fen[G3] = "g3"
	Sq2Fen[G4] = "g4"
	Sq2Fen[G5] = "g5"
	Sq2Fen[G6] = "g6"
	Sq2Fen[G7] = "g7"
	Sq2Fen[G8] = "g8"

	Sq2Fen[H1] = "h1"
	Sq2Fen[H2] = "h2"
	Sq2Fen[H3] = "h3"
	Sq2Fen[H4] = "h4"
	Sq2Fen[H5] = "h5"
	Sq2Fen[H6] = "h6"
	Sq2Fen[H7] = "h7"
	Sq2Fen[H8] = "h8"
	Sq2Fen[-1] = "-"
}

/*
*
*
*           MOVE CONSTANTS
*
*
 */
// AllMoves and OnlyCaptures flags
const (
	AllMoves     = 0
	OnlyCaptures = 1
)

/*
*
*
*           CASTLINGS CONSTANTS
*
*
 */
/*
- 0001 -- 1 -> white king can castle to the king side
- 0010 -- 2 -> white king can castle to the queen side
- 0100 -- 4 -> black king can castle to the king side
- 1000 -- 8 -> black king can castle to the queen side

ex.

	    1111 -> both side can castle in both directons

		1001 -> black king => queen side
		     -> white king => king side
*/
const (
	ShortW = uint(0x1)
	LongW  = uint(0x2)
	ShortB = uint(0x4)
	LongB  = uint(0x8)
)

/*
*
*
*           SEARCH CONSTANT
*
*
 */

// most valuable victim & less valuable attacker

/*
    (Victims) Pawn Knight Bishop   Rook  Queen   King
  (Attackers)
        Pawn   105    205    305    405    505    605
      Knight   104    204    304    404    504    604
      Bishop   103    203    303    403    503    603
        Rook   102    202    302    402    502    602
       Queen   101    201    301    401    501    601
        King   100    200    300    400    500    600

*/

// MvvLva [attacker][victim]
var MvvLva = [12][12]int{
	{105, 205, 305, 405, 505, 605, 105, 205, 305, 405, 505, 605},
	{104, 204, 304, 404, 504, 604, 104, 204, 304, 404, 504, 604},
	{103, 203, 303, 403, 503, 603, 103, 203, 303, 403, 503, 603},
	{102, 202, 302, 402, 502, 602, 102, 202, 302, 402, 502, 602},
	{101, 201, 301, 401, 501, 601, 101, 201, 301, 401, 501, 601},
	{100, 200, 300, 400, 500, 600, 100, 200, 300, 400, 500, 600},

	{105, 205, 305, 405, 505, 605, 105, 205, 305, 405, 505, 605},
	{104, 204, 304, 404, 504, 604, 104, 204, 304, 404, 504, 604},
	{103, 203, 303, 403, 503, 603, 103, 203, 303, 403, 503, 603},
	{102, 202, 302, 402, 502, 602, 102, 202, 302, 402, 502, 602},
	{101, 201, 301, 401, 501, 601, 101, 201, 301, 401, 501, 601},
	{100, 200, 300, 400, 500, 600, 100, 200, 300, 400, 500, 600},
}

// MaxPly is the maximum ply that we can reach
const MaxPly = 64

// KillerMove [id][ply]
var KillerMove [2][64]Move

// HistoryMove [piece][square]
var HistoryMove [12][64]int

/*
   ================================
         Triangular PV table
   --------------------------------
     PV line: e2e4 e7e5 g1f3 b8c6
   ================================

        0    1    2    3    4    5

   0    m1   m2   m3   m4   m5   m6

   1    0    m2   m3   m4   m5   m6

   2    0    0    m3   m4   m5   m6

   3    0    0    0    m4   m5   m6

   4    0    0    0    0    m5   m6

   5    0    0    0    0    0    m6
*/

// PvLength is the length of the pv table
var PvLength [64]int

// PvTable represents the pv table
var PvTable [64][64]Move

// FollowPv should follow the Priciple Variation
var FollowPv = false

// ScorePv should keep the score of the PV
var ScorePv = false
