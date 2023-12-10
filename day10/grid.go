package day10

import "fmt"

type Coord struct {
	r, c int
}

func NewCoord(r, c int) Coord {
	return Coord{
		r: r,
		c: c,
	}
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

func NewGrid(n int) *Grid {
	g := make([][]int, n)
	for r, _ := range g {
		g[r] = make([]int, n)
	}
	return &Grid{
		grid:  g,
		start: nil,
	}
}

func (g *Grid) Fill(val int) {
	for r, _ := range g.grid {
		for c, _ := range g.grid[r] {
			g.grid[r][c] = val
		}
	}
}

func (g *Grid) SetRow(r int, line string) {
	if len(g.grid[r]) != len(line) {
		panic("wrong line")
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
	n := len(g.grid)
	return r >= 0 && c >= 0 && r < n && c < n
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
		for c := range g.grid[r] {
			ret += fmt.Sprintf("%4d ", g.grid[r][c])
		}
		ret += fmt.Sprintln()
	}
	return ret
}
