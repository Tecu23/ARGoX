package main

import "fmt"

// HashSize is the size of the tt
const HashSize = 0x400000

// NoHashEntry => no hash entry was found
const NoHashEntry = 100000

// Hashf constants
const (
	HashfExact = iota // Exact score return from search
	HashfAlpha        // Node fails low
	HashfBeta         // Node fails high
)

// transposition table
type ttItem struct {
	Key   uint64
	Depth int
	Flags int // the type of node (fail-low/fail-high/PV)
	Score int // score (alpha/beta/PV)
}

type tt []ttItem

// TransTable is the global variable that keeps the values
var TransTable = make(tt, HashSize)

func (t *tt) Clear() {
	for _, v := range *t {
		v.Key = 0
		v.Depth = 0
		v.Flags = 0
		v.Score = 0
	}
}

func (t *tt) ReadEntry(alpha, beta, depth int, key uint64) int {
	hashEntry := (*t)[key%HashSize]

	if hashEntry.Key == key {
		if hashEntry.Depth >= depth {
			score := hashEntry.Score

			// extract score independent from the actual path from root to curr position
			if score < -MateScore {
				score += Ply
			}

			if score > MateScore {
				score -= Ply
			}

			if hashEntry.Flags == HashfExact {
				return score
			}

			if hashEntry.Flags == HashfAlpha && score <= alpha {
				return alpha
			}

			if hashEntry.Flags == HashfBeta && score >= beta {
				return beta
			}
		}
	}
	return NoHashEntry
}

func (t *tt) WriteEntry(score, depth, hashFlag int, key uint64) {
	hashEntry := &(*t)[key%HashSize]

	// store score independent from the actual path from root to curr position
	if score < -MateScore {
		score -= Ply
	}

	if score > MateScore {
		score += Ply
	}

	hashEntry.Key = key
	hashEntry.Score = score
	hashEntry.Flags = hashFlag
	hashEntry.Depth = depth
}

func (t *tt) PrintAll() {
	for _, v := range *t {
		fmt.Printf("%x, Sc:%d, Depth:%d, F:%d\n", v.Key, v.Score, v.Depth, v.Flags)
	}
}
