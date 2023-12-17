package day17

import "fmt"

type Grid struct {
	grid [][]int
}

func NewGrid(m, n int) *Grid {
	grid := make([][]int, m)
	for r := range grid {
		grid[r] = make([]int, n)
	}
	return &Grid{
		grid: grid,
	}
}

func (g *Grid) String() string {
	ret := ""
	for r := range g.grid {
		for _, val := range g.grid[r] {
			ret += fmt.Sprintf("%d", val)
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
