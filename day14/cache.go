package day14

type CacheLine struct {
	cache map[string]([]int)
}

func NewCache() *CacheLine {
	return &CacheLine{
		cache: make(map[string][]int, 0),
	}
}

func toKey(in []int) string {
	ret := ""
	for _, val := range in {
		ret += string(val)
	}
	return ret
}

func (c *CacheLine) Get(in []int) (out []int, found bool) {
	out, found = c.cache[toKey(in)]
	return out, found
}

func (c *CacheLine) Set(in, out []int) {
	c.cache[toKey(in)] = out
}

// ****

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
