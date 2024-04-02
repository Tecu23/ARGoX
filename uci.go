package main

import (
	"fmt"
	"strconv"
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

// ParseGo should parse the go command
func (b *BoardStruct) ParseGo(cmd string) error {
	cmd = strings.TrimSpace(strings.TrimPrefix(cmd, "go"))

	words := strings.Split(cmd, " ")

	if len(words) > 0 {
		words[0] = strings.TrimSpace(strings.ToLower(words[0]))

		switch words[0] {
		case "searchmoves":
			return fmt.Errorf("go searchmoves not implemented")
		case "ponder":
			return fmt.Errorf("go ponder not implemented")
		case "wtime":
			return fmt.Errorf("go wtime not implemented")
		case "btime":
			return fmt.Errorf("go btime not implemented")
		case "winc":
			return fmt.Errorf("go winc not implemented")
		case "binc":
			return fmt.Errorf("go binc not implemented")
		case "movestogo":
			return fmt.Errorf("go movestogo not implemented")
		case "depth":
			d, err := -1, error(nil)

			if len(words) >= 2 {
				d, err = strconv.Atoi(words[1])
			}

			if d < 0 || err != nil {
				return fmt.Errorf("depth not numeric")
			}

			fmt.Println(d)

		case "nodes":
			return fmt.Errorf("go nodes not implemented")
		case "movetime":
			return fmt.Errorf("go movetime not implemented")
		case "mate": // mate <x>  mate in x moves
			return fmt.Errorf("go mate not implemented")
		case "infinite":
			return fmt.Errorf("go infinite not implemented")
		case "register":
			return fmt.Errorf("go register not implemented")
		default:
			return fmt.Errorf("go ", words[1], " not implemented")
		}
	} else {
		fmt.Printf("suppose go infinite")
	}

	return nil
}
