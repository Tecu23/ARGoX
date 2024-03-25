package main

import (
	"fmt"
	"math/bits"
)

// Bitboard represent a 64 bit unsigned integer
type Bitboard uint64

// Count should count the number of ones in a 64 bit number
func (b Bitboard) Count() int {
	return bits.OnesCount64(uint64(b))
}

// Set should set the bit at position pos to 1
func (b *Bitboard) Set(pos int) {
	*b |= Bitboard(uint64(1) << uint(pos))
}

// Test should return whether the bit at position pos is a 1 or not
func (b *Bitboard) Test(pos int) bool {
	return *b&Bitboard(uint64(1)<<uint(pos)) != 0
}

// Clear should remove the bit at position pos
func (b *Bitboard) Clear(pos int) {
	*b &= Bitboard(^(uint64(1) << uint(pos)))
}

// FirstOne should retrieve the index of the least significant first bit
func (b *Bitboard) FirstOne() int {
	bit := bits.TrailingZeros64(uint64(*b))
	if bit == 64 {
		return 64
	}
	*b = (*b >> uint(bit+1)) << uint(bit+1)
	return bit
}

// LastOne should retrieve the index of the most significant first bit
func (b *Bitboard) LastOne() int {
	bit := bits.LeadingZeros64(uint64(*b))
	if bit == 64 {
		return 64
	}
	*b = (*b << uint(bit+1)) >> uint(bit+1)
	return 63 - bit
}

// returns the full bitstring (with leading zeroes) of the bitBoard
func (b Bitboard) String() string {
	zeroes := ""
	for i := 0; i < 64; i++ {
		zeroes = zeroes + "0"
	}

	bits := zeroes + fmt.Sprintf("%b", b)
	return bits[len(bits)-64:]
}

// PrintBitboard should print the bitboard
func (b Bitboard) PrintBitboard() {
	s := b.String()
	row := [8]string{}
	row[0] = s[0:8]
	row[1] = s[8:16]
	row[2] = s[16:24]
	row[3] = s[24:32]
	row[4] = s[32:40]
	row[5] = s[40:48]
	row[6] = s[48:56]
	row[7] = s[56:]
	for i, r := range row {
		fmt.Printf(
			"%d   %v %v %v %v %v %v %v %v\n", 8-i,
			r[7:8],
			r[6:7],
			r[5:6],
			r[4:5],
			r[3:4],
			r[2:3],
			r[1:2],
			r[0:1],
		)
	}
	fmt.Print("\n")
	fmt.Printf("    a b c d e f g h\n\n")
}
