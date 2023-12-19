package day19

import (
	"aoc2023/common"
	"regexp"
)

const (
	X string = "x"
	M string = "m"
	A string = "a"
	S string = "s"
)

type Part struct {
	xmas map[string]int // [xmas] -> value
}

func NewPart(x, m, a, s int) Part {
	xmas := map[string]int{
		X: x,
		M: m,
		A: a,
		S: s,
	}
	return Part{
		xmas: xmas,
	}
}

var rePart *regexp.Regexp = regexp.MustCompile(`^x=(\d+),m=(\d+),a=(\d+),s=(\d+)$`)

func ParsePart(str string) Part {
	tokens := rePart.FindStringSubmatch(str)
	if len(tokens) != 5 {
		panic("wrong part")
	}
	return NewPart(common.StringToInt(tokens[1]), common.StringToInt(tokens[2]), common.StringToInt(tokens[3]), common.StringToInt(tokens[4]))
}

func (p Part) Sum() int {
	sum := 0
	for _, v := range p.xmas {
		sum += v
	}
	return sum
}
