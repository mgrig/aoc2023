package day08

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

func Part1(lines []string) int {
	instr := lines[0]
	fmt.Println(instr)

	reNode := regexp.MustCompile(`^(.*?) = \((.*?), (.*?)\)`)
	g := NewGraph()
	for i := 1; i < len(lines); i++ {
		tokens := reNode.FindStringSubmatch(lines[i])
		g.Add(tokens[1], tokens[2], tokens[3])
	}
	// fmt.Println(*g)

	node := "AAA"
	steps := 0
	for node != "ZZZ" {
		i := steps % len(instr)
		dir := instr[i]
		if dir == 'L' {
			node = g.GetLeft(node)
		} else {
			node = g.GetRight(node)
		}
		steps++
	}

	return steps
}

func Part2(lines []string) int {
	instr := lines[0]
	// fmt.Println(instr)

	reNode := regexp.MustCompile(`^(.*?) = \((.*?), (.*?)\)`)
	g := NewGraph()
	starts := make(map[string]bool) // set
	// ends := make(map[string]bool)   // set
	for i := 1; i < len(lines); i++ {
		tokens := reNode.FindStringSubmatch(lines[i])
		g.Add(tokens[1], tokens[2], tokens[3])
		if strings.HasSuffix(tokens[1], "A") {
			starts[tokens[1]] = true
		}
		// else if strings.HasSuffix(tokens[1], "Z") {
		// 	ends[tokens[1]] = true
		// }
	}
	// fmt.Println(starts, ends)

	// compute the length of each path (from a start to an end) and get LCM for the final result
	var gcd *big.Int = nil
	var lcm *big.Int = nil
	for start := range starts {
		node := start
		steps := 0
		for !strings.HasSuffix(node, "Z") {
			i := steps % len(instr)
			dir := instr[i]
			if dir == 'L' {
				node = g.GetLeft(node)
			} else {
				node = g.GetRight(node)
			}
			steps++
		}
		fmt.Println(start, steps)

		if gcd == nil {
			gcd = big.NewInt(int64(steps))
			lcm = big.NewInt(int64(steps))
		} else {
			b := big.NewInt(int64(steps))
			gcd.GCD(nil, nil, gcd, b)
			lcm = lcm.Mul(lcm, b.Div(b, gcd))
		}
	}
	fmt.Println("gcd:", gcd)

	return int(lcm.Int64())
}
