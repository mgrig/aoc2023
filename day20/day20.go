package day20

import (
	"fmt"
	"regexp"
	"strings"
)

func Part1(lines []string) int {
	reLine := regexp.MustCompile(`^([%&]?)(\w+) -> (.*?)$`)

	modules := make(map[string]Module, len(lines))

	for _, line := range lines {
		tokens := reLine.FindStringSubmatch(line)
		// fmt.Println(len(tokens), tokens)
		if len(tokens) != 4 {
			panic("wrong line")
		}

		typ, from := tokens[1], tokens[2]
		dest := splitDest(tokens[3])

		switch typ {
		case "":
			if from == "broadcaster" {
				modules[from] = NewBroadcast(from, dest)
			} else {
				modules[from] = NewUntyped(from)
			}
		case "%":
			modules[from] = NewFlipFlop(from, dest)
		case "&":
			modules[from] = NewConj(from, dest)
		default:
			panic("wrong type")
		}
	}

	// Conj modules don't know their inputs yet. Loop over all dests to create them
	for _, module := range modules {
		from := module.GetName()
		dest := module.GetDest()
		for _, dst := range dest {
			fmt.Println(from, dst)
			m, exists := modules[dst]
			if !exists {
				continue // skip untyped (e.g. "output")
			}
			m.AddIn(from)
		}
	}

	pulses := make([]Pulse, 0)

	countLow := 0
	countHigh := 0
	for b := 1; b <= 1000; b++ {
		pulses = append(pulses, NewPulse("", "broadcaster", LOW))
		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]

			if pulse.typ == LOW {
				countLow++
			} else {
				countHigh++
			}
			m, exists := modules[pulse.to]
			if !exists {
				continue // skip untyped (e.g. "output")
			}
			emitted := m.ProcessPulse(pulse)
			pulses = append(pulses, emitted...)
		}
	}
	fmt.Println("low", countLow, "high", countHigh)

	return countLow * countHigh
}

func splitDest(destStr string) []string {
	tokens := strings.Split(destStr, ", ")
	ret := make([]string, len(tokens))
	for i, token := range tokens {
		ret[i] = strings.Trim(token, " ")
	}
	return ret
}
