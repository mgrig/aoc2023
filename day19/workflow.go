package day19

import "strings"

type Workflow struct {
	name  string
	rules []Rule
}

func NewWorkflow(name string, rules []Rule) Workflow {
	return Workflow{
		name:  name,
		rules: rules,
	}
}

func ParseRules(str string) []Rule {
	rules := make([]Rule, 0)
	tokens := strings.Split(str, ",")
	for _, token := range tokens {
		rules = append(rules, ParseRule(token))
	}
	return rules
}
