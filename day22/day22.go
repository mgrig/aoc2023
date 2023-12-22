package day22

import (
	"aoc2023/common"
	"fmt"
	"regexp"
	"sort"
)

func Part1(lines []string) int {
	reBrick := regexp.MustCompile(`^(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)$`)

	bricks := []Brick{}

	var minx, maxx, miny, maxy int
	for r, line := range lines {
		tokens := reBrick.FindStringSubmatch(line)
		if len(tokens) != 7 {
			panic("wrong line")
		}
		from := NewCoord(common.StringToInt(tokens[1]), common.StringToInt(tokens[2]), common.StringToInt(tokens[3]))
		to := NewCoord(common.StringToInt(tokens[4]), common.StringToInt(tokens[5]), common.StringToInt(tokens[6]))

		// fmt.Println(r, from, to)
		brick := NewBrick(from, to)
		bricks = append(bricks, brick)

		if r == 0 {
			minx, maxx, miny, maxy, _, _ = brick.OuterBox()
		} else {
			bminx, bmaxx, bminy, bmaxy, _, _ := brick.OuterBox()
			minx = common.IntMin(minx, bminx)
			maxx = common.IntMax(maxx, bmaxx)
			miny = common.IntMin(miny, bminy)
			maxy = common.IntMax(maxy, bmaxy)
		}
	}
	fmt.Println(minx, maxx, miny, maxy)

	// sort bricks by their min z value (ascending)
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].MinZ() < bricks[j].MinZ()
	})

	// let all bricks fall
	n := maxx - minx + 1 // we know from data that size is the same on x and y (can generalize if needed)
	surf := NewGrid(n)
	maxZ := 0
	for i := range bricks {
		hProj := bricks[i].HorizProjection()

		maxH := 0
		for _, coord := range hProj {
			h := surf.grid[coord.x][coord.y]
			maxH = common.IntMax(maxH, h)
		}
		// this brick will rest at level maxH+1

		minZ := bricks[i].MinZ()
		if minZ > maxH+1 {
			diff := minZ - (maxH + 1)
			bricks[i].Lower(diff)
			// fmt.Printf("lower brick %d with %d \n", i, diff)
		}
		maxZ = common.IntMax(maxZ, bricks[i].MaxZ())

		// update surf
		for _, coord := range hProj {
			surf.grid[coord.x][coord.y] = bricks[i].MaxZ()
		}
	}

	// build 3d grid - to know which coord belongs to which brick (id of brick will be the index in the above sorted slice)
	// fmt.Println("maxZ", maxZ)
	g3 := NewGrid3(n, maxZ)
	for i, brick := range bricks {
		coords := brick.AllCoords()
		for _, coord := range coords {
			g3.grid[coord.x][coord.y][coord.z] = i
		}
	}

	// build Nodes
	nodes := make([]Node, len(bricks))
	for i := range bricks {
		nodes[i] = *NewNode(i)
	}
	for i, brick := range bricks {
		hProj := brick.HorizProjection()
		z := brick.MaxZ() + 1
		for _, coord := range hProj {
			other := g3.grid[coord.x][coord.y][z]
			if other != 0 {
				nodes[i].Supports(other)
				nodes[other].SupportedBy(i)
			}
		}
	}

	// count nodes for which all next (Supports) have more than 1 SupportedBy
	count := 0
	for i := range nodes {
		if AllNextHaveMorePrev(&nodes, i) {
			// fmt.Println("can remove brick", i)
			count++
		}
	}

	return count
}

func AllNextHaveMorePrev(nodes *[]Node, index int) bool {
	node := (*nodes)[index]
	for next := range node.supports {
		if len((*nodes)[next].supportedBy) <= 1 {
			return false
		}
	}
	return true
}
