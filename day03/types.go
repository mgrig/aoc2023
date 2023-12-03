package day03

type Coord struct {
	r, c int // 0-based
}

func NewCoord(r, c int) Coord {
	return Coord{
		r: r,
		c: c,
	}
}

func (co Coord) Equals(r int, c int) bool {
	return co.r == r && co.c == c
}

// ******

type PosNumber struct {
	value    int
	startPos Coord
}

func NewPosNumber(value int, startPos Coord) *PosNumber {
	return &PosNumber{
		value:    value,
		startPos: startPos,
	}
}

func (pn *PosNumber) AppendDigit(digit int) {
	pn.value = 10*pn.value + digit
}
