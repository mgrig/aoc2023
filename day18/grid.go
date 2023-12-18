package day18

import "fmt"

type Grid struct {
	grid [][]int
}

func NewGrid(m, n int) *Grid {
	g := make([][]int, m)
	for r := range g {
		g[r] = make([]int, n)
	}
	return &Grid{
		grid: g,
	}
}

func (g *Grid) String() string {
	ret := ""
	for r := range g.grid {
		for _, val := range g.grid[r] {
			if val == DIG {
				ret += "#"
			} else {
				ret += "."
			}
		}
		ret += fmt.Sprintln()
	}
	return ret
}

func (g *Grid) Inside(r, c int) bool {
	m := len(g.grid)
	n := len(g.grid[0])
	return r >= 0 && c >= 0 && r < m && c < n
}

func (g *Grid) InsideCoord(coord Coord) bool {
	return g.Inside(coord.r, coord.c)
}

func (g *Grid) GetAt(r, c int) int {
	if !g.Inside(r, c) {
		return 0
	}
	return g.grid[r][c]
}

// ****

type BoolGrid struct {
	grid [][]bool
}

func NewBoolGrid(m, n int) *BoolGrid {
	g := make([][]bool, m)
	for r := range g {
		g[r] = make([]bool, n)
	}
	return &BoolGrid{
		grid: g,
	}
}

func (g *BoolGrid) Inside(r, c int) bool {
	m := len(g.grid)
	n := len(g.grid[0])
	return r >= 0 && c >= 0 && r < m && c < n
}

func (g *BoolGrid) InsideCoord(coord Coord) bool {
	return g.Inside(coord.r, coord.c)
}

func (g *BoolGrid) GetAt(r, c int) bool {
	if !g.Inside(r, c) {
		return false
	}
	return g.grid[r][c]
}

func (g *BoolGrid) String() string {
	ret := ""
	for r := range g.grid {
		for _, val := range g.grid[r] {
			if val {
				ret += "#"
			} else {
				ret += "."
			}
		}
		ret += fmt.Sprintln()
	}
	return ret
}
