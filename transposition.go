package main

const HASH_SIZE = 0x400000

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

var TransTable = make(tt, HASH_SIZE)

func (t *tt) Clear() {
	for _, v := range *t {
		v.Key = 0
		v.Depth = 0
		v.Flags = 0
		v.Score = 0
	}
}
