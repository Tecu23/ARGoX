package main

// Movelist will keep track of all the moves being played throught the game
type Movelist []Move

// AddMove should add a move to the movelist
func (m *Movelist) AddMove(move Move) {
	*m = append(*m, move)
}

// PrintMovelist should print all moves that happen during the game
func (m Movelist) PrintMovelist() {
	for _, move := range m {
		move.PrintMove()
	}
}
