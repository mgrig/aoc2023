package day24

import (
	"aoc2023/common"
	"fmt"
	"regexp"
)

const EPS float64 = 1e-8

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

	for l := 0; l <= 2; l++ {
		line := lines[l]
		time := fmt.Sprintf("t%d", l+1)
		fmt.Printf("ox + %s * dx = %d + %s * %d\n", time, line.start.x, time, line.dir.x)
		fmt.Printf("oy + %s * dy = %d + %s * %d\n", time, line.start.y, time, line.dir.y)
		fmt.Printf("oz + %s * dz = %d + %s * %d\n", time, line.start.z, time, line.dir.z)
	}

	// Solve with an external solver. The result is ox + oy + oz

	return 0

	// line1 := lines[0]
	// line2 := lines[1]
	// distToRest := make([]float64, len(lines)-2)
	// var solutionLine Line
	// var solutionT1 int
	// var solutionT2 int
	// foundSolution := false
	// for t1 := 0; !foundSolution; t1++ {
	// 	coord1 := line1.PositionAt(t1)

	// 	for t2 := 0; ; t2++ {
	// 		coord2 := line2.PositionAt(t2)
	// 		tryLine := NewLine(coord1.x, coord1.y, coord1.z, coord2.x-coord1.x, coord2.y-coord1.y, coord2.z-coord1.z)
	// 		if t2 == 0 {
	// 			// init dist to remaining n-2 lines
	// 			allZero := true
	// 			for i := range distToRest {
	// 				distToRest[i] = dist(tryLine, lines[i+2])
	// 				if !isApproxZero(distToRest[i]) {
	// 					allZero = false
	// 				}
	// 			}
	// 			if allZero {
	// 				// found solution
	// 				foundSolution = true
	// 				solutionLine = tryLine
	// 				solutionT1 = t1
	// 				solutionT2 = t2
	// 				break
	// 			}
	// 		} else {
	// 			// compare to previous dist
	// 			allZero := true
	// 			goodDirection := true
	// 			for i := range distToRest {
	// 				newDist := dist(tryLine, lines[i+2])
	// 				if newDist > distToRest[i] {
	// 					// going in the wrong direction > break
	// 					goodDirection = false
	// 					break
	// 				} else {
	// 					distToRest[i] = newDist
	// 				}
	// 				if !isApproxZero(distToRest[i]) {
	// 					allZero = false
	// 				}
	// 			}
	// 			if allZero {
	// 				// found solution
	// 				foundSolution = true
	// 				solutionLine = tryLine
	// 				solutionT1 = t1
	// 				solutionT2 = t2
	// 				break
	// 			}
	// 			if !goodDirection || foundSolution {
	// 				break
	// 			}
	// 		}
	// 	}
	// 	if t1%1000 == 0 {
	// 		fmt.Println(t1)
	// 	}
	// }
	// fmt.Println(solutionT1, solutionT2, solutionLine)

	// t2MinusT1 := solutionT2 - solutionT1
	// diffT2T1 := sub(line2.PositionAt(solutionT2), line1.PositionAt(solutionT1))
	// dirMinusOne := NewCoord(diffT2T1.x/t2MinusT1, diffT2T1.y/t2MinusT1, diffT2T1.z/t2MinusT1)
	// origin := sub(line1.PositionAt(solutionT1), NewCoord(solutionT1*dirMinusOne.x, solutionT1*dirMinusOne.y, solutionT1*dirMinusOne.z))
	// fmt.Println("origin:", origin)

	// return origin.x + origin.y + origin.z
}

// func isApproxZero(dist float64) bool {
// 	return dist < EPS
// }

// func dist(line1, line2 Line) float64 {
// 	bxd := crossProduct(line1.dir, line2.dir)
// 	dotProdBig := new(big.Float).Abs(new(big.Float).SetInt(dotProd(sub(line2.start, line1.start), bxd)))
// 	retBig := new(big.Float).Quo(dotProdBig, big.NewFloat(bxd.Length()))
// 	ret, acc := retBig.Float64()
// 	if ret == math.Inf(1) || ret == math.Inf(-1) {
// 		panic(fmt.Sprintf("too large for float64, retBig=%v, ret=%v, acc=%v", retBig, ret, acc))
// 	}
// 	return ret
// }

// func crossProduct(a, b Coord) Coord {
// 	return NewCoord(mulSafe(a.y, b.z)-mulSafe(a.z, b.y), mulSafe(a.z, b.x)-mulSafe(a.x, b.z), mulSafe(a.x, b.y)-mulSafe(a.y, b.x))
// }

// func sub(a, b Coord) Coord {
// 	return NewCoord(a.x-b.x, a.y-b.y, a.z-b.z)
// }

// func dotProd(a, b Coord) *big.Int {
// 	return sumBig(mulIntBig(a.x, b.x), mulIntBig(a.y, b.y), mulIntBig(a.z, b.z))
// }

// func sumBig(values ...*big.Int) *big.Int {
// 	sum := big.NewInt(0)
// 	for _, val := range values {
// 		sum.Add(sum, val)
// 	}
// 	return sum
// }

// func mulIntBig(x, y int) *big.Int {
// 	xbig := big.NewInt(int64(x))
// 	ybig := big.NewInt(int64(y))
// 	return big.NewInt(0).Mul(xbig, ybig)
// }

// // panic if overflow
// func mulSafe(a, b int) int {
// 	result := a * b
// 	if a == 0 || b == 0 || a == 1 || b == 1 {
// 		return result
// 	}
// 	if a == math.MinInt || b == math.MinInt {
// 		panic(fmt.Sprintf("Overflow 1 multiplying %v and %v", a, b))
// 	}
// 	if result/b != a {
// 		panic(fmt.Sprintf("Overflow 2 multiplying %v and %v", a, b))
// 	}
// 	return result
// }
