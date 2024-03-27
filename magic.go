package main

import (
	"fmt"
	"math/rand"
)

// masks
var (
	BishopMasks [64]Bitboard
	RookMasks   [64]Bitboard
)

// attacks
var (
	BishopAttacks [64][512]Bitboard
	RookAttacks   [64][4096]Bitboard
)

// InitMagic should retrieve the random magic numbers
func InitMagic() {
	fmt.Printf("const Bitboard rookMagics[64] = {\n")

	// loop over 64 board squares
	for sq := A1; sq <= H8; sq++ {
		fmt.Printf("    %x,\n", findMagicNumbers(sq, RookRelevantBits[sq], Rook))
	}

	fmt.Printf("};\n\nconst U64 bishop_magics[64] = {\n")

	// loop over 64 board squares
	for sq := A1; sq <= H8; sq++ {
		fmt.Printf("    %x,\n", findMagicNumbers(sq, BishopRelevantBits[sq], Bishop))
	}

	fmt.Printf("};\n\n")
}

func findMagicNumbers(square, relevantBits int, piece int) Bitboard {
	// define occupancies array
	occupancy := [4096]Bitboard{}

	// define attacks array
	attacks := [4096]Bitboard{}

	// define used indices array
	usedAttacks := [4096]Bitboard{}

	var maskAttacks Bitboard
	// mask piece attack
	if piece == Bishop {
		maskAttacks = GenerateBishopAttacks(square)
	} else {
		maskAttacks = GenerateRookAttacks(square)
	}

	// occupancy variations
	occupancyVariations := 1 << relevantBits

	// loop over the number of occupancy variations
	for count := 0; count < occupancyVariations; count++ {
		// init occupancies
		occupancy[count] = SetOccupancy(count, relevantBits, maskAttacks)

		// init attacks
		if piece == Bishop {
			attacks[count] = GenerateBishopAttacksOnTheFly(square, occupancy[count])
		} else {
			attacks[count] = GenerateRookAttacksOnTheFly(square, occupancy[count])
		}
	}

	// test magic numbers
	for randomCount := 0; randomCount < 100000000; randomCount++ {

		// init magic number candidate
		magic := Bitboard(rand.Uint64())

		// skip testing magic number if innappropriate
		if Bitboard((maskAttacks*magic)&0xFF00000000000000).Count() < 6 {
			continue
		}

		// reset used attacks array
		usedAttacks = [4096]Bitboard{}

		// init count & fail flag
		count, fail := 0, false

		// test magic index
		for count, fail = 0, false; !fail && count < occupancyVariations; count++ {
			// generate magic index
			magicIndex := int((occupancy[count] * magic) >> (64 - relevantBits))

			// if got free index
			if usedAttacks[magicIndex] == 0 {
				// assign corresponding attack map
				usedAttacks[magicIndex] = attacks[count]
			} else if usedAttacks[magicIndex] != attacks[count] {
				fail = true
			}
		}

		// return magic if it works
		if !fail {
			return magic
		}
	}

	// on fail
	fmt.Printf("***Failed***\n")
	return Bitboard(0)
}

// SetOccupancy should create an occupancy bitboard given an attack mask and a index
func SetOccupancy(index, bitsInMask int, attackMask Bitboard) Bitboard {
	// occupancy map
	occupancy := Bitboard(0)

	// loop over the range of bits withing attackMask
	for count := 0; count < bitsInMask; count++ {
		// get LSB index of attacks mask
		square := attackMask.FirstOne()

		// make sure occupancy is on board
		if index&(1<<count) != 0 {
			// populate occupancy map
			occupancy |= (1 << square)
		}
	}

	return occupancy
}

// FillOptimalMagicsB should fill the found optiomal magic numbers and relevant bits
func FillOptimalMagicsB() {
	BishopRelevantBits[A1] = 5
	BishopMagicNumbers[A1] = 0xffedf9fd7cfcffff
	BishopRelevantBits[B1] = 4
	BishopMagicNumbers[B1] = 0xfc0962854a77f576
	BishopRelevantBits[C1] = 5
	BishopMagicNumbers[C1] = 0xE433BF9FF9BD3C0D
	BishopRelevantBits[D1] = 5
	BishopMagicNumbers[D1] = 0x8F0BBE9CF98C0405
	BishopRelevantBits[E1] = 5
	BishopMagicNumbers[E1] = 0x7E11DFD9DDFBDBF0
	BishopRelevantBits[G1] = 4
	BishopMagicNumbers[G1] = 0xfc0a66c64a7ef576
	BishopRelevantBits[H1] = 5
	BishopMagicNumbers[H1] = 0x7ffdfdfcbd79ffff
	BishopRelevantBits[A2] = 4
	BishopMagicNumbers[A2] = 0xfc0846a64a34fff6
	BishopRelevantBits[B2] = 4
	BishopMagicNumbers[B2] = 0xfc087a874a3cf7f6
	BishopRelevantBits[C2] = 5
	BishopMagicNumbers[C2] = 0x0040020042188680
	BishopRelevantBits[D2] = 5
	BishopMagicNumbers[D2] = 0x0080000108D80200
	// BishopRelevantBits[E2] = 5
	// BishopMagicNumbers[E2] = 0xF2048D48B0240820
	BishopRelevantBits[F2] = 5
	BishopMagicNumbers[F2] = 0x810040B921030010
	BishopRelevantBits[G2] = 4
	BishopMagicNumbers[G2] = 0xfc0864ae59b4ff76
	BishopRelevantBits[H2] = 4
	BishopMagicNumbers[H2] = 0x3c0860af4b35ff76
	BishopRelevantBits[A3] = 4
	BishopMagicNumbers[A3] = 0x73C01AF56CF4CFFB
	BishopRelevantBits[B3] = 4
	BishopMagicNumbers[B3] = 0x41A01CFAD64AAFFC
	BishopRelevantBits[G3] = 4
	BishopMagicNumbers[G3] = 0x7c0c028f5b34ff76
	BishopRelevantBits[H3] = 4
	BishopMagicNumbers[H3] = 0xfc0a028e5ab4df76
	BishopRelevantBits[A6] = 4
	BishopMagicNumbers[A6] = 0xDCEFD9B54BFCC09F
	BishopRelevantBits[B6] = 4
	BishopMagicNumbers[B6] = 0xF95FFA765AFD602B
	BishopRelevantBits[G6] = 4
	BishopMagicNumbers[G6] = 0x43ff9a5cf4ca0c01
	BishopRelevantBits[H6] = 4
	BishopMagicNumbers[H6] = 0x4BFFCD8E7C587601
	BishopRelevantBits[A7] = 4
	BishopMagicNumbers[A7] = 0xfc0ff2865334f576
	BishopRelevantBits[B7] = 4
	BishopMagicNumbers[B7] = 0xfc0bf6ce5924f576
	BishopRelevantBits[G7] = 4
	BishopMagicNumbers[G7] = 0xc3ffb7dc36ca8c89
	BishopRelevantBits[H7] = 4
	BishopMagicNumbers[H7] = 0xc3ff8a54f4ca2c89
	BishopRelevantBits[A8] = 5
	BishopMagicNumbers[A8] = 0xfffffcfcfd79edff
	BishopRelevantBits[B8] = 4
	BishopMagicNumbers[B8] = 0xfc0863fccb147576
	BishopRelevantBits[G8] = 4
	BishopMagicNumbers[G8] = 0xfc087e8e4bb2f736
	BishopRelevantBits[H8] = 5
	BishopMagicNumbers[H8] = 0x43ff9e4ef4ca2c89
}

// FillOptimalMagicsR should fill the found optiomal magic numbers and relevant bits
func FillOptimalMagicsR() {
	RookRelevantBits[A7] = 10
	RookMagicNumbers[A7] = 0x48FFFE99FECFAA00
	RookRelevantBits[B7] = 9
	RookMagicNumbers[B7] = 0x48FFFE99FECFAA00
	RookRelevantBits[C7] = 9
	RookMagicNumbers[C7] = 0x497FFFADFF9C2E00
	RookRelevantBits[D7] = 9
	RookMagicNumbers[D7] = 0x613FFFDDFFCE9200
	RookRelevantBits[E7] = 9
	RookMagicNumbers[E7] = 0xffffffe9ffe7ce00
	RookRelevantBits[F7] = 9
	RookMagicNumbers[F7] = 0xfffffff5fff3e600
	RookRelevantBits[G7] = 9
	RookMagicNumbers[G7] = 0x3ff95e5e6a4c0
	RookRelevantBits[H7] = 10
	RookMagicNumbers[H7] = 0x510FFFF5F63C96A0
	RookRelevantBits[A8] = 11
	RookMagicNumbers[A8] = 0xEBFFFFB9FF9FC526
	RookRelevantBits[B8] = 10
	RookMagicNumbers[B8] = 0x61FFFEDDFEEDAEAE
	RookRelevantBits[C8] = 10
	RookMagicNumbers[C8] = 0x53BFFFEDFFDEB1A2
	RookRelevantBits[D8] = 10
	RookMagicNumbers[D8] = 0x127FFFB9FFDFB5F6
	RookRelevantBits[E8] = 10
	RookMagicNumbers[E8] = 0x411FFFDDFFDBF4D6
	RookRelevantBits[G8] = 10
	RookMagicNumbers[G8] = 0x0003ffef27eebe74
	RookRelevantBits[H8] = 11
	RookMagicNumbers[H8] = 0x7645FFFECBFEA79E
}

// BishopRelevantBits is the relevant occupancy bit count for every square on board
var BishopRelevantBits = [64]int{
	6, 5, 5, 5, 5, 5, 5, 6,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	6, 5, 5, 5, 5, 5, 5, 6,
}

// RookRelevantBits is the relevant bit count for every square on the board
var RookRelevantBits = [64]int{
	12, 11, 11, 11, 11, 11, 11, 12,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	12, 11, 11, 11, 11, 11, 11, 12,
}

// BishopMagicNumbers is the magic numbers for bishops
var BishopMagicNumbers = [64]Bitboard{
	0xc085080200420200,
	0x60014902028010,
	0x401240100c201,
	0x580ca104020080,
	0x8434052000230010,
	0x102080208820420,
	0x2188410410403024,
	0x40120805282800,
	0x4420410888208083,
	0x1049494040560,
	0x6090100400842200,
	0x1000090405002001,
	0x48044030808c409,
	0x20802080384,
	0x2012008401084008,
	0x9741088200826030,
	0x822000400204c100,
	0x14806004248220,
	0x30200101020090,
	0x148150082004004,
	0x6020402112104,
	0x4001000290080d22,
	0x2029100900400,
	0x804203145080880,
	0x60a10048020440,
	0xc08080b20028081,
	0x1009001420c0410,
	0x101004004040002,
	0x1004405014000,
	0x10029a0021005200,
	0x4002308000480800,
	0x301025015004800,
	0x2402304004108200,
	0x480110c802220800,
	0x2004482801300741,
	0x400400820a60200,
	0x410040040040,
	0x2828080020011000,
	0x4008020050040110,
	0x8202022026220089,
	0x204092050200808,
	0x404010802400812,
	0x422002088009040,
	0x180604202002020,
	0x400109008200,
	0x2420042000104,
	0x40902089c008208,
	0x4001021400420100,
	0x484410082009,
	0x2002051108125200,
	0x22e4044108050,
	0x800020880042,
	0xb2020010021204a4,
	0x2442100200802d,
	0x10100401c4040000,
	0x2004a48200c828,
	0x9090082014000,
	0x800008088011040,
	0x4000000a0900b808,
	0x900420000420208,
	0x4040104104,
	0x120208c190820080,
	0x4000102042040840,
	0x8002421001010100,
}

// RookMagicNumbers is the rook magic numbers
var RookMagicNumbers = [64]Bitboard{
	0x11800040001481a0,
	0x2040400010002000,
	0xa280200308801000,
	0x100082005021000,
	0x280280080040006,
	0x200080104100200,
	0xc00040221100088,
	0xe00072200408c01,
	0x2002045008600,
	0xa410804000200089,
	0x4081002000401102,
	0x2000c20420010,
	0x800800400080080,
	0x40060010041a0009,
	0x441004442000100,
	0x462800080004900,
	0x80004020004001,
	0x1840420021021081,
	0x8020004010004800,
	0x940220008420010,
	0x2210808008000400,
	0x24808002000400,
	0x803604001019a802,
	0x520000440081,
	0x802080004000,
	0x1200810500400024,
	0x8000100080802000,
	0x2008080080100480,
	0x8000404002040,
	0xc012040801104020,
	0xc015000900240200,
	0x20040200208041,
	0x1080004000802080,
	0x400081002110,
	0x30002000808010,
	0x2000100080800800,
	0x2c0800400800800,
	0x1004800400800200,
	0x818804000210,
	0x340082000a45,
	0x8520400020818000,
	0x2008900460020,
	0x100020008080,
	0x601001000a30009,
	0xc001000408010010,
	0x2040002008080,
	0x11008218018c0030,
	0x20c0080620011,
	0x400080002080,
	0x8810040002500,
	0x400801000200080,
	0x2402801000080480,
	0x204040280080080,
	0x31044090200801,
	0x40c10830020400,
	0x442800100004080,
	0x10080002d005041,
	0x134302820010a2c2,
	0x6202001080200842,
	0x1820041000210009,
	0x1002001008210402,
	0x2000108100402,
	0x10310090a00b824,
	0x800040100944822,
}
