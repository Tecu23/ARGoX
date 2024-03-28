package main

import (
	"fmt"
	"unicode"
)

/*
         binary move bits representaion                    hexadecimal constants

   0000 0000 0000 0000 0011 1111 source square              0x3f
   0000 0000 0000 1111 1100 0000 target square              0xfc0
   0000 0000 1111 0000 0000 0000 piece                      0xf000
   0000 1111 0000 0000 0000 0000 promoted piece             0xf0000
   0001 0000 0000 0000 0000 0000 capture flag               0x100000
   0010 0000 0000 0000 0000 0000 double push flag           0x200000
   0100 0000 0000 0000 0000 0000 enpassant capture flag     0x400000
   1000 0000 0000 0000 0000 0000 castling flag              0x800000
*/

// Shift and Mask Constants
const (
	SourceMask     = 0x3f
	TargetMask     = 0xfc0
	PieceMask      = 0xf000
	PromotedMask   = 0xf0000
	CaptureMask    = 0x100000
	DoublePushMask = 0x200000
	EnpassantMask  = 0x400000
	CastlingMask   = 0x800000

	SourceShift     = 0
	TargetShift     = 6
	PieceShift      = 12
	PromotedShift   = 16
	CaptureShift    = 20
	DoublePushShift = 21
	EnpassantShift  = 22
	CastlingShift   = 23
)

// Move is a representation of a move in binary format
type Move uint64

// EncodeMove creates a new move from every detail we need
func (m *Move) EncodeMove(
	source, target, piece, promoted, capture, doublePush, enpassant, castling int,
) {
	*m = Move(
		(source) | (target << TargetShift) | (piece << PieceShift) |
			(promoted << PromotedShift) | (capture << CaptureShift) |
			(doublePush << DoublePushShift) | (enpassant << EnpassantShift) |
			(castling << CastlingShift),
	)
}

// GetSource should retrieve the source square of a move
func (m Move) GetSource() int {
	return int(m & SourceMask)
}

// GetTarget should retrieve the target square of a move
func (m Move) GetTarget() int {
	return int(m&TargetMask) >> TargetShift
}

// GetPiece should retrieve the target square of a move
func (m Move) GetPiece() int {
	return int(m&PieceMask) >> PieceShift
}

// GetPromoted should retrieve the target square of a move
func (m Move) GetPromoted() int {
	return int(m&PromotedMask) >> PromotedShift
}

// GetCapture should retrieve the target square of a move
func (m Move) GetCapture() int {
	return int(m&CaptureMask) >> CaptureShift
}

// GetDoublePush should retrieve the target square of a move
func (m Move) GetDoublePush() int {
	return int(m&DoublePushMask) >> DoublePushShift
}

// GetEnpassant should retrieve the target square of a move
func (m Move) GetEnpassant() int {
	return int(m&EnpassantMask) >> EnpassantShift
}

// GetCastling should retrieve the target square of a move
func (m Move) GetCastling() int {
	return int(m&CastlingMask) >> CastlingShift
}

// PrintMove should print the a move
func (m Move) PrintMove() {
	fmt.Printf(
		"%s%s",
		Sq2Fen[m.GetSource()],
		Sq2Fen[m.GetTarget()],
	)

	if m.GetPromoted() != 0 {
		fmt.Printf("%c ", unicode.ToLower(rune(AciiPieces[m.GetPromoted()])))
	} else {
		fmt.Printf(" ")
	}

	// fmt.Printf("Promotion: %c", AciiPieces[m.GetPromoted()])
	// fmt.Printf("Capture: %t\n", m.GetCapture() != 0)
	// fmt.Printf("Double Push: %t\n", m.GetDoublePush() != 0)
	// fmt.Printf("Enpassant: %t\n", m.GetEnpassant() != 0)
	// fmt.Printf("Castling move: %t\n", m.GetCastling() != 0)
}
