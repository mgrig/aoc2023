package day24

import (
	"aoc2023/common"
	"fmt"
	"math"
	"math/big"
	"regexp"
)

func Part1(strLines []string) int {
	reLine2 := regexp.MustCompile(`^(\d+), +(\d+), +\d+ @ +(-?\d+), +(-?\d+), +-?\d+$`)

	lines := make([]Line, len(strLines))
	for i, strLine := range strLines {
		tokens := reLine2.FindStringSubmatch(strLine)
		if len(tokens) != 5 {
			panic(fmt.Sprintf("%d wrong line: %s", len(tokens), strLine))
		}
		x, y, dx, dy := common.StringToInt(tokens[1]), common.StringToInt(tokens[2]), common.StringToInt(tokens[3]), common.StringToInt(tokens[4])
		lines[i] = NewLine(x, y, dx, dy)
	}
	// fmt.Println(lines)

	min := 200000000000000
	max := 400000000000000
	// min := 7
	// max := 27
	count := 0
	for i := 0; i < len(lines)-1; i++ {
		line1 := lines[i]
		for j := i + 1; j < len(lines); j++ {
			line2 := lines[j]
			px, py, theyIntersect := intersectInFuture(line1, line2)
			if theyIntersect && inside(px, py, min, max) {
				// fmt.Println(px, py)
				count++
			}
		}
	}

	return count
}

func inside(px, py float64, min, max int) bool {
	return px >= float64(min) && px <= float64(max) && py >= float64(min) && py <= float64(max)
}

func intersectInFuture(line1, line2 Line) (px, py float64, theyIntersect bool) {
	x1 := line1.x
	x2 := line1.x2
	y1 := line1.y
	y2 := line1.y2

	x3 := line2.x
	x4 := line2.x2
	y3 := line2.y
	y4 := line2.y2

	josInt := line1.dx*line2.dy - line1.dy*line2.dx // checked it does not overflow with the input data
	if josInt == 0 {
		return 0, 0, false
	}
	jos := float64(josInt)

	x1big := big.NewInt(int64(x1))
	x2big := big.NewInt(int64(x2))
	y1big := big.NewInt(int64(y1))
	y2big := big.NewInt(int64(y2))
	t1big := big.NewInt(0).Sub(big.NewInt(0).Mul(x1big, y2big), big.NewInt(0).Mul(y1big, x2big))
	if !t1big.IsInt64() {
		panic("t1 is not an int64")
	}
	t1 := int(t1big.Int64())

	// t2 := x3*y4 - y3*x4 // x3 y4 - y3 x4
	x3big := big.NewInt(int64(x3))
	y3big := big.NewInt(int64(y3))
	x4big := big.NewInt(int64(x4))
	y4big := big.NewInt(int64(y4))
	t2big := big.NewInt(0).Sub(big.NewInt(0).Mul(x3big, y4big), big.NewInt(0).Mul(y3big, x4big))
	if !t2big.IsInt64() {
		panic("t2 is not an int64")
	}
	t2 := int(t2big.Int64())

	josBigFloat := big.NewFloat(jos)
	// px = float64(line1.dx*t2-line2.dx*t1) / jos
	diff, _ := big.NewInt(0).Sub(multiplyBig(line1.dx, t2), multiplyBig(line2.dx, t1)).Float64()
	px, _ = new(big.Float).Quo(big.NewFloat(diff), josBigFloat).Float64()

	// py = float64(line1.dy*t2-line2.dy*t1) / jos
	diff, _ = big.NewInt(0).Sub(multiplyBig(line1.dy, t2), multiplyBig(line2.dy, t1)).Float64()
	py, _ = new(big.Float).Quo(big.NewFloat(diff), josBigFloat).Float64()

	// check intersection is in the future
	if !inTheFuture(px, py, line1) || !inTheFuture(px, py, line2) {
		return 0, 0, false
	}

	return px, py, true
}

func multiplyBig(x, y int) *big.Int {
	xbig := big.NewInt(int64(x))
	ybig := big.NewInt(int64(y))
	return big.NewInt(0).Mul(xbig, ybig)
}

// panic if overflow
func multiplyExact(a, b int) int {
	result := a * b
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return result
	}
	if a == math.MinInt || b == math.MinInt {
		panic(fmt.Sprintf("Overflow 1 multiplying %v and %v", a, b))
	}
	if result/b != a {
		panic(fmt.Sprintf("Overflow 2 multiplying %v and %v", a, b))
	}
	return result
}

func inTheFuture(px, py float64, line Line) bool {
	checkY := sgn(py-float64(line.y)) == sgnInt(line.dy)
	if line.dx == 0 {
		return checkY
	}
	checkX := sgn(px-float64(line.x)) == sgnInt(line.dx)
	if line.dy == 0 {
		return checkX
	}
	return checkX && checkY
}

func sgnInt(val int) int {
	if val < 0 {
		return -1
	}
	if val > 0 {
		return 1
	}
	return 0
}

func sgn(val float64) int {
	if val < 0 {
		return -1
	}
	if val > 0 {
		return 1
	}
	return 0
}
