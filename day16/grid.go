package day16

type Grid struct {
	hist [][]([]int) // r, c ( -> is E: ray enters towards E, aka through W side!)
}

func NewGrid(n int) *Grid {
	hist := make([][][]int, n)
	for r := range hist {
		hist[r] = make([][]int, n)
		for c := range hist[r] {
			hist[r][c] = make([]int, 0)
		}
	}
	return &Grid{
		hist: hist,
	}
}

// func (h *Grid) IsKnownDir(r, c, dir int) bool {
// 	dirs := h.hist[r][c]
// 	for _, d := range dirs {
// 		if d == dir {
// 			return true
// 		}
// 	}
// 	return false
// }

func (h *Grid) AddDir(r, c, dir int) (wasKnown bool) {
	dirs := h.hist[r][c]
	for _, d := range dirs {
		if d == dir {
			return true
		}
	}
	h.hist[r][c] = append(h.hist[r][c], dir)
	return false
}
