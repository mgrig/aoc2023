package day17

import "fmt"

type State struct {
	r, c          int
	dir           int
	maxStepsInDir int
}

func NewState(r, c, dir, maxStepsInDir int) State {
	return State{
		r:             r,
		c:             c,
		dir:           dir,
		maxStepsInDir: maxStepsInDir,
	}
}

func (s State) String() string {
	dir := "n/a"
	switch s.dir {
	case N:
		dir = "N"
	case S:
		dir = "S"
	case E:
		dir = "E"
	case W:
		dir = "W"
	}
	return fmt.Sprintf("{(%d, %d) %s %d}", s.r, s.c, dir, s.maxStepsInDir)
}

// Can end up outside a grid!
func (s State) NextInDir() State {
	switch s.dir {
	case N:
		return NewState(s.r-1, s.c, s.dir, s.maxStepsInDir-1)
	case S:
		return NewState(s.r+1, s.c, s.dir, s.maxStepsInDir-1)
	case E:
		return NewState(s.r, s.c+1, s.dir, s.maxStepsInDir-1)
	case W:
		return NewState(s.r, s.c-1, s.dir, s.maxStepsInDir-1)
	default:
		panic("wrong dir")
	}
}

func (s State) NextLeft() State {
	switch s.dir {
	case E:
		return NewState(s.r-1, s.c, N, 9)
	case W:
		return NewState(s.r+1, s.c, S, 9)
	case S:
		return NewState(s.r, s.c+1, E, 9)
	case N:
		return NewState(s.r, s.c-1, W, 9)
	default:
		panic("wrong dir")
	}
}

func (s State) NextRight() State {
	switch s.dir {
	case W:
		return NewState(s.r-1, s.c, N, 9)
	case E:
		return NewState(s.r+1, s.c, S, 9)
	case N:
		return NewState(s.r, s.c+1, E, 9)
	case S:
		return NewState(s.r, s.c-1, W, 9)
	default:
		panic("wrong dir")
	}
}
