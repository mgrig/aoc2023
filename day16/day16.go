package day16

const (
	EMPTY     int = int('.')
	SLASH     int = int('/')
	BACKSLASH int = int('\\')
	PIPE      int = int('|')
	MINUS     int = int('-')

	N int = 0
	S int = 1
	E int = 2
	W int = 3
)

func Part1(lines []string) int {
	n := len(lines)
	hist := NewGrid(n)
	toPropagate := []Ray{*NewRay(0, 0, E)} // ray enters (0, 0), goes towards E

	for len(toPropagate) > 0 {
		ray := toPropagate[0]
		toPropagate = toPropagate[1:]

		propagate(ray, &lines, hist, &toPropagate)
	}
	// fmt.Println(hist)

	sum := 0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if len(hist.hist[r][c]) > 0 {
				sum++
			}
		}
	}

	return sum
}

var SlashNextDir map[int]int = map[int]int{
	N: E,
	S: W,
	E: N,
	W: S,
}

var BackslashNextDir map[int]int = map[int]int{
	N: W,
	S: E,
	E: S,
	W: N,
}

func propagate(ray Ray, lines *[]string, hist *Grid, toPropagate *[]Ray) {
	wasKnown := hist.AddDir(ray.r, ray.c, ray.dir)
	if wasKnown {
		return
	}
	n := len(*lines)
	var nextRay *Ray = nil
	cell := int((*lines)[ray.r][ray.c])
	switch cell {
	case EMPTY:
		nextRay = ray.GoToNext(n)
	case SLASH:
		nextRay = NewRay(ray.r, ray.c, SlashNextDir[ray.dir]).GoToNext(n)
	case BACKSLASH:
		nextRay = NewRay(ray.r, ray.c, BackslashNextDir[ray.dir]).GoToNext(n)
	case PIPE:
		if ray.dir == E || ray.dir == W {
			nextRay = NewRay(ray.r, ray.c, N).GoToNext(n)
			*toPropagate = append(*toPropagate, *NewRay(ray.r, ray.c, S))
		} else {
			nextRay = ray.GoToNext(n)
		}
	case MINUS:
		if ray.dir == N || ray.dir == S {
			nextRay = NewRay(ray.r, ray.c, E).GoToNext(n)
			*toPropagate = append(*toPropagate, *NewRay(ray.r, ray.c, W))
		} else {
			nextRay = ray.GoToNext(n)
		}
	}

	if nextRay != nil {
		*toPropagate = append(*toPropagate, *nextRay)
	}
}
