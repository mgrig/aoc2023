package day16

type Ray struct {
	r, c int
	dir  int // N, S, E, W
}

func NewRay(r, c, dir int) *Ray {
	return &Ray{
		r: r, c: c, dir: dir,
	}
}

// nil if next cell is outside
func (ray *Ray) GoToNext(n int) *Ray {
	nextR := ray.r
	nextC := ray.c
	switch ray.dir {
	case N:
		nextR--
	case S:
		nextR++
	case E:
		nextC++
	case W:
		nextC--
	default:
		panic("wrong dir")
	}
	if nextR < 0 || nextR >= n || nextC < 0 || nextC >= n {
		return nil
	}
	return NewRay(nextR, nextC, ray.dir)
}
