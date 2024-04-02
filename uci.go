package main

import (
	"fmt"
	"strings"
)

// ParsePosition should parse the position command
func (b *BoardStruct) ParsePosition(cmd string) error {
	cmd = strings.TrimSpace(strings.TrimPrefix(cmd, "position"))

	parts := strings.Split(cmd, "moves")

	if len(parts) == 0 || len(parts) > 2 {
		return fmt.Errorf("%v wrong length=%v", parts, len(parts))
	}

	pos := strings.Split(parts[0], " ")
	pos[0] = strings.TrimSpace(pos[0])

	if pos[0] == "startpos" {
		parts[0] = StartPosition
	} else if pos[0] == "fen" {
		parts[0] = strings.TrimSpace(strings.TrimPrefix(parts[0], "fen"))
	} else {
		return fmt.Errorf("%#v must be %#v or %#v", pos[0], "fen", "startpos")
	}

	ParseFEN(b, parts[0])

	if len(parts) == 2 {
		parts[1] = strings.ToLower(strings.TrimSpace(parts[1]))
		err := b.ParseMoves(parts[1])
		if err != nil {
			return err
		}
	}

	return nil
}

// ParseMoves should parse all the moves sent in a position command
func (b *BoardStruct) ParseMoves(cmd string) error {
	mvs := strings.Fields(strings.ToLower(cmd))

	for _, mv := range mvs {
		mv = strings.TrimSpace(mv)

		if len(mv) < 4 || len(mv) > 5 {
			return fmt.Errorf("%v is not a move", mv)
		}

		m := b.ParseMove(mv)

		if m == NoMove {
			return fmt.Errorf("%v is not a valid move", mv)
		}

		if !b.MakeMove(m, AllMoves) {
			return fmt.Errorf("%v is not a valid move", mv)
		}
	}

	return nil
}
