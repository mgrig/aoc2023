package day19

import (
	"fmt"
	"regexp"
)

func Part1(lines []string) int {
	reWorkflow := regexp.MustCompile(`^(.+)\{(.+)\}$`)
	rePart := regexp.MustCompile(`^\{(.+)\}$`)

	workflows := make(map[string]Workflow) // workflow name -> workflow
	parts := make([]Part, 0)

	for _, line := range lines {
		tokens := reWorkflow.FindStringSubmatch(line)
		if len(tokens) == 3 {
			name := tokens[1]
			workflows[name] = NewWorkflow(name, ParseRules(tokens[2]))
			continue
		}

		tokens = rePart.FindStringSubmatch(line)
		if len(tokens) != 2 {
			panic("wrong part")
		}
		parts = append(parts, ParsePart(tokens[1]))
	}
	fmt.Println("#workflows", len(workflows), "#parts", len(parts))

	sum := 0
	for _, part := range parts {
		if send(part, &workflows) {
			sum += part.Sum()
		}
	}

	return sum
}

func Part2(lines []string) int {
	reWorkflow := regexp.MustCompile(`^(.+)\{(.+)\}$`)

	workflows := make(map[string]Workflow) // workflow name -> workflow

	for _, line := range lines {
		tokens := reWorkflow.FindStringSubmatch(line)
		if len(tokens) == 3 {
			name := tokens[1]
			workflows[name] = NewWorkflow(name, ParseRules(tokens[2]))
		} else {
			break
		}
	}
	// fmt.Println("#workflows", len(workflows))

	return rec(NewRange4(), workflows["in"], &workflows)

}

func rec(r4 Range4, w Workflow, workflows *map[string]Workflow) (accepted int) {
	for _, rule := range w.rules {
		if rule.condition == nil {
			if rule.whereTo == "A" {
				accepted += r4.Prod()
				return
			}
			if rule.whereTo == "R" {
				return accepted
			}
			return accepted + rec(r4, (*workflows)[rule.whereTo], workflows)
		} else {
			cond := (*rule.condition).(SmallerLarger) // ugly "instance of". Maybe Condition interface not needed at all?
			r1 := r4.ranges[cond.category]
			// is split needed?
			if (cond.sign == "<" && cond.value > r1.min && cond.value <= r1.max) || (cond.sign == ">" && cond.value >= r1.min && cond.value < r1.max) {
				// range will split!
				switch cond.sign {
				case "<":
					left4, right4 := r4.SplitSmallerThan(cond.category, cond.value)
					if rule.whereTo == "A" {
						accepted += left4.Prod()
					} else if rule.whereTo == "R" {
						// nop: accepted += 0
					} else {
						accepted += rec(left4, (*workflows)[rule.whereTo], workflows)
					}
					r4 = right4 // continue with next rule
				case ">":
					left4, right4 := r4.SplitSmallerThan(cond.category, cond.value+1)
					if rule.whereTo == "A" {
						accepted += right4.Prod()
					} else if rule.whereTo == "R" {
						// nop: accepted += 0
					} else {
						accepted += rec(right4, (*workflows)[rule.whereTo], workflows)
					}
					r4 = left4 // continue with next rule
				default:
					panic("wrong sign")
				}
			} else {
				// no split needed
				switch cond.sign {
				case "<":
					if cond.value > r1.max {
						// cond applies to whole range
						if rule.whereTo == "A" {
							accepted += r4.Prod()
							return
						}
						if rule.whereTo == "R" {
							return accepted
						}
						return accepted + rec(r4, (*workflows)[rule.whereTo], workflows)
					} // else NOP, continue with next rule
				case ">":
					if cond.value < r1.min {
						// cond applies to whole range
						if rule.whereTo == "A" {
							accepted += r4.Prod()
							return
						}
						if rule.whereTo == "R" {
							return accepted
						}
						return accepted + rec(r4, (*workflows)[rule.whereTo], workflows)
					} // else NOP, continue with next rule
				default:
					panic("wrong sign")
				}
			}
		}
	}
	panic(fmt.Sprintf("no rule applies? %v, %v", r4, w))
}

func send(part Part, workflows *map[string]Workflow) (accepted bool) {
	w := (*workflows)["in"]
	for {
		for _, rule := range w.rules {
			if rule.condition == nil || (*rule.condition).Apply(part) {
				if rule.whereTo == "A" {
					return true
				}
				if rule.whereTo == "R" {
					return false
				}
				w = (*workflows)[rule.whereTo]
				break
			}
		}
	}
}
