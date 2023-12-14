package day14

type History struct {
	hist map[string](map[int]int) // grid as string -> (direction -> cycle)
}

func NewHistory() *History {
	return &History{
		hist: make(map[string](map[int]int), 0),
	}
}

func (cg *History) Get(in *Grid, dir int) (prevCycle int, found bool) {
	byDir, found := cg.hist[in.String()]
	if !found {
		return -1, false
	}
	value, found := byDir[dir]
	if found {
		// fmt.Println("inner cg hit")
		return value, true
	}
	return -1, false
}

func (cg *History) Set(in *Grid, dir int, cycle int) {
	_, found := cg.hist[in.String()]
	if !found {
		cg.hist[in.String()] = make(map[int]int)
	}
	cg.hist[in.String()][dir] = cycle
}
