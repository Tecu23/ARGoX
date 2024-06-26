package main

// Limits are the limits and settings regarding the engine sent via the command
var Limits searchLimits

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

func (s *searchLimits) Init() {
	s.MovesToGo = 30
	s.MoveTime = -1
	s.Time = -1
	s.Depth = -1
	s.Inc = 0
	s.StartTime = 0
	s.StopTime = 0
	s.Timeset = false
	s.Stop = false
	s.Infinite = false
}

// SetStop sets the value of the stop flag
func (s *searchLimits) SetStop(value bool) {
	s.Stop = value
}
