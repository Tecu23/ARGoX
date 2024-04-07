package main

var limits searchLimits

type searchLimits struct {
	Depth int // depth of the search

	MovesToGo int // "movestogo" command move counter

	Time     int // "time" command holder
	MoveTime int // "movetime" command time counter

	Inc int // "inc" command's time increment holder

	StartTime int64 // "starttime" command time holder
	StopTime  int64 // "stoptime" command time holder

	Timeset  bool // flag time control availability
	Infinite bool // whether the search is infinte or not

	Stop bool // flag to control when the time is up
}

func (s *searchLimits) init() {
	s.Depth = -1
	s.MovesToGo = 30
	s.MoveTime = -1
	s.Infinite = false
	s.Timeset = false
	s.Stop = false
}

func (s *searchLimits) setStop(st bool) {
	s.Stop = st
}

func (s *searchLimits) setDepth(d int) {
	s.Depth = d
}

func (s *searchLimits) setMoveTime(m int) {
	s.MoveTime = m
}

func (s *searchLimits) setTime(m int) {
	s.MoveTime = m
}

func (s *searchLimits) setInfinite(b bool) {
	s.Infinite = b
}
