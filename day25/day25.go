package day25

import (
	"fmt"
	"strings"
)

func Part1(lines []string) int {
	nodes := make(map[string]int)
	maxGroup := 0

	for _, line := range lines {
		tokens := strings.Split(line, ": ")
		from := tokens[0]
		tokens = strings.Split(tokens[1], " ")
		for _, to := range tokens {
			// fmt.Printf("  %s -- %s\n", from, to)
			if isEdge(from, to, "xxk", "cth") || isEdge(from, to, "nvt", "zdj") || isEdge(from, to, "bbm", "mzg") {
				continue
			}
			// if isEdge(from, to, "jqt", "nvd") || isEdge(from, to, "bvb", "cmg") || isEdge(from, to, "hfx", "pzl") {
			// 	continue
			// }

			if nodes[from] != 0 {
				if nodes[to] != 0 {
					// merge groups
					relabel(&nodes, nodes[to], nodes[from])
				} else {
					nodes[to] = nodes[from]
				}
			} else {
				if nodes[to] != 0 {
					nodes[from] = nodes[to]
				} else {
					maxGroup++
					nodes[from] = maxGroup
					nodes[to] = maxGroup
				}
			}
		}
	}
	// fmt.Println(nodes)

	countByVal := make(map[int]int)
	for _, v := range nodes {
		if countByVal[v] == 0 {
			countByVal[v] = 1
		} else {
			countByVal[v]++
		}
	}
	fmt.Println(countByVal)

	prod := 1
	for _, v := range countByVal {
		prod *= v
	}

	return prod
}

func isEdge(from, to, targetFrom, targetTo string) bool {
	return (from == targetFrom && to == targetTo) || (from == targetTo && to == targetFrom)
}

func relabel(nodes *map[string]int, from, to int) {
	for k, v := range *nodes {
		if v == from {
			(*nodes)[k] = to
		}
	}
}
