package main

import (
	"fmt"
)

func engine(b *BoardStruct) (toEngine chan bool, frEngine chan string) {
	frEngine = make(chan string)
	toEngine = make(chan bool)
	go root(toEngine, frEngine, b)

	return
}

func root(toEng chan bool, frEng chan string, b *BoardStruct) {
	for range toEng {
		if limits.MoveTime != -1 { // if move time is not available
			limits.Time = limits.MoveTime
			limits.MovesToGo = 1
		}

		limits.StartTime = GetTimeInMiliseconds()

		if limits.Time != -1 { // if time control is available
			limits.Timeset = true

			limits.Time /= limits.MovesToGo
			limits.Time -= 50
			limits.StopTime = limits.StartTime + int64(limits.Time) + int64(limits.Inc)
		}

		// if depth not available use 64 as default
		if limits.Depth == -1 {
			limits.Depth = 64
		}

		fmt.Printf(
			"time:%d, start:%d, stop:%d, depth:%d timeset:%t stopped:%t\n",
			limits.Time,
			limits.StartTime,
			limits.StopTime,
			limits.Depth,
			limits.Timeset,
			limits.Stop,
		)
		fmt.Println("Calling search position for depth ", limits.Depth)
		bm := b.SearchPosition(limits.Depth)

		frEng <- fmt.Sprintf("bestmove %s", bm)
	}
}
