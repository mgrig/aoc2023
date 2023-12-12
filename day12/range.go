package day12

type Range struct {
	start int
	size  int
}

func NewRange(start, size int) Range {
	if start < 0 || size <= 0 {
		panic("wrong input")
	}
	return Range{
		start: start,
		size:  size,
	}
}

func (r Range) fits(line string) bool {
	if r.start+r.size > len(line) {
		return false // goes past end
	}
	for i := r.start; i < r.start+r.size; i++ {
		if line[i] == byte(OPERATIONAL) {
			return false
		}
	}
	return true
}
