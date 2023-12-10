package day10

import (
	"aoc2023/common"
	"fmt"
)

type Coord struct {
	r, c int
}

func NewCoord(r, c int) Coord {
	return Coord{
		r: r,
		c: c,
	}
}

const (
	UP    int = 1
	DOWN  int = 2
	RIGHT int = 3
	LEFT  int = 4
)

func (c Coord) relativeTo(other Coord) int {
	dx := c.c - other.c
	dy := c.r - other.r
	if dx*dy != 0 || common.IntAbs(dx-dy) != 1 {
		panic("wrong relatives")
	}
	if dx == 0 {
		if dy == 1 {
			return DOWN
		}
		return UP
	}
	if dx == 1 {
		return RIGHT
	}
	return LEFT
}

// ****

type Visit struct {
	r, c int
	dist int // possible dist during visit
}

func NewVisit(r, c, dist int) Visit {
	return Visit{
		r:    r,
		c:    c,
		dist: dist,
	}
}

// ****

type Grid struct {
	grid  [][]int
	start *Coord
}

func NewGrid(m, n int) *Grid {
	g := make([][]int, m)
	for r := range g {
		g[r] = make([]int, n)
	}
	return &Grid{
		grid:  g,
		start: nil,
	}
}

func (g *Grid) Fill(val int) {
	for r := range g.grid {
		for c := range g.grid[r] {
			g.grid[r][c] = val
		}
	}
}

func (g *Grid) SetRow(r int, line string) {
	if len(g.grid[r]) != len(line) {
		panic(fmt.Sprintf("wrong line: %s", line))
	}
	for c, val := range line {
		g.grid[r][c] = int(val)
		if val == 'S' {
			g.start = &Coord{
				r: r,
				c: c,
			}
		}
	}
}

func (g *Grid) Inside(r, c int) bool {
	m := len(g.grid)
	n := len(g.grid[0])
	return r >= 0 && c >= 0 && r < m && c < n
}

func (g *Grid) InsideCoord(pos Coord) bool {
	return g.Inside(pos.r, pos.c)
}

func (g *Grid) GetAt(r, c int) int {
	if !g.Inside(r, c) {
		return int('.')
	}
	return g.grid[r][c]
}

func (g *Grid) String() string {
	// return fmt.Sprintf("%v, start:%v", g.grid, *g.start)
	ret := ""
	for r := range g.grid {
		for _, val := range g.grid[r] {
			if val == 0 {
				ret += " "
			} else {
				ret += fmt.Sprintf("%c", val)
			}
		}
		ret += fmt.Sprintln()
	}
	return ret
}
