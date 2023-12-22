package day22

type Grid struct {
	grid [][]int
}

func NewGrid(n int) *Grid {
	g := make([][]int, n)
	for x := range g {
		g[x] = make([]int, n)
	}
	return &Grid{
		grid: g,
	}
}

// ****
type Grid3 struct {
	grid [][][]int
}

func NewGrid3(n, maxZ int) *Grid3 {
	g := make([][][]int, n)
	for x := range g {
		g[x] = make([][]int, n)
		for y := range g[x] {
			g[x][y] = make([]int, maxZ+2)
		}
	}
	return &Grid3{
		grid: g,
	}
}
