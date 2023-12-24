package day24

import (
	"aoc2023/common"
	"fmt"
	"math"
	"regexp"
)

const EPS float64 = 1e-10

func Part2(strLines []string) int {
	reLine2 := regexp.MustCompile(`^(\d+), +(\d+), +(\d+) @ +(-?\d+), +(-?\d+), +(-?\d+)$`)

	lines := make([]Line, len(strLines))
	for i, strLine := range strLines {
		tokens := reLine2.FindStringSubmatch(strLine)
		if len(tokens) != 7 {
			panic(fmt.Sprintf("%d wrong line: %s", len(tokens), strLine))
		}
		x, y, z, dx, dy, dz := common.StringToInt(tokens[1]), common.StringToInt(tokens[2]), common.StringToInt(tokens[3]), common.StringToInt(tokens[4]), common.StringToInt(tokens[5]), common.StringToInt(tokens[6])
		lines[i] = NewLine(x, y, z, dx, dy, dz)
	}
	// fmt.Println(lines)

	line1 := lines[0]
	line2 := lines[1]
	distToRest := make([]float64, len(lines)-2)
	var solutionLine Line
	var solutionT1 int
	var solutionT2 int
	foundSolution := false
	for t1 := 0; !foundSolution; t1++ {
		coord1 := line1.PositionAt(t1)

		for t2 := 0; ; t2++ {
			coord2 := line2.PositionAt(t2)
			tryLine := NewLine(coord1.x, coord1.y, coord1.z, coord2.x-coord1.x, coord2.y-coord1.y, coord2.z-coord1.z)
			if t2 == 0 {
				// init dist to remaining n-2 lines
				allZero := true
				for i := range distToRest {
					distToRest[i] = dist(tryLine, lines[i+2])
					if !isApproxZero(distToRest[i]) {
						allZero = false
					}
				}
				if allZero {
					// found solution
					foundSolution = true
					solutionLine = tryLine
					solutionT1 = t1
					solutionT2 = t2
					break
				}
			} else {
				// compare to previous dist
				allZero := true
				goodDirection := true
				for i := range distToRest {
					newDist := dist(tryLine, lines[i+2])
					if newDist > distToRest[i] {
						// going in the wrong direction > break
						goodDirection = false
						break
					} else {
						distToRest[i] = newDist
					}
					if !isApproxZero(distToRest[i]) {
						allZero = false
					}
				}
				if allZero {
					// found solution
					foundSolution = true
					solutionLine = tryLine
					solutionT1 = t1
					solutionT2 = t2
					break
				}
				if !goodDirection || foundSolution {
					break
				}
			}
		}

		if t1 == 20 {
			break
		}
	}
	fmt.Println(solutionT1, solutionT2, solutionLine)

	t2MinusT1 := solutionT2 - solutionT1
	diffT2T1 := sub(line2.PositionAt(solutionT2), line1.PositionAt(solutionT1))
	dirMinusOne := NewCoord(diffT2T1.x/t2MinusT1, diffT2T1.y/t2MinusT1, diffT2T1.z/t2MinusT1)
	origin := sub(line1.PositionAt(solutionT1), NewCoord(solutionT1*dirMinusOne.x, solutionT1*dirMinusOne.y, solutionT1*dirMinusOne.z))
	fmt.Println(origin)

	return origin.x + origin.y + origin.z
}

func isApproxZero(dist float64) bool {
	return dist < EPS
}

func dist(line1, line2 Line) float64 {
	bxd := crossProduct(line1.dir, line2.dir)
	return math.Abs(float64(dotProd(sub(line2.start, line1.start), bxd))) / bxd.Length()
}

func crossProduct(a, b Coord) Coord {
	return NewCoord(a.y*b.z-a.z*b.y, a.z*b.x-a.x*b.z, a.x*b.y-a.y*b.x)
}

func sub(a, b Coord) Coord {
	return NewCoord(a.x-b.x, a.y-b.y, a.z-b.z)
}

func dotProd(a, b Coord) int {
	return a.x*b.x + a.y*b.y + a.z*b.z
}
