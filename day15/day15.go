package day15

import (
	"aoc2023/common"
	"fmt"
	"regexp"
	"strings"
)

func Part1(line string) uint32 {
	var sum uint32 = 0

	tokens := strings.Split(line, ",")
	for _, t := range tokens {
		sum += Hash(strings.Trim(t, " "))
	}

	return sum
}

func Part2(line string) int {
	reDash := regexp.MustCompile(`(.*?)-`)
	reEqual := regexp.MustCompile(`(.*?)=(\d)`)

	boxes := make([]*Box, 256)
	for i := range boxes {
		boxes[i] = NewBox()
	}

	lenses := make(map[string]Lens) // label -> lens

	tokens := strings.Split(line, ",")
	for _, t := range tokens {
		groupsDash := reDash.FindStringSubmatch(t)
		groupsEqual := reEqual.FindStringSubmatch(t)

		if len(groupsDash) == 2 { // found
			label := groupsDash[1]
			hash := Hash(label)
			oldLens, found := boxes[hash].Remove(label)
			if found {
				_, exists := lenses[oldLens.label]
				if !exists {
					lenses[oldLens.label] = oldLens
				}
			}
		} else if len(groupsEqual) == 3 { // found
			label := groupsEqual[1]
			focus := common.StringToInt(groupsEqual[2])
			hash := Hash(label)
			oldLens, found := boxes[hash].Append(NewLens(label, focus))
			if found {
				_, exists := lenses[oldLens.label]
				if !exists {
					lenses[oldLens.label] = oldLens
				}
			}
		} else {
			panic("wrong line")
		}
		// fmt.Println(t)
		// fmt.Println(BoxesToString(&boxes))
	}

	sum := 0
	for ib, b := range boxes {
		if len(b.lenses) > 0 {
			for il, lens := range b.lenses {
				power := (ib + 1) * (il + 1) * lens.focal
				sum += power
			}
		}
	}

	return sum
}

func BoxesToString(boxes *[]*Box) string {
	ret := ""
	for i, b := range *boxes {
		if len(b.lenses) > 0 {
			ret += fmt.Sprintf("box %d: %v\n", i, b.lenses)
		}
	}
	return ret
}

func Hash(line string) uint32 {
	var hash uint32 = 0
	for _, r := range line {
		hash += uint32(r)
		hash *= 17
		hash %= 256
	}
	return hash
}
