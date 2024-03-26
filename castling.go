package main

type Castlings uint

/*
- 0001 -- 1 -> white king can castle to the king side
- 0010 -- 2 -> white king can castle to the queen side
- 0100 -- 4 -> black king can castle to the king side
- 1000 -- 8 -> black king can castle to the queen side

ex.

	    1111 -> both side can castle in both directons

		1001 -> black king => queen side
		     -> white king => king side
*/
const (
	ShortW = uint(0x1)
	LongW  = uint(0x2)
	ShortB = uint(0x4)
	LongB  = uint(0x8)
)

func (c Castlings) String() string {
	flags := ""
	if uint(c)&ShortW != 0 {
		flags = "K"
	}
	if uint(c)&LongW != 0 {
		flags += "Q"
	}
	if uint(c)&ShortB != 0 {
		flags += "k"
	}
	if uint(c)&LongB != 0 {
		flags += "q"
	}
	if flags == "" {
		flags = "-"
	}
	return flags
}
