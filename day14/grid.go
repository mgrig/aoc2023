package day14

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

func (g *Grid) Clone() (other *Grid) {
	m := len(g.grid)
	n := len(g.grid[0])
	otherGrid := make([][]int, m)
	for r := range otherGrid {
		otherGrid[r] = make([]int, n)
		copy(otherGrid[r], g.grid[r])
	}
	return &Grid{
		grid: otherGrid,
	}
}

func (g *Grid) String() string {
	ret := ""
	for _, row := range g.grid {
		for _, val := range row {
			ret += string(val)
		}
		ret += fmt.Sprintln()
	}
	return ret
}
