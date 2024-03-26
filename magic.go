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
		fmt.Printf("    %x,\n", findMagicNumbers(sq, RookRelevantBits[sq], false))
	}

	fmt.Printf("};\n\nconst U64 bishop_magics[64] = {\n")

	// loop over 64 board squares
	for sq := A1; sq <= H8; sq++ {
		fmt.Printf("    %x,\n", findMagicNumbers(sq, BishopRelevantBits[sq], true))
	}

	fmt.Printf("};\n\n")
}

func findMagicNumbers(square, relevantBits int, isBishop bool) Bitboard {
	// define occupancies array
	occupancy := [4096]Bitboard{}

	// define attacks array
	attacks := [4096]Bitboard{}

	// define used indices array
	usedAttacks := [4096]Bitboard{}

	var maskAttacks Bitboard
	// mask piece attack
	if isBishop {
		maskAttacks = generateBishopAttacks(square)
	} else {
		maskAttacks = generateRookAttacks(square)
	}

	// occupancy variations
	occupancyVariations := 1 << relevantBits

	// loop over the number of occupancy variations
	for count := 0; count < occupancyVariations; count++ {
		// init occupancies
		occupancy[count] = setOccupancy(count, relevantBits, maskAttacks)

		// init attacks
		if isBishop {
			attacks[count] = generateBishopAttacksOnTheFly(square, occupancy[count])
		} else {
			attacks[count] = generateRookAttacksOnTheFly(square, occupancy[count])
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
var BishopMagicNumbers = [64]uint64{
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
var RookMagicNumbers = [64]uint64{
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
