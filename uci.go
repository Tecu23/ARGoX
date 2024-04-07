package main

// TODO: Add some sort of helper commands for debugging

// TODO: Make engine run in a separate channel so that it can be stopped

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var saveBm Move

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
func (b *BoardStruct) ParseGo(cmd string, stopped bool, toEng chan bool) error {
	startTime := int64(0) // "startime" command time holder
	stopTime := int64(0)  // "stoptime" command time holder
	movesToGo := 30       // UCI moves counter
	moveTime := -1        // UCI time counter
	time := -1            // time command holder (ms)
	inc := 0              // command's time increment holder
	d := -1               // depth of the search

	timeset := false // varibale to falg time control availability

	cmd = strings.TrimSpace(strings.TrimPrefix(cmd, "go"))

	words := strings.Split(cmd, " ")

	if len(words) > 0 {
		words[0] = strings.TrimSpace(strings.ToLower(words[0]))

		fmt.Println(words)

		switch words[0] {
		case "searchmoves":
			return fmt.Errorf("go searchmoves not implemented")
		case "ponder":
			return fmt.Errorf("go ponder not implemented")
		case "wtime":
			arg, err := 0, error(nil)

			if len(words) >= 2 {
				arg, err = strconv.Atoi(words[1])
			}

			if arg < 0 || err != nil {
				return fmt.Errorf("time value not numeric")
			}

			if b.SideToMove == WHITE {
				inc = arg
			}
		case "btime":
			arg, err := 0, error(nil)

			if len(words) >= 2 {
				arg, err = strconv.Atoi(words[1])
			}

			if arg < 0 || err != nil {
				return fmt.Errorf("time value not numeric")
			}

			if b.SideToMove == BLACK {
				inc = arg
			}
		case "winc":
			arg, err := 0, error(nil)

			if len(words) >= 2 {
				arg, err = strconv.Atoi(words[1])
			}

			if arg < 0 || err != nil {
				return fmt.Errorf("increment value not numeric")
			}

			if b.SideToMove == WHITE {
				inc = arg
			}
		case "binc":
			arg, err := 0, error(nil)

			if len(words) >= 2 {
				arg, err = strconv.Atoi(words[1])
			}

			if arg < 0 || err != nil {
				return fmt.Errorf("increment value not numeric")
			}

			if b.SideToMove == BLACK {
				inc = arg
			}

		case "movestogo":
			arg, err := 0, error(nil)

			if len(words) >= 2 {
				arg, err = strconv.Atoi(words[1])
			}

			if arg < 0 || err != nil {
				return fmt.Errorf("time value not numeric")
			}

			movesToGo = arg
		case "depth":
			arg, err := -1, error(nil)

			if len(words) >= 2 {
				arg, err = strconv.Atoi(words[1])
			}

			if arg < 0 || err != nil {
				return fmt.Errorf("depth not numeric")
			}

			d = arg

		case "nodes":
			return fmt.Errorf("go nodes not implemented")
		case "movetime":
			arg, err := -1, error(nil)

			if len(words) >= 2 {
				arg, err = strconv.Atoi(words[1])
			}

			if arg < 0 || err != nil {
				return fmt.Errorf("depth not numeric")
			}

			moveTime = arg
		case "mate": // mate <x>  mate in x moves
			return fmt.Errorf("go mate not implemented")
		case "infinite":
			return fmt.Errorf("go infinite not implemented")
		case "register":
			return fmt.Errorf("go register not implemented")
		case "perft":
			arg, err := -1, error(nil)

			if len(words) >= 2 {
				arg, err = strconv.Atoi(words[1])
			}

			if arg < 0 || err != nil {
				return fmt.Errorf("depth not numeric")
			}

			perftTest(b, arg)
			return nil
		default:
			return fmt.Errorf("go %v not implemented", words[0])
		}
	} else {
		fmt.Printf("suppose go infinite")
	}

	if moveTime != -1 { // if move time is not available
		time = moveTime
		movesToGo = 1
	}

	startTime = GetTimeInMiliseconds()

	if time != -1 { // if time control is available
		timeset = true

		time /= movesToGo
		time -= 50
		stopTime = startTime + int64(time) + int64(inc)
	}

	// if depth not available use 64 as default
	if d == -1 {
		d = 64
	}

	fmt.Printf(
		"time:%d, start:%d, stop:%d, depth:%d timeset:%t stopped:%t\n",
		time,
		startTime,
		stopTime,
		d,
		timeset,
		stopped,
	)

	limits.setDepth(d)

	toEng <- true

	return nil
}

func handleBm(bm string) {
	fmt.Println(bm)
}

// Uci is the main loop of the engine
func Uci(input chan string) {
	board := BoardStruct{}
	var cmd string
	var bm string

	toEng, frEng := engine(&board)

	quit := false

	stopped := false // flag when the time is up

	for !quit {
		select {
		case cmd = <-input:
		case bm = <-frEng:
			handleBm(bm)
			continue
		}

		fmt.Println(cmd)
		fmt.Println(bm)

		words := strings.Split(cmd, " ")
		words[0] = strings.TrimSpace(strings.ToLower(words[0]))

		switch words[0] {
		case "uci":
			fmt.Printf("id name ARGoX\n")
			fmt.Printf("id author Tecu23\n")

			fmt.Printf("uciok\n")
		case "setoption":
			fmt.Printf("setoption command not implemented yet\n")
		case "isready":
			fmt.Printf("readyok\n")
		case "ucinewgame":
			board.ParsePosition("position startpos")
		case "position":
			board.ParsePosition(cmd)
		case "debug":
			fmt.Printf("debug command not implemented yet\n")
		case "register":
			fmt.Printf("register command not implemented yet\n")
		case "go":
			board.ParseGo(cmd, stopped, toEng)
		case "ponderhit":
			fmt.Printf("ponderhit command not implemented yet\n")
		case "stop":
			if limits.Infinite {
				if saveBm != NoMove {
					fmt.Println(saveBm)
					saveBm = NoMove
				}

				limits.setInfinite(false)
			}
			limits.setStop(true)
			continue
		case "quit", "q":
			quit = true
			continue

			/* My Own Commands*/
		case "printboard":
			board.PrintBoard()
		case "eval":
			fmt.Printf("Score: %d\n", board.EvaluatePosition())
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
