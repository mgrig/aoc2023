package day10

import "fmt"

func Part1(lines []string) int {
	m := len(lines)
	n := len(lines[0])
	g := NewGrid(m, n)
	for r, line := range lines {
		g.SetRow(r, line)
	}
	// fmt.Println(g)

	dist := NewGrid(m, n)
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

const (
	PROCESSED int = int('p')
	OUTSIDE   int = int('o')
)

func Part2(lines []string) int {
	m := len(lines)
	n := len(lines[0])
	g := NewGrid(m, n)
	for r, line := range lines {
		g.SetRow(r, line)
	}

	start := *g.start
	right := NewCoord(start.r, start.c+1)
	left := NewCoord(start.r, start.c-1)
	up := NewCoord(start.r-1, start.c)
	down := NewCoord(start.r+1, start.c)
	var hist *[]Coord

	// Identify the loop
	if Advance(g, start, right) != nil {
		hist = FollowPipe(g, start, right)
	}

	if hist == nil && Advance(g, start, left) != nil {
		hist = FollowPipe(g, start, left)
	}

	if hist == nil && Advance(g, start, up) != nil {
		hist = FollowPipe(g, start, up)
	}
	if hist == nil && Advance(g, start, down) != nil {
		hist = FollowPipe(g, start, down)
	}

	// replace S with actual pipe piece
	first := (*hist)[0]
	last := (*hist)[len(*hist)-1]
	rep := computePipe(start, first, last)
	// fmt.Printf("replace start with %c\n", rep)
	g.grid[start.r][start.c] = rep

	// ... and add the start coord to hist
	*hist = append(*hist, start)
	// fmt.Println(len(*hist))

	// create a grid with just the loop ...
	loopGrid := NewGrid(m, n)
	for _, h := range *hist {
		loopGrid.grid[h.r][h.c] = int('*')
	}
	// fmt.Println(loopGrid)

	// ... so we can clear all pipes that are not in the loop in the main grid (g)
	for r := range g.grid {
		for c := range g.grid[r] {
			if loopGrid.grid[r][c] == 0 {
				g.grid[r][c] = '.'
			}
		}
	}
	fmt.Println(g)

	// Coord (0, 0) cannot be inside, so we start from it and propagate the outside info
	// along CORNERS!
	// A cell of grid g is only INNER if all its 4 corners are inner (aka not marked as OUTSIDE).
	// Propagating OUTSIDE info on corners, and not cells allows the propagation to "sqeeze"
	// between neighboring loop cells.
	/*
		  |      |
		--C------C--
		  |      |
		  | Cell |
		  |      |
		--C------C--
		  |      |
	*/
	corners := NewGrid(m+1, n+1) // 1 will mean outside
	x := make([]Coord, 0)
	toPropagate := &x
	*toPropagate = append(*toPropagate, NewCoord(0, 0)) // 0, 0 is certainly outside
	// count := 1
	for len(*toPropagate) > 0 {
		pos := (*toPropagate)[len(*toPropagate)-1]
		*toPropagate = (*toPropagate)[0 : len(*toPropagate)-1]

		propagateCorners(g, loopGrid, corners, toPropagate, pos)
		// count++
		// if count%100 == 0 {
		// 	fmt.Println(corners)
		// }
	}
	fmt.Println(corners)

	// count inner cells
	inner := NewGrid(m, n)
	sum := 0
	for r := range g.grid {
		for c := range g.grid[r] {
			if corners.GetAt(r, c) != OUTSIDE && corners.GetAt(r, c+1) != OUTSIDE && corners.GetAt(r+1, c+1) != OUTSIDE && corners.GetAt(r+1, c) != OUTSIDE {
				inner.grid[r][c] = int('i')
				sum += 1
			}
		}
	}
	fmt.Println(inner)

	return sum
}

func propagateCorners(g, loopGrid, corners *Grid, toPropagate *[]Coord, pos Coord) {
	if !corners.InsideCoord(pos) || corners.GetAt(pos.r, pos.c) != 0 {
		return
	}
	// if pos.r == 70 && pos.c == 70 {
	// 	fmt.Println(corners)
	// 	panic("bzzz")
	// }
	corners.grid[pos.r][pos.c] = OUTSIDE

	right := NewCoord(pos.r, pos.c+1)
	if corners.InsideCoord(right) && !( /*blocked*/ loopGrid.GetAt(pos.r, pos.c) == int('*') && loopGrid.GetAt(pos.r-1, pos.c) == int('*') && oneOf(g.GetAt(pos.r, pos.c), int('|'), int('L'), int('J'))) {
		// corners.grid[pos.r][pos.c+1] = OUTSIDE
		*toPropagate = append(*toPropagate, right)
	}

	left := NewCoord(pos.r, pos.c-1)
	if corners.InsideCoord(left) && !( /*blocked*/ loopGrid.GetAt(pos.r, pos.c-1) == int('*') && loopGrid.GetAt(pos.r-1, pos.c-1) == int('*') && oneOf(g.GetAt(pos.r, pos.c-1), int('|'), int('L'), int('J'))) {
		// corners.grid[pos.r][pos.c-1] = OUTSIDE
		*toPropagate = append(*toPropagate, left)
	}

	up := NewCoord(pos.r-1, pos.c)
	if corners.InsideCoord(up) && !( /*blocked*/ loopGrid.GetAt(pos.r-1, pos.c-1) == int('*') && loopGrid.GetAt(pos.r-1, pos.c) == int('*') && oneOf(g.GetAt(pos.r-1, pos.c), int('-'), int('7'), int('J'))) {
		// corners.grid[pos.r-1][pos.c] = OUTSIDE
		*toPropagate = append(*toPropagate, up)
	}

	down := NewCoord(pos.r+1, pos.c)
	if corners.InsideCoord(down) && !( /*blocked*/ loopGrid.GetAt(pos.r, pos.c-1) == int('*') && loopGrid.GetAt(pos.r, pos.c) == int('*') && oneOf(g.GetAt(pos.r, pos.c), int('-'), int('7'), int('J'))) {
		// corners.grid[pos.r+1][pos.c] = OUTSIDE
		*toPropagate = append(*toPropagate, down)
	}
}

func computePipe(start, first, last Coord) int {
	firstRel := first.relativeTo(start)
	lastRel := last.relativeTo(start)

	if firstRel == lastRel {
		panic("wrong first last")
	}

	switch firstRel {
	case UP:
		switch lastRel {
		case DOWN:
			return int('|')
		case RIGHT:
			return int('L')
		case LEFT:
			return int('J')
		}
	case DOWN:
		switch lastRel {
		case UP:
			return int('|')
		case RIGHT:
			return int('F')
		case LEFT:
			return int('7')
		}
	case RIGHT:
		switch lastRel {
		case UP:
			return int('L')
		case DOWN:
			return int('F')
		case LEFT:
			return int('-')
		}
	case LEFT:
		switch lastRel {
		case UP:
			return int('J')
		case DOWN:
			return int('7')
		case RIGHT:
			return int('-')
		}
	}
	panic("OOPS")
}

func oneOf(elem int, values ...int) bool {
	for _, v := range values {
		if elem == v {
			return true
		}
	}
	return false
}

// returns nil if broken pipe
func FollowPipe(g *Grid, prev, pos Coord) *[]Coord {
	hist := make([]Coord, 0)
	for {
		next := Advance(g, prev, pos)
		if next == nil || !g.Inside(next.r, next.c) {
			return nil // end of pipe
		}
		hist = append(hist, pos)
		if *next == *g.start {
			// fmt.Println("loop found!", hist)
			return &hist
		}
		prev = pos
		pos = *next
	}
}

// returns nil if it cannot advance
func Advance(g *Grid, prevPos, pos Coord) (newPos *Coord) { //TODO or just prevPos instead of full hist here?
	if !g.Inside(pos.r, pos.c) {
		return nil
	}
	pipe := g.GetAt(pos.r, pos.c)

	var coord Coord
	if prevPos.r == pos.r && prevPos.c+1 == pos.c { // coming from left
		switch pipe {
		case int('-'):
			coord = NewCoord(pos.r, pos.c+1)
		case int('J'):
			coord = NewCoord(pos.r-1, pos.c)
		case int('7'):
			coord = NewCoord(pos.r+1, pos.c)
		default:
			return nil
		}
	} else if prevPos.r == pos.r && prevPos.c-1 == pos.c { // coming from right
		switch pipe {
		case int('-'):
			coord = NewCoord(pos.r, pos.c-1)
		case int('L'):
			coord = NewCoord(pos.r-1, pos.c)
		case int('F'):
			coord = NewCoord(pos.r+1, pos.c)
		default:
			return nil
		}
	} else if prevPos.r-1 == pos.r && prevPos.c == pos.c { // coming from down
		switch pipe {
		case int('|'):
			coord = NewCoord(pos.r-1, pos.c)
		case int('F'):
			coord = NewCoord(pos.r, pos.c+1)
		case int('7'):
			coord = NewCoord(pos.r, pos.c-1)
		default:
			return nil
		}
	} else if prevPos.r+1 == pos.r && prevPos.c == pos.c { // coming from up
		switch pipe {
		case int('|'):
			coord = NewCoord(pos.r+1, pos.c)
		case int('L'):
			coord = NewCoord(pos.r, pos.c+1)
		case int('J'):
			coord = NewCoord(pos.r, pos.c-1)
		default:
			return nil
		}
	} else {
		return nil // helps with S
		// panic("wrong prevPos")
	}
	return &coord
}

func propagate(g *Grid, dist *Grid, toVisit *[]Visit, visit Visit) {
	pos := NewCoord(visit.r, visit.c)
	d := dist.GetAt(pos.r, pos.c)
	if d >= 0 && d <= visit.dist {
		// point already processed with same or better reach
		fmt.Println("already processed:", d, visit.dist)
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
