package main

import (
	"fmt"
)

// Engine represents the engine function
func Engine(b *BoardStruct) (toEngine chan bool, frEngine chan string) {
	frEngine = make(chan string)
	toEngine = make(chan bool)
	go root(toEngine, frEngine, b)

	return
}

func root(toEng chan bool, frEng chan string, b *BoardStruct) {
	for range toEng {
		if Limits.MoveTime != -1 { // if move time is not available
			Limits.Time = Limits.MoveTime
			Limits.MovesToGo = 1
		}

		Limits.StartTime = GetTimeInMiliseconds()

		if Limits.Time != -1 { // if time control is available
			Limits.Timeset = true

			Limits.Time /= Limits.MovesToGo
			if Limits.Time > 1500 {
				Limits.Time -= 50
			}
			Limits.StopTime = Limits.StartTime + int64(Limits.Time) + int64(Limits.Inc)
		}

		// if depth not available use 64 as default
		if Limits.Depth == -1 {
			Limits.Depth = 64
		}

		fmt.Printf(
			"time:%d, start:%d, stop:%d, depth:%d timeset:%t stopped:%t\n",
			Limits.Time,
			Limits.StartTime,
			Limits.StopTime,
			Limits.Depth,
			Limits.Timeset,
			Limits.Stop,
		)
		bm := b.SearchPosition(Limits.Depth)

		frEng <- fmt.Sprintf("bestmove %s", bm)
	}
}
