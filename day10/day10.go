package day10

func Part1(lines []string) int {
	n := len(lines)
	g := NewGrid(n)
	for r, line := range lines {
		g.SetRow(r, line)
	}
	// fmt.Println(g)

	dist := NewGrid(n)
	dist.Fill(-1)
	toVisit := make([]Visit, 0)
	AddToVisit(g, dist, &toVisit, g.start.r, g.start.c, 0)
	for len(toVisit) > 0 {
		visit := toVisit[len(toVisit)-1]
		toVisit = toVisit[0 : len(toVisit)-1]
		// fmt.Println(visit, toVisit)

		propagate(g, dist, &toVisit, visit)
	}
	// fmt.Println(dist)

	max := 0
	for r := range dist.grid {
		for _, val := range dist.grid[r] {
			if val > max {
				max = val
			}
		}
	}

	return max
}

func propagate(g *Grid, dist *Grid, toVisit *[]Visit, visit Visit) {
	pos := NewCoord(visit.r, visit.c)
	d := dist.GetAt(pos.r, pos.c)
	if d >= 0 && d <= visit.dist {
		// point already processed with same or better reach
		return
	} else {
		dist.grid[pos.r][pos.c] = visit.dist
		d = visit.dist
	}
	pipe := g.GetAt(pos.r, pos.c)
	switch pipe {
	case int('S'):
		p := g.GetAt(pos.r, pos.c-1)
		if p == int('-') || p == int('F') || p == int('L') {
			AddToVisit(g, dist, toVisit, pos.r, pos.c-1, d+1) // left
		}

		p = g.GetAt(pos.r, pos.c+1)
		if p == int('-') || p == int('7') || p == int('J') {
			AddToVisit(g, dist, toVisit, pos.r, pos.c+1, d+1) // right
		}

		p = g.GetAt(pos.r-1, pos.c)
		if p == int('|') || p == int('7') || p == int('F') {
			AddToVisit(g, dist, toVisit, pos.r-1, pos.c, d+1) // up
		}

		p = g.GetAt(pos.r+1, pos.c)
		if p == int('|') || p == int('J') || p == int('L') {
			AddToVisit(g, dist, toVisit, pos.r+1, pos.c, d+1) // down
		}
	case int('.'):
		return
	case int('|'):
		AddToVisit(g, dist, toVisit, pos.r-1, pos.c, d+1) // up
		AddToVisit(g, dist, toVisit, pos.r+1, pos.c, d+1) // down
	case int('-'):
		AddToVisit(g, dist, toVisit, pos.r, pos.c-1, d+1) // left
		AddToVisit(g, dist, toVisit, pos.r, pos.c+1, d+1) // right
	case int('F'):
		AddToVisit(g, dist, toVisit, pos.r+1, pos.c, d+1) // down
		AddToVisit(g, dist, toVisit, pos.r, pos.c+1, d+1) // right
	case int('7'):
		AddToVisit(g, dist, toVisit, pos.r, pos.c-1, d+1) // left
		AddToVisit(g, dist, toVisit, pos.r+1, pos.c, d+1) // down
	case int('J'):
		AddToVisit(g, dist, toVisit, pos.r, pos.c-1, d+1) // left
		AddToVisit(g, dist, toVisit, pos.r-1, pos.c, d+1) // up
	case int('L'):
		AddToVisit(g, dist, toVisit, pos.r-1, pos.c, d+1) // up
		AddToVisit(g, dist, toVisit, pos.r, pos.c+1, d+1) // right
	default:
		panic("wrong pipe")
	}
}

func AddToVisit(g *Grid, dist *Grid, toVisit *[]Visit, r, c int, newDist int) {
	if g.Inside(r, c) {
		*toVisit = append(*toVisit, NewVisit(r, c, newDist))
	}
}
