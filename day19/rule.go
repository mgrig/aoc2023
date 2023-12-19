package day19

import (
	"strings"
)

type Rule struct {
	condition *Condition // nil means "no condition" -> always applies
	whereTo   string
}

// a<2006:qkq or A or qkq
func ParseRule(str string) Rule {
	rule := Rule{}
	if strings.Contains(str, ":") {
		tokens := strings.Split(str, ":")
		cond := ParseCondition(tokens[0])
		rule.condition = &cond
		rule.whereTo = tokens[1]
	} else {
		rule.whereTo = str
	}
	return rule
}
