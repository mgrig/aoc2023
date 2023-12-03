package day03

type SymbolGrid struct {
	grid [][]bool
}

func NewSymbolGrid(N int) *SymbolGrid {
	sg := &SymbolGrid{
		grid: make([][]bool, N),
	}

	for i := 0; i < N; i++ {
		sg.grid[i] = make([]bool, N)
	}

	return sg
}

func (sg *SymbolGrid) SetSymbolAt(coord Coord) {
	sg.grid[coord.r][coord.c] = true
}

func (sg *SymbolGrid) isInside(r, c int) bool {
	N := len(sg.grid)
	return r >= 0 && r < N && c >= 0 && c < N
}

// coord allowed to be outside grid > returns false
func (sg *SymbolGrid) IsSymbolAt(r, c int) bool {
	if !sg.isInside(r, c) {
		return false
	}

	return sg.grid[r][c]
}
