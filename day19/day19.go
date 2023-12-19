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
