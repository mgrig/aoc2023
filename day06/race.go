package day06

import (
	"aoc2023/common"
	"math"
	"strings"
)

func Part1(lines []string) int {
	if len(lines) != 2 {
		panic("not 2 lines")
	}

	timeLine, _ := strings.CutPrefix(lines[0], "Time:")
	times := stringToIntArray(timeLine)

	distLine, _ := strings.CutPrefix(lines[1], "Distance:")
	dists := stringToIntArray(distLine)

	prod := 1
	for i, t := range times {
		d := dists[i]
		prod *= CountWins(t, d)
	}

	return prod
}

func Part2(lines []string) int {
	if len(lines) != 2 {
		panic("not 2 lines")
	}

	timeLine, _ := strings.CutPrefix(lines[0], "Time:")
	timeLine = strings.ReplaceAll(timeLine, " ", "")
	times := stringToIntArray(timeLine)

	distLine, _ := strings.CutPrefix(lines[1], "Distance:")
	distLine = strings.ReplaceAll(distLine, " ", "")
	dists := stringToIntArray(distLine)

	prod := 1
	for i, t := range times {
		d := dists[i]
		prod *= CountWins(t, d)
	}

	return prod
}

func CountWins(t, d int) int {
	// (t-p)*p > d  ->  p^2 - tp + d < 0

	// a x^2 + b x + c = 0
	/*
		x^2 + b/a x + c/a = 0
		x^2 + 2 b/2a x + (b/2a)^2 - (b/2a)^2 + c/a = 0
		(x + b/2a)^2 = (b/2a)^2 - c/a
	*/

	/*
		a = 1
		b = -t
		c = d

		delta = (b/2)^2 - c = t^2 / 4 - d

		x0 = -b/2 - sqrt(delta) = t/2 - sqrt(delta)
		x1 = -b/2 + sqrt(delta)
	*/

	tf := float64(t)
	df := float64(d)

	// compute square roots ...
	delta := tf*tf/4 - df
	sqDelta := math.Sqrt(delta)
	x0 := tf/2 - sqDelta
	x1 := tf/2 + sqDelta
	// ... and round towards the "inner" integers
	x0Int := int(math.Floor(x0 + 1))
	x1Int := int(math.Ceil(x1 - 1))

	// intersect with [1..t-1]
	if 1 > x1Int || (t-1) < x0Int {
		// no intersection
		return 0
	}
	min := common.IntMax(x0Int, 1)
	max := common.IntMin(x1Int, t-1)

	// our answer is the length of the intersection
	return max - min + 1
}

func stringToIntArray(str string) []int {
	tokens := strings.Split(strings.Trim(str, " "), " ")
	ret := make([]int, 0)
	for _, token := range tokens {
		if token == "" {
			continue
		}
		ret = append(ret, common.StringToInt(token))
	}
	return ret
}
