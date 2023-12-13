package day13

import "fmt"

type Grid struct {
	grid []string
}

func NewGrid() *Grid {
	return &Grid{
		grid: make([]string, 0),
	}
}

func (g *Grid) Transpose() *Grid {
	m := len(g.grid)
	n := len(g.grid[0])
	ret := make([]string, n)
	for r := 0; r < n; r++ {
		line := ""
		for c := 0; c < m; c++ {
			line += string(g.grid[c][r])
		}
		ret[r] = line
	}
	return &Grid{
		grid: ret,
	}
}

func (g *Grid) String() string {
	ret := ""
	for _, line := range g.grid {
		ret += fmt.Sprintln(line)
	}
	return ret
}
