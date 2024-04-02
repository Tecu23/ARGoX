package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

			b.SearchPosition(d)

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
			return fmt.Errorf("go %v not implemented", words[0])
		}
	} else {
		fmt.Printf("suppose go infinite")
	}

	return nil
}

// SearchPosition should search the current board position for the best move
func (b *BoardStruct) SearchPosition(depth int) {
	fmt.Printf("bestmove d2d4\n")
}

// Uci is the main loop of the engine
func Uci(input chan string) {
	board := BoardStruct{}
	var cmd string

	quit := false

	for !quit {
		select {
		case cmd = <-input:
		}

		words := strings.Split(cmd, " ")
		words[0] = strings.TrimSpace(strings.ToLower(words[0]))

		switch words[0] {
		case "uci":
			fmt.Printf("id name ARGoX\n")
			fmt.Printf("id author Tecu23\n")

			fmt.Printf("uciok\n")
		case "setoption":
			fmt.Printf("setoption command not implemented yet")
		case "isready":
			fmt.Printf("readyok")
		case "ucinewgame":
			board.ParsePosition("position startpos")
		case "position":
			board.ParsePosition(cmd)
		case "debug":
			fmt.Printf("debug command not implemented yet")
		case "register":
			fmt.Printf("register command not implemented yet")
		case "go":
			board.ParseGo(cmd)
		case "ponderhit":
			fmt.Printf("ponderhit command not implemented yet")
		case "stop":
			fmt.Printf("stop command not implemented yet")
		case "quit", "q":
			quit = true
			continue

			/* My Own Commands*/
		case "printboard":
			board.PrintBoard()
		default:
			fmt.Printf("unknown cmd %v", cmd)
		}
	}
}

func input() chan string {
	line := make(chan string)
	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)

	go func() {
		for {
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}
	}()

	return line
}
